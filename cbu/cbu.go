package cbu

import (
	"fmt"
	"log"

	ds "rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

// CBUData contains all the info for calculating CBU times
type CBUData struct {
	CBUs            int
	CBUBoxes        int
	CBUBoxesSkipped int

	CENTs            int
	CENTBoxes        int
	CENTBoxesSkipped int

	NPUs            int
	NPUBoxes        int
	NPUBoxesSkipped int

	DETs            int
	DETBoxes        int
	DETBoxesSkipped int

	CollectionBoxes int

	Bundle standards.Bundling
}

func (cbud *CBUData) Populate(dd ds.DailyData, fd ds.FixedData) {

	cbud.CBUs = fd.CBUs
	cbud.CBUBoxes = fd.CBUBoxes
	cbud.CBUBoxesSkipped = dd.CBUBoxesSkipped
	cbud.CENTs = fd.CENTs
	cbud.CENTBoxes = fd.CENTBoxes
	cbud.CENTBoxesSkipped = dd.CENTBoxesSkipped
	cbud.NPUs = fd.NPUs
	cbud.NPUBoxes = fd.NPUBoxes
	cbud.NPUBoxesSkipped = dd.NPUBoxesSkipped
	cbud.DETs = fd.DETs
	cbud.DETBoxes = fd.DETBoxes
	cbud.DETBoxesSkipped = dd.DETBoxesSkipped

	cbud.CollectionBoxes = fd.CollectionBoxes

	cbud.Bundle = fd.Bundle
}

func (cbud *CBUData) Report() string {
	out := "\n************ Centralized Box Time *******\n"
	out += fmt.Sprintf("CBU time: %4.2f\n", cbud.cbuTime())
	out += fmt.Sprintf("CENT time: %4.2f\n", cbud.centTime())
	out += fmt.Sprintf("NPU time: %4.2f\n", cbud.npuTime())
	out += fmt.Sprintf("DET time: %4.2f\n", cbud.detTime())
	out += fmt.Sprintf("Collection Time %4.2f\n", cbud.collectionTime())
	out += fmt.Sprintf("\nCentralized Box Time Total %4.2f\n", cbud.GetTime())

	return out

}

func (cbud *CBUData) collectionTime() float64 {
	return float64(cbud.CollectionBoxes) * standards.COLLECTION_TIME
}

func (cbud *CBUData) cbuTime() float64 {
	time := 0.0
	boxes := float64(cbud.CBUBoxes - cbud.CBUBoxesSkipped)
	if boxes < 0 {
		log.Printf("Skipped more boxes than on the route!\n")
		boxes = 0
	}
	time += float64(cbud.CBUs) * standards.UNIT_TIME_CBU
	time += float64(boxes) * standards.CLUSTER_BUNDLE_BOX_TIME[cbud.Bundle]
	return time
}

func (cbud *CBUData) centTime() float64 {
	time := 0.0
	boxes := float64(cbud.CENTBoxes - cbud.CENTBoxesSkipped)
	if boxes < 0 {
		log.Printf("Skipped more boxes than on the route!\n")
		boxes = 0
	}
	time += float64(cbud.CENTs) * standards.UNIT_TIME_CENT
	time += float64(boxes) * standards.CLUSTER_BUNDLE_BOX_TIME[cbud.Bundle]
	return time
}

func (cbud *CBUData) npuTime() float64 {
	time := 0.0
	boxes := float64(cbud.NPUBoxes - cbud.NPUBoxesSkipped)
	if boxes < 0 {
		log.Printf("Skipped more boxes than on the route!\n")
		boxes = 0
	}
	time += float64(cbud.NPUs) * standards.UNIT_TIME_NPU
	time += float64(boxes) * standards.CLUSTER_BUNDLE_BOX_TIME[cbud.Bundle]
	return time
}

func (cbud *CBUData) detTime() float64 {
	time := 0.0
	boxes := float64(cbud.DETBoxes - cbud.DETBoxesSkipped)
	if boxes < 0 {
		log.Printf("Skipped more boxes than on the route!\n")
		boxes = 0
	}
	time += float64(cbud.DETs) * (standards.UNIT_TIME_DET + standards.PRESORT_TIME_DET)
	time += float64(boxes) * standards.BOX_TIME_DET
	return time
}

func (cbud *CBUData) GetTime() float64 {
	time := 0.0
	time += cbud.cbuTime()
	time += cbud.centTime()
	time += cbud.npuTime()
	time += cbud.detTime()

	collectionTime := cbud.collectionTime()

	time += collectionTime
	return time
}
