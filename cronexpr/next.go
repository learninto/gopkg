package cronexpr

import (
	"errors"
	"time"
)

// NextN returns a slice of `n` closest time instants immediately following
// `fromTime` which match the cron expression `expr`.
//
// The time instants in the returned slice are in chronological ascending order.
// The `time.Location` of the returned time instants is the same as that of
// `fromTime`.
//
// The start time
// The `time.Location` of the returned time instants is the same as that of
// `startDate`.
//
// The end time
// The `time.Location` of the returned time instants is the same as that of
// `endDate`.
//
// A slice with len between [0-`n`] is returned, that is, if not enough existing
// matching time instants exist, the number of returned entries will be less
// than `n`.
func (expr *Expression) Next(fromTime, startDate, endDate time.Time) (time.Time, error) {
	if fromTime.Unix() > endDate.Unix() {
		return time.Time{}, errors.New("无记录")
	}

	for true {
		if fromTime.Unix() > endDate.Unix() {
			break
		}

		if _, b := expr.Validate(fromTime, startDate); b == nil {
			return fromTime, nil
		}
		fromTime = fromTime.AddDate(0, 0, 1)
	}

	return time.Time{}, errors.New("无记录")
}

// NextN returns a slice of `n` closest time instants immediately following
// `fromTime` which match the cron expression `expr`.
//
// The time instants in the returned slice are in chronological ascending order.
// The `time.Location` of the returned time instants is the same as that of
// `fromTime`.
//
// The start time
// The `time.Location` of the returned time instants is the same as that of
// `startDate`.
//
// A slice with len between [0-`n`] is returned, that is, if not enough existing
// matching time instants exist, the number of returned entries will be less
// than `n`.
func (expr *Expression) NextN(fromTime, startDate time.Time, n int) (listCron []time.Time, e error) {
	for i := 0; i < n; {
		if _, b := expr.Validate(fromTime, startDate); b == nil {
			i++
			listCron = append(listCron, fromTime)
		}
		fromTime = fromTime.AddDate(0, 0, 1)
	}

	return
}
