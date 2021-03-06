package facts_test

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"io"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/go-verifiable-data/facts"
	"github.com/monetha/go-verifiable-data/ipfs"
)

func TestPrivateDataExchange(t *testing.T) {
	var (
		ipfsURL       = "https://ipfs.infura.io:5001"
		randSeed      = int64(1)
		factKey       = [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
		factData      = []byte("this is a very secret message!")
		exchangeStake = big.NewInt(100000)
	)

	type testContext struct {
		context.Context
		*testing.T
		*passportWithActors
		Clock                *clockMock
		Rand                 io.Reader
		IPFS                 *ipfs.IPFS
		DataRequesterKey     *ecdsa.PrivateKey
		DataRequesterAddress common.Address
		AdjustTime           func(adjustment time.Duration)
	}

	type actAssertFun func(
		tc *testContext,
	)

	arrangeActAssert := func(t *testing.T, cassetteName string, actAssert actAssertFun) {
		// start http recorder
		r, err := recorder.New(cassetteName)
		if err != nil {
			panic(err)
		}
		defer func() { _ = r.Stop() }() // Make sure recorder is stopped once done with it
		c := &http.Client{Transport: r}

		ctx := context.Background()
		rnd := rand.New(rand.NewSource(randSeed))

		dataRequesterKey, dataRequesterAddress := generatePrivateKey(t, rnd)
		pa := newPassportWithActors(ctx, t, rnd, account{Address: dataRequesterAddress, Balance: math.MaxInt64})

		fs, err := ipfs.NewWithClient(ipfsURL, c)
		if err != nil {
			t.Fatalf("creating instance of IPFS: %v", err)
		}

		wr := facts.NewPrivateDataWriter(pa.FactProviderSession(), fs)

		_, err = wr.WritePrivateData(ctx, pa.PassportAddress, factKey, factData, rnd)
		if err != nil {
			t.Fatalf("WritePrivateData: %v", err)
		}

		clock := newClockMock()
		actAssert(&testContext{
			Context:              ctx,
			T:                    t,
			passportWithActors:   pa,
			Clock:                clock,
			Rand:                 rnd,
			IPFS:                 fs,
			DataRequesterKey:     dataRequesterKey,
			DataRequesterAddress: dataRequesterAddress,
			AdjustTime: func(adjustment time.Duration) {
				if err != pa.Eth.AdjustTime(adjustment) {
					t.Fatalf("AdjustTime: %v", err)
				}
				clock.Add(adjustment)
			},
		})
	}

	t.Run("data requester is able to propose private data exchange", func(t *testing.T) {
		arrangeActAssert(t, "fixtures/private-data-exchange-propose", func(tc *testContext) {
			_, err := facts.
				NewExchangeProposer(tc.NewSession(tc.DataRequesterKey)).
				ProposePrivateDataExchange(tc.Context, tc.PassportAddress, tc.FactProviderAddress, factKey, exchangeStake, tc.Rand)
			if err != nil {
				tc.Fatalf("ProposePrivateDataExchange: %v", err)
			}
		})
	})

	t.Run("data requester is able to timeout proposed private data exchange after 25 hours", func(t *testing.T) {
		arrangeActAssert(t, "fixtures/private-data-exchange-timeout", func(tc *testContext) {
			dataRequesterSession := tc.NewSession(tc.DataRequesterKey)
			propRes, err := facts.
				NewExchangeProposer(dataRequesterSession).
				ProposePrivateDataExchange(tc.Context, tc.PassportAddress, tc.FactProviderAddress, factKey, exchangeStake, tc.Rand)
			if err != nil {
				tc.Fatalf("ProposePrivateDataExchange: %v", err)
			}

			tc.AdjustTime(25 * time.Hour)

			if err := facts.
				NewExchangeTimeouter(dataRequesterSession, tc.Clock).
				TimeoutPrivateDataExchange(tc.Context, tc.PassportAddress, propRes.ExchangeIdx); err != nil {
				tc.Errorf("TimeoutPrivateDataExchange: %v", err)
			}
		})
	})

	t.Run("passport owner is able to accept proposed private data exchange", func(t *testing.T) {
		arrangeActAssert(t, "fixtures/private-data-exchange-accept", func(tc *testContext) {
			propRes, err := facts.
				NewExchangeProposer(tc.NewSession(tc.DataRequesterKey)).
				ProposePrivateDataExchange(tc.Context, tc.PassportAddress, tc.FactProviderAddress, factKey, exchangeStake, tc.Rand)
			if err != nil {
				tc.Fatalf("ProposePrivateDataExchange: %v", err)
			}

			if err := facts.
				NewExchangeAcceptor(tc.Eth, tc.PassportOwnerKey, tc.IPFS, tc.Clock).
				AcceptPrivateDataExchange(tc.Context, tc.PassportAddress, propRes.ExchangeIdx); err != nil {
				tc.Fatalf("AcceptPrivateDataExchange: %v", err)
			}
		})
	})

	t.Run("data requester is able to read an decrypt the data of accepted private data exchange", func(t *testing.T) {
		arrangeActAssert(t, "fixtures/private-data-exchange-read", func(tc *testContext) {
			propRes, err := facts.
				NewExchangeProposer(tc.NewSession(tc.DataRequesterKey)).
				ProposePrivateDataExchange(tc.Context, tc.PassportAddress, tc.FactProviderAddress, factKey, exchangeStake, tc.Rand)
			if err != nil {
				tc.Fatalf("ProposePrivateDataExchange: %v", err)
			}

			if err := facts.
				NewExchangeAcceptor(tc.Eth, tc.PassportOwnerKey, tc.IPFS, tc.Clock).
				AcceptPrivateDataExchange(tc.Context, tc.PassportAddress, propRes.ExchangeIdx); err != nil {
				tc.Fatalf("AcceptPrivateDataExchange: %v", err)
			}

			readFactData, err := facts.
				NewExchangeDataReader(tc.Eth, tc.IPFS).
				ReadPrivateData(tc.Context, tc.PassportAddress, propRes.ExchangeIdx, propRes.ExchangeKey)
			if err != nil {
				tc.Fatalf("ReadPrivateData: %v", err)
			}

			if !bytes.Equal(factData, readFactData) {
				tc.Fatalf("expected to read %v, but got %v", factData, readFactData)
			}
		})
	})

	t.Run("data requester is able to finish accepted private data exchange", func(t *testing.T) {
		arrangeActAssert(t, "fixtures/private-data-exchange-finish", func(tc *testContext) {
			propRes, err := facts.
				NewExchangeProposer(tc.NewSession(tc.DataRequesterKey)).
				ProposePrivateDataExchange(tc.Context, tc.PassportAddress, tc.FactProviderAddress, factKey, exchangeStake, tc.Rand)
			if err != nil {
				tc.Fatalf("ProposePrivateDataExchange: %v", err)
			}

			if err := facts.
				NewExchangeAcceptor(tc.Eth, tc.PassportOwnerKey, tc.IPFS, tc.Clock).
				AcceptPrivateDataExchange(tc.Context, tc.PassportAddress, propRes.ExchangeIdx); err != nil {
				tc.Fatalf("AcceptPrivateDataExchange: %v", err)
			}

			if err := facts.
				NewExchangeFinisher(tc.Eth.NewSession(tc.DataRequesterKey), tc.Clock).
				FinishPrivateDataExchange(tc.Context, tc.PassportAddress, propRes.ExchangeIdx); err != nil {
				tc.Fatalf("FinishPrivateDataExchange: %v", err)
			}
		})
	})

	t.Run("data requester is able to dispute fairly accepted private data exchange, but will be recognized as cheater", func(t *testing.T) {
		arrangeActAssert(t, "fixtures/private-data-exchange-finish", func(tc *testContext) {
			propRes, err := facts.
				NewExchangeProposer(tc.NewSession(tc.DataRequesterKey)).
				ProposePrivateDataExchange(tc.Context, tc.PassportAddress, tc.FactProviderAddress, factKey, exchangeStake, tc.Rand)
			if err != nil {
				tc.Fatalf("ProposePrivateDataExchange: %v", err)
			}

			if err := facts.
				NewExchangeAcceptor(tc.Eth, tc.PassportOwnerKey, tc.IPFS, tc.Clock).
				AcceptPrivateDataExchange(tc.Context, tc.PassportAddress, propRes.ExchangeIdx); err != nil {
				tc.Fatalf("AcceptPrivateDataExchange: %v", err)
			}

			dispRes, err := facts.
				NewExchangeDisputer(tc.Eth.NewSession(tc.DataRequesterKey), tc.Clock).
				DisputePrivateDataExchange(tc.Context, tc.PassportAddress, propRes.ExchangeIdx, propRes.ExchangeKey)
			if err != nil {
				tc.Fatalf("DisputePrivateDataExchange: %v", err)
			}

			if dispRes.Successful {
				tc.Error("expected unsuccessful dispute result")
			}

			if dispRes.Cheater != tc.DataRequesterAddress {
				tc.Errorf("expected cheater address %v, but got %v", tc.DataRequesterAddress.String(), dispRes.Cheater.String())
			}
		})
	})

	t.Run("data requester is able to read an decrypt the data of finished private data exchange", func(t *testing.T) {
		arrangeActAssert(t, "fixtures/private-data-exchange-read", func(tc *testContext) {
			propRes, err := facts.
				NewExchangeProposer(tc.NewSession(tc.DataRequesterKey)).
				ProposePrivateDataExchange(tc.Context, tc.PassportAddress, tc.FactProviderAddress, factKey, exchangeStake, tc.Rand)
			if err != nil {
				tc.Fatalf("ProposePrivateDataExchange: %v", err)
			}

			if err := facts.
				NewExchangeAcceptor(tc.Eth, tc.PassportOwnerKey, tc.IPFS, tc.Clock).
				AcceptPrivateDataExchange(tc.Context, tc.PassportAddress, propRes.ExchangeIdx); err != nil {
				tc.Fatalf("AcceptPrivateDataExchange: %v", err)
			}

			if err := facts.
				NewExchangeFinisher(tc.Eth.NewSession(tc.DataRequesterKey), tc.Clock).
				FinishPrivateDataExchange(tc.Context, tc.PassportAddress, propRes.ExchangeIdx); err != nil {
				tc.Fatalf("FinishPrivateDataExchange: %v", err)
			}

			readFactData, err := facts.
				NewExchangeDataReader(tc.Eth, tc.IPFS).
				ReadPrivateData(tc.Context, tc.PassportAddress, propRes.ExchangeIdx, propRes.ExchangeKey)
			if err != nil {
				tc.Fatalf("ReadPrivateData: %v", err)
			}

			if !bytes.Equal(factData, readFactData) {
				tc.Fatalf("expected to read %v, but got %v", factData, readFactData)
			}
		})
	})
}
