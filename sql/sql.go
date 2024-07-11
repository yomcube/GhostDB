package sql

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type Config struct {
	user         string
	secret       string
	databaseName string
	host         string
	port         string
}

type tables_data struct {
	table_schema *string // This is the same as Option<String>
	table_name   *string
}

const PATH_TO_CFG_FILE = "./sql.cfg"

func writeDefaultConf(cfgFile *os.File) {
	def_string := "# On the line below, type in the username of the postgers user\npostgres\n# On the line below, type in the client secret\npassword\n# On the line below, type in the preferred name for the database\nghostdb\n# On the line below, type in the preferred Postgres Host\nlocalhost\n# On the line below, type in the Postgres Client port\n5432"
	n, err := cfgFile.WriteString(def_string)
	if n != 0 && err != nil {
		panic(err)
	}
}

func SetupConfig() (Config, int) {
	// This works fine on Linux,
	// this software isn't really made for Windows after all.

	if _, err := os.Stat(PATH_TO_CFG_FILE); os.IsNotExist(err) {
		// If the file doesn't exist, write the default config to it.
		cfgFile, err := os.Create(PATH_TO_CFG_FILE)
		if err != nil {
			panic(err)
		}
		writeDefaultConf(cfgFile)
		// Then close the program
		return Config{}, 1
	}

	buf, err := os.ReadFile(PATH_TO_CFG_FILE)
	if err != nil {
		panic(err)
	}
	stringBuf := string(buf)

	split := strings.Split(stringBuf, "\n")

	return Config{
		user:         split[1],
		secret:       split[3],
		databaseName: split[5],
		host:         split[7],
		port:         split[9],
	}, 0
}

func setupGhostIndexTable(db *sql.DB) {
	// Code to create table here, structs haven't been completely decided
}

func setupPlayerTable(db *sql.DB) {
	_, err := db.Query(`
		CREATE TABLE public.players (
			uuid uuid NOT NULL DEFAULT gen_random_uuid (),
			player_name character varying(30) NOT NULL,
			mii_ids_array uuid[] NOT NULL,
			"regionID" character(1) NOT NULL DEFAULT 255,
			"provinceID" character(1) NOT NULL DEFAULT 255,
			last_modified timestamp without time zone NOT NULL
		);
	`)

	if err != nil {
		panic(err)
	}
}

func setupMiiTable(db *sql.DB) {
	_, err := db.Query(`
		CREATE TABLE public.miis (
			uuid uuid NOT NULL DEFAULT gen_random_uuid (),
			is_invalid bool NOT NULL DEFAULT false,
			is_girl bool NOT NULL DEFAULT false,
			birth_month character(1) NOT NULL,
			birth_day character(1) NOT NULL,
			favorite_color character(1) NOT NULL,
			is_favorite bool NOT NULL DEFAULT false,
			mii_name character varying(20) NOT NULL,
			body_height character(1) NOT NULL,
			body_weight character(1) NOT NULL,
			mii_id INT NOT NULL,
			creation_timestamp TIMESTAMP WITHOUT TIME ZONE NOT NULL,
			console_mac_address INT NOT NULL,
			face_type character(1) NOT NULL,
			face_color character(1) NOT NULL,
			facial_feature character(1) NOT NULL,
			hair_type character(1) NOT NULL,
			hair_color character(1) NOT NULL,
			hair_flip bool NOT NULL DEFAULT false,
			eyebrow_type character(1) NOT NULL,
			eyebrow_rotation character(1) NOT NULL,
			eyebrow_color character(1) NOT NULL,
			eyebrow_size character(1) NOT NULL,
			eyebrow_vertical character(1) NOT NULL,
			eyebrow_horizontal character(1) NOT NULL,
			eye_type character(1) NOT NULL,
			eye_rotation character(1) NOT NULL,
			eye_vertical character(1) NOT NULL,
			eye_color character(1) NOT NULL,
			eye_horizontal character(1) NOT NULL,
			nose_type character(1) NOT NULL,
			nose_size character(1) NOT NULL,
			nose_vertical character(1) NOT NULL,
			mouth_type character(1) NOT NULL,
			mouth_color character(1) NOT NULL,
			mouth_size character(1) NOT NULL,
			mouth_vertical character(1) NOT NULL,
			glasses_type character(1) NOT NULL,
			glasses_color character(1) NOT NULL,
			glasses_size character(1) NOT NULL,
			glasses_vertical character(1) NOT NULL,
			facial_hair_mustache character(1) NOT NULL,
			facial_hair_beard character(1) NOT NULL,
			facial_hair_color character(1) NOT NULL,
			facial_hair_size character(1) NOT NULL,
			facial_hair_vertical character(1) NOT NULL,
			mole_enable bool NOT NULL,
			mole_size character(1) NOT NULL,
			mole_vertical character(1) NOT NULL,
			mole_horizontal character(1) NOT NULL,
			creator_name character varying(20) NOT NULL,
			utf8_display_url character(94) NOT NULL
		);
	`)

	if err != nil {
		panic(err)
	}
}

func (cfg Config) SetupDatabase() {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", cfg.user, cfg.secret, cfg.host, cfg.port, cfg.databaseName))
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}

	hasTableGhosts := false
	hasTablePlayers := false
	hasTableMiis := false

	// Check if the correct tables exist
	tables, err := db.Query("SELECT table_schema, table_name FROM information_schema.tables WHERE table_type = 'BASE TABLE' AND table_schema NOT IN ('pg_catalog', 'information_schema');")
	if err = db.Ping(); err != nil {
		panic(err)
	}
	for tables.Next() {
		row := tables_data{}
		err := tables.Scan(&row.table_schema, &row.table_name)
		if err != nil {
			panic(err)
		}
		// Check for Ghosts index table.
		if *row.table_name == "ghost_data" && *row.table_schema == "public" {
			hasTableGhosts = true
		}
		if *row.table_name == "players" && *row.table_schema == "public" {
			hasTablePlayers = true
		}
		if *row.table_name == "miis" && *row.table_schema == "public" {
			hasTableMiis = true
		}
	}

	// Create Tables that don't exist yet.
	if !hasTableGhosts {
		setupGhostIndexTable(db)
	}

	if !hasTablePlayers {
		setupPlayerTable(db)
	}

	if !hasTableMiis {
		setupMiiTable(db)
	}
}
