package reload

import (
	"fmt"
	"log"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

type ReloadData struct {

	//cased letters
	RandomLetters    int
	WSSLetters       int
	BoxholderLetters int

	//cased flats
	RandomFlats    int
	CRFlats        int
	WSSFlats       int
	BoxholderFlats int

	//DPS
	DPSLetters int
	DPSFlats   int

	//parcels
	MailBoxParcels int
	LockerParcels  int
	DoorParcels    int

	Bundle standards.Bundling
}

func (rd *ReloadData) Populate(dd ds.DailyData, fd ds.FixedData) {

	//cased letters
	rd.RandomLetters = dd.RawLetters
	rd.WSSLetters = dd.WSSLetters
	rd.BoxholderLetters = dd.BoxholderLetters

	//cased flats
	rd.RandomFlats = dd.RawFlats
	rd.CRFlats = dd.PreBundledFlats
	rd.WSSFlats = dd.WSSFlats
	rd.BoxholderFlats = dd.BoxholderFlats

	//DPS
	rd.DPSLetters = dd.DPSLetters
	rd.DPSFlats = dd.DPSFlats

	//parcels
	rd.MailBoxParcels = dd.MailBoxParcels
	rd.LockerParcels = dd.LockerParcels
	rd.DoorParcels = dd.DoorParcels

	rd.Bundle = fd.Bundle
}

func (rd *ReloadData) Report() string {
	out := "\n************* Reload Time **************\n"
	out += fmt.Sprintf("Mail Tray Time: %4.2f\n", rd.GetMailTrayTime())
	out += fmt.Sprintf("Medium Parcel Time: %4.2f\n", rd.GetMediumParcelTime())
	out += fmt.Sprintf("Large Parcel Time: %4.2f\n", rd.GetLargeParcelTime())
	out += fmt.Sprintf("\nReload Total Time: %4.2f\n", rd.GetTime())

	return out
}

func (rd *ReloadData) GetMailTrayTime() float64 {

	CasedLetters := rd.RandomLetters + rd.WSSLetters + rd.BoxholderLetters
	// how many cased Letter trays
	CasedLettersTray := 0
	for CasedLetters > 0 {
		CasedLettersTray++
		CasedLetters -= standards.LETTERS_PER_TRAY
	}

	DPSLetters := rd.DPSLetters
	// how many DPS Letter trays
	DPSLettersTray := 0
	for DPSLetters > 0 {
		DPSLettersTray++
		DPSLetters -= standards.LETTERS_PER_TRAY
	}

	CasedFlats := rd.RandomFlats + rd.CRFlats + rd.WSSFlats + rd.BoxholderFlats
	CasedFlatsTray := 0
	for CasedFlats > 0 {
		CasedFlatsTray++
		CasedFlats -= standards.FLATS_PER_TRAY
	}

	DPSFlats := rd.DPSFlats
	DPSFlatsTray := 0
	for DPSFlats > 0 {
		DPSFlatsTray++
		DPSFlats -= standards.FLATS_PER_TRAY
	}

	notLargeParcels := rd.LockerParcels + rd.MailBoxParcels

	smallParcels := float64(notLargeParcels) * standards.SMALL_PARCEL_PERCENT
	SmallParcelsTray := 0
	for smallParcels > 0 {
		SmallParcelsTray++
		smallParcels -= float64(standards.SMALL_PARCELS_PER_TRAY)
	}

	totalTrays := CasedLettersTray
	totalTrays += DPSLettersTray
	totalTrays += CasedFlatsTray
	totalTrays += DPSFlatsTray
	totalTrays += SmallParcelsTray

	//assumes some trays preloaded in the front
	totalTrays = totalTrays - standards.FRONT_LOADED_MAIL_TRAYS

	//make sure we don't have negative number of trays!
	if totalTrays <= 0 {
		return 0.0
	}

	if totalTrays >= len(standards.MAIL_TRAY_TABLE) {
		log.Printf("More trays than in the mail tray table!!!\n")
		totalTrays = len(standards.MAIL_TRAY_TABLE) - 1
	}

	time := standards.MAIL_TRAY_TABLE[totalTrays][rd.Bundle]

	return time
}

func (rd *ReloadData) GetMediumParcelTime() float64 {
	notLargeParcels := rd.LockerParcels + rd.MailBoxParcels
	mediumParcels := float64(notLargeParcels) * standards.MEDIUM_PARCEL_PERCENT

	mediumParcelsTray := 0
	for mediumParcels > 0 {
		mediumParcelsTray++
		mediumParcels -= float64(standards.MEDIUM_PARCELS_PER_TRAY)
	}

	//assumes some trays are reachable from the front
	mediumParcelsTray -= standards.FRONT_LOADED_MED_PARCEL_TRAY

	if mediumParcelsTray <= 0 {
		return 0.0
	}

	// if we are outside of the lookup table
	if mediumParcelsTray >= len(standards.MEDIUM_PARCEL_TABLE) {
		time := standards.MEDIUM_PARCEL_TABLE[len(standards.MEDIUM_PARCEL_TABLE)-1]
		return time

	}

	time := standards.MEDIUM_PARCEL_TABLE[mediumParcelsTray]
	return time
}

func (rd *ReloadData) GetLargeParcelTime() float64 {

	//some large parcels reachable, don't get reload time
	reloadLargeParcels := rd.DoorParcels - standards.REACHABLE_LARGE_PARCEL

	if reloadLargeParcels <= 0 {
		return 0.0
	}

	return float64(reloadLargeParcels) * standards.RELOAD_LARGE_PARCEL_TIME
}

func (rd *ReloadData) GetTime() float64 {
	time := rd.GetMailTrayTime()
	time += rd.GetMediumParcelTime()
	time += rd.GetLargeParcelTime()
	return time
}
