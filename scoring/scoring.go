package scoring

import "trainstations_domain/trains"

func Score(trainsToScore []trains.TrainInfo) []trains.TrainInfo {
	var scoredTrains []trains.TrainInfo

	for c, train := range trainsToScore {
		train.Points = c
		scoredTrains = append(scoredTrains, train)
	}

	return scoredTrains
}
