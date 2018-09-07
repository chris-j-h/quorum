package qcheckpoint

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
)

const (
	BlockInserted = "BLOCK-INSERTED"
)

var (
	blockCreatedMeter = metrics.NewRegisteredMeter("quorum/block", nil)
	blockInsertedMeter = metrics.NewRegisteredMeter("quorum/block-inserted", nil)
	txCreatedMeter = metrics.NewRegisteredMeter("quorum/tx-created", nil)
	txAcceptedMeter = metrics.NewRegisteredMeter("quorum/tx-accepted", nil)
)

func Create(checkpointName string, logValues ...interface{}) {
	log.EmitCheckpoint(checkpointName, logValues...)
	record(checkpointName, logValues...)
}

func record(metricName string, logValues ...interface{}) {
	switch metricName {
	case log.TxCreated:
		txCreatedMeter.Mark(1)
	case log.TxAccepted:
		txAcceptedMeter.Mark(1)
	case log.BlockCreated:
		blockCreatedMeter.Mark(1)
	case BlockInserted:
		log.Info("BlockInserted Entry")
		//
		//blockInsertedMeter.Mark(1)
		//
		//log.Info("BlockInserted", "counter", big.NewInt(blockInsertedMeter.Count()),
		//	"block", logValues[1].(*big.Int),
		//	"comp", big.NewInt(blockInsertedMeter.Count()).Cmp(logValues[1].(*big.Int)))
		//
		//if big.NewInt(blockInsertedMeter.Count()).Cmp(logValues[1].(*big.Int)) == -1 {
		//	blockInsertedMeter.Mark(1)
		//
		//	log.Info("BlockInserted extra count", "counter", big.NewInt(blockInsertedMeter.Count()),
		//		"block", logValues[1].(*big.Int),
		//		"comp", big.NewInt(blockInsertedMeter.Count()).Cmp(logValues[1].(*big.Int)))
		//}

		//for ok := true; ok; ok = (big.NewInt(blockInsertedMeter.Count()).Cmp(logValues[1].(*big.Int)) == -1) {
		//	blockInsertedMeter.Mark(1)
		//	log.Info("BlockInserted", "counter", big.NewInt(blockInsertedMeter.Count()),
		//		"block", logValues[1].(*big.Int),
		//		"comp", big.NewInt(blockInsertedMeter.Count()).Cmp(logValues[1].(*big.Int)))
		//
		//}

		log.Info("BlockInserted Exit")
	}
}