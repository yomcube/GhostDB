package rkg_input_test

import (
	"os"
	"testing"

	"github.com/yomcube/GhostDB/common"
	"github.com/yomcube/GhostDB/input/rkg_input"
)

func TestRKG(t *testing.T) {
	genericTimeTest(
		rkg_input.Time{
			CourseSlot:  common.MarioCircuit,
			FinalTime:   70716,
			Vehicle:     common.MachBike,
			Character:   common.Daisy,
			Controller:  common.WiiWheel,
			CountryCode: common.USA,
			StateCode:   0x16,
			Laps:        [...]int32{10051, 10277, 9786, 9789, 10255, 10285, 10273, 0, 0, 0},
		},
		"./test_rkg/arjun_ctgp_gcn_baby_park_wr.rkg",
		t,
	)

	genericTimeTest(
		rkg_input.Time{
			CourseSlot:  common.LuigiCircuit,
			FinalTime:   85945,
			Vehicle:     common.Jetsetter,
			Character:   common.DryBowser,
			Controller:  common.WiiWheel,
			CountryCode: common.XXX,
			StateCode:   0xFF,
			Laps:        [...]int32{28645, 28817, 28483, 0, 0, 0, 0, 0, 0, 0},
		},
		"./test_rkg/italian_random_lc.rkg",
		t,
	)
}

func genericTimeTest(testTime rkg_input.Time, filePath string, t *testing.T) {
	byteData, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("os.ReadFile() error: %v", err)
		panic(err)
	}

	constructedTime, err := rkg_input.InitializeFromRKGFile(byteData)
	if err != nil {
		t.Errorf("%v: constructedTime.InitializeFromRKGFile() error: %v", filePath, err)
		panic(err)
	}

	if constructedTime.CourseSlot != testTime.CourseSlot {
		t.Errorf("%v: constructedTime.InitializeFromRKGFile() error: constructedTime.CourseSlot is wrong, value is %d", filePath, constructedTime.CourseSlot)
	}

	if constructedTime.FinalTime != testTime.FinalTime {
		t.Errorf("%v: constructedTime.InitializeFromRKGFile() error: constructedTime.FinalTime is wrong, value is %d", filePath, constructedTime.FinalTime)
	}

	if constructedTime.Vehicle != testTime.Vehicle {
		t.Errorf("%v: constructedTime.InitializeFromRKGFile() error: constructedTime.Vehicle is wrong, value is %d", filePath, constructedTime.Vehicle)
	}

	if constructedTime.Character != testTime.Character {
		t.Errorf("%v: constructedTime.InitializeFromRKGFile() error: constructedTime.Character is wrong, value is %d", filePath, constructedTime.Character)
	}

	if constructedTime.Controller != testTime.Controller {
		t.Errorf("%v: constructedTime.InitializeFromRKGFile() error: constructedTime.Controller is wrong, value is %d", filePath, constructedTime.Controller)
	}

	if constructedTime.CountryCode != testTime.CountryCode {
		t.Errorf("%v: constructedTime.InitializeFromRKGFile() error: constructedTime.CountryCode is wrong, value is %d", filePath, constructedTime.Controller)
	}

	if constructedTime.StateCode != testTime.StateCode {
		t.Errorf("%v: constructedTime.InitializeFromRKGFile() error: constructedTime.StateCode is wrong, value is %d", filePath, constructedTime.Controller)
	}

	for i := 0; i < len(constructedTime.Laps); i++ {
		if constructedTime.Laps[i] != testTime.Laps[i] {
			t.Errorf("%v: constructedTime.InitializeFromRKGFile() error: constructedTime.Laps[%d] is wrong, value is %d", filePath, i, constructedTime.Laps[i])
		}
	}
}
