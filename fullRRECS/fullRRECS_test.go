package fullRRECS

import (
	"math"
	"testing"

	"rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

func sampleDailyData() dataSets.DailyData {
	// data from 21 Jul
	dd := dataSets.DailyData{

		//letters
		RawLetters:       27,
		DPSLetters:       350,
		WSSLetters:       0,
		BoxholderLetters: 0,

		//flats
		RawFlats:        29,
		PreBundledFlats: 0,
		DPSFlats:        0,
		WSSFlats:        0,
		BoxholderFlats:  0,

		//door trips
		DoorTrips: 15,

		//driveways
		LessSwimPool:     2,
		SwimPool:         4,
		FootballField:    2,
		TwoFootballField: 1,
		QuarterMile:      1,
		HalfMile:         1,

		//parcels
		MailBoxParcels: 17,
		LockerParcels:  0,
		DoorParcels:    18,

		RegCurbBoxesSkipped: 15,

		StampSale:  0,
		RuralReach: 1,

		OneStepScan: 20,

		ExtraSteps: 744,

		Loading:    23,
		EndOfShift: 29,
	}
	return dd
}

func sampleFixedData() dataSets.FixedData {
	fd := dataSets.FixedData{
		RegCurbBox:       238,
		DrivePOV:         true,
		DismountOther:    2,
		LowVolumePouches: 2,
		Miles:            92,
		DriveTime:        192.1,
		MiscTime:         10,
		FixedFeet:        190,
		FeetPerStep:      2.5,
		Bundle:           standards.TWO_BUNDLE,
	}
	return fd
}

func TestFull(t *testing.T) {
	dd := sampleDailyData()
	fd := sampleFixedData()

	got := GetTime(fd, dd)
	want := 392.1

	if math.Abs(got-want) > 0.1 {
		t.Errorf("%.4f New Time does not match hand calc %.4f", got, want)
	}

}
