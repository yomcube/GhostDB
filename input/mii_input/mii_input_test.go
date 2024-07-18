package mii_input_test

import (
	"os"
	"testing"
	"time"

	"github.com/yomcube/GhostDB/input/mii_input"
)

func TestMain(t *testing.T) {
	genericMiiTest(
		mii_input.Mii{
			IsFemale:             false,
			BirthMonth:           0,
			BirthDay:             0,
			FavoriteColor:        mii_input.FOREST_GREEN,
			IsFavorite:           false,
			MiiName:              "no name",
			Height:               64,
			Weight:               64,
			MiiID:                0x80000000,
			MiiCreationTimestamp: time.Unix(1136073600, 0),
			ConsoleID:            0xECFF82D2,
			FaceType:             0,
			SkinTone:             4,
			FaceFeatures:         0,
			Mingling:             true,
			SourceType:           mii_input.LOCAL,
			HairType:             68,
			HairColor:            0,
			HairFlip:             false,
			EyebrowType:          6,
			EyebrowRotation:      6,
			EyebrowColor:         0,
			EyebrowSize:          4,
			EyebrowVertical:      10,
			EyebrowHorizontal:    2,
			EyeType:              2,
			EyeVertical:          12,
			EyeColor:             0,
			EyeSize:              4,
			EyeHorizontal:        2,
			NoseType:             1,
			NoseSize:             4,
			NoseVertical:         9,
			MouthType:            23,
			MouthColor:           0,
			MouthSize:            4,
			MouthVertical:        13,
			GlassesType:          0,
			GlassesColor:         0,
			GlassesSize:          4,
			GlassesVertical:      10,
			FacialHairMustache:   0,
			FacialHairBeard:      0,
			FacialHairColor:      0,
			FacialHairSize:       4,
			FacialHairVertical:   10,
			MoleType:             false,
			MoleSize:             4,
			MoleVertical:         20,
			MoleHorizontal:       2,
			CreatorName:          "",
		},
		"./test_mii/GuestA.miigx",
		t,
	)

	genericMiiTest(
		mii_input.Mii{
			IsFemale:             false,
			BirthMonth:           0,
			BirthDay:             0,
			FavoriteColor:        mii_input.ROYAL_BLUE,
			IsFavorite:           false,
			MiiName:              "no name",
			Height:               64,
			Weight:               64,
			MiiID:                0x80000001,
			MiiCreationTimestamp: time.Unix(1136073604, 0),
			ConsoleID:            0xECFF82D2,
			FaceType:             0,
			SkinTone:             0,
			FaceFeatures:         0,
			Mingling:             true,
			SourceType:           mii_input.LOCAL,
			HairType:             55,
			HairColor:            6,
			HairFlip:             false,
			EyebrowType:          6,
			EyebrowRotation:      6,
			EyebrowColor:         6,
			EyebrowSize:          4,
			EyebrowVertical:      10,
			EyebrowHorizontal:    2,
			EyeType:              2,
			EyeVertical:          12,
			EyeColor:             4,
			EyeSize:              4,
			EyeHorizontal:        2,
			NoseType:             1,
			NoseSize:             4,
			NoseVertical:         9,
			MouthType:            23,
			MouthColor:           0,
			MouthSize:            4,
			MouthVertical:        13,
			GlassesType:          0,
			GlassesColor:         0,
			GlassesSize:          4,
			GlassesVertical:      10,
			FacialHairMustache:   0,
			FacialHairBeard:      0,
			FacialHairColor:      0,
			FacialHairSize:       4,
			FacialHairVertical:   10,
			MoleType:             false,
			MoleSize:             4,
			MoleVertical:         20,
			MoleHorizontal:       2,
			CreatorName:          "",
		},
		"./test_mii/GuestB.miigx",
		t,
	)

	genericMiiTest(
		mii_input.Mii{
			IsFemale:             false,
			BirthMonth:           0,
			BirthDay:             0,
			FavoriteColor:        mii_input.RED,
			IsFavorite:           false,
			MiiName:              "no name",
			Height:               64,
			Weight:               64,
			MiiID:                0x80000002,
			MiiCreationTimestamp: time.Unix(1136073608, 0),
			ConsoleID:            0xECFF82D2,
			FaceType:             0,
			SkinTone:             1,
			FaceFeatures:         0,
			Mingling:             true,
			SourceType:           mii_input.LOCAL,
			HairType:             33,
			HairColor:            1,
			HairFlip:             false,
			EyebrowType:          6,
			EyebrowRotation:      6,
			EyebrowColor:         1,
			EyebrowSize:          4,
			EyebrowVertical:      10,
			EyebrowHorizontal:    2,
			EyeType:              2,
			EyeVertical:          12,
			EyeColor:             0,
			EyeSize:              4,
			EyeHorizontal:        2,
			NoseType:             1,
			NoseSize:             4,
			NoseVertical:         9,
			MouthType:            23,
			MouthColor:           0,
			MouthSize:            4,
			MouthVertical:        13,
			GlassesType:          0,
			GlassesColor:         0,
			GlassesSize:          4,
			GlassesVertical:      10,
			FacialHairMustache:   0,
			FacialHairBeard:      0,
			FacialHairColor:      0,
			FacialHairSize:       4,
			FacialHairVertical:   10,
			MoleType:             false,
			MoleSize:             4,
			MoleVertical:         20,
			MoleHorizontal:       2,
			CreatorName:          "",
		},
		"./test_mii/GuestC.miigx",
		t,
	)

	genericMiiTest(
		mii_input.Mii{
			IsFemale:             true,
			BirthMonth:           0,
			BirthDay:             0,
			FavoriteColor:        mii_input.YELLOW,
			IsFavorite:           false,
			MiiName:              "no name",
			Height:               64,
			Weight:               64,
			MiiID:                0x80000003,
			MiiCreationTimestamp: time.Unix(1136073612, 0),
			ConsoleID:            0xECFF82D2,
			FaceType:             0,
			SkinTone:             2,
			FaceFeatures:         0,
			Mingling:             true,
			SourceType:           mii_input.LOCAL,
			HairType:             24,
			HairColor:            0,
			HairFlip:             false,
			EyebrowType:          0,
			EyebrowRotation:      6,
			EyebrowColor:         0,
			EyebrowSize:          4,
			EyebrowVertical:      10,
			EyebrowHorizontal:    2,
			EyeType:              4,
			EyeVertical:          12,
			EyeColor:             0,
			EyeSize:              4,
			EyeHorizontal:        2,
			NoseType:             1,
			NoseSize:             4,
			NoseVertical:         9,
			MouthType:            23,
			MouthColor:           0,
			MouthSize:            4,
			MouthVertical:        13,
			GlassesType:          0,
			GlassesColor:         0,
			GlassesSize:          4,
			GlassesVertical:      10,
			FacialHairMustache:   0,
			FacialHairBeard:      0,
			FacialHairColor:      0,
			FacialHairSize:       4,
			FacialHairVertical:   10,
			MoleType:             false,
			MoleSize:             4,
			MoleVertical:         20,
			MoleHorizontal:       2,
			CreatorName:          "",
		},
		"./test_mii/GuestD.miigx",
		t,
	)

	genericMiiTest(
		mii_input.Mii{
			IsFemale:             true,
			BirthMonth:           0,
			BirthDay:             0,
			FavoriteColor:        mii_input.SKY_BLUE,
			IsFavorite:           false,
			MiiName:              "no name",
			Height:               64,
			Weight:               64,
			MiiID:                0x80000004,
			MiiCreationTimestamp: time.Unix(1136073616, 0),
			ConsoleID:            0xECFF82D2,
			FaceType:             0,
			SkinTone:             0,
			FaceFeatures:         0,
			Mingling:             true,
			SourceType:           mii_input.LOCAL,
			HairType:             14,
			HairColor:            7,
			HairFlip:             false,
			EyebrowType:          0,
			EyebrowRotation:      6,
			EyebrowColor:         7,
			EyebrowSize:          4,
			EyebrowVertical:      10,
			EyebrowHorizontal:    2,
			EyeType:              4,
			EyeVertical:          12,
			EyeColor:             5,
			EyeSize:              4,
			EyeHorizontal:        2,
			NoseType:             1,
			NoseSize:             4,
			NoseVertical:         9,
			MouthType:            23,
			MouthColor:           0,
			MouthSize:            4,
			MouthVertical:        13,
			GlassesType:          0,
			GlassesColor:         0,
			GlassesSize:          4,
			GlassesVertical:      10,
			FacialHairMustache:   0,
			FacialHairBeard:      0,
			FacialHairColor:      0,
			FacialHairSize:       4,
			FacialHairVertical:   10,
			MoleType:             false,
			MoleSize:             4,
			MoleVertical:         20,
			MoleHorizontal:       2,
			CreatorName:          "",
		},
		"./test_mii/GuestE.miigx",
		t,
	)

	genericMiiTest(
		mii_input.Mii{
			IsFemale:             true,
			BirthMonth:           0,
			BirthDay:             0,
			FavoriteColor:        mii_input.PINK,
			IsFavorite:           false,
			MiiName:              "no name",
			Height:               64,
			Weight:               64,
			MiiID:                0x80000005,
			MiiCreationTimestamp: time.Unix(1136073620, 0),
			ConsoleID:            0xECFF82D2,
			FaceType:             0,
			SkinTone:             0,
			FaceFeatures:         0,
			Mingling:             true,
			SourceType:           mii_input.LOCAL,
			HairType:             12,
			HairColor:            1,
			HairFlip:             false,
			EyebrowType:          0,
			EyebrowRotation:      6,
			EyebrowColor:         1,
			EyebrowSize:          4,
			EyebrowVertical:      10,
			EyebrowHorizontal:    2,
			EyeType:              4,
			EyeVertical:          12,
			EyeColor:             0,
			EyeSize:              4,
			EyeHorizontal:        2,
			NoseType:             1,
			NoseSize:             4,
			NoseVertical:         9,
			MouthType:            23,
			MouthColor:           0,
			MouthSize:            4,
			MouthVertical:        13,
			GlassesType:          0,
			GlassesColor:         0,
			GlassesSize:          4,
			GlassesVertical:      10,
			FacialHairMustache:   0,
			FacialHairBeard:      0,
			FacialHairColor:      0,
			FacialHairSize:       4,
			FacialHairVertical:   10,
			MoleType:             false,
			MoleSize:             4,
			MoleVertical:         20,
			MoleHorizontal:       2,
			CreatorName:          "",
		},
		"./test_mii/GuestF.miigx",
		t,
	)

	genericMiiTestFromRKG(
		mii_input.Mii{
			IsFemale:             true,
			BirthMonth:           9,
			BirthDay:             23,
			FavoriteColor:        mii_input.FOREST_GREEN,
			IsFavorite:           true,
			MiiName:              "sy FαlB",
			Height:               104,
			Weight:               66,
			MiiID:                0x86198DBA,
			MiiCreationTimestamp: time.Unix(1545425512, 0),
			ConsoleID:            0x6A24E862,
			FaceType:             1,
			SkinTone:             0,
			FaceFeatures:         0,
			Mingling:             false,
			SourceType:           mii_input.LOCAL,
			HairType:             8,
			HairColor:            1,
			HairFlip:             false,
			EyebrowType:          6,
			EyebrowRotation:      5,
			EyebrowColor:         1,
			EyebrowSize:          4,
			EyebrowVertical:      10,
			EyebrowHorizontal:    2,
			EyeType:              15,
			EyeVertical:          13,
			EyeColor:             0,
			EyeSize:              3,
			EyeHorizontal:        2,
			NoseType:             1,
			NoseSize:             4,
			NoseVertical:         8,
			MouthType:            18,
			MouthColor:           0,
			MouthSize:            2,
			MouthVertical:        13,
			GlassesType:          3,
			GlassesColor:         0,
			GlassesSize:          4,
			GlassesVertical:      11,
			FacialHairMustache:   0,
			FacialHairBeard:      0,
			FacialHairColor:      0,
			FacialHairSize:       4,
			FacialHairVertical:   10,
			MoleType:             false,
			MoleSize:             4,
			MoleVertical:         20,
			MoleHorizontal:       2,
			CreatorName:          "FαlB",
		},
		"../rkg_input/test_rkg/falb_ct.rkg",
		t,
	)
}

func compareTwoValues[T comparable](constructVal T, testVal T, filePath string, valName string, t *testing.T) {
	if testVal != constructVal {
		t.Errorf("%v: constructedMii.InitializeFromMiigx() error: constructedMii.%v is wrong, extracted value is %v, expected %v", filePath, valName, constructVal, testVal)
	}
}

func genericMiiTestFromRKG(testMii mii_input.Mii, filePath string, t *testing.T) {
	byteData, err := os.ReadFile(filePath)
	byteData = byteData[0x3C : 0x4A+0x3C]
	if err != nil {
		t.Errorf("os.ReadFile() error: %v", err)
		panic(err)
	}

	constructedMii, err := mii_input.InitializeFromMiigx(byteData)
	if err != nil {
		t.Errorf("%v: constructedMii.InitializeFromMiigx() error: %v", filePath, err)
		panic(err)
	}

	compareTwoValues(constructedMii.IsFemale, testMii.IsFemale, filePath, "IsFemale", t)
	compareTwoValues(constructedMii.BirthMonth, testMii.BirthMonth, filePath, "BirthMonth", t)
	compareTwoValues(constructedMii.BirthDay, testMii.BirthDay, filePath, "BirthDay", t)
	compareTwoValues(constructedMii.FavoriteColor, testMii.FavoriteColor, filePath, "FavoriteColor", t)
	compareTwoValues(constructedMii.IsFavorite, testMii.IsFavorite, filePath, "IsFavorite", t)
	compareTwoValues(constructedMii.MiiName, testMii.MiiName, filePath, "MiiName", t)
	compareTwoValues(constructedMii.Height, testMii.Height, filePath, "Height", t)
	compareTwoValues(constructedMii.Weight, testMii.Weight, filePath, "Weight", t)
	compareTwoValues(constructedMii.MiiID, testMii.MiiID, filePath, "MiiID", t)
	compareTwoValues(constructedMii.MiiCreationTimestamp, testMii.MiiCreationTimestamp, filePath, "MiiCreationTimestamp", t)
	compareTwoValues(constructedMii.ConsoleID, testMii.ConsoleID, filePath, "ConsoleID", t)
	compareTwoValues(constructedMii.FaceType, testMii.FaceType, filePath, "FaceType", t)
	compareTwoValues(constructedMii.SkinTone, testMii.SkinTone, filePath, "SkinTone", t)
	compareTwoValues(constructedMii.FaceFeatures, testMii.FaceFeatures, filePath, "FaceFeatures", t)
	compareTwoValues(constructedMii.Mingling, testMii.Mingling, filePath, "Mingling", t)
	compareTwoValues(constructedMii.SourceType, testMii.SourceType, filePath, "SourceType", t)
	compareTwoValues(constructedMii.HairType, testMii.HairType, filePath, "HairType", t)
	compareTwoValues(constructedMii.HairColor, testMii.HairColor, filePath, "HairColor", t)
	compareTwoValues(constructedMii.HairFlip, testMii.HairFlip, filePath, "HairFlip", t)
	compareTwoValues(constructedMii.EyebrowType, testMii.EyebrowType, filePath, "EyebrowType", t)
	compareTwoValues(constructedMii.EyebrowRotation, testMii.EyebrowRotation, filePath, "EyebrowRotation", t)
	compareTwoValues(constructedMii.EyebrowColor, testMii.EyebrowColor, filePath, "EyebrowColor", t)
	compareTwoValues(constructedMii.EyebrowSize, testMii.EyebrowSize, filePath, "EyebrowSize", t)
	compareTwoValues(constructedMii.EyebrowVertical, testMii.EyebrowVertical, filePath, "EyebrowVertical", t)
	compareTwoValues(constructedMii.EyebrowHorizontal, testMii.EyebrowHorizontal, filePath, "EyebrowHorizontal", t)
	compareTwoValues(constructedMii.EyeType, testMii.EyeType, filePath, "EyeType", t)
	compareTwoValues(constructedMii.EyeVertical, testMii.EyeVertical, filePath, "EyeVertical", t)
	compareTwoValues(constructedMii.EyeColor, testMii.EyeColor, filePath, "EyeColor", t)
	compareTwoValues(constructedMii.EyeSize, testMii.EyeSize, filePath, "EyeSize", t)
	compareTwoValues(constructedMii.EyeHorizontal, testMii.EyeHorizontal, filePath, "EyeHorizontal", t)
	compareTwoValues(constructedMii.NoseType, testMii.NoseType, filePath, "NoseType", t)
	compareTwoValues(constructedMii.NoseSize, testMii.NoseSize, filePath, "NoseSize", t)
	compareTwoValues(constructedMii.NoseVertical, testMii.NoseVertical, filePath, "NoseVertical", t)
	compareTwoValues(constructedMii.MouthType, testMii.MouthType, filePath, "MouthType", t)
	compareTwoValues(constructedMii.MouthColor, testMii.MouthColor, filePath, "MouthColor", t)
	compareTwoValues(constructedMii.MouthSize, testMii.MouthSize, filePath, "MouthSize", t)
	compareTwoValues(constructedMii.MouthVertical, testMii.MouthVertical, filePath, "MouthVertical", t)
	compareTwoValues(constructedMii.GlassesType, testMii.GlassesType, filePath, "GlassesType", t)
	compareTwoValues(constructedMii.GlassesColor, testMii.GlassesColor, filePath, "GlassesColor", t)
	compareTwoValues(constructedMii.GlassesSize, testMii.GlassesSize, filePath, "GlassesSize", t)
	compareTwoValues(constructedMii.GlassesVertical, testMii.GlassesVertical, filePath, "GlassesVertical", t)
	compareTwoValues(constructedMii.FacialHairMustache, testMii.FacialHairMustache, filePath, "FacialHairMustache", t)
	compareTwoValues(constructedMii.FacialHairBeard, testMii.FacialHairBeard, filePath, "FacialHairBeard", t)
	compareTwoValues(constructedMii.FacialHairColor, testMii.FacialHairColor, filePath, "FacialHairColor", t)
	compareTwoValues(constructedMii.FacialHairSize, testMii.FacialHairSize, filePath, "FacialHairSize", t)
	compareTwoValues(constructedMii.FacialHairVertical, testMii.FacialHairVertical, filePath, "FacialHairVertical", t)
	compareTwoValues(constructedMii.MoleType, testMii.MoleType, filePath, "MoleType", t)
	compareTwoValues(constructedMii.MoleSize, testMii.MoleSize, filePath, "MoleSize", t)
	compareTwoValues(constructedMii.MoleVertical, testMii.MoleVertical, filePath, "MoleVertical", t)
	compareTwoValues(constructedMii.MoleHorizontal, testMii.MoleHorizontal, filePath, "MoleHorizontal", t)
	compareTwoValues(constructedMii.CreatorName, testMii.CreatorName, filePath, "CreatorName", t)
}

func genericMiiTest(testMii mii_input.Mii, filePath string, t *testing.T) {
	byteData, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("os.ReadFile() error: %v", err)
		panic(err)
	}

	constructedMii, err := mii_input.InitializeFromMiigx(byteData)
	if err != nil {
		t.Errorf("%v: constructedMii.InitializeFromMiigx() error: %v", filePath, err)
		panic(err)
	}

	compareTwoValues(constructedMii.IsFemale, testMii.IsFemale, filePath, "IsFemale", t)
	compareTwoValues(constructedMii.BirthMonth, testMii.BirthMonth, filePath, "BirthMonth", t)
	compareTwoValues(constructedMii.BirthDay, testMii.BirthDay, filePath, "BirthDay", t)
	compareTwoValues(constructedMii.FavoriteColor, testMii.FavoriteColor, filePath, "FavoriteColor", t)
	compareTwoValues(constructedMii.IsFavorite, testMii.IsFavorite, filePath, "IsFavorite", t)
	compareTwoValues(constructedMii.MiiName, testMii.MiiName, filePath, "MiiName", t)
	compareTwoValues(constructedMii.Height, testMii.Height, filePath, "Height", t)
	compareTwoValues(constructedMii.Weight, testMii.Weight, filePath, "Weight", t)
	compareTwoValues(constructedMii.MiiID, testMii.MiiID, filePath, "MiiID", t)
	compareTwoValues(constructedMii.MiiCreationTimestamp, testMii.MiiCreationTimestamp, filePath, "MiiCreationTimestamp", t)
	compareTwoValues(constructedMii.ConsoleID, testMii.ConsoleID, filePath, "ConsoleID", t)
	compareTwoValues(constructedMii.FaceType, testMii.FaceType, filePath, "FaceType", t)
	compareTwoValues(constructedMii.SkinTone, testMii.SkinTone, filePath, "SkinTone", t)
	compareTwoValues(constructedMii.FaceFeatures, testMii.FaceFeatures, filePath, "FaceFeatures", t)
	compareTwoValues(constructedMii.Mingling, testMii.Mingling, filePath, "Mingling", t)
	compareTwoValues(constructedMii.SourceType, testMii.SourceType, filePath, "SourceType", t)
	compareTwoValues(constructedMii.HairType, testMii.HairType, filePath, "HairType", t)
	compareTwoValues(constructedMii.HairColor, testMii.HairColor, filePath, "HairColor", t)
	compareTwoValues(constructedMii.HairFlip, testMii.HairFlip, filePath, "HairFlip", t)
	compareTwoValues(constructedMii.EyebrowType, testMii.EyebrowType, filePath, "EyebrowType", t)
	compareTwoValues(constructedMii.EyebrowRotation, testMii.EyebrowRotation, filePath, "EyebrowRotation", t)
	compareTwoValues(constructedMii.EyebrowColor, testMii.EyebrowColor, filePath, "EyebrowColor", t)
	compareTwoValues(constructedMii.EyebrowSize, testMii.EyebrowSize, filePath, "EyebrowSize", t)
	compareTwoValues(constructedMii.EyebrowVertical, testMii.EyebrowVertical, filePath, "EyebrowVertical", t)
	compareTwoValues(constructedMii.EyebrowHorizontal, testMii.EyebrowHorizontal, filePath, "EyebrowHorizontal", t)
	compareTwoValues(constructedMii.EyeType, testMii.EyeType, filePath, "EyeType", t)
	compareTwoValues(constructedMii.EyeVertical, testMii.EyeVertical, filePath, "EyeVertical", t)
	compareTwoValues(constructedMii.EyeColor, testMii.EyeColor, filePath, "EyeColor", t)
	compareTwoValues(constructedMii.EyeSize, testMii.EyeSize, filePath, "EyeSize", t)
	compareTwoValues(constructedMii.EyeHorizontal, testMii.EyeHorizontal, filePath, "EyeHorizontal", t)
	compareTwoValues(constructedMii.NoseType, testMii.NoseType, filePath, "NoseType", t)
	compareTwoValues(constructedMii.NoseSize, testMii.NoseSize, filePath, "NoseSize", t)
	compareTwoValues(constructedMii.NoseVertical, testMii.NoseVertical, filePath, "NoseVertical", t)
	compareTwoValues(constructedMii.MouthType, testMii.MouthType, filePath, "MouthType", t)
	compareTwoValues(constructedMii.MouthColor, testMii.MouthColor, filePath, "MouthColor", t)
	compareTwoValues(constructedMii.MouthSize, testMii.MouthSize, filePath, "MouthSize", t)
	compareTwoValues(constructedMii.MouthVertical, testMii.MouthVertical, filePath, "MouthVertical", t)
	compareTwoValues(constructedMii.GlassesType, testMii.GlassesType, filePath, "GlassesType", t)
	compareTwoValues(constructedMii.GlassesColor, testMii.GlassesColor, filePath, "GlassesColor", t)
	compareTwoValues(constructedMii.GlassesSize, testMii.GlassesSize, filePath, "GlassesSize", t)
	compareTwoValues(constructedMii.GlassesVertical, testMii.GlassesVertical, filePath, "GlassesVertical", t)
	compareTwoValues(constructedMii.FacialHairMustache, testMii.FacialHairMustache, filePath, "FacialHairMustache", t)
	compareTwoValues(constructedMii.FacialHairBeard, testMii.FacialHairBeard, filePath, "FacialHairBeard", t)
	compareTwoValues(constructedMii.FacialHairColor, testMii.FacialHairColor, filePath, "FacialHairColor", t)
	compareTwoValues(constructedMii.FacialHairSize, testMii.FacialHairSize, filePath, "FacialHairSize", t)
	compareTwoValues(constructedMii.FacialHairVertical, testMii.FacialHairVertical, filePath, "FacialHairVertical", t)
	compareTwoValues(constructedMii.MoleType, testMii.MoleType, filePath, "MoleType", t)
	compareTwoValues(constructedMii.MoleSize, testMii.MoleSize, filePath, "MoleSize", t)
	compareTwoValues(constructedMii.MoleVertical, testMii.MoleVertical, filePath, "MoleVertical", t)
	compareTwoValues(constructedMii.MoleHorizontal, testMii.MoleHorizontal, filePath, "MoleHorizontal", t)
	compareTwoValues(constructedMii.CreatorName, testMii.CreatorName, filePath, "CreatorName", t)
}
