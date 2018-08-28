package qmetrics

import (
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/log"
)

var (
	blockMeter = metrics.NewRegisteredMeter("quorum/block", nil)
	txCreatedMeter = metrics.NewRegisteredMeter("quorum/tx-created", nil)
	txAcceptedMeter = metrics.NewRegisteredMeter("quorum/tx-accepted", nil)
)

func Emit(metricName string) {
	switch metricName {
	case log.TxCreated:
		txCreatedMeter.Mark(1)
	case log.TxAccepted:
		txAcceptedMeter.Mark(1)
	case log.BlockCreated:
		blockMeter.Mark(1)
	}
}