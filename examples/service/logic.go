package service

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -package=mocks -destination=../mocks/storage.go github.com/martinsirbe/go-metrics-poc/examples/service Storage
type Storage interface {
	Insert(o string) error
	Delete(o string) error
}

type LogicLayer struct {
	client Storage
}

func NewLogicLayer(client Storage) *LogicLayer {
	return &LogicLayer{client}
}

func (ll *LogicLayer) Run() error {
	errs := make(chan error, 1)

	go func() {
		for {
			if err := ll.client.Insert(fmt.Sprintf("%v", uuid.Must(uuid.NewV4()))); err != nil {
				logrus.WithError(err).Error("stopped inserting stuff")
				errs <- err
			}
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			if err := ll.client.Delete(fmt.Sprintf("%v", uuid.Must(uuid.NewV4()))); err != nil {
				logrus.WithError(err).Error("stopped deleting stuff")
				errs <- err
			}
			time.Sleep(time.Second * 5)
		}
	}()

	if err, open := <-errs; open {
		return err
	}

	return nil
}
