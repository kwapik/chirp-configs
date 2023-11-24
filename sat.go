package main

import "fmt"

var (
	entries = []*Entry{
		{NameSuffix: "-10", DownlinkOffset: -0.01, UplinkOffset: 0},
		{NameSuffix: "-5", DownlinkOffset: -0.005, UplinkOffset: 0},
		{NameSuffix: "--", DownlinkOffset: 0, UplinkOffset: 0},
		{NameSuffix: "+5", DownlinkOffset: +0.005, UplinkOffset: 0},
		{NameSuffix: "+10", DownlinkOffset: +0.01, UplinkOffset: 0},
	}
)

type Entry struct {
	NameSuffix     string
	DownlinkOffset float64
	UplinkOffset   float64
}

func main() {
	sats := []*Sat{
		{"ISS", 437.800, 145.990, 67},
		{"FX1B", 145.960, 435.250, 67},
		{"FX1D", 145.880, 435.350, 67},
		{"SO50", 436.795, 145.850, 67},
		{"DWT2", 145.900, 437.500, 141.3},
		{"LIL2", 437.200, 144.350, 88.5},
		{"UVSQ", 437.020, 145.905, 88.5},
		{"CS5A", 435.600, 145.925, 88.5},
		// Commented out until they are launched/enabled.
		// {"CS7A", 435.455, 145.950, 88.5},
		// {"CS7C", 435.690, 145.900, 88.5},
		{"TEVL", 435.400, 145.970, 88.5},
		// Commented out until they are launched/enabled.
		// {"CSS1", 436.510, 145.875, 88.5},
		// {"CSS2", 145.985, 435.075, 88.5},
		// {"CRC1", 145.980, 435.900, 88.5},
		// {"BY70", 145.950, 435.950, 88.5},
		// {"INS8", 435.200, 145.830, 88.5},
	}
	fmt.Println(CSVHeaderRow())
	row := 1
	for _, s := range sats {
		lines := s.CSVRows(row)
		row += len(lines)
		for _, l := range lines {
			fmt.Println(l)
		}
	}
}

type Sat struct {
	// Name can have max 4 letters
	Name     string
	Downlink float64
	Uplink   float64
	// Tone ignored if set to 0
	Tone float64
}

func (s *Sat) CSVRows(offset int) []string {
	toneMode := ""
	tone := s.Tone
	// TODO: I have no idea why `88.5` is "zero" value.
	if s.Tone != 0 && s.Tone != 88.5 {
		toneMode = "Tone"
		tone = s.Tone
	}
	rows := make([]string, len(entries))
	for i, e := range entries {
		rows[i] = fmt.Sprintf("%v,%v,%.4f,split,%.4f,%v,%v,88.5,23,NN,23,Tone->Tone,FM,5,,4.0W,,,,,",
			offset+i, s.Name+e.NameSuffix, s.Downlink+e.DownlinkOffset, s.Uplink+e.UplinkOffset, toneMode, tone)
	}
	return rows
}

func CSVHeaderRow() string {
	return "Location,Name,Frequency,Duplex,Offset,Tone,rToneFreq,cToneFreq,DtcsCode,DtcsPolarity,RxDtcsCode,CrossMode,Mode,TStep,Skip,Power,Comment,URCALL,RPT1CALL,RPT2CALL,DVCODE"
}
