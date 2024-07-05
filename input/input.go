// A lot from this package will have to be moved to reorganize, eventually.
// this was to throw some stuff together.
package input

import (
	"errors"
	"fmt"
)

// Can't declare const Array in Golang
// Arrays are always evaluated at runtime, for no reason
var RKGDMagicNumbers = [...]byte{0x52, 0x4B, 0x47, 0x44}
var CKGDMagicNumbers = [...]byte{0x43, 0x4B, 0x47, 0x44}

// There is no need to decode to an "RKGD" type, it would be pointless.
// Though this was to be decided yet, so any data here is temporary
type Time struct {
	FinalTime   int32
	Laps        [10]int32 // max number of laps possible to write on file
	Vehicle     Vehicle
	Character   Character
	Controller  Controller
	AutoDrift   bool
	CountryCode uint8 // Someone will have to map these out
	StateCode   uint8 // ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
}

// constructor from file
// tested on https://www.chadsoft.co.uk/time-trials/rkgd/BE/B0/2FAD030C44088A385FC834AF02A50D7C2BB1.rkg
func (outputTime Time) InitializeFromRKGFile(inputBytes []byte) (Time, error) {
	if [4]byte(inputBytes) != RKGDMagicNumbers {
		return outputTime, errors.New("not an RKGD file, missing RKGD headers")
	}

	// -9 because len()-CRC Checksum (4bytes) - CKGD magic numbers (also 4bytes)
	isCKGD := [4]byte(inputBytes[len(inputBytes)-8:]) == CKGDMagicNumbers

	outputTime.FinalTime = readTimeFromRKGFormat([3]byte(inputBytes[4:]))
	outputTime.Vehicle = Vehicle(inputBytes[8] >> 2)
	outputTime.Character = Character(((inputBytes[8] & 0b00000011) << 4) | (inputBytes[9] >> 4))
	outputTime.Controller = Controller(inputBytes[0xB] & 0b00001111)
	outputTime.AutoDrift = (inputBytes[0xB] & 0b00001000) != 0b000000000

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
	// Go up 8 lines for the breakdown of why
	seconds := int32(inputBytes[1] >> 2)

	// I don't think these casts might be faster than just doing OR on a int32
	milliseconds := int32(((int16(inputBytes[1]) & 0b00000011) << 8) | int16(inputBytes[2]))

	// (Minutes * 60000) + (Seconds * 1000) + Milliseconds =
	// = (2*Minutes * 30000) + (Seconds * 1000) + Milliseconds
	return ((minutes) * 30000) + ((seconds) * 1000) + milliseconds
}

// Character Enum
type Character uint8

const (
	Mario            Character = 0
	BabyPeach        Character = 1
	Waluigi          Character = 2
	Bowser           Character = 3
	BabyDaisy        Character = 4
	DryBones         Character = 5
	BabyMario        Character = 6
	Luigi            Character = 7
	Toad             Character = 8
	DonkeyKong       Character = 9
	Yoshi            Character = 10
	Wario            Character = 11
	BabyLuigi        Character = 12
	Toadette         Character = 13
	Koopa            Character = 14
	Daisy            Character = 15
	Peach            Character = 16
	Birdo            Character = 17
	DiddyKong        Character = 18
	KingBoo          Character = 19
	BowserJr         Character = 20
	DryBowser        Character = 21
	FunkyKong        Character = 22
	Rosalina         Character = 23
	MiiAMaleSmall    Character = 24
	MiiAFemaleSmall  Character = 25
	MiiBMaleSmall    Character = 26
	MiiBFemaleSmall  Character = 27
	MiiCMaleSmall    Character = 28
	MiiCFemaleSmall  Character = 29
	MiiAMaleMedium   Character = 30
	MiiAFemaleMedium Character = 31
	MiiBMaleMedium   Character = 32
	MiiBFemaleMedium Character = 33
	MiiCMaleMedium   Character = 34
	MiiCFemaleMedium Character = 35
	MiiAMaleHeavy    Character = 36
	MiiAFemaleHeavy  Character = 37
	MiiBMaleHeavy    Character = 38
	MiiBFemaleHeavy  Character = 39
	MiiCMaleHeavy    Character = 40
	MiiCFemaleHeavy  Character = 41
	MediumMii        Character = 42
	SmallMii         Character = 43
	LargeMii         Character = 44
	BikerPeach       Character = 45
	BikerDaisy       Character = 46
	BikerRosalina    Character = 47
)

// Vehicle Enum
type Vehicle uint8

const (
	StandardKartS   Vehicle = 0
	StandardKartM   Vehicle = 1
	StandardKartL   Vehicle = 2
	BoosterSeat     Vehicle = 3
	ClassicDragster Vehicle = 4
	Offroader       Vehicle = 5
	MiniBeast       Vehicle = 6
	WildWing        Vehicle = 7
	FlameFlyer      Vehicle = 8
	CheepCharger    Vehicle = 9
	SuperBlooper    Vehicle = 10
	PiranhaProwler  Vehicle = 11
	TinyTitan       Vehicle = 12
	Daytripper      Vehicle = 13
	Jetsetter       Vehicle = 14
	BlueFalcon      Vehicle = 15
	Sprinter        Vehicle = 16
	Honeycoupe      Vehicle = 17
	StandardBikeS   Vehicle = 18
	StandardBikeM   Vehicle = 19
	StandardBikeL   Vehicle = 20
	BulletBike      Vehicle = 21
	MachBike        Vehicle = 22
	FlameRunner     Vehicle = 23
	BitBike         Vehicle = 24
	Sugarscoot      Vehicle = 25
	WarioBike       Vehicle = 26
	Quacker         Vehicle = 27
	ZipZip          Vehicle = 28
	ShootingStar    Vehicle = 29
	Magikruiser     Vehicle = 30
	Sneakster       Vehicle = 31
	Spear           Vehicle = 32
	JetBubble       Vehicle = 33
	DolphinDasher   Vehicle = 34
	Phantom         Vehicle = 35
)

// Controller Enum
type Controller uint8

const (
	WiiWheel          = 0
	Nunchuck          = 1
	ClassicController = 2
	GameCube          = 3
)
