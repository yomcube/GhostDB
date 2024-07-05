package input

import (
	"errors"
	"fmt"
	"encoding/binary"
)


// https://wiki.tockdom.com/wiki/RKG_(File_Format)
// TODO: clean up this package and add more RKG data


var RKGDMagicNumbers = [...]byte{0x52, 0x4B, 0x47, 0x44}
var CKGDMagicNumbers = [...]byte{0x43, 0x4B, 0x47, 0x44}

type RKGTime struct {
	Minutes byte
	Seconds byte
	Milliseconds uint16
}

type RKG struct { // Magic isn't necessary
	FinalTime   int32 // Times are stored as milliseconds
	Laps        [10]int32 // Lap splits
	Track       byte
	Vehicle     byte
	Character   byte
	Controller  byte
	AutoDrift   bool
	CountryCode byte
	StateCode   byte // TODO: Map these
}

// Constructor from file
func (outputRkg RKG) InitializeFromRKGFile(inputBytes []byte) (RKG, error) {
	if [4]byte(inputBytes) != RKGDMagicNumbers {
		return outputRkg, errors.New("Not a valid RKG file")
	}

	// -8 because len()-CRC Checksum (4bytes) - CKGD magic numbers (also 4bytes)
	isCKGD := [4]byte(inputBytes[len(inputBytes)-8:]) == CKGDMagicNumbers

	outputRkg.FinalTime = readTimeFromRKGFormat(inputBytes[4:7])
	outputRkg.Track = inputBytes[7] >> 2
	outputRkg.Vehicle = inputBytes[8] >> 1
	outputRkg.Character = ((inputBytes[8] & 0x3) << 2) | (inputBytes[9] >> 4)
	outputRkg.Controller = inputBytes[0xB] & 0xf
	outputRkg.AutoDrift = inputBytes[0xB] & 0x10 == 1

	for i := 0; i < 10; i++ {
		if i > 5 && !isCKGD { // If the file is not a CKGD, further laps will just be junk data
			break
		}
		outputRkg.Laps[i] = readTimeFromRKGFormat([3]byte(inputBytes[(0x11 + i*3):]))
	}

	return outputRkg, nil
}

func readTimeFromRKGFormat(inputBytes [3]byte) int32 {
	asU32 := binary.BigEndian.Uint32(inputBytes)
	minutes := (asU32 >> 17) & 0x7F
	seconds := (asU32 >> 10) & 0x7F
	milliseconds := asU32 & 0x3FF

	return (minutes * 60000) + (seconds * 1000) + milliseconds
}

// Formats an int32 time into a string like 00:00.000
func FormatTimeToString(milliseconds int32) string {
	var minutes = milliseconds / 60000
	milliseconds %= 60000
	var seconds = milliseconds / 1000
	milliseconds %= 1000
	return fmt.Sprintf("%02d:%02d.%03d", minutes, seconds, milliseconds)
}