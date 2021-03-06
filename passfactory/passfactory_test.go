package passfactory_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/deployer"
	"github.com/monetha/go-verifiable-data/eth"
	"github.com/monetha/go-verifiable-data/eth/backend"
	sdklog "github.com/monetha/go-verifiable-data/log"
	"github.com/monetha/go-verifiable-data/passfactory"
)

func TestReader_FilterPassports(t *testing.T) {
	t.Run("nil filtering options returns all passes", func(t *testing.T) {
		passCreation, ps := prepareReaderFilterPassports(t, func(result passportCreationResult) *passfactory.PassportFilterOpts {
			return nil
		})

		if len(ps) != 1 {
			t.Errorf("Expected to get exactly one passport, got %v", len(ps))
		}

		rPass := ps[0]
		if rPass.ContractAddress != passCreation.PassportAddress {
			t.Errorf("expected to get passport address %v, got %v", passCreation.PassportAddress.Hex(), rPass.ContractAddress)
		}
	})

	t.Run("empty filtering options returns all passes", func(t *testing.T) {
		passCreation, ps := prepareReaderFilterPassports(t, func(result passportCreationResult) *passfactory.PassportFilterOpts {
			return &passfactory.PassportFilterOpts{}
		})

		if len(ps) != 1 {
			t.Errorf("Expected to get exactly one passport, got %v", len(ps))
		}

		rPass := ps[0]
		if rPass.ContractAddress != passCreation.PassportAddress {
			t.Errorf("expected to get passport address %v, got %v", passCreation.PassportAddress.Hex(), rPass.ContractAddress)
		}
	})

	t.Run("filtering by passport returns pass", func(t *testing.T) {
		passCreation, ps := prepareReaderFilterPassports(t, func(result passportCreationResult) *passfactory.PassportFilterOpts {
			return &passfactory.PassportFilterOpts{Passport: []common.Address{result.PassportAddress}}
		})

		if len(ps) != 1 {
			t.Errorf("Expected to get exactly one passport, got %v", len(ps))
		}

		rPass := ps[0]
		if rPass.ContractAddress != passCreation.PassportAddress {
			t.Errorf("expected to get passport address %v, got %v", passCreation.PassportAddress.Hex(), rPass.ContractAddress)
		}
	})

	t.Run("filtering by owner returns pass", func(t *testing.T) {
		passCreation, ps := prepareReaderFilterPassports(t, func(result passportCreationResult) *passfactory.PassportFilterOpts {
			return &passfactory.PassportFilterOpts{Owner: []common.Address{result.PassportOwnerAddress}}
		})

		if len(ps) != 1 {
			t.Errorf("Expected to get exactly one passport, got %v", len(ps))
		}

		rPass := ps[0]
		if rPass.ContractAddress != passCreation.PassportAddress {
			t.Errorf("expected to get passport address %v, got %v", passCreation.PassportAddress.Hex(), rPass.ContractAddress)
		}
	})

	t.Run("filtering by unknown passport address returns nothing", func(t *testing.T) {
		_, ps := prepareReaderFilterPassports(t, func(result passportCreationResult) *passfactory.PassportFilterOpts {
			return &passfactory.PassportFilterOpts{Passport: []common.Address{result.PassportOwnerAddress}}
		})

		if len(ps) != 0 {
			t.Errorf("Expected to get nothing, got %v", len(ps))
		}
	})
}

func prepareReaderFilterPassports(t *testing.T, optsFun func(passportCreationResult) *passfactory.PassportFilterOpts) (passportCreationResult, []*passfactory.Passport) {
	ctx := context.Background()
	passCreation, e := createPassport(ctx, t)

	opts := optsFun(passCreation)

	pf := passfactory.NewReader(e)
	if opts != nil {
		opts.Context = ctx
	}
	it, err := pf.FilterPassports(opts, passCreation.PassportFactoryAddress)
	if err != nil {
		t.Errorf("FilterPassports: %v", err)
	}

	ps, err := it.ToSlice()
	if err != nil {
		t.Errorf("ToSlice: %v", err)
	}

	return passCreation, ps
}

type passportCreationResult struct {
	PassportAddress        common.Address
	PassportOwnerAddress   common.Address
	PassportFactoryAddress common.Address
}

func createPassport(ctx context.Context, t *testing.T) (passportCreationResult, *eth.Eth) {
	monethaKey, err := crypto.GenerateKey()
	if err != nil {
		t.Errorf("crypto.GenerateKey() error = %v", err)
	}
	monethaAddress := bind.NewKeyedTransactor(monethaKey).From

	passportOwnerKey, err := crypto.GenerateKey()
	if err != nil {
		t.Errorf("crypto.GenerateKey() error = %v", err)
	}
	passportOwnerAddress := bind.NewKeyedTransactor(passportOwnerKey).From

	factProviderKey, err := crypto.GenerateKey()
	if err != nil {
		t.Errorf("crypto.GenerateKey() error = %v", err)
	}
	factProviderAddress := bind.NewKeyedTransactor(factProviderKey).From
	alloc := core.GenesisAlloc{
		monethaAddress:       {Balance: big.NewInt(deployer.PassportFactoryGasLimit)},
		passportOwnerAddress: {Balance: big.NewInt(deployer.PassportGasLimit)},
		factProviderAddress:  {Balance: big.NewInt(10000000000000)},
	}
	sim := backend.NewSimulatedBackendExtended(alloc, 10000000)
	sim.Commit()

	if sdklog.IsTestLogSet() {
		glogHandler := log.NewGlogHandler(log.StreamHandler(tWriter{}, log.LogfmtFormat()))
		glogHandler.Verbosity(log.LvlWarn)
		log.Root().SetHandler(glogHandler)
	}

	e := eth.New(sim, log.Warn)
	if err := e.UpdateSuggestedGasPrice(ctx); err != nil {
		t.Fatalf("Eth.UpdateSuggestedGasPrice() error = %v", err)
	}

	monethaSession := e.NewSession(monethaKey)
	// deploying passport factory with all dependencies: passport logic, passport logic registry
	res, err := deployer.New(monethaSession).Bootstrap(ctx)
	if err != nil {
		t.Errorf("Deploy.Bootstrap() error = %v", err)
	}
	passportFactoryAddress := res.PassportFactoryAddress

	passportOwnerSession := e.NewSession(passportOwnerKey)
	// deploying passport
	passportAddress, err := deployer.New(passportOwnerSession).DeployPassport(ctx, passportFactoryAddress)
	if err != nil {
		t.Errorf("Deploy.DeployPassport() error = %v", err)
	}

	return passportCreationResult{
		PassportAddress:        passportAddress,
		PassportOwnerAddress:   passportOwnerAddress,
		PassportFactoryAddress: passportFactoryAddress,
	}, e
}

type tWriter struct{}

func (t tWriter) Write(p []byte) (n int, err error) {
	fmt.Print(string(p))
	return len(p), nil
}
