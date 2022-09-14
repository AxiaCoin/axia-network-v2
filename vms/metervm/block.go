// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metervm

import "github.com/axiacoin/axia-network-v2/snow/consensus/kleroterion"

var (
	_ kleroterion.Block       = &meterBlock{}
	_ kleroterion.OracleBlock = &meterBlock{}
)

type meterBlock struct {
	kleroterion.Block

	vm *blockVM
}

func (mb *meterBlock) Verify() error {
	start := mb.vm.clock.Time()
	err := mb.Block.Verify()
	end := mb.vm.clock.Time()
	duration := float64(end.Sub(start))
	if err != nil {
		mb.vm.blockMetrics.verifyErr.Observe(duration)
	} else {
		mb.vm.verify.Observe(duration)
	}
	return err
}

func (mb *meterBlock) Accept() error {
	start := mb.vm.clock.Time()
	err := mb.Block.Accept()
	end := mb.vm.clock.Time()
	duration := float64(end.Sub(start))
	mb.vm.blockMetrics.accept.Observe(duration)
	return err
}

func (mb *meterBlock) Reject() error {
	start := mb.vm.clock.Time()
	err := mb.Block.Reject()
	end := mb.vm.clock.Time()
	duration := float64(end.Sub(start))
	mb.vm.blockMetrics.reject.Observe(duration)
	return err
}

func (mb *meterBlock) Options() ([2]kleroterion.Block, error) {
	oracleBlock, ok := mb.Block.(kleroterion.OracleBlock)
	if !ok {
		return [2]kleroterion.Block{}, kleroterion.ErrNotOracle
	}

	blks, err := oracleBlock.Options()
	if err != nil {
		return [2]kleroterion.Block{}, err
	}
	return [2]kleroterion.Block{
		&meterBlock{
			Block: blks[0],
			vm:    mb.vm,
		},
		&meterBlock{
			Block: blks[1],
			vm:    mb.vm,
		},
	}, nil
}
