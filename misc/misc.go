package misc

import (
	"fmt"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

// data for other events.  Catchalls that only have 1 simple standard
type MiscData struct {
	Form3982 int
	BlueBox  int

	//spm scans included here
	OneStepScan int
	TwoStepScan int

	StampSale  int
	RuralReach int

	//From mail survey
	MiscTime float64

	//TODO: should move drivetime to a separate package?
	DriveTime float64
}

func (md *MiscData) Populate(dd ds.DailyData, fd ds.FixedData) {
	md.Form3982 = dd.Form3982
	md.BlueBox = fd.BlueBox
	md.OneStepScan = dd.OneStepScan
	md.TwoStepScan = dd.TwoStepScan
	md.StampSale = dd.StampSale
	md.RuralReach = dd.RuralReach
	md.DriveTime = fd.DriveTime
	md.MiscTime = fd.MiscTime
}

func (md *MiscData) Report() string {
	out := "\n*********** Drive Time *************\n"
	out += fmt.Sprintf("Drive Time: %4.2f\n", md.DriveTime)
	out += "\n*********** Misc Time *************\n"
	out += fmt.Sprintf("Trip Report: %4.2f\n", md.GetTripReportTime())
	out += fmt.Sprintf("Tray Storage: %4.2f\n", md.GetTrayStorageTime())
	out += fmt.Sprintf("Scanner Setup: %4.2f\n", md.GetScannerSetupTime())
	out += fmt.Sprintf("Stamp Sale: %4.2f\n", md.GetStampSaleTime())
	out += fmt.Sprintf("Rural Reach: %4.2f\n", md.GetRuralReachTime())
	out += fmt.Sprintf("Form3982: %4.2f\n", md.GetForm3982Time())
	out += fmt.Sprintf("Blue Box: %4.2f\n", md.GetBlueBoxTime())
	out += fmt.Sprintf("One Step Scan: %4.2f\n", md.GetOneStepScanTime())
	out += fmt.Sprintf("Two Step Scan: %4.2f\n", md.GetTwoStepScanTime())
	out += fmt.Sprintf("\nMisc Time from Mail Survey: %4.2f\n", md.MiscTime)
	out += fmt.Sprintf("\nMisc Total Time: %4.2f\n", md.GetTime()-md.DriveTime)
	return out

}

func (md *MiscData) GetTime() float64 {
	time := 0.0
	time += md.DriveTime
	time += md.GetTripReportTime()
	time += md.GetTrayStorageTime()
	time += md.GetScannerSetupTime()
	time += md.GetStampSaleTime()
	time += md.GetRuralReachTime()
	time += md.GetForm3982Time()
	time += md.GetBlueBoxTime()
	time += md.GetOneStepScanTime()
	time += md.GetTwoStepScanTime()
	time += md.MiscTime

	return time
}

func (md *MiscData) GetTripReportTime() float64 {
	// 1 trip report per day
	return 1 * standards.TRIP_REPORT_TIME
}

func (md *MiscData) GetTrayStorageTime() float64 {
	//1 tray storage time per day
	return 1 * standards.MOVE_TRAYS_STORAGE_TIME
}

func (md *MiscData) GetScannerSetupTime() float64 {
	// only get 1 scanner setup time per day
	return 1 * standards.SCANNER_SETUP_TIME
}

func (md *MiscData) GetStampSaleTime() float64 {
	return float64(md.StampSale) * standards.SELL_STAMP_STOCK_TIME
}
func (md *MiscData) GetRuralReachTime() float64 {
	return float64(md.RuralReach) * standards.RURAL_REACH_TIME
}

func (md *MiscData) GetForm3982Time() float64 {
	return float64(md.Form3982) * standards.FORM3982_PARS_LABEL
}

func (md *MiscData) GetBlueBoxTime() float64 {
	return float64(md.BlueBox) * standards.SERVICE_BLUE_BOX
}

func (md *MiscData) GetOneStepScanTime() float64 {
	return float64(md.OneStepScan) * standards.ONE_STEP_RRECS_SCAN
}

func (md *MiscData) GetTwoStepScanTime() float64 {
	return float64(md.TwoStepScan) * standards.TWO_STEP_RRECS_SCAN
}
