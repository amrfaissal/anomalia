package anomalia

import (
	"math/rand"
	"testing"
	"time"
)

func TestRunWithBitmap(t *testing.T) {
	//
	// Generate the data set
	//
	datasetSize := 2000
	timestamps := make([]float64, datasetSize)
	for i := 0; i < datasetSize; i++ {
		timestamps[i] = float64(i) + 1
	}

	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	values := make([]float64, datasetSize)
	for i := 0; i < datasetSize; i++ {
		values[i] = generator.Float64() * 10
	}
	timeSeries := &TimeSeries{
		Timestamps: timestamps,
		Values:     values,
	}

	//
	// Run the bitmap algorithm
	//
	bitmap := NewBitmap()
	scoreList := bitmap.Run(timeSeries)
	if scoreList == nil {
		t.Fatalf("score list cannot be nil")
	}

	if len(scoreList.Scores) != len(timeSeries.Timestamps) {
		t.Fatalf("both time series and score list dimensions do not match")
	}

	//
	// Use Case: Not enough data points
	//
	timeSeries = &TimeSeries{
		Timestamps: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		Values:     []float64{1, 5, 52, 49, 49, 1.5, 48, 50, 53, 44},
	}
	scoreList = bitmap.Run(timeSeries)
	if scoreList != nil {
		t.Fatalf("score list must be nil (not enough data points)")
	}
}
