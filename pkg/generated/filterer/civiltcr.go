// Code generated by 'gen/eventhandlergen.go'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'gen/filterergen_template.go' for more details
// File was generated at 2018-08-08 20:43:49.4427766 +0000 UTC
package filterer

import (
	"fmt"
	log "github.com/golang/glog"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/joincivil/civil-events-crawler/pkg/generated/contract"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/utils"

	"math/big"
)

// TODO(IS): Need to move this to a central place, use it outside the package
var eventTypesCivilTCRContract = []string{
	"AppealGranted",
	"AppealRequested",
	"Application",
	"ApplicationRemoved",
	"ApplicationWhitelisted",
	"Challenge",
	"ChallengeFailed",
	"ChallengeSucceeded",
	"Deposit",
	"FailedChallengeOverturned",
	"GovernmentTransfered",
	"GrantedAppealChallenged",
	"GrantedAppealConfirmed",
	"GrantedAppealOverturned",
	"ListingRemoved",
	"ListingWithdrawn",
	"RewardClaimed",
	"SuccessfulChallengeOverturned",
	"TouchAndRemoved",
	"Withdrawal",
}

// TODO(IS): Need to move this to a central place, use it outside the package
func EventTypesCivilTCRContract() []string {
	tmp := make([]string, len(eventTypesCivilTCRContract))
	copy(tmp, eventTypesCivilTCRContract)
	return tmp
}

func NewCivilTCRContractFilterers(contractAddress common.Address) *CivilTCRContractFilterers {
	contractFilterers := &CivilTCRContractFilterers{
		contractAddress:   contractAddress,
		eventTypes:        eventTypesCivilTCRContract,
		eventToStartBlock: make(map[string]uint64),
		lastEvents:        make([]*model.Event, 0),
	}
	for _, eventType := range contractFilterers.eventTypes {
		contractFilterers.eventToStartBlock[eventType] = 0
	}
	return contractFilterers
}

type CivilTCRContractFilterers struct {
	contractAddress   common.Address
	contract          *contract.CivilTCRContract
	eventTypes        []string
	eventToStartBlock map[string]uint64
	lastEvents        []*model.Event
}

func (f *CivilTCRContractFilterers) ContractName() string {
	return "CivilTCRContract"
}

func (f *CivilTCRContractFilterers) ContractAddress() common.Address {
	return f.contractAddress
}

func (f *CivilTCRContractFilterers) StartFilterers(client bind.ContractBackend, pastEvents []*model.Event) (error, []*model.Event) {
	return f.StartCivilTCRContractFilterers(client, pastEvents)
}

func (f *CivilTCRContractFilterers) EventTypes() []string {
	return f.eventTypes
}

func (f *CivilTCRContractFilterers) UpdateStartBlock(eventType string, startBlock uint64) {
	f.eventToStartBlock[eventType] = startBlock
}

func (f *CivilTCRContractFilterers) LastEvents() []*model.Event {
	return f.lastEvents
}

// StartCivilTCRContractFilterers retrieves events for CivilTCRContract
func (f *CivilTCRContractFilterers) StartCivilTCRContractFilterers(client bind.ContractBackend, pastEvents []*model.Event) (error, []*model.Event) {
	contract, err := contract.NewCivilTCRContract(f.contractAddress, client)
	if err != nil {
		log.Errorf("Error initializing StartCivilTCRContract: err: %v", err)
		return err, pastEvents
	}
	f.contract = contract
	var startBlock uint64
	prevEventsLength := len(pastEvents)

	startBlock = f.eventToStartBlock["AppealGranted"]
	err, pastEvents = f.startFilterAppealGranted(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving AppealGranted: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	startBlock = f.eventToStartBlock["AppealRequested"]
	err, pastEvents = f.startFilterAppealRequested(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving AppealRequested: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	startBlock = f.eventToStartBlock["Application"]
	err, pastEvents = f.startFilterApplication(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving Application: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	startBlock = f.eventToStartBlock["ApplicationRemoved"]
	err, pastEvents = f.startFilterApplicationRemoved(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving ApplicationRemoved: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	startBlock = f.eventToStartBlock["ApplicationWhitelisted"]
	err, pastEvents = f.startFilterApplicationWhitelisted(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving ApplicationWhitelisted: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	startBlock = f.eventToStartBlock["Challenge"]
	err, pastEvents = f.startFilterChallenge(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving Challenge: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	startBlock = f.eventToStartBlock["ChallengeFailed"]
	err, pastEvents = f.startFilterChallengeFailed(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving ChallengeFailed: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	startBlock = f.eventToStartBlock["ChallengeSucceeded"]
	err, pastEvents = f.startFilterChallengeSucceeded(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving ChallengeSucceeded: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	startBlock = f.eventToStartBlock["Deposit"]
	err, pastEvents = f.startFilterDeposit(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving Deposit: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	startBlock = f.eventToStartBlock["FailedChallengeOverturned"]
	err, pastEvents = f.startFilterFailedChallengeOverturned(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving FailedChallengeOverturned: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	startBlock = f.eventToStartBlock["GovernmentTransfered"]
	err, pastEvents = f.startFilterGovernmentTransfered(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving GovernmentTransfered: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	startBlock = f.eventToStartBlock["GrantedAppealChallenged"]
	err, pastEvents = f.startFilterGrantedAppealChallenged(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving GrantedAppealChallenged: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	startBlock = f.eventToStartBlock["GrantedAppealConfirmed"]
	err, pastEvents = f.startFilterGrantedAppealConfirmed(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving GrantedAppealConfirmed: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	startBlock = f.eventToStartBlock["GrantedAppealOverturned"]
	err, pastEvents = f.startFilterGrantedAppealOverturned(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving GrantedAppealOverturned: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	startBlock = f.eventToStartBlock["ListingRemoved"]
	err, pastEvents = f.startFilterListingRemoved(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving ListingRemoved: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	startBlock = f.eventToStartBlock["ListingWithdrawn"]
	err, pastEvents = f.startFilterListingWithdrawn(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving ListingWithdrawn: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	startBlock = f.eventToStartBlock["RewardClaimed"]
	err, pastEvents = f.startFilterRewardClaimed(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving RewardClaimed: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	startBlock = f.eventToStartBlock["SuccessfulChallengeOverturned"]
	err, pastEvents = f.startFilterSuccessfulChallengeOverturned(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving SuccessfulChallengeOverturned: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	startBlock = f.eventToStartBlock["TouchAndRemoved"]
	err, pastEvents = f.startFilterTouchAndRemoved(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving TouchAndRemoved: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	startBlock = f.eventToStartBlock["Withdrawal"]
	err, pastEvents = f.startFilterWithdrawal(startBlock, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving Withdrawal: err: %v", err), pastEvents
	}
	if len(pastEvents) > prevEventsLength {
		f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents)-1])
		prevEventsLength = len(pastEvents)
	}

	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterAppealGranted(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for AppealGranted for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterAppealGranted(
		opts,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event AppealGranted: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("AppealGranted", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterAppealRequested(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for AppealRequested for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterAppealRequested(
		opts,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event AppealRequested: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("AppealRequested", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterApplication(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for Application for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterApplication(
		opts,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event Application: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("Application", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterApplicationRemoved(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for ApplicationRemoved for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterApplicationRemoved(
		opts,
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event ApplicationRemoved: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("ApplicationRemoved", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterApplicationWhitelisted(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for ApplicationWhitelisted for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterApplicationWhitelisted(
		opts,
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event ApplicationWhitelisted: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("ApplicationWhitelisted", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterChallenge(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for Challenge for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterChallenge(
		opts,
		[]common.Address{},
		[]*big.Int{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event Challenge: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("Challenge", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterChallengeFailed(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for ChallengeFailed for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterChallengeFailed(
		opts,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event ChallengeFailed: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("ChallengeFailed", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterChallengeSucceeded(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for ChallengeSucceeded for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterChallengeSucceeded(
		opts,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event ChallengeSucceeded: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("ChallengeSucceeded", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterDeposit(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for Deposit for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterDeposit(
		opts,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event Deposit: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("Deposit", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterFailedChallengeOverturned(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for FailedChallengeOverturned for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterFailedChallengeOverturned(
		opts,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event FailedChallengeOverturned: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("FailedChallengeOverturned", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterGovernmentTransfered(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for GovernmentTransfered for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterGovernmentTransfered(
		opts,
	)
	if err != nil {
		log.Errorf("Error getting event GovernmentTransfered: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("GovernmentTransfered", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterGrantedAppealChallenged(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for GrantedAppealChallenged for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterGrantedAppealChallenged(
		opts,
		[]common.Address{},
		[]*big.Int{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event GrantedAppealChallenged: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("GrantedAppealChallenged", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterGrantedAppealConfirmed(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for GrantedAppealConfirmed for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterGrantedAppealConfirmed(
		opts,
		[]common.Address{},
		[]*big.Int{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event GrantedAppealConfirmed: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("GrantedAppealConfirmed", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterGrantedAppealOverturned(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for GrantedAppealOverturned for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterGrantedAppealOverturned(
		opts,
		[]common.Address{},
		[]*big.Int{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event GrantedAppealOverturned: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("GrantedAppealOverturned", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterListingRemoved(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for ListingRemoved for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterListingRemoved(
		opts,
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event ListingRemoved: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("ListingRemoved", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterListingWithdrawn(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for ListingWithdrawn for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterListingWithdrawn(
		opts,
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event ListingWithdrawn: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("ListingWithdrawn", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterRewardClaimed(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for RewardClaimed for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterRewardClaimed(
		opts,
		[]*big.Int{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event RewardClaimed: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("RewardClaimed", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterSuccessfulChallengeOverturned(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for SuccessfulChallengeOverturned for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterSuccessfulChallengeOverturned(
		opts,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event SuccessfulChallengeOverturned: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("SuccessfulChallengeOverturned", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterTouchAndRemoved(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for TouchAndRemoved for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterTouchAndRemoved(
		opts,
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event TouchAndRemoved: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("TouchAndRemoved", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}

func (f *CivilTCRContractFilterers) startFilterWithdrawal(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for Withdrawal for contract %v", f.contractAddress.Hex())
	itr, err := f.contract.FilterWithdrawal(
		opts,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event Withdrawal: %v", err)
		return err, pastEvents
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("Withdrawal", f.ContractName(), f.contractAddress, itr.Event, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return nil, pastEvents
}
