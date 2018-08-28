package qcheckpoint

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
)

var (
	blockMeter = metrics.NewRegisteredMeter("quorum/block", nil)
	txCreatedMeter = metrics.NewRegisteredMeter("quorum/tx-created", nil)
	txAcceptedMeter = metrics.NewRegisteredMeter("quorum/tx-accepted", nil)
)

func Create(checkpointName string, logValues ...interface{}) {
	log.EmitCheckpoint(checkpointName, logValues...)
	record(checkpointName)
}

func record(metricName string) {
	switch metricName {
	case log.TxCreated:
		txCreatedMeter.Mark(1)
	case log.TxAccepted:
		txAcceptedMeter.Mark(1)
	case log.BlockCreated:
		blockMeter.Mark(1)
	}
}