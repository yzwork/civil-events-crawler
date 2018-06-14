// Code generated by 'gen/eventhandlergen.go'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'gen/filterergen_template.go' for more details
// File was generated at 2018-06-14 21:01:05.81749563 +0000 UTC
package filterer

import (
	"fmt"
	log "github.com/golang/glog"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/joincivil/civil-events-crawler/pkg/generated/contract"
	"github.com/joincivil/civil-events-crawler/pkg/model"

	"math/big"
)

var eventNamesNewsroomContract = []string{
	"ContentPublished",
	"NameChanged",
	"OwnershipTransferred",
	"RevisionSigned",
	"RevisionUpdated",
	"RoleAdded",
	"RoleRemoved",
}

func NewNewsroomContractFilterers(contractAddress common.Address) *NewsroomContractFilterers {
	var c NewsroomContractFilterers
	c.contractAddress = contractAddress
	c.eventNames = eventNamesNewsroomContract
	c.eventToStartBlock = make(map[string]uint64)
	for _, eventName := range c.eventNames {
		c.eventToStartBlock[eventName] = 0
	}
	return &c
}

type NewsroomContractFilterers struct {
	contractAddress   common.Address
	contract          *contract.NewsroomContract
	eventNames        []string
	eventToStartBlock map[string]uint64
}

func (f *NewsroomContractFilterers) ContractName() string {
	return "NewsroomContract"
}

func (f *NewsroomContractFilterers) StartFilterers(client bind.ContractBackend, pastEvents *[]model.CivilEvent) error {
	return f.StartNewsroomContractFilterers(client, pastEvents)
}

func (f *NewsroomContractFilterers) EventNames() []string {
	return f.eventNames
}

func (f *NewsroomContractFilterers) UpdateStartBlock(eventName string, startBlock int) {
	f.eventToStartBlock[eventName] = uint64(startBlock)
}

// StartNewsroomContractFilterers retrieves events for NewsroomContract
func (f *NewsroomContractFilterers) StartNewsroomContractFilterers(client bind.ContractBackend, pastEvents *[]model.CivilEvent) error {
	contract, err := contract.NewNewsroomContract(f.contractAddress, client)
	if err != nil {
		log.Errorf("Error initializing StartNewsroomContract: err: %v", err)
		return err
	}
	f.contract = contract
	var startBlock uint64

	startBlock = f.eventToStartBlock["ContentPublished"]
	err = f.startFilterContentPublished(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving ContentPublished: err: %v", err)
	}
	startBlock = f.eventToStartBlock["NameChanged"]
	err = f.startFilterNameChanged(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving NameChanged: err: %v", err)
	}
	startBlock = f.eventToStartBlock["OwnershipTransferred"]
	err = f.startFilterOwnershipTransferred(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving OwnershipTransferred: err: %v", err)
	}
	startBlock = f.eventToStartBlock["RevisionSigned"]
	err = f.startFilterRevisionSigned(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving RevisionSigned: err: %v", err)
	}
	startBlock = f.eventToStartBlock["RevisionUpdated"]
	err = f.startFilterRevisionUpdated(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving RevisionUpdated: err: %v", err)
	}
	startBlock = f.eventToStartBlock["RoleAdded"]
	err = f.startFilterRoleAdded(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving RoleAdded: err: %v", err)
	}
	startBlock = f.eventToStartBlock["RoleRemoved"]
	err = f.startFilterRoleRemoved(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving RoleRemoved: err: %v", err)
	}

	return nil
}

func (f *NewsroomContractFilterers) startFilterContentPublished(startBlock uint64, pastEvents *[]model.CivilEvent) error {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}
	itr, err := f.contract.FilterContentPublished(
		opts,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event ContentPublished: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent, err := model.NewCivilEvent("ContentPublished", r.contractAddress, itr.Event)
		if err != nil {
			log.Errorf("Error creating new civil event: event: %v, err: %v", itr.Event, err)
			continue
		}
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (f *NewsroomContractFilterers) startFilterNameChanged(startBlock uint64, pastEvents *[]model.CivilEvent) error {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}
	itr, err := f.contract.FilterNameChanged(
		opts,
	)
	if err != nil {
		log.Errorf("Error getting event NameChanged: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent, err := model.NewCivilEvent("NameChanged", r.contractAddress, itr.Event)
		if err != nil {
			log.Errorf("Error creating new civil event: event: %v, err: %v", itr.Event, err)
			continue
		}
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (f *NewsroomContractFilterers) startFilterOwnershipTransferred(startBlock uint64, pastEvents *[]model.CivilEvent) error {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}
	itr, err := f.contract.FilterOwnershipTransferred(
		opts,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event OwnershipTransferred: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent, err := model.NewCivilEvent("OwnershipTransferred", r.contractAddress, itr.Event)
		if err != nil {
			log.Errorf("Error creating new civil event: event: %v, err: %v", itr.Event, err)
			continue
		}
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (f *NewsroomContractFilterers) startFilterRevisionSigned(startBlock uint64, pastEvents *[]model.CivilEvent) error {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}
	itr, err := f.contract.FilterRevisionSigned(
		opts,
		[]*big.Int{},
		[]*big.Int{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event RevisionSigned: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent, err := model.NewCivilEvent("RevisionSigned", r.contractAddress, itr.Event)
		if err != nil {
			log.Errorf("Error creating new civil event: event: %v, err: %v", itr.Event, err)
			continue
		}
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (f *NewsroomContractFilterers) startFilterRevisionUpdated(startBlock uint64, pastEvents *[]model.CivilEvent) error {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}
	itr, err := f.contract.FilterRevisionUpdated(
		opts,
		[]common.Address{},
		[]*big.Int{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event RevisionUpdated: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent, err := model.NewCivilEvent("RevisionUpdated", r.contractAddress, itr.Event)
		if err != nil {
			log.Errorf("Error creating new civil event: event: %v, err: %v", itr.Event, err)
			continue
		}
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (f *NewsroomContractFilterers) startFilterRoleAdded(startBlock uint64, pastEvents *[]model.CivilEvent) error {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}
	itr, err := f.contract.FilterRoleAdded(
		opts,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event RoleAdded: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent, err := model.NewCivilEvent("RoleAdded", r.contractAddress, itr.Event)
		if err != nil {
			log.Errorf("Error creating new civil event: event: %v, err: %v", itr.Event, err)
			continue
		}
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (f *NewsroomContractFilterers) startFilterRoleRemoved(startBlock uint64, pastEvents *[]model.CivilEvent) error {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}
	itr, err := f.contract.FilterRoleRemoved(
		opts,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event RoleRemoved: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent, err := model.NewCivilEvent("RoleRemoved", r.contractAddress, itr.Event)
		if err != nil {
			log.Errorf("Error creating new civil event: event: %v, err: %v", itr.Event, err)
			continue
		}
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}
