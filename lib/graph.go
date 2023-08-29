package lib


func makeGraph (maxDate int64, minDate int64, interval int64) []Datapoint {
	numIntervals := (maxDate - minDate) / interval
	intervals := make([]*Interval, numIntervals)
	for _,value := range feedback_store {
		if value.Timestamp > minDate && value.Timestamp < maxDate {
			intervNum := (value.Timestamp - minDate) / interval
			if intervals[intervNum] == nil {
				var fb []Feedback 
				interv := Interval {feedback: append(fb, value),}
				intervals[intervNum] = interv
			} else {
				intervals[intervNum].feedback = append(intervals[intervNum].feedback, value)
			}
		}
	}

	var datapoints []Datapoint

	for idx,i := range intervals {
		var avg int64
		avg = 0
		for _,feedback := range i.feedback {
			avg = avg + feedback.Rating
		}

		avg = avg / len(i.feedback)
		x := minDate + idx * interval
		datapoints = append(datapoints, Datapoint{x: x, y: avg})

	}

	return datapoints
}