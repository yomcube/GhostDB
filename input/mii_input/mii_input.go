package mii_input

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

type Mii struct {
	IsFemale             bool
	BirthMonth           uint8
	BirthDay             uint8
	FavoriteColor        mii_fav_colors
	IsFavorite           bool
	MiiName              string // utf-16be encoded null-terminated-ish string with max 10 char.
	Height               uint8
	Width                uint8
	MiiID                uint32
	MiiCreationTimestamp time.Time // it is the rightmost 28 of the 32 bits in the mii id
	ConsoleID            uint32    // the last three bytes are the last three bytes of the mac address, first byte is a checksum
	FaceType             uint8
	SkinTone             uint8
	FaceFeatures         uint8
	Mingling             bool
	SourceType           mii_source_type
	HairType             uint8
	HairColor            uint8
	HairFlip             bool
	EyebrowType          uint8
	EyebrowRotation      uint8
	EyebrowColor         uint8
	EyebrowSize          uint8
	EyebrowVertical      uint8
	EyebrowHorizontal    uint8
	EyeType              uint8
	EyeRotation          uint8
	EyeVertical          uint8
	EyeColor             uint8
	EyeSize              uint8
	EyeHorizontal        uint8
	NoseType             uint8
	NoseSize             uint8
	NoseVertical         uint8
	MouthType            uint8
	MouthColor           uint8
	MouthSize            uint8
	MouthVertical        uint8
	GlassesType          uint8
	GlassesColor         uint8
	GlassesSize          uint8
	GlassesVertical      uint8
	FacialHairMustache   uint8
	FacialHairBeard      uint8
	FacialHairColor      uint8
	FacialHairSize       uint8
	FacialHairVertical   uint8
	MoleType             bool
	MoleSize             uint8
	MoleVertical         uint8
	MoleHorizontal       uint8
	CreatorName          string // utf-16be encoded null-terminated-ish string with max 10 char.
}

func (mii *Mii) SetBirth(month byte, day byte) error {
	switch month {
	case 0:
		return nil // Both have to be unknown.
	case 2:
		if day > 29 {
			return errors.New("february has at most 29 days")
		}
		fallthrough
	case 4, 6, 9, 11:
		if day > 30 {
			return errors.New(fmt.Sprintf("%s has at most 30 days", strings.ToLower(mii_birthmonth(month).ToString())))
		}
		fallthrough
	case 1, 3, 5, 7, 8, 10, 12:
		if day > 31 {
			return errors.New(fmt.Sprintf("%s has at most 30 days", strings.ToLower(mii_birthmonth(month).ToString())))
		}
	}
	mii.BirthMonth = month
	mii.BirthDay = day
	return nil
}

func (mii *Mii) SetFavoriteColor(color byte) error {
	if color > 11 {
		return errors.New("favorite color has to be under 11")
	}
	mii.FavoriteColor = mii_fav_colors(color)
	return nil
}

func decodeUTF16BE(data []byte) (string, error) {
	preset := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	strBytes, err := io.ReadAll(transform.NewReader(bytes.NewReader([]byte(data)), preset.NewDecoder()))
	if err != nil {
		return "", err
	}
	var outInd int = 21
	for ind, strByte := range strBytes {
		if strByte == 0 {
			outInd = ind
			break
		}
	}
	return string(strBytes[:outInd]), nil
}

func (mii *Mii) SetNameFromUTF16BEBytes(name [20]byte) error {
	outName, err := decodeUTF16BE([]byte(name[:]))
	if err != nil {
		return err
	}
	mii.MiiName = outName
	return nil
}

func (mii *Mii) SetCreatorNameFromUTF16BEBytes(creatorName [20]byte) error {
	outCreatorName, err := decodeUTF16BE([]byte(creatorName[:]))
	if err != nil {
		return err
	}
	mii.CreatorName = outCreatorName
	return nil
}

func max7bits(n byte) byte {
	return (n & 0b01111111)
}

func (mii *Mii) SetHeight(height byte) {
	mii.Height = max7bits(height)
}

func (mii *Mii) SetWidth(width byte) {
	mii.Width = max7bits(width)
}

func (mii *Mii) SetSkinTone(skinTone byte) error {
	if skinTone > 5 {
		return errors.New("skin tone value more than 5")
	}
	mii.SkinTone = skinTone
	return nil
}

func (mii *Mii) SetFaceFeatures(faceFeatures byte) error {
	if faceFeatures > 11 {
		return errors.New("face features value more than 11")
	}
	mii.FaceFeatures = faceFeatures
	return nil
}

func (mii *Mii) SetHairType(hairType byte) error {
	if hairType > 71 {
		return errors.New("hair type value more than 71")
	}
	mii.HairType = hairType
	return nil
}

func (mii *Mii) SetEyebrowType(eyebrowType byte) error {
	if eyebrowType > 23 {
		return errors.New("eyebrow type value more than 23")
	}
	mii.EyebrowType = eyebrowType
	return nil
}

func (mii *Mii) SetEyebrowRotation(eyebrowRotation byte) error {
	if eyebrowRotation > 11 {
		return errors.New("eyebrow rotation value more than 11")
	}
	mii.EyebrowRotation = eyebrowRotation
	return nil
}

func (mii *Mii) SetEyebrowSize(eyebrowSize byte) error {
	if eyebrowSize > 8 {
		return errors.New("eyebrow size value more than 8")
	}
	mii.EyebrowSize = eyebrowSize
	return nil
}

func (mii *Mii) SetEyebrowVertical(eyebrowVertical byte) error {
	if eyebrowVertical > 18 {
		return errors.New("eyebrow vertical value more than 18")
	}
	if eyebrowVertical < 3 {
		return errors.New("eyebrow vertical value less than 3")
	}
	mii.EyebrowVertical = eyebrowVertical
	return nil
}

func (mii *Mii) SetEyebrowHorizontal(eyebrowHorizontal byte) error {
	if eyebrowHorizontal > 12 {
		return errors.New("eyebrow horizontal value more than 12")
	}
	mii.EyebrowHorizontal = eyebrowHorizontal
	return nil
}

func (mii *Mii) SetEyeType(eyeType byte) error {
	if eyeType > 47 {
		return errors.New("eye type value more than 47")
	}
	mii.EyeType = eyeType
	return nil
}

func (mii *Mii) SetEyeVertical(eyeVertical byte) error {
	if eyeVertical > 18 {
		return errors.New("eye vertical value more than 18")
	}
	mii.EyeVertical = eyeVertical
	return nil
}

func (mii *Mii) SetEyeColor(eyeColor byte) error {
	if eyeColor > 5 {
		return errors.New("eye color value more than 5")
	}
	mii.EyeColor = eyeColor
	return nil
}

func (mii *Mii) SetEyeHorizontal(eyeHorizontal byte) error {
	if eyeHorizontal > 12 {
		return errors.New("eye horizontal value more than 12")
	}
	mii.EyeHorizontal = eyeHorizontal
	return nil
}

func (mii *Mii) SetNoseType(noseType byte) error {
	if noseType > 11 {
		return errors.New("nose type value more than 11")
	}
	mii.NoseType = noseType
	return nil
}

func (mii *Mii) SetNoseSize(noseSize byte) error {
	if noseSize > 8 {
		return errors.New("nose size value more than 8")
	}
	mii.NoseSize = noseSize
	return nil
}

func (mii *Mii) SetNoseVertical(noseVertical byte) error {
	if noseVertical > 18 {
		return errors.New("nose vertical value more than 18")
	}
	mii.NoseVertical = noseVertical
	return nil
}

func (mii *Mii) SetMouthType(mouthType byte) error {
	if mouthType > 23 {
		return errors.New("mouth type value more than 23")
	}
	mii.MouthType = mouthType
	return nil
}

func (mii *Mii) SetMouthColor(mouthColor byte) error {
	if mouthColor > 2 {
		return errors.New("mouth color value more than 2")
	}
	mii.MouthColor = mouthColor
	return nil
}

func (mii *Mii) SetMouthSize(mouthSize byte) error {
	if mouthSize > 8 {
		return errors.New("mouth size value more than 8")
	}
	mii.MouthSize = mouthSize
	return nil
}

func (mii *Mii) SetMouthVertical(mouthVertical byte) error {
	if mouthVertical > 18 {
		return errors.New("mouth vertical value more than 18")
	}
	mii.MouthVertical = mouthVertical
	return nil
}

func (mii *Mii) SetGlassesType(glassesType byte) error {
	if glassesType > 8 {
		return errors.New("glasses type value more than 8")
	}
	mii.GlassesType = glassesType
	return nil
}

func (mii *Mii) SetGlassesColor(glassesColor byte) error {
	if glassesColor > 5 {
		return errors.New("glasses color value more than 5")
	}
	mii.GlassesColor = glassesColor
	return nil
}

func (mii *Mii) SetGlassesVertical(glassesVertical byte) error {
	if glassesVertical > 20 {
		return errors.New("glasses vertical value more than 20")
	}
	mii.GlassesVertical = glassesVertical
	return nil
}

func (mii *Mii) SetFacialHairSize(facialHairSize byte) error {
	if facialHairSize > 8 {
		return errors.New("facial hair size value more than 8")
	}
	mii.FacialHairSize = facialHairSize
	return nil
}

func (mii *Mii) SetFacialHairVertical(facialHairVertical byte) error {
	if facialHairVertical > 10 {
		return errors.New("glasses vertical value more than 10")
	}
	mii.FacialHairVertical = facialHairVertical
	return nil
}

func (mii *Mii) SetMoleSize(moleSize byte) error {
	if moleSize > 8 {
		return errors.New("mole size value more than 8")
	}
	mii.MoleSize = moleSize
	return nil
}

func (mii *Mii) SetMoleVertical(moleVertical byte) error {
	if moleVertical > 30 {
		return errors.New("mole vertical value more than 30")
	}
	mii.MoleVertical = moleVertical
	return nil
}

func (mii *Mii) SetMoleHorizontal(moleHorizontal byte) error {
	if moleHorizontal > 16 {
		return errors.New("mole horizontal value more than 16")
	}
	mii.MoleHorizontal = moleHorizontal
	return nil
}

func InitializeFromMiigx(inputBytes []byte) (Mii, error) {
	outMii := Mii{}
	if inputBytes[0]>>7 == 1 {
		return outMii, errors.New("mii is invalid")
	}

	outMii.IsFemale = ((inputBytes[0] >> 6) & 0b00000001) > 0

	if err := outMii.SetBirth((inputBytes[0]<<2)&0b11110000, ((inputBytes[0]<<6)|(inputBytes[1]>>2))>>3); err != nil {
		return outMii, err
	}

	if err := outMii.SetFavoriteColor((inputBytes[1] >> 1) & 0b00001111); err != nil {
		return outMii, err
	}

	outMii.IsFavorite = (inputBytes[1] & 0b00000001) > 0

	if err := outMii.SetNameFromUTF16BEBytes([20]byte(inputBytes[2:])); err != nil {
		return outMii, err
	}

	outMii.SetHeight(inputBytes[22])
	outMii.SetWidth(inputBytes[23])
	outMii.MiiID = uint32(inputBytes[24])<<24 | uint32(inputBytes[25])<<16 | uint32(inputBytes[26])<<8 | uint32(inputBytes[27])
	outMii.MiiCreationTimestamp = time.Unix(1136073600+int64((outMii.MiiID<<4)>>2), 0)
	outMii.ConsoleID = uint32(inputBytes[28])<<24 | uint32(inputBytes[29])<<16 | uint32(inputBytes[30])<<8 | uint32(inputBytes[31])
	outMii.FaceType = inputBytes[32] >> 5
	outMii.SkinTone = (inputBytes[32] >> 2) & 0b00000111

	if err := outMii.SetFaceFeatures(((inputBytes[32] << 2) | (inputBytes[33] >> 6)) & 0b00001111); err != nil {
		return outMii, err
	}

	outMii.Mingling = (inputBytes[33] & 0b00000100) > 0
	outMii.SourceType = mii_source_type(inputBytes[33] & 0b00000011)

	if err := outMii.SetHairType(inputBytes[34] >> 1); err != nil {
		return outMii, err
	}

	outMii.HairColor = ((inputBytes[34] << 2) | (inputBytes[35] >> 6)) & 0b00000111
	outMii.HairFlip = (inputBytes[35] & 0b00100000) > 0

	if err := outMii.SetEyebrowType(inputBytes[36] >> 3); err != nil {
		return outMii, err
	}

	if err := outMii.SetEyebrowRotation(((inputBytes[36] << 1) | (inputBytes[37] >> 7)) & 0b00000111); err != nil {
		return outMii, err
	}

	outMii.EyebrowColor = inputBytes[38] >> 5

	if err := outMii.SetEyebrowSize((inputBytes[38] >> 1) & 0b00001111); err != nil {
		return outMii, err
	}

	if err := outMii.SetEyebrowVertical(((inputBytes[38] << 4) | (inputBytes[39] >> 4)) & 0b00011111); err != nil {
		return outMii, err
	}

	if err := outMii.SetEyebrowHorizontal(inputBytes[39] & 0b00001111); err != nil {
		return outMii, err
	}

	if err := outMii.SetEyeType(inputBytes[40] >> 2); err != nil {
		return outMii, err
	}

	outMii.EyeRotation = inputBytes[41] >> 5

	if err := outMii.SetEyeVertical(inputBytes[41] & 0b00001111); err != nil {
		return outMii, err
	}

	if err := outMii.SetEyeColor(inputBytes[42] >> 5); err != nil {
		return outMii, err
	}

	outMii.EyeSize = (inputBytes[42] >> 1) & 0b00000111

	if err := outMii.SetEyeHorizontal(((inputBytes[42] << 3) | (inputBytes[43] >> 5)) & 0b00001111); err != nil {
		return outMii, err
	}

	if err := outMii.SetNoseType(inputBytes[44] >> 4); err != nil {
		return outMii, err
	}

	if err := outMii.SetNoseSize(inputBytes[44] & 0b00001111); err != nil {
		return outMii, err
	}

	if err := outMii.SetNoseVertical(inputBytes[45] >> 3); err != nil {
		return outMii, err
	}

	if err := outMii.SetMouthType(inputBytes[46] >> 3); err != nil {
		return outMii, err
	}

	if err := outMii.SetMouthColor((inputBytes[46] >> 1) & 0b00000011); err != nil {

		return outMii, err
	}
	if err := outMii.SetMouthSize(((inputBytes[46] << 3) | (inputBytes[47] >> 3)) & 0b00001111); err != nil {
		return outMii, err
	}

	if err := outMii.SetMouthVertical(inputBytes[47] & 0b00011111); err != nil {
		return outMii, err
	}

	if err := outMii.SetGlassesType(inputBytes[48] >> 4); err != nil {
		return outMii, err
	}

	if err := outMii.SetGlassesColor((inputBytes[48] >> 1) % 0b00000111); err != nil {
		return outMii, err
	}

	outMii.GlassesSize = ((inputBytes[48] << 3) | (inputBytes[49] >> 5)) & 0b00001111

	if err := outMii.SetGlassesVertical(inputBytes[49] & 0b00001111); err != nil {
		return outMii, err
	}

	outMii.FacialHairMustache = inputBytes[50] >> 6
	outMii.FacialHairBeard = (inputBytes[50] >> 4) & 0b00000011
	outMii.FacialHairColor = (inputBytes[50] >> 1) & 0b00000111

	if err := outMii.SetFacialHairSize(((inputBytes[50] << 3) | (inputBytes[51] >> 5)) & 0b00001111); err != nil {
		return outMii, err
	}

	if err := outMii.SetFacialHairVertical(inputBytes[51] & 0b00011111); err != nil {
		return outMii, err
	}

	outMii.MoleType = (inputBytes[52] >> 7) > 0

	if err := outMii.SetMoleSize((inputBytes[52] >> 3) & 0b00001111); err != nil {
		return outMii, err
	}

	if err := outMii.SetMoleVertical(((inputBytes[52] << 2) | (inputBytes[53] >> 6)) & 0b00011111); err != nil {
		return outMii, err
	}

	if err := outMii.SetMoleHorizontal((inputBytes[53] >> 1) & 0b00011111); err != nil {
		return outMii, err
	}

	if err := outMii.SetCreatorNameFromUTF16BEBytes([20]byte(inputBytes[0x36:])); err != nil {
		return outMii, err
	}

	return outMii, nil
}

type mii_source_type uint8

const (
	LOCAL mii_source_type = iota
	DOWNLOADED
	UNKNOWN_0
	UNKNOWN_1
)

type mii_fav_colors uint8

const (
	RED mii_fav_colors = iota
	ORANGE
	YELLOW
	LIME_GREEN
	FOREST_GREEN
	ROYAL_BLUE
	SKY_BLUE
	PINK
	PURPLE
	BROWN
	WHITE
	BLACK
)

type mii_birthmonth uint8

const (
	NO_BIRTHDAY mii_birthmonth = iota
	JANUARY
	FEBRUARY
	MARCH
	APRIL
	MAY
	JUNE
	JULY
	AUGUST
	SEPTEMBER
	OCTOBER
	NOVEMBER
	DECEMBER
)

func (birthmonth mii_birthmonth) ToString() string {
	switch birthmonth {
	case JANUARY:
		return "January"
	case FEBRUARY:
		return "February"
	case MARCH:
		return "March"
	case APRIL:
		return "April"
	case MAY:
		return "May"
	case JUNE:
		return "June"
	case JULY:
		return "July"
	case AUGUST:
		return "August"
	case SEPTEMBER:
		return "September"
	case OCTOBER:
		return "October"
	case NOVEMBER:
		return "November"
	case DECEMBER:
		return "December"
	default:
		return "No Birthmonth"
	}
}
