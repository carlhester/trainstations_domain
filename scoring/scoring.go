package scoring

import "trainstations_domain/trains"

func Score(trainsToScore []trains.TrainInfo) []trains.TrainInfo {
	var scoredTrains []trains.TrainInfo

	scoredTrains = soonerBetterThanLater(trainsToScore)
	scoredTrains = withinXMinutesAddYPts(10, 10, scoredTrains)
	scoredTrains = withinXMinutesAddYPts(5, 20, scoredTrains)
	scoredTrains = withinXMinutesAddYPts(3, 30, scoredTrains)

	return scoredTrains
}

// The earlier train is worth more than a later one
func soonerBetterThanLater(targets []trains.TrainInfo) []trains.TrainInfo {
	for c, _ := range targets {
		targets[c].Points = len(targets) - c
	}

	return targets
}

// The closer the previous train, the more points this train is worth
func withinXMinutesAddYPts(x int, y int, targets []trains.TrainInfo) []trains.TrainInfo {
	for i, _ := range targets {
		if i > 0 {
			if (targets[i].Minutes - targets[i-1].Minutes) < x {
				targets[i].Points += y
			}
		}

	}

	return targets
}
