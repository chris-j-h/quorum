package quorumcheckpoint

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
)

const (
	TxCreated          = "TX-CREATED"
	TxAccepted         = "TX-ACCEPTED"
	BecameMinter       = "BECAME-MINTER"
	BecameVerifier     = "BECAME-VERIFIER"
	RaftBlockCreated   = "RAFT-BLOCK-CREATED"
    BlockInserted	   = "BLOCK-INSERTED"
	BlockVotingStarted = "BLOCK-VOTING-STARTED"
)

var (
	blockCreatedMeter = metrics.NewRegisteredMeter("quorum/raft-block-created", nil)
	blockInsertedMeter = metrics.NewRegisteredMeter("quorum/block-inserted", nil)
	txCreatedMeter = metrics.NewRegisteredMeter("quorum/tx-created", nil)
	txAcceptedMeter = metrics.NewRegisteredMeter("quorum/tx-accepted", nil)

 	DoEmitCheckpoints = false
)

func Create(checkpointName string, logValues ...interface{}) {
	emitCheckpoint(checkpointName, logValues...)
	updateMetric(checkpointName)
}

func emitCheckpoint(checkpointName string, logValues ...interface{}) {
	args := []interface{}{"name", checkpointName}
	args = append(args, logValues...)
	if DoEmitCheckpoints {
		log.Info("QUORUM-CHECKPOINT", args...)
	}
}

func updateMetric(metricName string) {
	switch metricName {
	case TxCreated:
		txCreatedMeter.Mark(1)
	case TxAccepted:
		txAcceptedMeter.Mark(1)
	case RaftBlockCreated:
		blockCreatedMeter.Mark(1)
	case BlockInserted:
		blockInsertedMeter.Mark(1)
	}

}