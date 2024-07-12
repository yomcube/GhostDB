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

// Kaitai Struct provided by HEYimHeroic: https://mii.tools/studio/ksy/mii_data_wii.ksy
// This gets out all of the Wii data, and should resemble the SQL table
// Skipped bits, or `unused_xx` according to the yaml file, are junk data.
type Mii struct { // 						Index	Len	Info
	IsFemale             bool            //	0x00.1	0.1
	BirthMonth           uint8           //	0x00.2	0.4
	BirthDay             uint8           //	0x00.6	0.5
	FavoriteColor        mii_fav_colors  //	0x01.3	0.4
	IsFavorite           bool            //	0x01.7	0.1
	MiiName              string          //	0x02	20	utf-16be encoded null-terminated-ish string with max 10 char.
	Height               uint8           //	0x16.1	0.7
	Width                uint8           //	0x17.1	0.7
	MiiID                uint32          //	0x18	4
	MiiCreationTimestamp time.Time       //	0x18.4	3.4
	ConsoleID            uint32          //	0x1C	4	the last three bytes are the last three bytes of the mac address, first byte is a checksum
	FaceType             uint8           //	0x20	0.3
	SkinTone             uint8           //	0x20.3	0.3	max = 5
	FaceFeatures         uint8           //	0x20.6	0.4	max = 11
	Mingling             bool            //	0x21.5	0.1
	SourceType           mii_source_type //	0x21.6	0.2
	HairType             uint8           //	0x22	0.7	max = 71
	HairColor            uint8           //	0x22.7	0.3
	HairFlip             bool            //	0x23.2	0.1
	EyebrowType          uint8           //	0x24	0.5	max = 23
	EyebrowRotation      uint8           //	0x24.5	0.5	max = 11
	EyebrowColor         uint8           //	0x26	0.3
	EyebrowSize          uint8           //	0x26.3	0.4 max = 8
	EyebrowVertical      uint8           //	0x26.7	0.5	max = 18
	EyebrowHorizontal    uint8           //	0x27.4	0.4	max = 12
	EyeType              uint8           //	0x28	0.6	max = 47
	EyeRotation          uint8           //	0x28.6	0.5 max = 7
	EyeVertical          uint8           //	0x29.3	0.5	max = 18
	EyeColor             uint8           //	0x2A	0.3	max = 5
	EyeSize              uint8           //	0x2A.3	0.4 max = 7
	EyeHorizontal        uint8           //	0x2A.7	0.4 max = 12
	NoseType             uint8           //	0x2C	0.4	max = 11
	NoseSize             uint8           //	0x2C.4	0.4	max = 8
	NoseVertical         uint8           //	0x2D	0.5	max = 18
	MouthType            uint8           //	0x2E	0.5	max = 23
	MouthColor           uint8           //	0x2E.5	0.2	max = 2
	MouthSize            uint8           //	0x2E.7	0.4	max = 8
	MouthVertical        uint8           //	0x2F.3	0.5	max = 18
	GlassesType          uint8           //	0x30	0.4	max = 8
	GlassesColor         uint8           //	0x30.4	0.3	max = 5
	GlassesSize          uint8           //	0x30.7	0.4	max = 7
	GlassesVertical      uint8           //	0x31.3	0.5	max = 20
	FacialHairMustache   uint8           //	0x32	0.2
	FacialHairBeard      uint8           //	0x32.2	0.2
	FacialHairColor      uint8           //	0x32.4	0.3
	FacialHairSize       uint8           //	0x32.7	0.4	max = 8
	FacialHairVertical   uint8           //	0x33.3	0.5	max = 16
	MoleType             bool            //	0x34	0.1
	MoleSize             uint8           //	0x34.1	0.4	max = 8
	MoleVertical         uint8           //	0x34.5	0.5	max = 30
	MoleHorizontal       uint8           //	0x35.2	0.5	max = 16
	CreatorName          string          //	0x36	20	utf-16be encoded null-terminated-ish string with max 10 char.
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

	outMii.IsFemale = ((inputBytes[0x00] >> 6) & 0b00000001) > 0

	if err := outMii.SetBirth((inputBytes[0x00]>>2)&0b00001111, ((inputBytes[0x00]<<6)|(inputBytes[0x01]>>2))>>3); err != nil {
		return outMii, err
	}

	if err := outMii.SetFavoriteColor((inputBytes[0x01] >> 1) & 0b00001111); err != nil {
		return outMii, err
	}

	outMii.IsFavorite = (inputBytes[0x01] & 0b00000001) > 0

	if err := outMii.SetNameFromUTF16BEBytes([20]byte(inputBytes[0x02:])); err != nil {
		return outMii, err
	}

	outMii.SetHeight(inputBytes[0x16])
	outMii.SetWidth(inputBytes[0x17])
	outMii.MiiID = uint32(inputBytes[0x18])<<24 | uint32(inputBytes[0x19])<<16 | uint32(inputBytes[0x1A])<<8 | uint32(inputBytes[0x1B])
	outMii.MiiCreationTimestamp = time.Unix(1136073600+int64((outMii.MiiID<<4)>>2), 0)
	outMii.ConsoleID = uint32(inputBytes[0x1C])<<24 | uint32(inputBytes[0x1D])<<16 | uint32(inputBytes[0x1E])<<8 | uint32(inputBytes[0x1F])
	outMii.FaceType = inputBytes[0x20] >> 5
	outMii.SkinTone = (inputBytes[0x20] >> 2) & 0b00000111

	if err := outMii.SetFaceFeatures(((inputBytes[0x20] << 2) | (inputBytes[0x21] >> 6)) & 0b00001111); err != nil {
		return outMii, err
	}

	outMii.Mingling = (inputBytes[0x21] & 0b00000100) > 0
	outMii.SourceType = mii_source_type(inputBytes[0x21] & 0b00000011)

	if err := outMii.SetHairType(inputBytes[0x22] >> 1); err != nil {
		return outMii, err
	}

	outMii.HairColor = ((inputBytes[0x22] << 2) | (inputBytes[0x23] >> 6)) & 0b00000111
	outMii.HairFlip = (inputBytes[0x23] & 0b00100000) > 0

	if err := outMii.SetEyebrowType(inputBytes[0x24] >> 3); err != nil {
		return outMii, err
	}

	if err := outMii.SetEyebrowRotation(((inputBytes[0x24] << 2) | (inputBytes[0x25] >> 6)) & 0b00001111); err != nil {
		return outMii, err
	}

	outMii.EyebrowColor = inputBytes[0x26] >> 5

	if err := outMii.SetEyebrowSize((inputBytes[0x26] >> 1) & 0b00001111); err != nil {
		return outMii, err
	}

	if err := outMii.SetEyebrowVertical(((inputBytes[0x26] << 4) | (inputBytes[0x27] >> 4)) & 0b00011111); err != nil {
		return outMii, err
	}

	if err := outMii.SetEyebrowHorizontal(inputBytes[0x27] & 0b00001111); err != nil {
		return outMii, err
	}

	if err := outMii.SetEyeType(inputBytes[0x28] >> 2); err != nil {
		return outMii, err
	}

	outMii.EyeRotation = inputBytes[0x29] >> 5

	if err := outMii.SetEyeVertical(inputBytes[0x29] & 0b00001111); err != nil {
		return outMii, err
	}

	if err := outMii.SetEyeColor(inputBytes[0x2A] >> 5); err != nil {
		return outMii, err
	}

	outMii.EyeSize = (inputBytes[0x2A] >> 1) & 0b00000111

	if err := outMii.SetEyeHorizontal(((inputBytes[0x2A] << 3) | (inputBytes[0x2B] >> 5)) & 0b00001111); err != nil {
		return outMii, err
	}

	if err := outMii.SetNoseType(inputBytes[0x2C] >> 4); err != nil {
		return outMii, err
	}

	if err := outMii.SetNoseSize(inputBytes[0x2C] & 0b00001111); err != nil {
		return outMii, err
	}

	if err := outMii.SetNoseVertical(inputBytes[0x2D] >> 3); err != nil {
		return outMii, err
	}

	if err := outMii.SetMouthType(inputBytes[0x2E] >> 3); err != nil {
		return outMii, err
	}

	if err := outMii.SetMouthColor((inputBytes[0x2E] >> 1) & 0b00000011); err != nil {

		return outMii, err
	}

	if err := outMii.SetMouthSize(((inputBytes[0x2E] << 3) | (inputBytes[0x2F] >> 5)) & 0b00001111); err != nil {
		return outMii, err
	}

	if err := outMii.SetMouthVertical(inputBytes[0x2F] & 0b00011111); err != nil {
		return outMii, err
	}

	if err := outMii.SetGlassesType(inputBytes[0x30] >> 4); err != nil {
		return outMii, err
	}

	if err := outMii.SetGlassesColor((inputBytes[0x30] >> 1) & 0b00000111); err != nil {
		return outMii, err
	}

	outMii.GlassesSize = ((inputBytes[0x30] << 3) | (inputBytes[0x31] >> 5)) & 0b00001111

	if err := outMii.SetGlassesVertical(inputBytes[0x31] & 0b00011111); err != nil {
		return outMii, err
	}

	outMii.FacialHairMustache = inputBytes[0x32] >> 6
	outMii.FacialHairBeard = (inputBytes[0x32] >> 4) & 0b00000011
	outMii.FacialHairColor = (inputBytes[0x32] >> 1) & 0b00000111

	if err := outMii.SetFacialHairSize(((inputBytes[0x32] << 3) | (inputBytes[0x33] >> 5)) & 0b00001111); err != nil {
		return outMii, err
	}

	if err := outMii.SetFacialHairVertical(inputBytes[0x33] & 0b00011111); err != nil {
		return outMii, err
	}

	outMii.MoleType = (inputBytes[0x34] >> 7) > 0

	if err := outMii.SetMoleSize((inputBytes[0x34] >> 3) & 0b00001111); err != nil {
		return outMii, err
	}

	if err := outMii.SetMoleVertical(((inputBytes[0x34] << 2) | (inputBytes[0x35] >> 6)) & 0b00011111); err != nil {
		return outMii, err
	}

	if err := outMii.SetMoleHorizontal((inputBytes[0x35] >> 1) & 0b00011111); err != nil {
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
