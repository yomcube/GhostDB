package sql

import (
	"github.com/yomcube/GhostDB/common"
	"github.com/yomcube/GhostDB/input/mii_input"

	"time"
)

type uuid [16]uint8

type SQLGhost struct {
	UUID uuid
	PlayerID uuid
	MiiID uuid
	Date time.Time
	IsCT bool
	CourseID common.CourseID
	CharacterID common.CharacterID
	VehicleID common.VehicleID
	FinishTime uint32 // Total milliseconds
	Controller common.ControllerID
}

type SQLPlayer struct {
	UUID uuid
	PlayerName string
	MiiIDs []uuid
	RegionID common.CountryCode
	ProvinceID byte
	LastModified time.Time
}

type SQLMii struct {
	UUID uuid
	Mii mii_input.Mii
	RenderStr string
}
