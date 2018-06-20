// Code generated by 'gen/watchergen'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'listener/watchergen' for more details
// File was generated at 2018-06-20 18:19:20.694088542 +0000 UTC
package watcher

import (
	"fmt"
	log "github.com/golang/glog"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/joincivil/civil-events-crawler/pkg/generated/contract"
	"github.com/joincivil/civil-events-crawler/pkg/model"

	"math/big"
)

func NewNewsroomContractWatchers(contractAddress common.Address) *NewsroomContractWatchers {
	return &NewsroomContractWatchers{
		contractAddress: contractAddress,
	}
}

type NewsroomContractWatchers struct {
	contractAddress common.Address
	contract        *contract.NewsroomContract
}

func (w *NewsroomContractWatchers) ContractName() string {
	return "NewsroomContract"
}

func (w *NewsroomContractWatchers) StartWatchers(client bind.ContractBackend,
	eventRecvChan chan model.CivilEvent) ([]event.Subscription, error) {
	return w.StartNewsroomContractWatchers(client, eventRecvChan)
}

// StartNewsroomContractWatchers starts up the event watchers for NewsroomContract
func (w *NewsroomContractWatchers) StartNewsroomContractWatchers(client bind.ContractBackend,
	eventRecvChan chan model.CivilEvent) ([]event.Subscription, error) {
	contract, err := contract.NewNewsroomContract(w.contractAddress, client)
	if err != nil {
		log.Errorf("Error initializing StartNewsroomContract: err: %v", err)
		return nil, err
	}
	w.contract = contract

	var sub event.Subscription
	subs := []event.Subscription{}

	sub, err = w.startWatchContentPublished(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startContentPublished: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchNameChanged(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startNameChanged: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchOwnershipTransferred(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startOwnershipTransferred: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchRevisionSigned(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startRevisionSigned: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchRevisionUpdated(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startRevisionUpdated: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchRoleAdded(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startRoleAdded: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchRoleRemoved(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startRoleRemoved: err: %v", err)
	}
	subs = append(subs, sub)

	return subs, nil
}

func (w *NewsroomContractWatchers) startWatchContentPublished(eventRecvChan chan model.CivilEvent) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *contract.NewsroomContractContentPublished, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.NewsroomContractContentPublished)
				sub, err := w.contract.WatchContentPublished(
					opts,
					recvChan,
					[]common.Address{},
					[]*big.Int{},
				)
				if err != nil {
					if sub != nil {
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("Retrying start WatchContentPublished: %v", err)
					continue
				}
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchContentPublished: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchContentPublished for contract %v", w.contractAddress.Hex())
		for {
			select {
			case event := <-recvChan:
				civilEvent, err := model.NewCivilEvent("ContentPublished", w.contractAddress, event)
				if err != nil {
					log.Errorf("Error creating new civil event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- *civilEvent:
				case err := <-sub.Err():
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting WatchContentPublished, fatal (a): %v", err)
						return err
					}
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error restarting WatchContentPublished, fatal (b): %v", err)
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchNameChanged(eventRecvChan chan model.CivilEvent) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *contract.NewsroomContractNameChanged, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.NewsroomContractNameChanged)
				sub, err := w.contract.WatchNameChanged(
					opts,
					recvChan,
				)
				if err != nil {
					if sub != nil {
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("Retrying start WatchNameChanged: %v", err)
					continue
				}
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchNameChanged: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchNameChanged for contract %v", w.contractAddress.Hex())
		for {
			select {
			case event := <-recvChan:
				civilEvent, err := model.NewCivilEvent("NameChanged", w.contractAddress, event)
				if err != nil {
					log.Errorf("Error creating new civil event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- *civilEvent:
				case err := <-sub.Err():
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting WatchNameChanged, fatal (a): %v", err)
						return err
					}
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error restarting WatchNameChanged, fatal (b): %v", err)
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchOwnershipTransferred(eventRecvChan chan model.CivilEvent) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *contract.NewsroomContractOwnershipTransferred, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.NewsroomContractOwnershipTransferred)
				sub, err := w.contract.WatchOwnershipTransferred(
					opts,
					recvChan,
					[]common.Address{},
					[]common.Address{},
				)
				if err != nil {
					if sub != nil {
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("Retrying start WatchOwnershipTransferred: %v", err)
					continue
				}
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchOwnershipTransferred: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchOwnershipTransferred for contract %v", w.contractAddress.Hex())
		for {
			select {
			case event := <-recvChan:
				civilEvent, err := model.NewCivilEvent("OwnershipTransferred", w.contractAddress, event)
				if err != nil {
					log.Errorf("Error creating new civil event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- *civilEvent:
				case err := <-sub.Err():
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting WatchOwnershipTransferred, fatal (a): %v", err)
						return err
					}
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error restarting WatchOwnershipTransferred, fatal (b): %v", err)
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchRevisionSigned(eventRecvChan chan model.CivilEvent) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *contract.NewsroomContractRevisionSigned, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.NewsroomContractRevisionSigned)
				sub, err := w.contract.WatchRevisionSigned(
					opts,
					recvChan,
					[]*big.Int{},
					[]*big.Int{},
					[]common.Address{},
				)
				if err != nil {
					if sub != nil {
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("Retrying start WatchRevisionSigned: %v", err)
					continue
				}
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchRevisionSigned: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchRevisionSigned for contract %v", w.contractAddress.Hex())
		for {
			select {
			case event := <-recvChan:
				civilEvent, err := model.NewCivilEvent("RevisionSigned", w.contractAddress, event)
				if err != nil {
					log.Errorf("Error creating new civil event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- *civilEvent:
				case err := <-sub.Err():
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting WatchRevisionSigned, fatal (a): %v", err)
						return err
					}
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error restarting WatchRevisionSigned, fatal (b): %v", err)
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchRevisionUpdated(eventRecvChan chan model.CivilEvent) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *contract.NewsroomContractRevisionUpdated, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.NewsroomContractRevisionUpdated)
				sub, err := w.contract.WatchRevisionUpdated(
					opts,
					recvChan,
					[]common.Address{},
					[]*big.Int{},
					[]*big.Int{},
				)
				if err != nil {
					if sub != nil {
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("Retrying start WatchRevisionUpdated: %v", err)
					continue
				}
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchRevisionUpdated: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchRevisionUpdated for contract %v", w.contractAddress.Hex())
		for {
			select {
			case event := <-recvChan:
				civilEvent, err := model.NewCivilEvent("RevisionUpdated", w.contractAddress, event)
				if err != nil {
					log.Errorf("Error creating new civil event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- *civilEvent:
				case err := <-sub.Err():
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting WatchRevisionUpdated, fatal (a): %v", err)
						return err
					}
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error restarting WatchRevisionUpdated, fatal (b): %v", err)
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchRoleAdded(eventRecvChan chan model.CivilEvent) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *contract.NewsroomContractRoleAdded, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.NewsroomContractRoleAdded)
				sub, err := w.contract.WatchRoleAdded(
					opts,
					recvChan,
					[]common.Address{},
					[]common.Address{},
				)
				if err != nil {
					if sub != nil {
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("Retrying start WatchRoleAdded: %v", err)
					continue
				}
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchRoleAdded: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchRoleAdded for contract %v", w.contractAddress.Hex())
		for {
			select {
			case event := <-recvChan:
				civilEvent, err := model.NewCivilEvent("RoleAdded", w.contractAddress, event)
				if err != nil {
					log.Errorf("Error creating new civil event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- *civilEvent:
				case err := <-sub.Err():
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting WatchRoleAdded, fatal (a): %v", err)
						return err
					}
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error restarting WatchRoleAdded, fatal (b): %v", err)
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchRoleRemoved(eventRecvChan chan model.CivilEvent) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *contract.NewsroomContractRoleRemoved, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.NewsroomContractRoleRemoved)
				sub, err := w.contract.WatchRoleRemoved(
					opts,
					recvChan,
					[]common.Address{},
					[]common.Address{},
				)
				if err != nil {
					if sub != nil {
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("Retrying start WatchRoleRemoved: %v", err)
					continue
				}
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchRoleRemoved: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchRoleRemoved for contract %v", w.contractAddress.Hex())
		for {
			select {
			case event := <-recvChan:
				civilEvent, err := model.NewCivilEvent("RoleRemoved", w.contractAddress, event)
				if err != nil {
					log.Errorf("Error creating new civil event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- *civilEvent:
				case err := <-sub.Err():
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting WatchRoleRemoved, fatal (a): %v", err)
						return err
					}
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error restarting WatchRoleRemoved, fatal (b): %v", err)
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}
