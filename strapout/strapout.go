package strapout

import (
	"fmt"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/flat"
	"rrecsulator.com/letter"
	"rrecsulator.com/standards"
)

type StrapoutData struct {

	//cased letters
	RawLetters       int
	WSSLetters       int
	BoxholderLetters int
	CasedDPSLetters  int

	//cased flats
	RawFlats        int
	PreBundledFlats int
	WSSFlats        int
	BoxholderFlats  int
	CasedDPSFlats   int

	RegCurbBox  int
	RegSdwkBox  int
	RegOtherBox int

	CBUBoxes  int
	CENTBoxes int
	NPUBoxes  int
	DETBoxes  int

	Bundle standards.Bundling
}

func (sd *StrapoutData) Populate(dd ds.DailyData, fd ds.FixedData) {
	//cased letters
	sd.RawLetters = dd.RawLetters
	sd.WSSLetters = dd.WSSLetters
	sd.BoxholderLetters = dd.BoxholderLetters

	sd.CasedDPSLetters = dd.CasedDPSLetters
	if letter.CaseDPSLetters(dd, fd) {
		sd.CasedDPSLetters += dd.DPSLetters
	}

	//cased flats
	sd.RawFlats = dd.RawFlats
	sd.PreBundledFlats = dd.PreBundledFlats
	sd.WSSFlats = dd.WSSFlats
	sd.BoxholderFlats = dd.BoxholderFlats
	sd.CasedDPSFlats = dd.CasedDPSFlats

	if flat.CaseDPSFlats(dd, fd) {
		sd.CasedDPSFlats += dd.DPSFlats
	}

	// all address (case slots)
	sd.RegCurbBox = fd.RegCurbBox
	sd.RegSdwkBox = fd.RegSdwkBox
	sd.RegOtherBox = fd.RegOtherBox

	sd.CBUBoxes = fd.CBUBoxes
	sd.CENTBoxes = fd.CENTBoxes
	sd.NPUBoxes = fd.NPUBoxes
	sd.DETBoxes = fd.DETBoxes

	sd.Bundle = fd.Bundle
}

//Total slots in the case
func (sd *StrapoutData) TotalSlots() int {
	out := 0
	out += sd.RegCurbBox
	out += sd.RegSdwkBox
	out += sd.RegOtherBox
	out += sd.CBUBoxes
	out += sd.CENTBoxes
	out += sd.NPUBoxes
	out += sd.DETBoxes
	return out
}

func (sd *StrapoutData) TotalPieces() int {
	out := 0
	out += sd.RawLetters
	out += sd.WSSLetters
	out += sd.BoxholderLetters
	out += sd.CasedDPSLetters
	out += sd.RawFlats
	out += sd.PreBundledFlats
	out += sd.WSSFlats
	out += sd.BoxholderFlats
	out += sd.CasedDPSFlats
	return out
}

func (sd *StrapoutData) Report() string {
	out := "\n************* Strapout Time **************\n"
	out += fmt.Sprintf("Total Cased Pieces %d\n", sd.TotalPieces())
	out += fmt.Sprintf("Total Case Slots %d\n\n", sd.TotalSlots())

	out += fmt.Sprintf("Reach Time %4.2f\n", sd.GetReachTime())
	out += fmt.Sprintf("Handful Time %4.2f\n", sd.GetHandfulTime())
	out += fmt.Sprintf("RubberBand Time %4.2f\n", sd.GetRubberBandTime())
	out += fmt.Sprintf("Tray Time: %4.2f\n", sd.GetTrayTime())
	out += fmt.Sprintf("\nStrapout Total Time: %4.2f\n", sd.GetTime())

	return out
}

func (sd *StrapoutData) GetReachTime() float64 {
	return float64(sd.numReaches()) * standards.REACH_TIME
}

func (sd *StrapoutData) numReaches() int {
	if sd.Bundle == standards.ONE_BUNDLE {
		return sd.TotalSlots()
	}
	return sd.TotalSlots() / 2
}

func (sd *StrapoutData) GetHandfulTime() float64 {
	return float64(sd.numHandfuls()) * standards.HANDFUL_TIME
}

func (sd *StrapoutData) numHandfuls() int {
	temp := sd.TotalPieces()
	hf := 0
	for temp > 0 {
		hf++
		temp -= standards.HANDFUL_PIECES
	}
	return hf
}

func (sd *StrapoutData) GetRubberBandTime() float64 {
	return float64(sd.numRubberBands()) * standards.RUBBER_BAND_TIME
}

func (sd *StrapoutData) numRubberBands() int {
	temp := sd.TotalPieces()
	rb := 0
	for temp > 0 {
		rb++
		temp -= standards.RUBBER_BAND_PIECES
	}
	return rb
}

func (sd *StrapoutData) GetTrayTime() float64 {
	return float64(sd.numTrays()) * standards.TRAY_TIME
}

func (sd *StrapoutData) numTrays() int {
	temp := sd.TotalPieces()
	nt := 0
	for temp > 0 {
		nt++
		temp -= standards.TRAY_PIECES
	}
	return nt
}

func (sd *StrapoutData) GetTime() float64 {
	time := sd.GetTrayTime()
	time += sd.GetHandfulTime()
	time += sd.GetReachTime()
	time += sd.GetRubberBandTime()
	return time
}
