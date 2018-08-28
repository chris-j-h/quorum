package qmetrics

import (
	"github.com/ethereum/go-ethereum/metrics"
)

var blockMeter = metrics.NewRegisteredMeter("quorum/checkpoint/block", nil)

func UpdateCheckpointMetric() {
	blockMeter.Mark(1)
}