package state

var (
	Users *UserList
	Conf  *Config
)

func Init(configFile string) error {
	// initialize UserList and Config structs
	users := UserList{
		Users: map[string]User{},
	}
	config := Config{}

	// Parse configuration file
	if err := config.ParseConfig(configFile); err != nil {
		return err
	}

	// Initialize user list
	if err := users.ReadDatabase(config.Files.UsersDb); err != nil {
		return err
	}

	Users = &users
	Conf = &config

	return nil
}
