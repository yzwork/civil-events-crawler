// Code generated by 'gen/watchergen'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'listener/watchergen' for more details
// File was generated at 2018-06-10 05:20:29.092019048 +0000 UTC
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

func NewCivilTCRContractFilterers(contractAddress common.Address) *CivilTCRContractFilterers {
	return &CivilTCRContractFilterers{
		contractAddress: contractAddress,
	}
}

type CivilTCRContractFilterers struct {
	contractAddress common.Address
	contract        *contract.CivilTCRContract
}

func (r *CivilTCRContractFilterers) ContractName() string {
	return "CivilTCRContract"
}

func (r *CivilTCRContractFilterers) StartFilterers(client bind.ContractBackend,
	pastEvents *[]model.CivilEvent, startBlock uint64) error {
	return r.StartCivilTCRContractFilterers(client, pastEvents, startBlock)
}

// StartCivilTCRContractFilterers retrieves events for CivilTCRContract
func (r *CivilTCRContractFilterers) StartCivilTCRContractFilterers(client bind.ContractBackend,
	pastEvents *[]model.CivilEvent, startBlock uint64) error {
	contract, err := contract.NewCivilTCRContract(r.contractAddress, client)
	if err != nil {
		log.Errorf("Error initializing StartCivilTCRContract: err: %v", err)
		return err
	}
	r.contract = contract

	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	err = r.startFilterAppealGranted(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving AppealGranted: err: %v", err)
	}

	err = r.startFilterAppealRequested(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving AppealRequested: err: %v", err)
	}

	err = r.startFilterApplication(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving Application: err: %v", err)
	}

	err = r.startFilterApplicationRemoved(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving ApplicationRemoved: err: %v", err)
	}

	err = r.startFilterApplicationWhitelisted(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving ApplicationWhitelisted: err: %v", err)
	}

	err = r.startFilterChallenge(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving Challenge: err: %v", err)
	}

	err = r.startFilterChallengeFailed(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving ChallengeFailed: err: %v", err)
	}

	err = r.startFilterChallengeSucceeded(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving ChallengeSucceeded: err: %v", err)
	}

	err = r.startFilterDeposit(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving Deposit: err: %v", err)
	}

	err = r.startFilterFailedChallengeOverturned(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving FailedChallengeOverturned: err: %v", err)
	}

	err = r.startFilterGovernmentTransfered(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving GovernmentTransfered: err: %v", err)
	}

	err = r.startFilterGrantedAppealChallenged(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving GrantedAppealChallenged: err: %v", err)
	}

	err = r.startFilterGrantedAppealConfirmed(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving GrantedAppealConfirmed: err: %v", err)
	}

	err = r.startFilterGrantedAppealOverturned(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving GrantedAppealOverturned: err: %v", err)
	}

	err = r.startFilterListingRemoved(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving ListingRemoved: err: %v", err)
	}

	err = r.startFilterListingWithdrawn(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving ListingWithdrawn: err: %v", err)
	}

	err = r.startFilterRewardClaimed(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving RewardClaimed: err: %v", err)
	}

	err = r.startFilterSuccessfulChallengeOverturned(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving SuccessfulChallengeOverturned: err: %v", err)
	}

	err = r.startFilterTouchAndRemoved(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving TouchAndRemoved: err: %v", err)
	}

	err = r.startFilterWithdrawal(opts, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving Withdrawal: err: %v", err)
	}

	return nil
}

func (r *CivilTCRContractFilterers) startFilterAppealGranted(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterAppealGranted(
		opts,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event AppealGranted: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("AppealGranted", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (r *CivilTCRContractFilterers) startFilterAppealRequested(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterAppealRequested(
		opts,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event AppealRequested: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("AppealRequested", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (r *CivilTCRContractFilterers) startFilterApplication(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterApplication(
		opts,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event Application: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("Application", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (r *CivilTCRContractFilterers) startFilterApplicationRemoved(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterApplicationRemoved(
		opts,
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event ApplicationRemoved: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("ApplicationRemoved", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (r *CivilTCRContractFilterers) startFilterApplicationWhitelisted(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterApplicationWhitelisted(
		opts,
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event ApplicationWhitelisted: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("ApplicationWhitelisted", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (r *CivilTCRContractFilterers) startFilterChallenge(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterChallenge(
		opts,
		[]common.Address{},
		[]*big.Int{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event Challenge: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("Challenge", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (r *CivilTCRContractFilterers) startFilterChallengeFailed(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterChallengeFailed(
		opts,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event ChallengeFailed: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("ChallengeFailed", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (r *CivilTCRContractFilterers) startFilterChallengeSucceeded(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterChallengeSucceeded(
		opts,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event ChallengeSucceeded: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("ChallengeSucceeded", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (r *CivilTCRContractFilterers) startFilterDeposit(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterDeposit(
		opts,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event Deposit: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("Deposit", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (r *CivilTCRContractFilterers) startFilterFailedChallengeOverturned(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterFailedChallengeOverturned(
		opts,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event FailedChallengeOverturned: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("FailedChallengeOverturned", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (r *CivilTCRContractFilterers) startFilterGovernmentTransfered(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterGovernmentTransfered(
		opts,
	)
	if err != nil {
		log.Errorf("Error getting event GovernmentTransfered: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("GovernmentTransfered", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (r *CivilTCRContractFilterers) startFilterGrantedAppealChallenged(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterGrantedAppealChallenged(
		opts,
		[]common.Address{},
		[]*big.Int{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event GrantedAppealChallenged: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("GrantedAppealChallenged", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (r *CivilTCRContractFilterers) startFilterGrantedAppealConfirmed(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterGrantedAppealConfirmed(
		opts,
		[]common.Address{},
		[]*big.Int{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event GrantedAppealConfirmed: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("GrantedAppealConfirmed", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (r *CivilTCRContractFilterers) startFilterGrantedAppealOverturned(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterGrantedAppealOverturned(
		opts,
		[]common.Address{},
		[]*big.Int{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event GrantedAppealOverturned: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("GrantedAppealOverturned", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (r *CivilTCRContractFilterers) startFilterListingRemoved(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterListingRemoved(
		opts,
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event ListingRemoved: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("ListingRemoved", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (r *CivilTCRContractFilterers) startFilterListingWithdrawn(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterListingWithdrawn(
		opts,
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event ListingWithdrawn: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("ListingWithdrawn", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (r *CivilTCRContractFilterers) startFilterRewardClaimed(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterRewardClaimed(
		opts,
		[]*big.Int{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event RewardClaimed: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("RewardClaimed", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (r *CivilTCRContractFilterers) startFilterSuccessfulChallengeOverturned(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterSuccessfulChallengeOverturned(
		opts,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event SuccessfulChallengeOverturned: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("SuccessfulChallengeOverturned", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (r *CivilTCRContractFilterers) startFilterTouchAndRemoved(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterTouchAndRemoved(
		opts,
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event TouchAndRemoved: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("TouchAndRemoved", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func (r *CivilTCRContractFilterers) startFilterWithdrawal(opts *bind.FilterOpts, pastEvents *[]model.CivilEvent) error {
	itr, err := r.contract.FilterWithdrawal(
		opts,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event Withdrawal: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("Withdrawal", r.contractAddress, itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}
