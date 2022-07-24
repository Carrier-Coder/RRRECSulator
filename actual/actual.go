package actual

import (
	"fmt"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

// ActualData contains all the info for calculating
type ActualData struct {
	SafetyServiceTalks float64
	Loading            float64
	EndOfShift         float64
	Deviation          float64
	Fudge              float64
}

func (ad *ActualData) Populate(dd ds.DailyData, fd ds.FixedData) {
	ad.SafetyServiceTalks = dd.SafetyServiceTalks
	ad.Loading = dd.Loading
	ad.EndOfShift = dd.EndOfShift
	ad.Deviation = dd.Deviation
	ad.Fudge = dd.FudgeTime
}

func (ad *ActualData) GetSafetyServiceTalkTime() float64 {
	return ad.SafetyServiceTalks * (1 + standards.PERSONAL_TIME_PERCENT)
}
func (ad *ActualData) GetLoadingTime() float64 {
	return ad.Loading * (1 + standards.PERSONAL_TIME_PERCENT)
}
func (ad *ActualData) GetDeviationTime() float64 {
	return ad.Deviation * (1 + standards.PERSONAL_TIME_PERCENT)
}

//Guide pg 14 Personal Time doesn't apply to end of shift
func (ad *ActualData) GetEndOfShiftTime() float64 {
	return ad.EndOfShift
}

func (ad *ActualData) Report() string {
	out := "\n**************** Actual Time ****************\n"
	out += fmt.Sprintf("Service Talks: %4.2f\n", ad.GetSafetyServiceTalkTime())
	out += fmt.Sprintf("Loading: %4.2f\n", ad.GetLoadingTime())
	out += fmt.Sprintf("Deviation: %4.2f\n", ad.GetDeviationTime())
	out += fmt.Sprintf("End Of Shift: %4.2f\n", ad.GetEndOfShiftTime())
	out += fmt.Sprintf("Fudge Time Added: %4.2f\n", ad.Fudge)
	out += fmt.Sprintf("Actual Time Total: %4.2f\n", ad.GetTime())
	return out
}

func (ad *ActualData) GetTime() float64 {

	time := 0.0
	time += ad.GetSafetyServiceTalkTime()
	time += ad.GetLoadingTime()
	time += ad.GetDeviationTime()
	time += ad.GetEndOfShiftTime()
	time += ad.Fudge
	return time
}
