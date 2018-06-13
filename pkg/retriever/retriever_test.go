// Package retriever_test contains tests for retriever package
package retriever_test

import (
	"errors"
	log "github.com/golang/glog"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/joincivil/civil-events-crawler/pkg/generated/contract"
	"github.com/joincivil/civil-events-crawler/pkg/generated/filterer"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/retriever"
)

// Using rinkeby for now
// TODO(IS) change to simulated backend, write more tests
const (
	rinkebyAddress = "https://rinkeby.infura.io"
	testTCRAddress = "0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d"
	startBlock     = 2335623
)

func setupRinkebyClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial(rinkebyAddress)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		return nil, err
	}
	return client, nil
}

type testErrorFilterer struct{}

func (t *testErrorFilterer) ContractName() string {
	return "TestErrorContract"
}

func (t *testErrorFilterer) StartFilterers(client bind.ContractBackend,
	pastEvents *[]model.CivilEvent, startBlock uint64) error {
	return errors.New("This is an error starting filterer")
}

// TestEventCollection tests that events are being collected,
func TestEventCollection(t *testing.T) {
	client, err := setupRinkebyClient()
	if err != nil {
		t.Errorf("Error connecting to rinkeby: %v", err)
	}
	filterers := []model.ContractFilterers{
		filterer.NewCivilTCRContractFilterers(common.HexToAddress(testTCRAddress)),
	}
	retrieve := retriever.NewCivilEventRetriever(client, startBlock, filterers)
	err = retrieve.Retrieve()
	if err != nil {
		t.Errorf("Error retrieving events: %v", err)
	}
	pastEvents := retrieve.PastEvents
	if len(pastEvents) == 0 {
		t.Error("No events collected")
	}
}

func TestErrorFilterer(t *testing.T) {
	client, err := setupRinkebyClient()
	if err != nil {
		t.Errorf("Error connecting to rinkeby: %v", err)
	}
	filterers := []model.ContractFilterers{
		&testErrorFilterer{},
	}
	retrieve := retriever.NewCivilEventRetriever(client, startBlock, filterers)
	err = retrieve.Retrieve()
	if err == nil {
		t.Errorf("Should have failed to retrieve due to bad filterer: %v", err)
	}
}

// TestSorting tests that sorting is happening by block number
func TestSorting(t *testing.T) {
	testEvent1 := &contract.CivilTCRContractApplicationWhitelisted{
		ListingAddress: common.HexToAddress(testTCRAddress),
		Raw: types.Log{
			Address:     common.HexToAddress(testTCRAddress),
			Topics:      []common.Hash{},
			Data:        []byte{},
			BlockNumber: 8888888,
		},
	}
	testEvent2 := &contract.CivilTCRContractApplication{
		ListingAddress: common.HexToAddress(testTCRAddress),
		Deposit:        big.NewInt(1000),
		AppEndDate:     big.NewInt(1653860896),
		Data:           "DATA",
		Applicant:      common.HexToAddress(testTCRAddress),
		Raw: types.Log{
			Address:     common.HexToAddress(testTCRAddress),
			Topics:      []common.Hash{},
			Data:        []byte{},
			BlockNumber: 8888886,
		},
	}
	client, err := setupRinkebyClient()
	if err != nil {
		t.Errorf("Error connecting to rinkeby: %v", err)
	}
	filterers := []model.ContractFilterers{
		filterer.NewCivilTCRContractFilterers(common.HexToAddress(testTCRAddress)),
	}
	retrieve := retriever.NewCivilEventRetriever(client, startBlock, filterers)
	model1, _ := model.NewCivilEvent("ApplicationWhitelisted", common.HexToAddress(testTCRAddress), testEvent1)
	model2, _ := model.NewCivilEvent("Application", common.HexToAddress(testTCRAddress), testEvent2)
	retrieve.PastEvents = append(retrieve.PastEvents, *model1, *model2)
	ok := retrieve.SortEventsByBlock()
	if !ok {
		t.Error("Sorting didn't happen")
	}
}

// TestSorting tests that sorting is happening by block number
func TestSortingFail(t *testing.T) {
	testEvent1 := &contract.CivilTCRContractApplicationWhitelisted{
		ListingAddress: common.HexToAddress(testTCRAddress),
		Raw: types.Log{
			Address: common.HexToAddress(testTCRAddress),
			Topics:  []common.Hash{},
			Data:    []byte{},
		},
	}
	testEvent2 := &contract.CivilTCRContractApplication{
		ListingAddress: common.HexToAddress(testTCRAddress),
		Deposit:        big.NewInt(1000),
		AppEndDate:     big.NewInt(1653860896),
		Data:           "DATA",
		Applicant:      common.HexToAddress(testTCRAddress),
		Raw: types.Log{
			Address:     common.HexToAddress(testTCRAddress),
			Topics:      []common.Hash{},
			Data:        []byte{},
			BlockNumber: 8888886,
		},
	}
	client, err := setupRinkebyClient()
	if err != nil {
		t.Errorf("Error connecting to rinkeby: %v", err)
	}
	filterers := []model.ContractFilterers{
		filterer.NewCivilTCRContractFilterers(common.HexToAddress(testTCRAddress)),
	}
	retrieve := retriever.NewCivilEventRetriever(client, startBlock, filterers)
	model1, _ := model.NewCivilEvent("ApplicationWhitelisted", common.HexToAddress(testTCRAddress), testEvent1)
	model2, _ := model.NewCivilEvent("Application", common.HexToAddress(testTCRAddress), testEvent2)
	retrieve.PastEvents = append(retrieve.PastEvents, *model1, *model2)
	ok := retrieve.SortEventsByBlock()
	if ok {
		t.Error("Sorting happened when it shouldn't have")
	}

}
