<div align="center">
  <img src="https://axia.global/images/logotype/axia-logo.svg?raw=true">
</div>

---

Node implementation for the [Axia](https://axc.network) network -
a blockchains platform with high throughput, and blazing fast transactions.

## Installation

Axia is an incredibly lightweight protocol, so the minimum computer requirements are quite modest.
Note that as network usage increases, hardware requirements may change.

The minimum recommended hardware specification for nodes connected to Mainnet is:

- CPU: Equivalent of 8 AWS vCPU
- RAM: 16 GiB
- Storage: 512 GiB
- OS: Ubuntu 18.04/20.04 or macOS >= 10.15 (Catalina)
- Network: Reliable IPv4 or IPv6 network connection, with an open public port.

If you plan to build Axia from source, you will also need the following software:

- [Go](https://golang.org/doc/install) version >= 1.17.9
- [gcc](https://gcc.gnu.org/)
- g++

### Native Install

Clone the Axia repository:

```sh
git clone git@github.com:axiacoin/axia.git
cd axia
```

This will clone and checkout to `master` branch.

#### Building the Axia Executable

Build Axia using the build script:

```sh
./scripts/build.sh
```

The Axia binary, named `axia`, is in the `build` directory.

### Binary Repository

Install Axia using an `apt` repository.

#### Adding the APT Repository

If you have already added the APT repository, you do not need to add it again.

To add the repository on Ubuntu 18.04 (Bionic), run:

```sh
sudo su -
wget -O - https://downloads.axc.network/axia.gpg.key | apt-key add -
echo "deb https://downloads.axc.network/apt bionic main" > /etc/apt/sources.list.d/axia.list
exit
```

To add the repository on Ubuntu 20.04 (Focal), run:

```sh
sudo su -
wget -O - https://downloads.axc.network/axia.gpg.key | apt-key add -
echo "deb https://downloads.axc.network/apt focal main" > /etc/apt/sources.list.d/axia.list
exit
```

#### Installing the Latest Version

After adding the APT repository, install axia by running:

```sh
sudo apt update
sudo apt install axia
```

### Binary Install

Download the [latest build](https://github.com/axiacoin/axia-network-v2/releases/latest) for your operating system and architecture.

The Axia binary to be executed is named `axia`.

### Docker Install

Make sure docker is installed on the machine - so commands like `docker run` etc. are available.

Building the docker image of latest axia branch can be done by running:

```sh
./scripts/build_image.sh
```

To check the built image, run:

```sh
docker image ls
```

The image should be tagged as `axiasystem/axia:xxxxxxxx`, where `xxxxxxxx` is the shortened commit of the Axia source it was built from. To run the axia node, run:

```sh
docker run -ti -p 9650:9650 -p 9651:9651 axiasystem/axia:xxxxxxxx /axia/build/axgo
```

## Running Axia

### Connecting to Mainnet

To connect to the Axia Mainnet, run:

```sh
./build/axgo
```

You should see some pretty ASCII art and log messages.

You can use `Ctrl+C` to kill the node.

### Connecting to Test

To connect to the Test Testnet, run:

```sh
./build/axgo --network-id=test
```

### Creating a Local Testnet

See [this tutorial.](https://docs.axc.network/build/tutorials/platform/create-a-local-test-network/)

## Bootstrapping

A node needs to catch up to the latest network state before it can participate in consensus and serve API calls. This process, called bootstrapping, currently takes several days for a new node connected to Mainnet.

A node will not [report healthy](https://docs.axc.network/build/axgo-apis/health) until it is done bootstrapping.

Improvements that reduce the amount of time it takes to bootstrap are under development.

The bottleneck during bootstrapping is typically database IO. Using a more powerful CPU or increasing the database IOPS on the computer running a node will decrease the amount of time bootstrapping takes.

## Generating Code

Axia uses multiple tools to generate efficient and boilerplate code.

### Running protobuf codegen

To regenerate the protobuf go code, run `scripts/protobuf_codegen.sh` from the root of the repo.

This should only be necessary when upgrading protobuf versions or modifying .proto definition files.

To use this script, you must have [buf](https://docs.buf.build/installation) (v1.0.0-rc12), protoc-gen-go (v1.27.1) and protoc-gen-go-grpc (v1.2.0) installed.

To install the buf dependencies:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
```

If you have not already, you may need to add `$GOPATH/bin` to your `$PATH`:

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

If you extract buf to ~/software/buf/bin, the following should work:

```sh
export PATH=$PATH:~/software/buf/bin/:~/go/bin
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/protobuf/cmd/protoc-gen-go-grpc
scripts/protobuf_codegen.sh
```

For more information, refer to the [GRPC Golang Quick Start Guide](https://grpc.io/docs/languages/go/quickstart/).

### Running protobuf codegen from docker

```sh
docker build -t axia:protobuf_codegen -f api/Dockerfile.buf .
docker run -t -i -v $(pwd):/opt/axia -w/opt/axia axia:protobuf_codegen bash -c "scripts/protobuf_codegen.sh"
```

## Supported Platforms

Axia can run on different platforms, with different support tiers:

- **Tier 1**: Fully supported by the maintainers, guaranteed to pass all tests including e2e and stress tests.
- **Tier 2**: Passes all unit and integration tests but not necessarily e2e tests.
- **Tier 3**: Builds but lightly tested (or not), considered _experimental_.
- **Not supported**: May not build and not tested, considered _unsafe_. To be supported in the future.

The following table lists currently supported platforms and their corresponding
Axia support tiers:

| Architecture | Operating system | Support tier  |
| :----------: | :--------------: | :-----------: |
|    amd64     |      Linux       |       1       |
|    arm64     |      Linux       |       2       |
|    amd64     |      Darwin      |       2       |
|    amd64     |     Windows      |       3       |
|     arm      |      Linux       | Not supported |
|     i386     |      Linux       | Not supported |
|    arm64     |      Darwin      | Not supported |

To officially support a new platform, one must satisfy the following requirements:

| Axia continuous integration | Tier 1  | Tier 2  | Tier 3  |
| ---------------------------------- | :-----: | :-----: | :-----: |
| Build passes                       | &check; | &check; | &check; |
| Unit and integration tests pass    | &check; | &check; |         |
| End-to-end and stress tests pass   | &check; |         |         |

## Security Bugs

**We and our community welcome responsible disclosures.**

If you've discovered a security vulnerability, please report it via our [bug bounty program](https://hackenproof.com/axia/). Valid reports will be eligible for a reward (terms and conditions apply).
