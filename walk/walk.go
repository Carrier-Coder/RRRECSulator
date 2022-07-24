package walk

import (
	"fmt"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

// Time for taking things to the door.
// check walking and driveway for those specific elements
type WalkData struct {
	Steps         int
	FixedSteps    int
	DismountSteps int
	FixedFeet     int
	FeetPerStep   float64
}

func (ww *WalkData) Populate(dd ds.DailyData, fd ds.FixedData) {
	ww.Steps = dd.ExtraSteps + fd.FixedSteps + fd.DismountSteps
	ww.FixedFeet = fd.FixedFeet
	ww.FeetPerStep = fd.FeetPerStep
}

func (ww *WalkData) Report() string {
	out := "\n************ Walking Time **************\n"
	out += fmt.Sprintf("Package Deliveries: %d steps\n", ww.Steps)
	out += fmt.Sprintf("Fixed Walking: %d steps\n", ww.FixedSteps)
	out += fmt.Sprintf("Dismount Walking: %d steps\n", ww.DismountSteps)
	out += fmt.Sprintf("Stride length: %4.2f feet per step\n", ww.FeetPerStep)
	out += fmt.Sprintf("Measured Fixed feet: %d feet\n", ww.FixedFeet)
	out += fmt.Sprintf("\nWalking Total Time %4.2f\n", ww.GetTime())
	return out
}

func (ww *WalkData) GetTime() float64 {
	out := 0.0
	out += float64(ww.Steps+ww.FixedSteps) * ww.FeetPerStep
	out += float64(ww.FixedFeet)
	return out * standards.WALKING_STANDARD
}
