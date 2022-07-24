package averaged

import (
	"fmt"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

//Parcel accepted
//Registered certified accepted
//money order accepted

// These values are based off of averages and applied by number of boxes
type AveragedData struct {
	Boxes int
}

func (ad *AveragedData) Populate(dd ds.DailyData, fd ds.FixedData) {
	ad.Boxes = fd.RegCurbBox + fd.RegSdwkBox + fd.RegOtherBox
	ad.Boxes += fd.CBUBoxes + fd.CENTBoxes + fd.NPUBoxes + fd.DETBoxes
}

func (ad *AveragedData) Report() string {
	out := "\n************** Averaged Time *******************\n"
	out += fmt.Sprintf("Parcel Accepted Time: %4.2f\n", ad.GetParcelAcceptedTime())
	out += fmt.Sprintf("Reg/Cert Accepted Time: %4.2f\n", ad.GetRegCertAcceptedTime())
	out += fmt.Sprintf("Money Order Accepted Time:%4.2f\n", ad.GetMoneyOrderAcceptedTime())
	out += fmt.Sprintf("Averaged Time Total: %4.2f\n", ad.GetTime())
	return out
}

func (ad *AveragedData) GetTime() float64 {
	out := ad.GetParcelAcceptedTime()
	out += ad.GetRegCertAcceptedTime()
	out += ad.GetMoneyOrderAcceptedTime()
	return out
}

func (ad *AveragedData) GetParcelAcceptedTime() float64 {
	num := standards.PARCEL_ACCEPTED_PROB * float64(ad.Boxes)
	return num * standards.PARCEL_ACCEPTED_TIME
}

func (ad *AveragedData) GetRegCertAcceptedTime() float64 {
	num := standards.REG_CERT_ACCEPTED_PROB * float64(ad.Boxes)
	return num * standards.REG_CERT_ACCEPTED_TIME
}

func (ad *AveragedData) GetMoneyOrderAcceptedTime() float64 {
	num := standards.MONEY_ORDER_ACCEPTED_PROB * float64(ad.Boxes)
	return num * standards.MONEY_ORDER_ACCEPTED_TIME
}
