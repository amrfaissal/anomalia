package anomalia

type NormalDistribution struct {
	EpsilonThreshold float64
}

func (nd *NormalDistribution) Run(timeSeries *TimeSeries) *ScoreList {
	scoreList, _ := nd.computeScores(timeSeries)
	return scoreList
}

func (nd *NormalDistribution) computeScores(timeSeries *TimeSeries) (*ScoreList, error) {
	mean := timeSeries.Average()
	std := timeSeries.Stdev()

	scores := mapSlice(timeSeries.Values, func(value float64) float64 {
		score := Pdf(mean, std)(value)
		if score < nd.EpsilonThreshold {
			return score
		} else {
			return 0.0
		}
	})

	scoreList := &ScoreList{
		Timestamps: timeSeries.Timestamps,
		Scores:     scores,
	}
	return scoreList, nil
}
