// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"errors"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/axiacoin/axia-network-v2/cache"
	"github.com/axiacoin/axia-network-v2/cache/metercacher"
	"github.com/axiacoin/axia-network-v2/database"
	"github.com/axiacoin/axia-network-v2/database/linkeddb"
	"github.com/axiacoin/axia-network-v2/database/prefixdb"
	"github.com/axiacoin/axia-network-v2/database/versiondb"
	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/snow/choices"
	"github.com/axiacoin/axia-network-v2/snow/uptime"
	"github.com/axiacoin/axia-network-v2/utils/constants"
	"github.com/axiacoin/axia-network-v2/utils/hashing"
	"github.com/axiacoin/axia-network-v2/utils/uint128"
	"github.com/axiacoin/axia-network-v2/utils/wrappers"
	"github.com/axiacoin/axia-network-v2/vms/components/axc"
	"github.com/axiacoin/axia-network-v2/vms/platformvm/status"

	safemath "github.com/axiacoin/axia-network-v2/utils/math"
)

var (
	validatorsPrefix         = []byte("validators")
	currentPrefix            = []byte("current")
	pendingPrefix            = []byte("pending")
	validatorPrefix          = []byte("validator")
	nominatorPrefix          = []byte("nominator")
	allychainValidatorPrefix = []byte("allychainValidator")
	validatorDiffsPrefix     = []byte("validatorDiffs")
	blockPrefix              = []byte("block")
	txPrefix                 = []byte("tx")
	rewardUTXOsPrefix        = []byte("rewardUTXOs")
	utxoPrefix               = []byte("utxo")
	allychainPrefix          = []byte("allychain")
	chainPrefix              = []byte("chain")
	singletonPrefix          = []byte("singleton")

	timestampKey     = []byte("timestamp")
	currentSupplyKey = []byte("current supply")
	lastAcceptedKey  = []byte("last accepted")
	initializedKey   = []byte("initialized")

	errWrongNetworkID = errors.New("tx has wrong network ID")

	_ InternalState = &internalStateImpl{}
)

const (
	// priority values are used as part of the keys in the pending/current
	// validator state to ensure they are sorted in the order that they should
	// be added/removed.
	lowPriority byte = iota
	mediumPriority
	topPriority

	validatorDiffsCacheSize = 2048
	blockCacheSize          = 2048
	txCacheSize             = 2048
	rewardUTXOsCacheSize    = 2048
	chainCacheSize          = 2048
	chainDBCacheSize        = 2048
)

type InternalState interface {
	MutableState
	uptime.State
	axc.UTXOReader

	SetHeight(height uint64)

	AddCurrentStaker(tx *Tx, potentialReward uint64)
	DeleteCurrentStaker(tx *Tx)
	GetValidatorWeightDiffs(height uint64, allychainID ids.ID) (map[ids.ShortID]*ValidatorWeightDiff, error)

	AddPendingStaker(tx *Tx)
	DeletePendingStaker(tx *Tx)

	SetCurrentStakerChainState(currentStakerChainState)
	SetPendingStakerChainState(pendingStakerChainState)

	GetLastAccepted() ids.ID
	SetLastAccepted(ids.ID)

	GetBlock(blockID ids.ID) (Block, error)
	AddBlock(block Block)

	Abort()
	Commit() error
	CommitBatch() (database.Batch, error)
	Close() error
}

/*
 * VMDB
 * |-. validators
 * | |-. current
 * | | |-. validator
 * | | | '-. list
 * | | |   '-- txID -> uptime + potential reward
 * | | |-. nominator
 * | | | '-. list
 * | | |   '-- txID -> potential reward
 * | | '-. allychainValidator
 * | |   '-. list
 * | |     '-- txID -> nil
 * | |-. pending
 * | | |-. validator
 * | | | '-. list
 * | | |   '-- txID -> nil
 * | | |-. nominator
 * | | | '-. list
 * | | |   '-- txID -> nil
 * | | '-. allychainValidator
 * | |   '-. list
 * | |     '-- txID -> nil
 * | '-. diffs
 * |   '-. height+allychain
 * |     '-. list
 * |       '-- nodeID -> weightChange
 * |-. blocks
 * | '-- blockID -> block bytes
 * |-. txs
 * | '-- txID -> tx bytes + tx status
 * |- rewardUTXOs
 * | '-. txID
 * |   '-. list
 * |     '-- utxoID -> utxo bytes
 * |- utxos
 * | '-- utxoDB
 * |-. allychains
 * | '-. list
 * |   '-- txID -> nil
 * |-. chains
 * | '-. allychainID
 * |   '-. list
 * |     '-- txID -> nil
 * '-. singletons
 *   |-- initializedKey -> nil
 *   |-- timestampKey -> timestamp
 *   |-- currentSupplyKey -> currentSupply
 *   '-- lastAcceptedKey -> lastAccepted
 */
type internalStateImpl struct {
	vm *VM

	baseDB *versiondb.Database

	currentStakerChainState currentStakerChainState
	pendingStakerChainState pendingStakerChainState

	currentHeight         uint64
	addedCurrentStakers   []*validatorReward
	deletedCurrentStakers []*Tx
	addedPendingStakers   []*Tx
	deletedPendingStakers []*Tx
	uptimes               map[ids.ShortID]*currentValidatorState // nodeID -> uptimes
	updatedUptimes        map[ids.ShortID]struct{}               // nodeID -> nil

	validatorsDB                    database.Database
	currentValidatorsDB             database.Database
	currentValidatorBaseDB          database.Database
	currentValidatorList            linkeddb.LinkedDB
	currentNominatorBaseDB          database.Database
	currentNominatorList            linkeddb.LinkedDB
	currentAllychainValidatorBaseDB database.Database
	currentAllychainValidatorList   linkeddb.LinkedDB
	pendingValidatorsDB             database.Database
	pendingValidatorBaseDB          database.Database
	pendingValidatorList            linkeddb.LinkedDB
	pendingNominatorBaseDB          database.Database
	pendingNominatorList            linkeddb.LinkedDB
	pendingAllychainValidatorBaseDB database.Database
	pendingAllychainValidatorList   linkeddb.LinkedDB

	validatorDiffsCache cache.Cacher // cache of heightWithAllychain -> map[ids.ShortID]*ValidatorWeightDiff
	validatorDiffsDB    database.Database

	addedBlocks map[ids.ID]Block // map of blockID -> Block
	blockCache  cache.Cacher     // cache of blockID -> Block, if the entry is nil, it is not in the database
	blockDB     database.Database

	addedTxs map[ids.ID]*txStatusImpl // map of txID -> {*Tx, Status}
	txCache  cache.Cacher             // cache of txID -> {*Tx, Status} if the entry is nil, it is not in the database
	txDB     database.Database

	addedRewardUTXOs map[ids.ID][]*axc.UTXO // map of txID -> []*UTXO
	rewardUTXOsCache cache.Cacher           // cache of txID -> []*UTXO
	rewardUTXODB     database.Database

	modifiedUTXOs map[ids.ID]*axc.UTXO // map of modified UTXOID -> *UTXO if the UTXO is nil, it has been removed
	utxoDB        database.Database
	utxoState     axc.UTXOState

	cachedAllychains []*Tx // nil if the allychains haven't been loaded
	addedAllychains  []*Tx
	allychainBaseDB  database.Database
	allychainDB      linkeddb.LinkedDB

	addedChains  map[ids.ID][]*Tx // maps allychainID -> the newly added chains to the allychain
	chainCache   cache.Cacher     // cache of allychainID -> the chains after all local modifications []*Tx
	chainDBCache cache.Cacher     // cache of allychainID -> linkedDB
	chainDB      database.Database

	originalTimestamp, timestamp         time.Time
	originalCurrentSupply, currentSupply uint128.Uint128
	originalLastAccepted, lastAccepted   ids.ID
	singletonDB                          database.Database
}

type ValidatorWeightDiff struct {
	Decrease bool   `serialize:"true"`
	Amount   uint64 `serialize:"true"`
}

type heightWithAllychain struct {
	Height      uint64 `serialize:"true"`
	AllychainID ids.ID `serialize:"true"`
}

type stateTx struct {
	Tx     []byte        `serialize:"true"`
	Status status.Status `serialize:"true"`
}

type stateBlk struct {
	Blk    []byte         `serialize:"true"`
	Status choices.Status `serialize:"true"`
}

func newInternalStateDatabases(vm *VM, db database.Database) *internalStateImpl {
	baseDB := versiondb.New(db)

	validatorsDB := prefixdb.New(validatorsPrefix, baseDB)

	currentValidatorsDB := prefixdb.New(currentPrefix, validatorsDB)
	currentValidatorBaseDB := prefixdb.New(validatorPrefix, currentValidatorsDB)
	currentNominatorBaseDB := prefixdb.New(nominatorPrefix, currentValidatorsDB)
	currentAllychainValidatorBaseDB := prefixdb.New(allychainValidatorPrefix, currentValidatorsDB)

	pendingValidatorsDB := prefixdb.New(pendingPrefix, validatorsDB)
	pendingValidatorBaseDB := prefixdb.New(validatorPrefix, pendingValidatorsDB)
	pendingNominatorBaseDB := prefixdb.New(nominatorPrefix, pendingValidatorsDB)
	pendingAllychainValidatorBaseDB := prefixdb.New(allychainValidatorPrefix, pendingValidatorsDB)

	validatorDiffsDB := prefixdb.New(validatorDiffsPrefix, validatorsDB)

	rewardUTXODB := prefixdb.New(rewardUTXOsPrefix, baseDB)
	utxoDB := prefixdb.New(utxoPrefix, baseDB)
	allychainBaseDB := prefixdb.New(allychainPrefix, baseDB)
	return &internalStateImpl{
		vm: vm,

		baseDB: baseDB,

		uptimes:        make(map[ids.ShortID]*currentValidatorState),
		updatedUptimes: make(map[ids.ShortID]struct{}),

		validatorsDB:                    validatorsDB,
		currentValidatorsDB:             currentValidatorsDB,
		currentValidatorBaseDB:          currentValidatorBaseDB,
		currentValidatorList:            linkeddb.NewDefault(currentValidatorBaseDB),
		currentNominatorBaseDB:          currentNominatorBaseDB,
		currentNominatorList:            linkeddb.NewDefault(currentNominatorBaseDB),
		currentAllychainValidatorBaseDB: currentAllychainValidatorBaseDB,
		currentAllychainValidatorList:   linkeddb.NewDefault(currentAllychainValidatorBaseDB),
		pendingValidatorsDB:             pendingValidatorsDB,
		pendingValidatorBaseDB:          pendingValidatorBaseDB,
		pendingValidatorList:            linkeddb.NewDefault(pendingValidatorBaseDB),
		pendingNominatorBaseDB:          pendingNominatorBaseDB,
		pendingNominatorList:            linkeddb.NewDefault(pendingNominatorBaseDB),
		pendingAllychainValidatorBaseDB: pendingAllychainValidatorBaseDB,
		pendingAllychainValidatorList:   linkeddb.NewDefault(pendingAllychainValidatorBaseDB),
		validatorDiffsDB:                validatorDiffsDB,

		addedBlocks: make(map[ids.ID]Block),
		blockDB:     prefixdb.New(blockPrefix, baseDB),

		addedTxs: make(map[ids.ID]*txStatusImpl),
		txDB:     prefixdb.New(txPrefix, baseDB),

		addedRewardUTXOs: make(map[ids.ID][]*axc.UTXO),
		rewardUTXODB:     rewardUTXODB,

		modifiedUTXOs: make(map[ids.ID]*axc.UTXO),
		utxoDB:        utxoDB,

		allychainBaseDB: allychainBaseDB,
		allychainDB:     linkeddb.NewDefault(allychainBaseDB),

		addedChains: make(map[ids.ID][]*Tx),
		chainDB:     prefixdb.New(chainPrefix, baseDB),

		singletonDB: prefixdb.New(singletonPrefix, baseDB),
	}
}

func (st *internalStateImpl) initCaches() {
	st.validatorDiffsCache = &cache.LRU{Size: validatorDiffsCacheSize}
	st.blockCache = &cache.LRU{Size: blockCacheSize}
	st.txCache = &cache.LRU{Size: txCacheSize}
	st.rewardUTXOsCache = &cache.LRU{Size: rewardUTXOsCacheSize}
	st.utxoState = axc.NewUTXOState(st.utxoDB, GenesisCodec)
	st.chainCache = &cache.LRU{Size: chainCacheSize}
	st.chainDBCache = &cache.LRU{Size: chainDBCacheSize}
}

func (st *internalStateImpl) initMeteredCaches(metrics prometheus.Registerer) error {
	validatorDiffsCache, err := metercacher.New(
		"validator_diffs_cache",
		metrics,
		&cache.LRU{Size: validatorDiffsCacheSize},
	)
	if err != nil {
		return err
	}

	blockCache, err := metercacher.New(
		"block_cache",
		metrics,
		&cache.LRU{Size: blockCacheSize},
	)
	if err != nil {
		return err
	}

	txCache, err := metercacher.New(
		"tx_cache",
		metrics,
		&cache.LRU{Size: txCacheSize},
	)
	if err != nil {
		return err
	}

	rewardUTXOsCache, err := metercacher.New(
		"reward_utxos_cache",
		metrics,
		&cache.LRU{Size: rewardUTXOsCacheSize},
	)
	if err != nil {
		return err
	}

	utxoState, err := axc.NewMeteredUTXOState(st.utxoDB, GenesisCodec, metrics)
	if err != nil {
		return err
	}

	chainCache, err := metercacher.New(
		"chain_cache",
		metrics,
		&cache.LRU{Size: chainCacheSize},
	)
	if err != nil {
		return err
	}

	chainDBCache, err := metercacher.New(
		"chain_db_cache",
		metrics,
		&cache.LRU{Size: chainDBCacheSize},
	)
	st.validatorDiffsCache = validatorDiffsCache
	st.blockCache = blockCache
	st.txCache = txCache
	st.rewardUTXOsCache = rewardUTXOsCache
	st.utxoState = utxoState
	st.chainCache = chainCache
	st.chainDBCache = chainDBCache
	return err
}

func (st *internalStateImpl) sync(genesis []byte) error {
	shouldInit, err := st.shouldInit()
	if err != nil {
		return fmt.Errorf(
			"failed to check if the database is initialized: %w",
			err,
		)
	}

	// If the database is empty, create the platform chain anew using the
	// provided genesis state
	if shouldInit {
		if err := st.init(genesis); err != nil {
			return fmt.Errorf(
				"failed to initialize the database: %w",
				err,
			)
		}
	}

	if err := st.load(); err != nil {
		return fmt.Errorf(
			"failed to load the database state: %w",
			err,
		)
	}
	return nil
}

func NewInternalState(vm *VM, db database.Database, genesis []byte) (InternalState, error) {
	is := newInternalStateDatabases(vm, db)
	is.initCaches()

	if err := is.sync(genesis); err != nil {
		// Drop any errors on close to return the first error
		_ = is.Close()

		return nil, err
	}
	return is, nil
}

func NewMeteredInternalState(vm *VM, db database.Database, genesis []byte, metrics prometheus.Registerer) (InternalState, error) {
	is := newInternalStateDatabases(vm, db)
	if err := is.initMeteredCaches(metrics); err != nil {
		// Drop any errors on close to return the first error
		_ = is.Close()

		return nil, err
	}

	if err := is.sync(genesis); err != nil {
		// Drop any errors on close to return the first error
		_ = is.Close()

		return nil, err
	}
	return is, nil
}

func (st *internalStateImpl) GetTimestamp() time.Time          { return st.timestamp }
func (st *internalStateImpl) SetTimestamp(timestamp time.Time) { st.timestamp = timestamp }

func (st *internalStateImpl) GetCurrentSupply() uint128.Uint128 { return st.currentSupply }
func (st *internalStateImpl) SetCurrentSupply(currentSupply uint128.Uint128) {
	st.currentSupply = currentSupply
}

func (st *internalStateImpl) GetLastAccepted() ids.ID             { return st.lastAccepted }
func (st *internalStateImpl) SetLastAccepted(lastAccepted ids.ID) { st.lastAccepted = lastAccepted }

func (st *internalStateImpl) GetAllychains() ([]*Tx, error) {
	if st.cachedAllychains != nil {
		return st.cachedAllychains, nil
	}

	allychainDBIt := st.allychainDB.NewIterator()
	defer allychainDBIt.Release()

	txs := []*Tx(nil)
	for allychainDBIt.Next() {
		allychainIDBytes := allychainDBIt.Key()
		allychainID, err := ids.ToID(allychainIDBytes)
		if err != nil {
			return nil, err
		}
		allychainTx, _, err := st.GetTx(allychainID)
		if err != nil {
			return nil, err
		}
		txs = append(txs, allychainTx)
	}
	if err := allychainDBIt.Error(); err != nil {
		return nil, err
	}
	txs = append(txs, st.addedAllychains...)
	st.cachedAllychains = txs
	return txs, nil
}

func (st *internalStateImpl) AddAllychain(createAllychainTx *Tx) {
	st.addedAllychains = append(st.addedAllychains, createAllychainTx)
	if st.cachedAllychains != nil {
		st.cachedAllychains = append(st.cachedAllychains, createAllychainTx)
	}
}

func (st *internalStateImpl) GetChains(allychainID ids.ID) ([]*Tx, error) {
	if chainsIntf, cached := st.chainCache.Get(allychainID); cached {
		return chainsIntf.([]*Tx), nil
	}
	chainDB := st.getChainDB(allychainID)
	chainDBIt := chainDB.NewIterator()
	defer chainDBIt.Release()

	txs := []*Tx(nil)
	for chainDBIt.Next() {
		chainIDBytes := chainDBIt.Key()
		chainID, err := ids.ToID(chainIDBytes)
		if err != nil {
			return nil, err
		}
		chainTx, _, err := st.GetTx(chainID)
		if err != nil {
			return nil, err
		}
		txs = append(txs, chainTx)
	}
	if err := chainDBIt.Error(); err != nil {
		return nil, err
	}
	txs = append(txs, st.addedChains[allychainID]...)
	st.chainCache.Put(allychainID, txs)
	return txs, nil
}

func (st *internalStateImpl) AddChain(createChainTxIntf *Tx) {
	createChainTx := createChainTxIntf.UnsignedTx.(*UnsignedCreateChainTx)
	allychainID := createChainTx.AllychainID
	st.addedChains[allychainID] = append(st.addedChains[allychainID], createChainTxIntf)
	if chainsIntf, cached := st.chainCache.Get(allychainID); cached {
		chains := chainsIntf.([]*Tx)
		chains = append(chains, createChainTxIntf)
		st.chainCache.Put(allychainID, chains)
	}
}

func (st *internalStateImpl) getChainDB(allychainID ids.ID) linkeddb.LinkedDB {
	if chainDBIntf, cached := st.chainDBCache.Get(allychainID); cached {
		return chainDBIntf.(linkeddb.LinkedDB)
	}
	rawChainDB := prefixdb.New(allychainID[:], st.chainDB)
	chainDB := linkeddb.NewDefault(rawChainDB)
	st.chainDBCache.Put(allychainID, chainDB)
	return chainDB
}

func (st *internalStateImpl) GetTx(txID ids.ID) (*Tx, status.Status, error) {
	if tx, exists := st.addedTxs[txID]; exists {
		return tx.tx, tx.status, nil
	}
	if txIntf, cached := st.txCache.Get(txID); cached {
		if txIntf == nil {
			return nil, status.Unknown, database.ErrNotFound
		}
		tx := txIntf.(*txStatusImpl)
		return tx.tx, tx.status, nil
	}
	txBytes, err := st.txDB.Get(txID[:])
	if err == database.ErrNotFound {
		st.txCache.Put(txID, nil)
		return nil, status.Unknown, database.ErrNotFound
	} else if err != nil {
		return nil, status.Unknown, err
	}

	stx := stateTx{}
	if _, err := GenesisCodec.Unmarshal(txBytes, &stx); err != nil {
		return nil, status.Unknown, err
	}

	tx := Tx{}
	if _, err := GenesisCodec.Unmarshal(stx.Tx, &tx); err != nil {
		return nil, status.Unknown, err
	}
	if err := tx.Sign(GenesisCodec, nil); err != nil {
		return nil, status.Unknown, err
	}

	ptx := &txStatusImpl{
		tx:     &tx,
		status: stx.Status,
	}

	st.txCache.Put(txID, ptx)
	return ptx.tx, ptx.status, nil
}

func (st *internalStateImpl) AddTx(tx *Tx, status status.Status) {
	st.addedTxs[tx.ID()] = &txStatusImpl{
		tx:     tx,
		status: status,
	}
}

func (st *internalStateImpl) GetRewardUTXOs(txID ids.ID) ([]*axc.UTXO, error) {
	if utxos, exists := st.addedRewardUTXOs[txID]; exists {
		return utxos, nil
	}
	if utxos, exists := st.rewardUTXOsCache.Get(txID); exists {
		return utxos.([]*axc.UTXO), nil
	}

	rawTxDB := prefixdb.New(txID[:], st.rewardUTXODB)
	txDB := linkeddb.NewDefault(rawTxDB)
	it := txDB.NewIterator()
	defer it.Release()

	utxos := []*axc.UTXO(nil)
	for it.Next() {
		utxo := &axc.UTXO{}
		if _, err := Codec.Unmarshal(it.Value(), utxo); err != nil {
			return nil, err
		}
		utxos = append(utxos, utxo)
	}
	if err := it.Error(); err != nil {
		return nil, err
	}

	st.rewardUTXOsCache.Put(txID, utxos)
	return utxos, nil
}

func (st *internalStateImpl) AddRewardUTXO(txID ids.ID, utxo *axc.UTXO) {
	st.addedRewardUTXOs[txID] = append(st.addedRewardUTXOs[txID], utxo)
}

func (st *internalStateImpl) GetUTXO(utxoID ids.ID) (*axc.UTXO, error) {
	if utxo, exists := st.modifiedUTXOs[utxoID]; exists {
		if utxo == nil {
			return nil, database.ErrNotFound
		}
		return utxo, nil
	}
	return st.utxoState.GetUTXO(utxoID)
}

func (st *internalStateImpl) AddUTXO(utxo *axc.UTXO) {
	st.modifiedUTXOs[utxo.InputID()] = utxo
}

func (st *internalStateImpl) DeleteUTXO(utxoID ids.ID) {
	st.modifiedUTXOs[utxoID] = nil
}

func (st *internalStateImpl) GetBlock(blockID ids.ID) (Block, error) {
	if blk, exists := st.addedBlocks[blockID]; exists {
		return blk, nil
	}
	if blkIntf, cached := st.blockCache.Get(blockID); cached {
		if blkIntf == nil {
			return nil, database.ErrNotFound
		}
		return blkIntf.(Block), nil
	}

	blkBytes, err := st.blockDB.Get(blockID[:])
	if err == database.ErrNotFound {
		st.blockCache.Put(blockID, nil)
		return nil, database.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	blkStatus := stateBlk{}
	if _, err := GenesisCodec.Unmarshal(blkBytes, &blkStatus); err != nil {
		return nil, err
	}

	var blk Block
	if _, err := GenesisCodec.Unmarshal(blkStatus.Blk, &blk); err != nil {
		return nil, err
	}
	if err := blk.initialize(st.vm, blkStatus.Blk, blkStatus.Status, blk); err != nil {
		return nil, err
	}

	st.blockCache.Put(blockID, blk)
	return blk, nil
}

func (st *internalStateImpl) AddBlock(block Block) {
	st.addedBlocks[block.ID()] = block
}

func (st *internalStateImpl) UTXOIDs(addr []byte, start ids.ID, limit int) ([]ids.ID, error) {
	return st.utxoState.UTXOIDs(addr, start, limit)
}

func (st *internalStateImpl) CurrentStakerChainState() currentStakerChainState {
	return st.currentStakerChainState
}

func (st *internalStateImpl) PendingStakerChainState() pendingStakerChainState {
	return st.pendingStakerChainState
}

func (st *internalStateImpl) SetCurrentStakerChainState(cs currentStakerChainState) {
	st.currentStakerChainState = cs
}

func (st *internalStateImpl) SetPendingStakerChainState(ps pendingStakerChainState) {
	st.pendingStakerChainState = ps
}

func (st *internalStateImpl) GetUptime(nodeID ids.ShortID) (upDuration time.Duration, lastUpdated time.Time, err error) {
	uptime, exists := st.uptimes[nodeID]
	if !exists {
		return 0, time.Time{}, database.ErrNotFound
	}
	return uptime.UpDuration, uptime.lastUpdated, nil
}

func (st *internalStateImpl) SetUptime(nodeID ids.ShortID, upDuration time.Duration, lastUpdated time.Time) error {
	uptime, exists := st.uptimes[nodeID]
	if !exists {
		return database.ErrNotFound
	}
	uptime.UpDuration = upDuration
	uptime.lastUpdated = lastUpdated
	st.updatedUptimes[nodeID] = struct{}{}
	return nil
}

func (st *internalStateImpl) GetStartTime(nodeID ids.ShortID) (time.Time, error) {
	currentValidator, err := st.currentStakerChainState.GetValidator(nodeID)
	if err != nil {
		return time.Time{}, err
	}
	return currentValidator.AddValidatorTx().StartTime(), nil
}

func (st *internalStateImpl) SetHeight(height uint64) {
	st.currentHeight = height
}

func (st *internalStateImpl) AddCurrentStaker(tx *Tx, potentialReward uint64) {
	st.addedCurrentStakers = append(st.addedCurrentStakers, &validatorReward{
		addStakerTx:     tx,
		potentialReward: potentialReward,
	})
}

func (st *internalStateImpl) DeleteCurrentStaker(tx *Tx) {
	st.deletedCurrentStakers = append(st.deletedCurrentStakers, tx)
}

func (st *internalStateImpl) AddPendingStaker(tx *Tx) {
	st.addedPendingStakers = append(st.addedPendingStakers, tx)
}

func (st *internalStateImpl) DeletePendingStaker(tx *Tx) {
	st.deletedPendingStakers = append(st.deletedPendingStakers, tx)
}

func (st *internalStateImpl) GetValidatorWeightDiffs(height uint64, allychainID ids.ID) (map[ids.ShortID]*ValidatorWeightDiff, error) {
	prefixStruct := heightWithAllychain{
		Height:      height,
		AllychainID: allychainID,
	}
	prefixBytes, err := GenesisCodec.Marshal(CodecVersion, prefixStruct)
	if err != nil {
		return nil, err
	}
	prefixStr := string(prefixBytes)

	if weightDiffsIntf, ok := st.validatorDiffsCache.Get(prefixStr); ok {
		return weightDiffsIntf.(map[ids.ShortID]*ValidatorWeightDiff), nil
	}

	rawDiffDB := prefixdb.New(prefixBytes, st.validatorDiffsDB)
	diffDB := linkeddb.NewDefault(rawDiffDB)
	diffIter := diffDB.NewIterator()

	weightDiffs := make(map[ids.ShortID]*ValidatorWeightDiff)
	for diffIter.Next() {
		nodeID, err := ids.ToShortID(diffIter.Key())
		if err != nil {
			return nil, err
		}

		weightDiff := ValidatorWeightDiff{}
		_, err = GenesisCodec.Unmarshal(diffIter.Value(), &weightDiff)
		if err != nil {
			return nil, err
		}

		weightDiffs[nodeID] = &weightDiff
	}

	st.validatorDiffsCache.Put(prefixStr, weightDiffs)
	return weightDiffs, nil
}

func (st *internalStateImpl) Abort() {
	st.baseDB.Abort()
}

func (st *internalStateImpl) Commit() error {
	defer st.Abort()
	batch, err := st.CommitBatch()
	if err != nil {
		return err
	}
	return batch.Write()
}

func (st *internalStateImpl) CommitBatch() (database.Batch, error) {
	if err := st.writeCurrentStakers(); err != nil {
		return nil, fmt.Errorf("failed to write current stakers with: %w", err)
	}
	if err := st.writePendingStakers(); err != nil {
		return nil, fmt.Errorf("failed to write pending stakers with: %w", err)
	}
	if err := st.writeUptimes(); err != nil {
		return nil, fmt.Errorf("failed to write uptimes with: %w", err)
	}
	if err := st.writeBlocks(); err != nil {
		return nil, fmt.Errorf("failed to write blocks with: %w", err)
	}
	if err := st.writeTXs(); err != nil {
		return nil, fmt.Errorf("failed to write txs with: %w", err)
	}
	if err := st.writeRewardUTXOs(); err != nil {
		return nil, fmt.Errorf("failed to write reward UTXOs with: %w", err)
	}
	if err := st.writeUTXOs(); err != nil {
		return nil, fmt.Errorf("failed to write UTXOs with: %w", err)
	}
	if err := st.writeAllychains(); err != nil {
		return nil, fmt.Errorf("failed to write current allychains with: %w", err)
	}
	if err := st.writeChains(); err != nil {
		return nil, fmt.Errorf("failed to write chains with: %w", err)
	}
	if err := st.writeSingletons(); err != nil {
		return nil, fmt.Errorf("failed to write singletons with: %w", err)
	}
	return st.baseDB.CommitBatch()
}

func (st *internalStateImpl) Close() error {
	errs := wrappers.Errs{}
	errs.Add(
		st.pendingAllychainValidatorBaseDB.Close(),
		st.pendingNominatorBaseDB.Close(),
		st.pendingValidatorBaseDB.Close(),
		st.pendingValidatorsDB.Close(),
		st.currentAllychainValidatorBaseDB.Close(),
		st.currentNominatorBaseDB.Close(),
		st.currentValidatorBaseDB.Close(),
		st.currentValidatorsDB.Close(),
		st.validatorsDB.Close(),
		st.blockDB.Close(),
		st.txDB.Close(),
		st.rewardUTXODB.Close(),
		st.utxoDB.Close(),
		st.allychainBaseDB.Close(),
		st.chainDB.Close(),
		st.singletonDB.Close(),
		st.baseDB.Close(),
	)
	return errs.Err
}

type currentValidatorState struct {
	txID        ids.ID
	lastUpdated time.Time

	UpDuration      time.Duration `serialize:"true"`
	LastUpdated     uint64        `serialize:"true"` // Unix time in seconds
	PotentialReward uint64        `serialize:"true"`
}

func (st *internalStateImpl) writeCurrentStakers() error {
	weightDiffs := make(map[ids.ID]map[ids.ShortID]*ValidatorWeightDiff) // allychainID -> nodeID -> weightDiff
	for _, currentStaker := range st.addedCurrentStakers {
		txID := currentStaker.addStakerTx.ID()
		potentialReward := currentStaker.potentialReward

		var (
			allychainID ids.ID
			nodeID      ids.ShortID
			weight      uint64
		)
		switch tx := currentStaker.addStakerTx.UnsignedTx.(type) {
		case *UnsignedAddValidatorTx:
			startTime := tx.StartTime()
			vdr := &currentValidatorState{
				txID:        txID,
				lastUpdated: startTime,

				UpDuration:      0,
				LastUpdated:     uint64(startTime.Unix()),
				PotentialReward: potentialReward,
			}

			vdrBytes, err := GenesisCodec.Marshal(CodecVersion, vdr)
			if err != nil {
				return err
			}

			if err := st.currentValidatorList.Put(txID[:], vdrBytes); err != nil {
				return err
			}
			st.uptimes[tx.Validator.NodeID] = vdr

			allychainID = constants.PrimaryNetworkID
			nodeID = tx.Validator.NodeID
			weight = tx.Validator.Wght
		case *UnsignedAddNominatorTx:
			if err := database.PutUInt64(st.currentNominatorList, txID[:], potentialReward); err != nil {
				return err
			}

			allychainID = constants.PrimaryNetworkID
			nodeID = tx.Validator.NodeID
			weight = tx.Validator.Wght
		case *UnsignedAddAllychainValidatorTx:
			if err := st.currentAllychainValidatorList.Put(txID[:], nil); err != nil {
				return err
			}

			allychainID = tx.Validator.Allychain
			nodeID = tx.Validator.NodeID
			weight = tx.Validator.Wght
		default:
			return errWrongTxType
		}

		allychainDiffs, ok := weightDiffs[allychainID]
		if !ok {
			allychainDiffs = make(map[ids.ShortID]*ValidatorWeightDiff)
			weightDiffs[allychainID] = allychainDiffs
		}

		nodeDiff, ok := allychainDiffs[nodeID]
		if !ok {
			nodeDiff = &ValidatorWeightDiff{}
			allychainDiffs[nodeID] = nodeDiff
		}

		newWeight, err := safemath.Add64(nodeDiff.Amount, weight)
		if err != nil {
			return err
		}
		nodeDiff.Amount = newWeight
	}
	st.addedCurrentStakers = nil

	for _, tx := range st.deletedCurrentStakers {
		var (
			db          database.KeyValueDeleter
			allychainID ids.ID
			nodeID      ids.ShortID
			weight      uint64
		)
		switch tx := tx.UnsignedTx.(type) {
		case *UnsignedAddValidatorTx:
			db = st.currentValidatorList

			delete(st.uptimes, tx.Validator.NodeID)
			delete(st.updatedUptimes, tx.Validator.NodeID)

			allychainID = constants.PrimaryNetworkID
			nodeID = tx.Validator.NodeID
			weight = tx.Validator.Wght
		case *UnsignedAddNominatorTx:
			db = st.currentNominatorList

			allychainID = constants.PrimaryNetworkID
			nodeID = tx.Validator.NodeID
			weight = tx.Validator.Wght
		case *UnsignedAddAllychainValidatorTx:
			db = st.currentAllychainValidatorList

			allychainID = tx.Validator.Allychain
			nodeID = tx.Validator.NodeID
			weight = tx.Validator.Wght
		default:
			return errWrongTxType
		}

		txID := tx.ID()
		if err := db.Delete(txID[:]); err != nil {
			return err
		}

		allychainDiffs, ok := weightDiffs[allychainID]
		if !ok {
			allychainDiffs = make(map[ids.ShortID]*ValidatorWeightDiff)
			weightDiffs[allychainID] = allychainDiffs
		}

		nodeDiff, ok := allychainDiffs[nodeID]
		if !ok {
			nodeDiff = &ValidatorWeightDiff{}
			allychainDiffs[nodeID] = nodeDiff
		}

		if nodeDiff.Decrease {
			newWeight, err := safemath.Add64(nodeDiff.Amount, weight)
			if err != nil {
				return err
			}
			nodeDiff.Amount = newWeight
		} else {
			nodeDiff.Decrease = nodeDiff.Amount < weight
			nodeDiff.Amount = safemath.Diff64(nodeDiff.Amount, weight)
		}
	}
	st.deletedCurrentStakers = nil

	for allychainID, nodeUpdates := range weightDiffs {
		prefixStruct := heightWithAllychain{
			Height:      st.currentHeight,
			AllychainID: allychainID,
		}
		prefixBytes, err := GenesisCodec.Marshal(CodecVersion, prefixStruct)
		if err != nil {
			return err
		}
		rawDiffDB := prefixdb.New(prefixBytes, st.validatorDiffsDB)
		diffDB := linkeddb.NewDefault(rawDiffDB)
		for nodeID, nodeDiff := range nodeUpdates {
			if nodeDiff.Amount == 0 {
				delete(nodeUpdates, nodeID)
				continue
			}

			if allychainID == constants.PrimaryNetworkID || st.vm.WhitelistedAllychains.Contains(allychainID) {
				var err error
				if nodeDiff.Decrease {
					err = st.vm.Validators.RemoveWeight(allychainID, nodeID, nodeDiff.Amount)
				} else {
					err = st.vm.Validators.AddWeight(allychainID, nodeID, nodeDiff.Amount)
				}
				if err != nil {
					return err
				}
			}

			nodeDiffBytes, err := GenesisCodec.Marshal(CodecVersion, nodeDiff)
			if err != nil {
				return err
			}

			// Copy so value passed into [Put] doesn't get overwritten next iteration
			nodeID := nodeID
			if err := diffDB.Put(nodeID[:], nodeDiffBytes); err != nil {
				return err
			}
		}
		st.validatorDiffsCache.Put(string(prefixBytes), nodeUpdates)
	}

	// Attempt to update the stake metrics
	primaryValidators, ok := st.vm.Validators.GetValidators(constants.PrimaryNetworkID)
	if !ok {
		return nil
	}
	weight, _ := primaryValidators.GetWeight(st.vm.ctx.NodeID)
	st.vm.localStake.Set(float64(weight))
	st.vm.totalStake.Set(float64(primaryValidators.Weight()))
	return nil
}

func (st *internalStateImpl) writePendingStakers() error {
	for _, tx := range st.addedPendingStakers {
		var db database.KeyValueWriter
		switch tx.UnsignedTx.(type) {
		case *UnsignedAddValidatorTx:
			db = st.pendingValidatorList
		case *UnsignedAddNominatorTx:
			db = st.pendingNominatorList
		case *UnsignedAddAllychainValidatorTx:
			db = st.pendingAllychainValidatorList
		default:
			return errWrongTxType
		}

		txID := tx.ID()
		if err := db.Put(txID[:], nil); err != nil {
			return err
		}
	}
	st.addedPendingStakers = nil

	for _, tx := range st.deletedPendingStakers {
		var db database.KeyValueDeleter
		switch tx.UnsignedTx.(type) {
		case *UnsignedAddValidatorTx:
			db = st.pendingValidatorList
		case *UnsignedAddNominatorTx:
			db = st.pendingNominatorList
		case *UnsignedAddAllychainValidatorTx:
			db = st.pendingAllychainValidatorList
		default:
			return errWrongTxType
		}

		txID := tx.ID()
		if err := db.Delete(txID[:]); err != nil {
			return err
		}
	}
	st.deletedPendingStakers = nil
	return nil
}

func (st *internalStateImpl) writeUptimes() error {
	for nodeID := range st.updatedUptimes {
		delete(st.updatedUptimes, nodeID)

		uptime := st.uptimes[nodeID]
		uptime.LastUpdated = uint64(uptime.lastUpdated.Unix())

		uptimeBytes, err := GenesisCodec.Marshal(CodecVersion, uptime)
		if err != nil {
			return err
		}

		if err := st.currentValidatorList.Put(uptime.txID[:], uptimeBytes); err != nil {
			return err
		}
	}
	return nil
}

func (st *internalStateImpl) writeBlocks() error {
	for blkID, blk := range st.addedBlocks {
		blkID := blkID

		sblk := stateBlk{
			Blk:    blk.Bytes(),
			Status: blk.Status(),
		}
		btxBytes, err := GenesisCodec.Marshal(CodecVersion, &sblk)
		if err != nil {
			return err
		}

		delete(st.addedBlocks, blkID)
		st.blockCache.Put(blkID, blk)
		if err := st.blockDB.Put(blkID[:], btxBytes); err != nil {
			return err
		}
	}
	return nil
}

func (st *internalStateImpl) writeTXs() error {
	for txID, txStatus := range st.addedTxs {
		txID := txID

		stx := stateTx{
			Tx:     txStatus.tx.Bytes(),
			Status: txStatus.status,
		}
		txBytes, err := GenesisCodec.Marshal(CodecVersion, &stx)
		if err != nil {
			return err
		}

		delete(st.addedTxs, txID)
		st.txCache.Put(txID, txStatus)
		if err := st.txDB.Put(txID[:], txBytes); err != nil {
			return err
		}
	}
	return nil
}

func (st *internalStateImpl) writeRewardUTXOs() error {
	for txID, utxos := range st.addedRewardUTXOs {
		delete(st.addedRewardUTXOs, txID)

		st.rewardUTXOsCache.Put(txID, utxos)

		rawTxDB := prefixdb.New(txID[:], st.rewardUTXODB)
		txDB := linkeddb.NewDefault(rawTxDB)

		for _, utxo := range utxos {
			utxoBytes, err := GenesisCodec.Marshal(CodecVersion, utxo)
			if err != nil {
				return err
			}
			utxoID := utxo.InputID()
			if err := txDB.Put(utxoID[:], utxoBytes); err != nil {
				return err
			}
		}
	}
	return nil
}

func (st *internalStateImpl) writeUTXOs() error {
	for utxoID, utxo := range st.modifiedUTXOs {
		delete(st.modifiedUTXOs, utxoID)

		if utxo == nil {
			if err := st.utxoState.DeleteUTXO(utxoID); err != nil {
				return err
			}
			continue
		}
		if err := st.utxoState.PutUTXO(utxoID, utxo); err != nil {
			return err
		}
	}
	return nil
}

func (st *internalStateImpl) writeAllychains() error {
	for _, allychain := range st.addedAllychains {
		allychainID := allychain.ID()

		if err := st.allychainDB.Put(allychainID[:], nil); err != nil {
			return err
		}
	}
	st.addedAllychains = nil
	return nil
}

func (st *internalStateImpl) writeChains() error {
	for allychainID, chains := range st.addedChains {
		for _, chain := range chains {
			chainDB := st.getChainDB(allychainID)

			chainID := chain.ID()
			if err := chainDB.Put(chainID[:], nil); err != nil {
				return err
			}
		}
		delete(st.addedChains, allychainID)
	}
	return nil
}

func (st *internalStateImpl) writeSingletons() error {
	if !st.originalTimestamp.Equal(st.timestamp) {
		if err := database.PutTimestamp(st.singletonDB, timestampKey, st.timestamp); err != nil {
			return err
		}
		st.originalTimestamp = st.timestamp
	}
	if st.originalCurrentSupply != st.currentSupply {
		if err := database.PutUInt128(st.singletonDB, currentSupplyKey, st.currentSupply); err != nil {
			return err
		}
		st.originalCurrentSupply = st.currentSupply
	}
	if st.originalLastAccepted != st.lastAccepted {
		if err := database.PutID(st.singletonDB, lastAcceptedKey, st.lastAccepted); err != nil {
			return err
		}
		st.originalLastAccepted = st.lastAccepted
	}
	return nil
}

func (st *internalStateImpl) load() error {
	if err := st.loadSingletons(); err != nil {
		return err
	}
	if err := st.loadCurrentValidators(); err != nil {
		return err
	}
	return st.loadPendingValidators()
}

func (st *internalStateImpl) loadSingletons() error {
	timestamp, err := database.GetTimestamp(st.singletonDB, timestampKey)
	if err != nil {
		return err
	}
	st.originalTimestamp = timestamp
	st.timestamp = timestamp

	currentSupply, err := database.GetUInt128(st.singletonDB, currentSupplyKey)
	if err != nil {
		return err
	}
	st.originalCurrentSupply = currentSupply
	st.currentSupply = currentSupply

	lastAccepted, err := database.GetID(st.singletonDB, lastAcceptedKey)
	if err != nil {
		return err
	}
	st.originalLastAccepted = lastAccepted
	st.lastAccepted = lastAccepted

	return nil
}

func (st *internalStateImpl) loadCurrentValidators() error {
	cs := &currentStakerChainStateImpl{
		validatorsByNodeID: make(map[ids.ShortID]*currentValidatorImpl),
		validatorsByTxID:   make(map[ids.ID]*validatorReward),
	}

	validatorIt := st.currentValidatorList.NewIterator()
	defer validatorIt.Release()
	for validatorIt.Next() {
		txIDBytes := validatorIt.Key()
		txID, err := ids.ToID(txIDBytes)
		if err != nil {
			return err
		}
		tx, _, err := st.GetTx(txID)
		if err != nil {
			return err
		}

		uptimeBytes := validatorIt.Value()
		uptime := &currentValidatorState{
			txID: txID,
		}
		if _, err := GenesisCodec.Unmarshal(uptimeBytes, uptime); err != nil {
			return err
		}
		uptime.lastUpdated = time.Unix(int64(uptime.LastUpdated), 0)

		addValidatorTx, ok := tx.UnsignedTx.(*UnsignedAddValidatorTx)
		if !ok {
			return errWrongTxType
		}

		cs.validators = append(cs.validators, tx)
		cs.validatorsByNodeID[addValidatorTx.Validator.NodeID] = &currentValidatorImpl{
			validatorImpl: validatorImpl{
				allychains: make(map[ids.ID]*UnsignedAddAllychainValidatorTx),
			},
			addValidatorTx:  addValidatorTx,
			potentialReward: uptime.PotentialReward,
		}
		cs.validatorsByTxID[txID] = &validatorReward{
			addStakerTx:     tx,
			potentialReward: uptime.PotentialReward,
		}

		st.uptimes[addValidatorTx.Validator.NodeID] = uptime
	}

	if err := validatorIt.Error(); err != nil {
		return err
	}

	nominatorIt := st.currentNominatorList.NewIterator()
	defer nominatorIt.Release()
	for nominatorIt.Next() {
		txIDBytes := nominatorIt.Key()
		txID, err := ids.ToID(txIDBytes)
		if err != nil {
			return err
		}
		tx, _, err := st.GetTx(txID)
		if err != nil {
			return err
		}

		potentialRewardBytes := nominatorIt.Value()
		potentialReward, err := database.ParseUInt64(potentialRewardBytes)
		if err != nil {
			return err
		}

		addNominatorTx, ok := tx.UnsignedTx.(*UnsignedAddNominatorTx)
		if !ok {
			return errWrongTxType
		}

		cs.validators = append(cs.validators, tx)
		vdr, exists := cs.validatorsByNodeID[addNominatorTx.Validator.NodeID]
		if !exists {
			return errNominatorSubset
		}
		vdr.nominatorWeight += addNominatorTx.Validator.Wght
		vdr.nominators = append(vdr.nominators, addNominatorTx)
		cs.validatorsByTxID[txID] = &validatorReward{
			addStakerTx:     tx,
			potentialReward: potentialReward,
		}
	}
	if err := nominatorIt.Error(); err != nil {
		return err
	}

	allychainValidatorIt := st.currentAllychainValidatorList.NewIterator()
	defer allychainValidatorIt.Release()
	for allychainValidatorIt.Next() {
		txIDBytes := allychainValidatorIt.Key()
		txID, err := ids.ToID(txIDBytes)
		if err != nil {
			return err
		}
		tx, _, err := st.GetTx(txID)
		if err != nil {
			return err
		}

		addAllychainValidatorTx, ok := tx.UnsignedTx.(*UnsignedAddAllychainValidatorTx)
		if !ok {
			return errWrongTxType
		}

		cs.validators = append(cs.validators, tx)
		vdr, exists := cs.validatorsByNodeID[addAllychainValidatorTx.Validator.NodeID]
		if !exists {
			return errDSValidatorSubset
		}
		vdr.allychains[addAllychainValidatorTx.Validator.Allychain] = addAllychainValidatorTx
		cs.validatorsByTxID[txID] = &validatorReward{
			addStakerTx: tx,
		}
	}
	if err := allychainValidatorIt.Error(); err != nil {
		return err
	}

	for _, vdr := range cs.validatorsByNodeID {
		sortNominatorsByRemoval(vdr.nominators)
	}
	sortValidatorsByRemoval(cs.validators)
	cs.setNextStaker()

	st.currentStakerChainState = cs
	return nil
}

func (st *internalStateImpl) loadPendingValidators() error {
	ps := &pendingStakerChainStateImpl{
		validatorsByNodeID:      make(map[ids.ShortID]*UnsignedAddValidatorTx),
		validatorExtrasByNodeID: make(map[ids.ShortID]*validatorImpl),
	}

	validatorIt := st.pendingValidatorList.NewIterator()
	defer validatorIt.Release()
	for validatorIt.Next() {
		txIDBytes := validatorIt.Key()
		txID, err := ids.ToID(txIDBytes)
		if err != nil {
			return err
		}
		tx, _, err := st.GetTx(txID)
		if err != nil {
			return err
		}

		addValidatorTx, ok := tx.UnsignedTx.(*UnsignedAddValidatorTx)
		if !ok {
			return errWrongTxType
		}

		ps.validators = append(ps.validators, tx)
		ps.validatorsByNodeID[addValidatorTx.Validator.NodeID] = addValidatorTx
	}
	if err := validatorIt.Error(); err != nil {
		return err
	}

	nominatorIt := st.pendingNominatorList.NewIterator()
	defer nominatorIt.Release()
	for nominatorIt.Next() {
		txIDBytes := nominatorIt.Key()
		txID, err := ids.ToID(txIDBytes)
		if err != nil {
			return err
		}
		tx, _, err := st.GetTx(txID)
		if err != nil {
			return err
		}

		addNominatorTx, ok := tx.UnsignedTx.(*UnsignedAddNominatorTx)
		if !ok {
			return errWrongTxType
		}

		ps.validators = append(ps.validators, tx)
		if vdr, exists := ps.validatorExtrasByNodeID[addNominatorTx.Validator.NodeID]; exists {
			vdr.nominators = append(vdr.nominators, addNominatorTx)
		} else {
			ps.validatorExtrasByNodeID[addNominatorTx.Validator.NodeID] = &validatorImpl{
				nominators: []*UnsignedAddNominatorTx{addNominatorTx},
				allychains: make(map[ids.ID]*UnsignedAddAllychainValidatorTx),
			}
		}
	}
	if err := nominatorIt.Error(); err != nil {
		return err
	}

	allychainValidatorIt := st.pendingAllychainValidatorList.NewIterator()
	defer allychainValidatorIt.Release()
	for allychainValidatorIt.Next() {
		txIDBytes := allychainValidatorIt.Key()
		txID, err := ids.ToID(txIDBytes)
		if err != nil {
			return err
		}
		tx, _, err := st.GetTx(txID)
		if err != nil {
			return err
		}

		addAllychainValidatorTx, ok := tx.UnsignedTx.(*UnsignedAddAllychainValidatorTx)
		if !ok {
			return errWrongTxType
		}

		ps.validators = append(ps.validators, tx)
		if vdr, exists := ps.validatorExtrasByNodeID[addAllychainValidatorTx.Validator.NodeID]; exists {
			vdr.allychains[addAllychainValidatorTx.Validator.Allychain] = addAllychainValidatorTx
		} else {
			ps.validatorExtrasByNodeID[addAllychainValidatorTx.Validator.NodeID] = &validatorImpl{
				allychains: map[ids.ID]*UnsignedAddAllychainValidatorTx{
					addAllychainValidatorTx.Validator.Allychain: addAllychainValidatorTx,
				},
			}
		}
	}
	if err := allychainValidatorIt.Error(); err != nil {
		return err
	}

	for _, vdr := range ps.validatorExtrasByNodeID {
		sortNominatorsByAddition(vdr.nominators)
	}
	sortValidatorsByAddition(ps.validators)

	st.pendingStakerChainState = ps
	return nil
}

func (st *internalStateImpl) shouldInit() (bool, error) {
	has, err := st.singletonDB.Has(initializedKey)
	return !has, err
}

func (st *internalStateImpl) init(genesisBytes []byte) error {
	genesis := &Genesis{}
	if _, err := GenesisCodec.Unmarshal(genesisBytes, genesis); err != nil {
		return err
	}
	if err := genesis.Initialize(); err != nil {
		return err
	}

	// Persist UTXOs that exist at genesis
	for _, utxo := range genesis.UTXOs {
		st.AddUTXO(&utxo.UTXO)
	}

	// Persist the platform chain's timestamp at genesis
	genesisTime := time.Unix(int64(genesis.Timestamp), 0)
	st.SetTimestamp(genesisTime)
	st.SetCurrentSupply(genesis.InitialSupply)

	// Persist primary network validator set at genesis
	for _, vdrTx := range genesis.Validators {
		tx, ok := vdrTx.UnsignedTx.(*UnsignedAddValidatorTx)
		if !ok {
			return errWrongTxType
		}

		stakeAmount := tx.Validator.Wght
		stakeDuration := tx.Validator.Duration()
		currentSupply := st.GetCurrentSupply()

		r := st.vm.rewards.Calculate(
			stakeDuration,
			stakeAmount,
			currentSupply,
		)
		newCurrentSupply := currentSupply.Add64(r)

		st.AddCurrentStaker(vdrTx, r)
		st.AddTx(vdrTx, status.Committed)
		st.SetCurrentSupply(newCurrentSupply)
	}

	for _, chain := range genesis.Chains {
		unsignedChain, ok := chain.UnsignedTx.(*UnsignedCreateChainTx)
		if !ok {
			return errWrongTxType
		}

		// Ensure all chains that the genesis bytes say to create have the right
		// network ID
		if unsignedChain.NetworkID != st.vm.ctx.NetworkID {
			return errWrongNetworkID
		}

		st.AddChain(chain)
		st.AddTx(chain, status.Committed)
	}

	// Create the genesis block and save it as being accepted (We don't just
	// do genesisBlock.Accept() because then it'd look for genesisBlock's
	// non-existent parent)
	genesisID := hashing.ComputeHash256Array(genesisBytes)
	genesisBlock, err := st.vm.newCommitBlock(genesisID, 0, true)
	if err != nil {
		return err
	}
	genesisBlock.status = choices.Accepted
	st.AddBlock(genesisBlock)
	st.SetLastAccepted(genesisBlock.ID())

	if err := st.singletonDB.Put(initializedKey, nil); err != nil {
		return err
	}

	return st.Commit()
}
