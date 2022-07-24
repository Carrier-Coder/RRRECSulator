package pouch

import (
	"fmt"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

type PouchData struct {

	// low is one trip, high is multiple trips
	LowVolume  int
	HighVolume int
}

func (pd *PouchData) Populate(dd ds.DailyData, fd ds.FixedData) {
	pd.LowVolume = fd.LowVolumePouches
	pd.HighVolume = fd.HighVolumePouches
}

func (pd *PouchData) Report() string {
	out := "\n************* Lock Pouch Data *******************\n"
	out += fmt.Sprintf("Low Volume: %d  High Volume: %d\n", pd.LowVolume, pd.HighVolume)
	out += fmt.Sprintf("Lock Pouch Total Time %4.2f\n", pd.GetTime())
	return out
}
func (pd *PouchData) GetTime() float64 {

	lowTime := float64(pd.LowVolume) * standards.LOW_VOLUME_POUCH_STOP_TIME
	highTime := float64(pd.HighVolume) * standards.HIGH_VOLUME_POUCH_STOP_TIME

	return lowTime + highTime
}
