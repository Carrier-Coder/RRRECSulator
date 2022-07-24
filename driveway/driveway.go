package driveway

import (
	"fmt"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

// Time for taking things to the door.
// check walking and driveway for those specific elements
type DrivewayData struct {
	LessSwimPool     int
	SwimPool         int
	FootballField    int
	TwoFootballField int
	QuarterMile      int
	HalfMile         int
}

func (dw *DrivewayData) Populate(dd ds.DailyData, fd ds.FixedData) {

	dw.LessSwimPool = dd.LessSwimPool
	dw.SwimPool = dd.SwimPool
	dw.FootballField = dd.FootballField
	dw.TwoFootballField = dd.TwoFootballField
	dw.QuarterMile = dd.QuarterMile
	dw.HalfMile = dd.HalfMile
}

func (dw *DrivewayData) Report() string {
	out := "\n********** Driveway Time ***********\n"
	out += fmt.Sprintf("Driveway Time Total: %4.2f\n", dw.GetTime())
	return out
}

func (dw *DrivewayData) GetTime() float64 {

	time := 0.0
	time += float64(dw.LessSwimPool) * standards.LESS_THAN_A_POOL_TIME
	time += float64(dw.SwimPool) * standards.POOL_TIME
	time += float64(dw.FootballField) * standards.FOOTBALL_FIELD_TIME
	time += float64(dw.TwoFootballField) * standards.TWO_FOOTBALL_FIELD_TIME
	time += float64(dw.QuarterMile) * standards.QUARTER_MILE_TIME
	time += float64(dw.HalfMile) * standards.HALF_MILE_TIME
	return time
}
