package boxholder

import (
	"fmt"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

// BoxholderData contains all the info for calculating boxholder time
type BoxholderData struct {

	//these should be number delivered (ie 350 boxes got a bh on a given day)
	BoxholderFlats   int
	BoxholderLetters int
}

func (bhd *BoxholderData) Populate(dd ds.DailyData, fd ds.FixedData) {
	bhd.BoxholderFlats = dd.BoxholderFlats
	bhd.BoxholderLetters = dd.BoxholderLetters
}

func (bhd *BoxholderData) Report() string {
	out := "\n*************** Boxholder Time ********\n"
	out += fmt.Sprintf("Total Boxholders: %d\n", bhd.TotalBoxholders())
	out += fmt.Sprintf("Boxholder Time Total: %4.2f\n", bhd.GetTime())
	return out
}

func (bhd *BoxholderData) GetTime() float64 {

	//different times for letter vs flat in RRECS
	time := 0.0
	time += float64(bhd.BoxholderFlats) * standards.BOXHOLDER_FLATS_TIME
	time += float64(bhd.BoxholderLetters) * standards.BOXHOLDER_LETTERS_TIME

	return time

}

// The total number of boxholders recieved for the route
func (bhd *BoxholderData) TotalBoxholders() int {
	return bhd.BoxholderFlats + bhd.BoxholderLetters
}
