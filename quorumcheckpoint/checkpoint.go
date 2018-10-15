package quorumcheckpoint

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"math/big"
)

const (
	TxCreated          = "TX-CREATED"
	CanonTxAccepted    = "CANON-TX-ACCEPTED"
	RaftTxAccepted     = "RAFT-TX-ACCEPTED"
	BecameMinter       = "BECAME-MINTER"
	BecameVerifier     = "BECAME-VERIFIER"
	RaftBlockCreated   = "RAFT-BLOCK-CREATED"
    BlockInserted	   = "BLOCK-INSERTED"
	BlockVotingStarted = "BLOCK-VOTING-STARTED"
)

var (
	raftBlockCreatedMeter = metrics.NewRegisteredMeter("quorum/raft-block-created", nil)
	blockInsertedMeter = metrics.NewRegisteredMeter("quorum/block-inserted", nil)
	blockInsertedGauge = metrics.NewRegisteredGauge("quorum/block-inserted-gauge", nil)
	txCreatedMeter = metrics.NewRegisteredMeter("quorum/tx-created", nil)
	raftTxAcceptedMeter = metrics.NewRegisteredMeter("quorum/raft-tx-accepted", nil)
	canonTxAcceptedMeter = metrics.NewRegisteredMeter("quorum/canon-tx-accepted", nil)

	DoEmitCheckpoints = false
)

func Create(checkpointName string, logValues ...interface{}) {
	emitCheckpoint(checkpointName, logValues...)
	updateMetric(checkpointName, logValues...)
}

func emitCheckpoint(checkpointName string, logValues ...interface{}) {
	args := []interface{}{"name", checkpointName}
	args = append(args, logValues...)
	if DoEmitCheckpoints {
		log.Info("QUORUM-CHECKPOINT", args...)
	}
}

func updateMetric(metricName string, logValues ...interface{}) {
	switch metricName {
	case TxCreated:
		txCreatedMeter.Mark(1)
	case RaftTxAccepted:
		raftTxAcceptedMeter.Mark(1)
	case CanonTxAccepted:
		canonTxAcceptedMeter.Mark(1)
	case RaftBlockCreated:
		raftBlockCreatedMeter.Mark(1)
	case BlockInserted:
		blockInsertedMeter.Mark(1)

		for i, value := range logValues {
			if value == "number" {
				bigNum := logValues[i+1]
				blockNum := bigNum.(*big.Int).Int64()

				blockInsertedGauge.Update(blockNum)
				break
			}
		}
	}
}