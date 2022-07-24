package markup

import (
	"fmt"
	"math"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

//Markuptime just a percentage, based on totalmailpieces
type MarkupData struct {
	AddressedMailPieces int
}

//ASSUMPTION: parcels do NOT count towards markup mail volume
func (md *MarkupData) Populate(dd ds.DailyData, fd ds.FixedData) {
	total := dd.RawLetters + dd.DPSLetters + dd.WSSLetters
	total += dd.RawFlats + dd.PreBundledFlats + dd.DPSFlats + dd.WSSFlats
	total += dd.CasedDPSLetters
	md.AddressedMailPieces = total
}

func (md *MarkupData) Report() string {
	out := "\n*********** Markup Data ***********\n"
	out += fmt.Sprintf("Total mail pieces for markup: %d\n", md.AddressedMailPieces)
	out += fmt.Sprintf("Markup Total Time: %4.2f\n", md.GetTime())
	return out
}

func (md *MarkupData) GetTime() float64 {
	markupBundles := float64(md.AddressedMailPieces) * standards.MARKUP_BUNDLE_RATE
	markupBundles = math.Ceil(markupBundles)
	return markupBundles * standards.MARKUP_TIME
}
