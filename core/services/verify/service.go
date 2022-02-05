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

func (svc Service) Trigger(item verification.FoodEvent) {
	svc.stream <- item
}

// Start spins up the listener, listening for
// new events to be verified
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
