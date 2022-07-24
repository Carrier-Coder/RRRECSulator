package pickup

import (
	"fmt"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

// data for pickup events.  To the door part handled in door trips
type PickupData struct {
	EventForms int
	Events     int // can have events w/o form, eg not in mailbox

	PackageNum int
}

func (pd *PickupData) Populate(dd ds.DailyData, fd ds.FixedData) {
	pd.EventForms = dd.EventForms
	pd.Events = dd.Events
	pd.PackageNum = dd.PackageNum

}

func (pd *PickupData) Report() string {
	out := "\n************* Package Pickup Time **************\n"
	out += fmt.Sprintf("Manifests: %d\n", pd.EventForms)
	out += fmt.Sprintf("Events: %d\n", pd.Events)
	out += fmt.Sprintf("Packages: %d\n", pd.PackageNum)
	out += fmt.Sprintf("\nPickup Total Time: %4.2f\n", pd.GetTime())
	return out
}

func (pd *PickupData) GetTime() float64 {
	time := 0.0

	//process pickup form in office
	time += float64(pd.EventForms) * standards.PICKUP_PROCESS_FORM_TIME

	//each pickup on the route
	time += float64(pd.Events) * standards.PICKUP_EVENT_TIME

	//each parcel picked up
	time += float64(pd.PackageNum) * standards.PICKUP_ITEM_TIME

	//Full scan on each event (possible to have event with no scans, but is small potatoes anyway)
	time += float64(pd.Events) * standards.FULL_SCAN_TIME

	//partial scan on all the others
	if pd.PackageNum-pd.Events > 0 {
		time += float64(pd.PackageNum-pd.Events) * standards.PARTIAL_SCAN_TIME
	}

	return time
}
