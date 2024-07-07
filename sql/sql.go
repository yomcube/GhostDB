package sql

import (
	"os"
	"strings"
)

type Config struct {
	user   string
	secret string
	databaseName string
	host string
	port   string
}

const PATH_TO_CFG_FILE = "./sql.cfg"

func writeDefaultConf(cfg *os.File) {
	def_string := "# On the line below, type in the username of the postgers user\npostgres\n# On the line below, type in the client secret\npassword\n# On the line below, type in the preferred name for the database\nGHOSTDB\n# On the line below, type in the preferred Postgres Host\nlocalhost\n# On the line below, type in the Postgres Client port\n5432"
	n, err := cfg.WriteString(def_string)
	if n != 0 && err != nil {
		panic(err)
	}
}

func SetupConfig() (Config, int) {
	// This works fine on Linux,
	// this software isn't really made for Windows after all.

	if _, err := os.Stat(PATH_TO_CFG_FILE); os.IsNotExist(err) {
		// If the file doesn't exist, write the default config to it.
		cfg, err := os.Create(PATH_TO_CFG_FILE)
		if err != nil {
			panic(err)
		}
		writeDefaultConf(cfg)
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
		user: split[1],
		secret: split[3],
		databaseName: split[5],
		host: split[7],
		port: split[9],
	}, 0
}
