package sql

import (
	"github.com/yomcube/GhostDB/common"
	"github.com/yomcube/GhostDB/input/mii_input"

	"time"
)

type uuid [16]uint8

type SQLGhost struct {
	UUID        uuid                `json:"uuid"`
	PlayerID    uuid                `json:"player_id"`
	MiiID       uuid                `json:"mii_id"`
	TrackUUID   uuid                `json:"track_uuid"`
	Date        time.Time           `json:"date_set"`
	CharacterID common.CharacterID  `json:"character_id"`
	VehicleID   common.VehicleID    `json:"vehicle_id"`
	FinishTime  uint32              `json:"finish_time"` // Total milliseconds
	Controller  common.ControllerID `json:"controller"`
}

type SQLPlayer struct {
	UUID         uuid               `json:"uuid"`
	PlayerName   string             `json:"player_name"`
	MiiIDs       []uuid             `json:"mii_ids_array"`
	RegionID     common.CountryCode `json:"region"`
	ProvinceID   byte               `json:"province"`
	LastModified time.Time          `json:"last_modified"`
}

type SQLMii struct {
	UUID      uuid          `json:"uuid"`
	Mii       mii_input.Mii `json:"mii"`
	RenderStr string        `json:"render_string"`
}
