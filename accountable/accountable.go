package accountable

import (
	"fmt"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

// AccountableData contains all the info for calculating Accountable item time.
type AccountableData struct {
	Certified           int
	Registered          int
	SignatureExpress    int
	NonSignatureExpress int

	CODs             int
	PostageDueCustom int

	PostageDue int

	//ASSUMPTION: parcels requiring a delivery signature credited as accountables
	//Do *NOT* count towards parcel credit
	SignatureParcel int
}

// Populate the data from the two data structs
func (ad *AccountableData) Populate(dd ds.DailyData, fd ds.FixedData) {
	ad.Certified = dd.Certified
	ad.Registered = dd.Registered
	ad.SignatureExpress = dd.SignatureExpress
	ad.NonSignatureExpress = dd.NonSignatureExpress
	ad.CODs = dd.CODs
	ad.PostageDueCustom = dd.PostageDueCustom
	ad.PostageDue = dd.PostageDue
	ad.SignatureParcel = dd.SignatureParcel
}

// Report generates a detailed plain text report.
func (ad *AccountableData) Report() string {
	out := "\n**************** Accountable Time ****************\n"
	out += fmt.Sprintf("Total Accountables: %d\n", ad.Total())
	out += fmt.Sprintf("Signature Accountables: %d\n", ad.Signature())
	out += fmt.Sprintf("Accountables Time Total: %4.2f\n", ad.GetTime())
	return out
}

//Total returns number of accountables for the day
func (ad AccountableData) Total() int {
	out := ad.Certified + ad.Registered + ad.SignatureExpress
	out += ad.NonSignatureExpress + ad.CODs
	out += ad.PostageDue + ad.PostageDueCustom
	out += ad.SignatureParcel
	return out
}

//Signature returns accountables requiring a signature
func (ad AccountableData) Signature() int {
	out := ad.Certified + ad.Registered
	out += ad.SignatureExpress + ad.SignatureParcel
	return out
}

//GetTime calculates the evaluation for accountable items
func (ad *AccountableData) GetTime() float64 {

	time := 0.0
	// only paid for 1 cage fight a day, as long as at leat 1 accountable
	if ad.Total() >= 1 {
		time += standards.FIXED_ACCOUNTABLE_CAGE
	}

	// process signature items, COD, customs due, same time
	temp := ad.Signature()
	temp += ad.CODs + ad.PostageDueCustom
	time += float64(temp) * standards.PROCESS_ACCOUNTABLE_MAIL

	//ASSUMPTION:  non-signature express mail gets standard s038
	time += float64(ad.NonSignatureExpress) *
		standards.DELIVER_ACCOUNTABLE_SIG_TIME

	//delivering accountable signature item
	time += float64(ad.Signature()) *
		standards.DELIVER_ACCOUNTABLE_SIG_TIME

	// delivering customs due and cod
	time += float64(ad.CODs) * standards.DELIVER_COD_TIME
	time += float64(ad.PostageDueCustom) *
		standards.DELIVER_CUSTOMS_DUE_TIME

	// delivering postage due
	time += float64(ad.PostageDue) * standards.PROCESS_POSTAGE_DUE
	time += float64(ad.PostageDue) * standards.COLLECT_POSTAGE_DUE

	return time
}
