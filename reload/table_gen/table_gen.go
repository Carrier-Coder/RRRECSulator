package main

import "fmt"

// Generate the MailTray time table and Medium parcel time table
// the one in the comprehensive guide was not long enough imo
// times are funky, don't exactly follow the values from the 4241-M
// Look at the excel sheet in folder for details

//Will calculate mail tray and medium parcel look up tables

const NUM_ROWS = 100

func main() {

	//mail tray lookup table
	fmt.Println("********** Mail Tray Table ***************")
	fmt.Println("")
	for i := 0; i < NUM_ROWS; i++ {
		line := "{"
		line += fmt.Sprintf("%2.4f, ", OneBundleTable(i))
		line += fmt.Sprintf("%2.4f, ", TwoBundleTable(i))
		line += fmt.Sprintf("%2.4f },", ThreeBundleTable(i))
		fmt.Println(line)
	}

	fmt.Println("********** Medium Parcel Table ***********")
	fmt.Println("")
	for i := 0; i < NUM_ROWS; i++ {
		line := fmt.Sprintf("%2.4f,", MediumParcelTable(i))
		fmt.Println(line)
	}

}

// Number of not reachable trays, returns time
func OneBundleTable(trays int) float64 {

	//first few values are funky, assumes preloaded upfont
	if trays <= 0 {
		return 0.0
	}
	if trays == 1 {
		return 0.6957
	}
	if trays == 2 {
		return 0.8689
	}
	if trays == 3 {
		return 2.282
	}
	if trays == 4 {
		return 2.4023
	}
	time := 2.4023
	trays -= 4

	// repeating pattern, 1 rear door time, 2 additional rear time
	for {
		trays -= 1
		time += 1.7544
		if trays == 0 {
			break
		}

		trays -= 1
		time += 0.2156
		if trays == 0 {
			break
		}

		trays -= 1
		time += 0.2156
		if trays == 0 {
			break
		}
	}
	return time
}

// 2-bundle, Number of not reachable trays, returns time
func TwoBundleTable(trays int) float64 {

	//first few values are funky, assumes preloaded upfont
	if trays <= 0 {
		return 0.0
	}
	if trays == 1 {
		return 0.6957
	}
	if trays == 2 {
		return 0.8689
	}
	if trays == 3 {
		return 2.282
	}

	if trays == 4 {
		return 3.6951
	}
	time := 3.6951
	trays -= 4

	// repeating pattern, 1 rear door time, 1 additional rear time
	for {
		trays -= 1
		time += 1.7544
		if trays == 0 {
			break
		}

		trays -= 1
		time += 0.2156
		if trays == 0 {
			break
		}
	}
	return time
}

// 3-bundle, Number of not reachable trays, returns time
func ThreeBundleTable(trays int) float64 {

	//first few values are funky, assumes preloaded upfont
	if trays <= 0 {
		return 0.0
	}
	if trays == 1 {
		return 0.6957
	}
	if trays == 2 {
		return 1.3914
	}
	if trays == 3 {
		return 2.8045
	}

	if trays == 4 {
		return 4.2176
	}
	time := 4.2176
	trays -= 4

	// repeating pattern, 1 rear door time
	for {
		trays -= 1
		time += 1.7544
		if trays == 0 {
			break
		}
	}
	return time
}

// Medium Parcels
func MediumParcelTable(tubs int) float64 {

	//first few values are funky, assumes preloaded upfont
	if tubs <= 0 {
		return 0.0
	}
	if tubs == 1 {
		return 0.7357
	}
	if tubs == 2 {
		return 0.9304
	}
	if tubs == 3 {
		return 1.661
	}
	if tubs == 4 {
		return 1.8606
	}

	time := 1.8606
	tubs -= 4

	// repeating pattern, 1 rear door time, 1 additional rear time
	for {
		tubs -= 1
		time += 1.2172
		if tubs == 0 {
			break
		}

		tubs -= 1
		time += 0.2156
		if tubs == 0 {
			break
		}
	}
	return time
}
