package drive

import (
	"rrecsulator.com/standards"
)

// DriveData for calculating new drive time
type DriveData struct {

	// Total miles in route
	StopSigns     float64 `json:"stop-signs,string"`
	YieldSigns    float64 `json:"yield-signs,string"`
	TrafficLights float64 `json:"traffic-lights,string"`
	Crosswalks    float64 `json:"crosswalks,string"`
	RRCrossings   float64 `json:"rr-crossings,string"`
	AccessGates   float64 `json:"access-gates,string"`

	// Each space b/n two boxes less than 5ft is a creep
	Creeps int `json:"creeps,string"`

	// array of distances, in ft
	Distances []float64 `json:"distances"`
}

func (dd *DriveData) stopSignTime() float64 {
	return float64(dd.StopSigns) * standards.STOP_SIGN_TIME
}
func (dd *DriveData) yieldSignTime() float64 {
	return float64(dd.YieldSigns) * standards.YIELD_SIGN_TIME
}
func (dd *DriveData) trafficLightTime() float64 {
	return float64(dd.TrafficLights) * standards.TRAFFIC_LIGHT_TIME
}
func (dd *DriveData) crosswalkTime() float64 {
	return float64(dd.Crosswalks) * standards.CROSSWALK_TIME
}
func (dd *DriveData) rrCrossingTime() float64 {
	return float64(dd.RRCrossings) * standards.RRCROSSING_TIME
}
func (dd *DriveData) accessGateTime() float64 {
	return float64(dd.AccessGates) * standards.ACCESS_GATE_TIME
}
func (dd *DriveData) creepTime() float64 {
	return float64(dd.Creeps) * standards.CREEP_TIME
}

func (dd *DriveData) GetTime() float64 {

	time := 0.0

	//fixed times for different traffic dealies
	time += dd.stopSignTime()
	time += dd.yieldSignTime()
	time += dd.trafficLightTime()
	time += dd.crosswalkTime()
	time += dd.rrCrossingTime()
	time += dd.accessGateTime()
	time += dd.creepTime()

	//get time from Drive Speed Matrix
	time += applyDSM(dd.Distances)

	return time

}

// find drive time from an array of distances.
func applyDSM(distances []float64) float64 {
	time := 0.0
	speed := 0.0
	for _, d := range distances {
		speed = findSpeed(d)
		time += d * speed
	}
	return time
}

// Given distance in feet, find app. speed from drive speed matrix
func findSpeed(dist float64) float64 {
	// DSM = drive speed matrix
	out := standards.DSM[len(standards.DSM)-1].Speed
	for _, d := range standards.DSM {
		if d.Min < dist && d.Max >= dist {
			out = d.Speed
		}
	}
	return out
}
