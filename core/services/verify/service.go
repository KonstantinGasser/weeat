package verify

import (
	"context"

	"github.com/KonstantinGasser/weeat/core/domain/verification"
	"github.com/sirupsen/logrus"
)

type Service struct {
	repo      VerifyRepo
	stream    chan verification.FoodEvent
	terminate chan struct{}
}

func New(repo VerifyRepo) *Service {
	return &Service{
		repo:      repo,
		stream:    make(chan verification.FoodEvent),
		terminate: make(chan struct{}),
	}
}

// Trigger pushes the an event to the processor stream
//
// an event can be anything which needs to be verified
func (svc Service) Trigger(item verification.FoodEvent) {
	svc.stream <- item
}

// Start spins up the listener, listening for
// new events to be verified
//
// the verification process:
// 1) a new food item which gets created initially will be caught by the stream
// 2) the food item gets inserted in the database with the verification status :INIT: and default values for ack etc
// 3) the food item will be queued to be serves through a SSE connection
// 4)
//
func (svc Service) Start() {

	for {
		select {
		case event := <-svc.stream:
			// persist
			if err := svc.repo.InitVerifyItem(context.Background(), verification.ToDAO(event)); err != nil {
				logrus.Errorf("[verify.Process] could not store init event of item: %v\n", err)

			}

		case <-svc.terminate:
			return
		}
	}
}

func (svc Service) 
