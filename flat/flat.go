package flat

import (
	"fmt"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

// FlatData contains all the info for calculating flat times
type FlatData struct {
	RandomFlats       int
	CarrierRouteFlats int
	WSSFlats          int

	DPSFlats int

	// Do you drive a personally owned vehicle or left hand drive gov machine?
	DrivePOV   bool
	DriveGOVLH bool
}

func (f *FlatData) Populate(dd ds.DailyData, fd ds.FixedData) {
	f.RandomFlats = dd.RawFlats
	f.CarrierRouteFlats = dd.PreBundledFlats
	f.WSSFlats = dd.WSSFlats
	f.DPSFlats = dd.DPSFlats
	f.DrivePOV = fd.DrivePOV
	f.DriveGOVLH = fd.DriveGOVLH
}

func (fd *FlatData) Report() string {
	out := "\n********** Flat Time ***********\n"
	out += fmt.Sprintf("Case DPS Flats? %t\n", fd.caseDPS())
	out += fmt.Sprintf("Cased Flats: %d\n", fd.CasedFlats())
	out += fmt.Sprintf("\nTotal Flats Time: %4.2f\n", fd.GetTime())
	return out
}

func (fd *FlatData) caseDPS() bool {
	return fd.DrivePOV || fd.DriveGOVLH
}

//CaseDPSFlats checks if a route is paid to case their DPSflats or not.
func CaseDPSFlats(dd ds.DailyData, fd ds.FixedData) bool {
	return fd.DrivePOV || fd.DriveGOVLH
}

func (fd *FlatData) GetTime() float64 {

	// if you drive a POV or a LHD government vehicle, paid to case fss

	time := 0.0

	time += float64(fd.RandomFlats) * standards.RAW_FLATS_TIME
	time += float64(fd.CarrierRouteFlats) * standards.CR_FLATS_TIME
	time += float64(fd.WSSFlats) * standards.WSS_FLATS_TIME
	if fd.caseDPS() {
		time += float64(fd.DPSFlats) * standards.CASE_DPS_FLATS_TIME
	}

	return time
}

// The total number of flats recieved for the route
func (fd *FlatData) TotalFlats() int {
	return fd.RandomFlats + fd.CarrierRouteFlats + fd.WSSFlats + fd.DPSFlats
}

// The total number of flats paid to case
func (fd *FlatData) CasedFlats() int {
	if fd.caseDPS() {
		return fd.TotalFlats()
	}
	return fd.TotalFlats() - fd.DPSFlats
}

// Time for verifying address, makes more sense here
//func (fd *FlatData) GetVerifyTime() float64 {
//return float64(fd.TotalFlats()) * standards.VERIFY_FLAT
//}
