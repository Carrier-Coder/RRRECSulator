package withdraw

import (
	"fmt"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

// if route doesn't recieve DPS Letters or Flats, set the value to zero
type WithdrawData struct {
	DPSLetters int
	DPSFlats   int

	Withdraw bool
}

func (wd *WithdrawData) Populate(dd ds.DailyData, fd ds.FixedData) {
	wd.DPSLetters = dd.DPSLetters
	wd.DPSFlats = dd.DPSFlats
	wd.Withdraw = fd.WithdrawMail
}

func (wd *WithdrawData) Report() string {
	out := "\n********* Withdraw Data ***********\n"
	out += fmt.Sprintf("Withdraw Total Time: %4.2f\n", wd.GetTime())
	return out
}

func (wd *WithdrawData) GetTime() float64 {
	if wd.Withdraw == false {
		return 0.0
	}

	letterTrays := 0
	temp := wd.DPSLetters
	for temp > 0 {
		letterTrays++
		temp -= standards.LETTERS_PER_TRAY
	}

	flatsTrays := 0
	temp = wd.DPSFlats
	for temp > 0 {
		flatsTrays++
		temp -= standards.FLATS_PER_TRAY
	}

	time := 0.0
	time += standards.GATHER_RANDOM_MAIL_TIME

	if wd.DPSLetters > 0 {
		time += standards.ACCESS_DPS_LETTER_TIME
		time += float64(letterTrays) * standards.GATHER_DPS_LETTER_TRAY_TIME
	}

	if wd.DPSFlats > 0 {
		time += standards.ACCESS_DPS_FLAT_TIME
		time += float64(flatsTrays) * standards.GATHER_DPS_FLAT_TRAY_TIME
	}

	return time
}
