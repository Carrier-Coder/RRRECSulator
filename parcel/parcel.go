package parcel

import (
	"fmt"
	"math"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

// ParcelData contains all the info for calculating parcel times
type ParcelData struct {
	MailBoxParcels int
	LockerParcels  int
	DoorParcels    int
}

func (pd *ParcelData) Populate(dd ds.DailyData, fd ds.FixedData) {
	pd.MailBoxParcels = dd.MailBoxParcels
	pd.LockerParcels = dd.LockerParcels
	pd.DoorParcels = dd.DoorParcels
}

func (pd *ParcelData) Report() string {
	out := "\n********* Parcel Time *******************\n"
	out += "Note: door parcel delivery time included in door trips\n"
	out += fmt.Sprintf("Total Parcels: %d\n", pd.TotalParcels())
	out += fmt.Sprintf("Large Parcels: %d\n", pd.LargeParcels())
	out += fmt.Sprintf("Medium Parcels: %d\n", pd.MediumParcels())
	out += fmt.Sprintf("Small Parcels: %d\n", pd.SmallParcels())
	out += fmt.Sprintf("Carts: %d\n", pd.howManyCarts())
	out += fmt.Sprintf("Cart Time: %4.2f\n\n", pd.GetCartTime())
	out += fmt.Sprintf("Parcel Total Time: %4.2f\n", pd.GetTime())
	return out
}

// TotalParcels returns count of every piece with scan code
func (pd ParcelData) TotalParcels() int {
	return pd.MailBoxParcels + pd.LockerParcels + pd.DoorParcels
}

// LargeParcels returns large (gravy train) parcels
func (pd ParcelData) LargeParcels() int {
	return pd.DoorParcels
}

//ASSUMPTION: Medium round up, small round down
// Medium and small are based off of fractional percentage
// round the medium up and the small down, since there are more likely to
// be medium than small.  Not sure what the actual system will do...

// MediumParcels are smaller than a breadbox
func (pd ParcelData) MediumParcels() int {
	var temp float64 = standards.MEDIUM_PARCEL_PERCENT
	temp = temp * float64(pd.TotalParcels()-pd.LargeParcels())
	temp = math.Ceil(temp)
	return int(temp)
}

// SmallParacels returns count of caseable parcels
func (pd ParcelData) SmallParcels() int {
	var temp float64 = standards.SMALL_PARCEL_PERCENT
	return int(temp * float64((pd.TotalParcels() - pd.LargeParcels())))
}

func (pd *ParcelData) howManyCarts() int {
	totalParcels := pd.TotalParcels()
	// cart trips
	carts := 0
	for totalParcels > 0 {
		carts++
		totalParcels = totalParcels - standards.PARCEL_PER_CART
	}
	return carts
}

func (pd *ParcelData) GetCartTime() float64 {
	cartTime := standards.GATHER_1ST_CART_TIME
	carts := pd.howManyCarts() - 1
	if carts < 0 {
		carts = 0
	}
	cartTime += float64(carts) * standards.GATHER_ADD_CART_TIME
	return cartTime
}

func (pd *ParcelData) GetTime() float64 {

	// cart trips
	cartTime := pd.GetCartTime()

	// large parcels
	// sort time
	// door delivery time
	// rest is door trip, count separately
	largeParcelTime := 0.0
	largeParcelTime = float64(pd.DoorParcels) *
		float64(standards.ORGANIZE_LARGE_TIME)

	largeParcelTime += float64(pd.DoorParcels) *
		float64(standards.DELIVER_DOOR_TIME)

	// small & medium parcels
	smallMediumParcelTime := 0.0

	// organize time
	smParcels := pd.MailBoxParcels + pd.LockerParcels
	smallMediumParcelTime = float64(smParcels) *
		float64(standards.ORGANIZE_SMALL_MED_TIME)

	// deliver time (large parcels included in door trips)
	smallMediumParcelTime += float64(pd.MailBoxParcels) *
		float64(standards.DELIVER_MAILBOX_TIME)
	smallMediumParcelTime += float64(pd.LockerParcels) *
		float64(standards.DELIVER_LOCKER_TIME)

	return cartTime + largeParcelTime + smallMediumParcelTime
}
