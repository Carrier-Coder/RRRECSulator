package dataSets

import "rrecsulator.com/standards"

type DailyData struct {

	// Letters
	RawLetters       int `json:"RawLetters,string"`
	DPSLetters       int `json:"DPSLetters,string"`
	WSSLetters       int `json:"WSSLetters,string"`
	BoxholderLetters int `json:"BoxholderLetters,string"`
	CasedDPSLetters  int `json:"CasedDPSLetters,string"`

	// Flats
	RawFlats        int `json:"RawFlats,string"`
	PreBundledFlats int `json:"PreBundledFlats,string"`
	DPSFlats        int `json:"DPSFlats,string"`
	WSSFlats        int `json:"WSSFlats,string"`
	BoxholderFlats  int `json:"BoxholderFlats,string"`
	CasedDPSFlats   int `json:"CasedDPSFlats,string"`

	//DoorTrips
	DoorTrips int `json:"DoorTrips,string"`

	//Driveways
	LessSwimPool     int `json:"LessSwimPool,string"`
	SwimPool         int `json:"SwimPool,string"`
	FootballField    int `json:"FootballField,string"`
	TwoFootballField int `json:"TwoFootballField,string"`
	QuarterMile      int `json:"QuarterMile,string"`
	HalfMile         int `json:"HalfMile,string"`

	//Parcels
	MailBoxParcels int `json:"MailBoxParcels,string"`
	LockerParcels  int `json:"LockerParcels,string"`
	DoorParcels    int `json:"DoorParcels,string"`

	//Accountables
	Certified           int `json:"Certified,string"`
	Registered          int `json:"Registered,string"`
	SignatureExpress    int `json:"SignatureExpress,string"`
	NonSignatureExpress int `json:"NonSignatureExpress,string"`
	CODs                int `json:"CODs,string"`
	PostageDueCustom    int `json:"PostageDueCustom,string"`
	PostageDue          int `json:"PostageDue,string"`

	SignatureParcel int `json:"SignatureParcel,string"`

	//Skipped boxes (box coverage factor)
	RegCurbBoxesSkipped  int `json:"RegCurbBoxesSkipped,string"`
	RegSdwkBoxesSkipped  int `json:"RegSdwkBoxesSkipped,string"`
	RegOtherBoxesSkipped int `json:"RegOtherBoxesSkipped,string"`

	CBUBoxesSkipped  int `json:"CBUBoxesSkipped,string"`
	CENTBoxesSkipped int `json:"CENTBoxesSkipped,string"`
	NPUBoxesSkipped  int `json:"NPUBoxesSkipped,string"`
	DETBoxesSkipped  int `json:"DETBoxesSkipped,string"`

	// miscellaneous trips to door (mdd entry)
	DoorMisc         int `json:"DoorMisc,string"`
	DoorPickupEvents int `json:"DoorPickupEvents,string"`

	//Additional Trips
	AddTripOther    int `json:"AddTripOther,string"`
	AddTripSidewalk int `json:"AddTripSidewalk,string"`
	AddTripCbu      int `json:"AddTripCbu,string"`
	AddTripCent     int `json:"AddTripCent,string"`
	AddTripNpu      int `json:"AddTripNpu,string"`
	AddTripDetVpo   int `json:"AddTripDetVpo,string"`

	//Scans
	//SPMs count as one step scans?
	OneStepScan int `json:"OneStepScan,string"`
	TwoStepScan int `json:"TwoStepScan,string"`
	StampSale   int `json:"StampSale,string"`
	RuralReach  int `json:"RuralReach,string"`

	//Package Pickups
	EventForms int `json:"EventForms,string"`
	Events     int `json:"Events,string"`
	PackageNum int `json:"PackageNum,string"`

	Form3982 int `json:"Form3982,string"`

	ExtraSteps int `json:"ExtraSteps,string"`

	//actual time in standard minutes
	Loading            float64 `json:"Loading,string"`
	EndOfShift         float64 `json:"EndOfShift,string"`
	Deviation          float64 `json:"Deviation,string"`
	SafetyServiceTalks float64 `json:"SafetyServiceTalks,string"`
	FudgeTime          float64 `json:"FudgeTime,string"`
}

type FixedData struct {
	//number of mailboxes on the route
	RegCurbBox  int `json:"RegCurbBox,string"`
	RegSdwkBox  int `json:"RegSdwkBox,string"`
	RegOtherBox int `json:"RegOtherBox,string"`

	//vehicle information
	DrivePOV   bool `json:"DrivePOV,string"`
	DriveGOVLH bool `json:"DriveGOVLH,string"`

	BlueBox int `json:"BlueBox,string"`

	//Centralized boxes on the route
	CBUs            int `json:"CBUs,string"`
	CBUBoxes        int `json:"CBUBoxes,string"`
	CENTs           int `json:"CENTs,string"`
	CENTBoxes       int `json:"CENTBoxes,string"`
	NPUs            int `json:"NPUs,string"`
	NPUBoxes        int `json:"NPUBoxes,string"`
	DETs            int `json:"DETs,string"`
	DETBoxes        int `json:"DETBoxes,string"`
	CollectionBoxes int `json:"CollectionBoxes,string"`

	//Authorized fixed dismounts
	DismountOther    int `json:"DismountOther,string"`
	DismountSidewalk int `json:"DismountSidewalk,string"`
	DismountCbu      int `json:"DismountCbu,string"`
	DismountCent     int `json:"DismountCent,string"`
	DismountNpu      int `json:"DismountNpu,string"`
	DismountDetVpo   int `json:"DismountDetVpo,string"`

	LowVolumePouches  int `json:"LowVolumePouches,string"`
	HighVolumePouches int `json:"HighVolumePouches,string"`

	//Route mileage and drive time
	Miles     float64 `json:"Miles,string"`
	DriveTime float64 `json:"DriveTime,string"`

	//Miscellaneous time from the mini mail survey
	MiscTime float64 `json:"MiscTime,string"`

	//Daily steps or walking distance in feet
	FixedSteps    int     `json:"FixedSteps,string"`
	DismountSteps int     `json:"DismountSteps,string"`
	FeetPerStep   float64 `json:"FeetPerStep,string"`
	FixedFeet     int     `json:"FixedFeet,string"`

	Bundle standards.Bundling `json:"Bundle,string"`

	WithdrawMail bool `json:"WithdrawMail,string"`
}
