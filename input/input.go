// A lot from this package will have to be moved to reorganize, eventually.
// this was to throw some stuff together.
package input

import (
	"errors"
	"fmt"

	"github.com/yomcube/GhostDB/common"
)

// Can't declare const Array in Golang
// Arrays are always evaluated at runtime, for no reason
var RKGDMagicNumbers = [...]byte{0x52, 0x4B, 0x47, 0x44}
var CKGDMagicNumbers = [...]byte{0x43, 0x4B, 0x47, 0x44}

// There is no need to decode to an "RKGD" type, it would be pointless.
// Though this was to be decided yet, so any data here is temporary
type Time struct {
	CourseSlot  common.CourseID // This will have, in the future, to be handled for CTs.
	FinalTime   int32
	Laps        [10]int32 // max number of laps possible to write on file
	Vehicle     common.VehicleID
	Character   common.CharacterID
	Controller  common.ControllerID
	AutoDrift   bool
	CountryCode common.CountryCode
	StateCode   uint8
}

func (outputTime *Time) setStateCode(inputByte uint8) error {
	if !outputTime.CountryCode.StatecodeValid(inputByte) {
		return errors.New("tried to set invalid state code for time")
	} else {
		outputTime.StateCode = inputByte
		return nil
	}
}

// constructor from file
func InitializeFromRKGFile(inputBytes []byte) (Time, error) {
	outputTime := Time{}
	if [4]byte(inputBytes) != RKGDMagicNumbers {
		return outputTime, errors.New("not an RKGD file, missing RKGD headers")
	}

	ghostType := (inputBytes[0xC]<<7 | inputBytes[0xD]>>2)
	isCKGD := ghostType == 0x26

	outputTime.CourseSlot = common.CourseID(inputBytes[7] >> 2)
	outputTime.FinalTime = readTimeFromRKGFormat([3]byte(inputBytes[4:]))
	outputTime.Vehicle = common.VehicleID(inputBytes[8] >> 2)
	outputTime.Character = common.CharacterID(((inputBytes[8] & 0b00000011) << 4) | (inputBytes[9] >> 4))
	outputTime.Controller = common.ControllerID(inputBytes[0xB] & 0b00001111)
	outputTime.AutoDrift = (inputBytes[0xB] & 0b00001000) != 0b000000000
	outputTime.CountryCode = common.CountryCode(inputBytes[0x34])

	err := outputTime.setStateCode(inputBytes[0x35])
	if err != nil {
		return outputTime, err
	}

	for i := 0; i < 10; i++ {
		if !isCKGD && i > 5 { // If the file is not a CKGD, further laps will just be junk data
			break
		}
		outputTime.Laps[i] = readTimeFromRKGFormat([3]byte(inputBytes[(0x11 + i*3):]))
	}

	return outputTime, nil
}

// Formats an int32 time into a string like 00:00.000
func FormatTimeToString(milliseconds int32) string {
	var minutes = milliseconds / 60000
	milliseconds %= 60000
	var seconds = milliseconds / 1000
	milliseconds %= 1000
	return fmt.Sprintf("%02d:%02d.%03d", minutes, seconds, milliseconds)
}

func readTimeFromRKGFormat(inputBytes [3]byte) int32 {
	// 3 Bytes, where M = Minutes, S = Seconds and C = Millis.
	// 1. 0bMMMMMMMS
	// 2. 0bSSSSSSCC
	// 3. 0bCCCCCCCC

	// max M = 5    // 0b0000101
	// max S = 59   // 0b0111011
	// max C = 999  // 0b1111100111

	// 1. 0b00001010
	// 2. 0b11101111
	// 3. 0b11100111
	minutes := int32(inputBytes[0]) // The minutes are doubled here because of the extra bit

	// You don't need the first bit, because 59 = 6 bits
	// Go up 7 lines for the breakdown of why
	seconds := int32(inputBytes[1] >> 2)

	// I don't think these casts might be faster than just doing OR on a int32
	milliseconds := int32(((int16(inputBytes[1]) & 0b00000011) << 8) | int16(inputBytes[2]))

	// (Minutes * 60000) + (Seconds * 1000) + Milliseconds =
	// = (2*Minutes * 30000) + (Seconds * 1000) + Milliseconds
	return ((minutes) * 30000) + ((seconds) * 1000) + milliseconds
}
