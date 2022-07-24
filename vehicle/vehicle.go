package vehicle

import (
	"fmt"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

// if route doesn't recieve DPS Letters or Flats, set the value to zero
type VehicleData struct {
	Miles float64
	POV   bool
}

func (vd *VehicleData) Populate(dd ds.DailyData, fd ds.FixedData) {
	vd.Miles = fd.Miles
	vd.POV = fd.DrivePOV
}

func (vd *VehicleData) Report() string {
	out := "\n************ Vehicle Time *************\n"
	out += fmt.Sprintf("Gov Vehicle Time: %4.2f\n", vd.GetTime())
	return out
}

func (vd *VehicleData) GetTime() float64 {
	//only applies for government vehicles
	if vd.POV {
		return 0.0
	}

	time := 0.0

	time += standards.INSPECT_GOV_VEHICLE_TIME

	refuels := (vd.Miles) / standards.FUEL_VEHICLE_MILES

	// for one day
	time += refuels * standards.FUEL_VEHICLE_TIME

	return time
}
