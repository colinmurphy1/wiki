package state

import (
	"path/filepath"
)

var (
	Users *UserList
	Conf  *Config

	usersDbPath string // Path to the users database
)

// Constants
const (
	CONFIG_FILE   = "config.yaml"
	USERS_DB_FILE = "users.db"
)

func Init(wikiDir string) error {
	// initialize UserList and Config structs
	users := UserList{
		Users: map[string]User{},
	}
	config := Config{}

	// Get absolute path to wiki base directory
	baseDir, _ := filepath.Abs(wikiDir)

	// Parse configuration file
	if err := config.ParseConfig(baseDir + "/" + CONFIG_FILE); err != nil {
		return err
	}

	// Export configuration from state package
	Conf = &config

	// Set absolute paths for document root and base root
	Conf.Files.BaseDir = baseDir
	Conf.Wiki.DocumentRoot = baseDir + "/" + Conf.Wiki.DocumentRoot
	usersDbPath = Conf.Files.BaseDir + "/" + USERS_DB_FILE

	// Initialize user list
	if err := users.ReadDatabase(usersDbPath); err != nil {
		return err
	}

	// Export users
	Users = &users

	return nil
}
