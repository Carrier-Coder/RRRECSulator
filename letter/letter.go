package letter

import (
	"fmt"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

// LetterData contains all the info for calculating box times
type LetterData struct {
	RandomLetters int
	WSSLetters    int

	DPSLetters      int
	CasedDPSLetters int
}

func (ld *LetterData) Populate(dd ds.DailyData, fd ds.FixedData) {
	ld.RandomLetters = dd.RawLetters
	ld.WSSLetters = dd.WSSLetters
	ld.DPSLetters = dd.DPSLetters
	ld.CasedDPSLetters = dd.CasedDPSLetters
}

func (ld *LetterData) Report() string {
	out := "\n********** Letter Time ************\n"
	out += fmt.Sprintf("DPS as Raw? %t\n", ld.caseDPS())
	out += fmt.Sprintf("Cased Letters: %d\n", ld.CasedLetters())
	out += fmt.Sprintf("\nLetters Total Time: %4.2f\n", ld.GetTime())
	return out
}

func (ld *LetterData) caseDPS() bool {
	return caseDPSLetters(ld.DPSLetters)
}

//CaseDPSLetters checks if carrier is paid for casing their DPS.
func CaseDPSLetters(dd ds.DailyData, fd ds.FixedData) bool {
	return caseDPSLetters(dd.DPSLetters)
}

func caseDPSLetters(numLetters int) bool {
	return numLetters <= standards.DPS_THRESHOLD
}

func (ld *LetterData) GetTime() float64 {

	time := 0.0
	time += float64(ld.RandomLetters) * standards.RAW_LETTERS_NEW
	time += float64(ld.WSSLetters) * standards.WSS_LETTERS_NEW
	time += float64(ld.CasedDPSLetters) * standards.CASE_DPS_LETTERS

	//ASSUMPTION: Less than 400 pieces of DPS is at the case DPS standard (not raw)
	if ld.caseDPS() {
		time += float64(ld.DPSLetters) * standards.CASE_DPS_LETTERS
	}

	return time
}

// TotalLetters returns the total number of letters recieved for the route
func (ld *LetterData) TotalLetters() int {
	return (ld.RandomLetters + ld.WSSLetters + ld.DPSLetters + ld.CasedDPSLetters)
}

// CasedLeters returns the  total number of letters paid to case
func (ld *LetterData) CasedLetters() int {
	if ld.caseDPS() {
		return ld.TotalLetters()
	}
	return ld.TotalLetters() - ld.DPSLetters
}
