package dismount

import (
	"fmt"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

// Dismount Data contains all info for calculating authorized dismounts
// Lump all the walk time together
type DismountData struct {
	Other    int
	Sidewalk int
	Cbu      int
	Cent     int
	Npu      int
	DetVpo   int

	AddTripOther    int
	AddTripSidewalk int
	AddTripCbu      int
	AddTripCent     int
	AddTripNpu      int
	AddTripDetVpo   int
}

func (dis *DismountData) Populate(dd ds.DailyData, fd ds.FixedData) {
	dis.Other = fd.DismountOther
	dis.Sidewalk = fd.DismountSidewalk
	dis.Cbu = fd.DismountCbu
	dis.Cent = fd.DismountCent
	dis.Npu = fd.DismountNpu
	dis.DetVpo = fd.DismountDetVpo

	dis.AddTripOther = dd.AddTripOther
	dis.AddTripSidewalk = dd.AddTripSidewalk
	dis.AddTripCbu = dd.AddTripCbu
	dis.AddTripCent = dd.AddTripCent
	dis.AddTripNpu = dd.AddTripNpu
	dis.AddTripDetVpo = dd.AddTripDetVpo
}

func (dd *DismountData) Report() string {
	out := "\n********** Authorized Dismount Time ********\n"
	out += fmt.Sprintf("Total Authorized Dismounts: %d\n", dd.Total())
	out += fmt.Sprintf("Total Additional Trips: %d\n", dd.TotalAddTrips())
	out += fmt.Sprintf("\nAuthorized Dismount Total Time: %4.2f\n", dd.GetTime())
	return out
}

func (dd *DismountData) Total() int {
	total := dd.Other + dd.Sidewalk + dd.Cbu + dd.Cent
	total += dd.Npu + dd.DetVpo
	return total
}

func (dd *DismountData) TotalAddTrips() int {
	total := dd.AddTripOther + dd.AddTripSidewalk + dd.AddTripCbu
	total += dd.AddTripCent + dd.AddTripNpu + dd.AddTripDetVpo
	return total
}

func (dd *DismountData) GetTime() float64 {
	time := 0.0

	time += float64(dd.Other) * standards.OTHER_DISMOUNT_PREP_TIME
	time += float64(dd.AddTripOther) * standards.OTHER_ADD_TRIP_TIME

	time += float64(dd.Sidewalk) * standards.SDWK_DISMOUNT_PREP_TIME
	time += float64(dd.AddTripSidewalk) * standards.SDWK_ADD_TRIP_TIME

	time += float64(dd.Cbu) * standards.CBU_DISMOUNT_PREP_TIME
	time += float64(dd.AddTripCbu) * standards.CBU_ADD_TRIP_TIME

	time += float64(dd.Cent) * standards.CENT_DISMOUNT_PREP_TIME
	time += float64(dd.AddTripCent) * standards.CENT_ADD_TRIP_TIME

	time += float64(dd.Npu) * standards.NPU_DISMOUNT_PREP_TIME
	time += float64(dd.AddTripNpu) * standards.NPU_ADD_TRIP_TIME

	time += float64(dd.DetVpo) * standards.VPO_DISMOUNT_PREP_TIME
	time += float64(dd.AddTripDetVpo) * standards.VPO_ADD_TRIP_TIME

	return time
}
