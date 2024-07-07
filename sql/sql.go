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

func (cfg Config) SetupDatabase() {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", cfg.user, cfg.secret, cfg.host, cfg.port, cfg.databaseName))
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}

	hasTableGhosts := false

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
	}

	// Create Ghosts index table if it doesn't exist.
	if !hasTableGhosts {
		setupGhostIndexTable(db)
	}
}
