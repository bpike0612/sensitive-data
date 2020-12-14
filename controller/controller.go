package controller

import (
	"github.com/pkg/errors"
	"math/rand"
	"sensitive-data/logger"
	"sensitive-data/protocol"
)

func Process(company *protocol.Company) error {
	if err := doStuff(); err != nil {

		// note the logging done here and company provided as an argument here
		logger.Error("failed to process company", company)

		return errors.Wrap(err, "failed to process")
	}

	return nil
}

func doStuff() error {
	if rand.Int63n(10) < 5 {
		return errors.New("random error")
	}

	return nil
}
