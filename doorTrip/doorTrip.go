package doorTrip

import (
	"fmt"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

// Time for taking things to the door.
// check walking and driveway sections for their respective times
type DoorData struct {

	//total number of trips to the door
	Trips int

	//parcels to the door
	Parcels int
	//accountables to the door
	Accountables int

	//misc to the door, w/o anything else to scan (assume nothing scanned on trip)
	Misc int

	//trips to the door to pickup packages
	PickupEvents int
}

func (d *DoorData) Populate(dd ds.DailyData, fd ds.FixedData) {
	d.Trips = dd.DoorTrips
	d.Parcels = dd.DoorParcels

	d.Accountables = dd.Certified
	d.Accountables += dd.Registered
	d.Accountables += dd.SignatureExpress
	d.Accountables += dd.CODs
	d.Accountables += dd.PostageDueCustom

	d.Misc = dd.DoorMisc
	d.PickupEvents = dd.DoorPickupEvents
}

func (dd *DoorData) Report() string {
	out := "\n*********** Door Trip Time **********\n"
	out += fmt.Sprintf("Total Door Trips: %d\n", dd.Trips)
	out += fmt.Sprintf("Multiple Item Deliveries: %d\n", dd.multiItems())
	out += fmt.Sprintf("Misc Time: %4.2f\n", dd.doorMiscTime())
	out += fmt.Sprintf("\nDoor Trips Total Time: %4.2f\n", dd.GetTime())
	return out
}

//Time for delivering misc to the door
func (dd *DoorData) doorMiscTime() float64 {
	return float64(dd.Misc) * standards.DELIVER_MISC_TO_DOOR_TIME
}

//Find trips where delivering more than 1 item
func (dd *DoorData) multiItems() int {
	packageAndAccountableTrips := dd.Trips - dd.Misc - dd.PickupEvents
	return (dd.Parcels + dd.Accountables) - packageAndAccountableTrips
}

func (dd *DoorData) GetTime() float64 {

	time := 0.0

	time += dd.doorMiscTime()

	//number of parcels that aren't the only ones at the door.
	multiItems := dd.multiItems()
	// prep time for trip to door
	time += float64(dd.Trips) * standards.PRELIM_DOOR_TRIP_TIME
	//more time for first item to door
	time += float64(dd.Trips-dd.Misc-dd.PickupEvents) * standards.GATHER_DOOR_FIXED_TIME
	//less time for additional items to door
	time += float64(multiItems) * standards.GATHER_DOOR_VAR_TIME

	return time
}
