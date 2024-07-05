package input_test

import (
	"os"
	"testing"

	"github.com/yomcube/GhostDB/input"
)

func TestArjunCTGPGCNBabyPark(t *testing.T) {
	byteData, err := os.ReadFile("./test/arjun_ctgp_gcn_baby_park_wr.rkg")
	if err != nil {
		t.Errorf("os.ReadFile() error: %v", err)
		panic(err)
	}

	testTime, err := input.Time{}.InitializeFromRKGFile(byteData)
	if err != nil {
		t.Errorf("test_arjun: testTime.InitializeFromRKGFile() error: %v", err)
		panic(err)
	}

	if testTime.FinalTime != 70716 {
		t.Errorf("test_arjun: testTime.InitializeFromRKGFile() error: testTime.FinalTime is wrong, value is %d", testTime.FinalTime)
	}

	if testTime.Vehicle != input.Vehicle(input.MachBike) {
		t.Errorf("test_arjun: testTime.InitializeFromRKGFile() error: testTime.Vehicle is wrong, value is %d", testTime.Vehicle)
	}

	if testTime.Character != input.Character(input.Daisy) {
		t.Errorf("test_arjun: testTime.InitializeFromRKGFile() error: testTime.Character is wrong, value is %d", testTime.Character)
	}

	if testTime.Controller != input.Controller(input.WiiWheel) {
		t.Errorf("test_arjun: testTime.InitializeFromRKGFile() error: testTime.Controller is wrong, value is %d", testTime.Controller)
	}

	expectedLaps := [...]int32{10051, 10277, 9786, 9789, 10255, 10285, 10273, 0, 0, 0}

	for i := 0; i < len(testTime.Laps); i++ {
		if testTime.Laps[i] != expectedLaps[i] {
			t.Errorf("test_arjun: testTime.InitializeFromRKGFile() error: testTime.Laps[%d] is wrong, value is %d", i, testTime.Laps[i])
		}
	}

}