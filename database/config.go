package database

type DbConfig struct {
	Host         string
	Name         string
	Port         string
	User         string
	Password     string
	AppUser      string
	AppPassword  string
	RunBaseLine  string
	StartupLocal string // Startup local database when app starts
}

func NewDbConfig(
	host string,
	name string,
	port string,
	user string,
	password string,
	appUser string,
	appPassword string,
	runBaseLine string,
	startupLocal string,
) *DbConfig {
	return &DbConfig{
		Host:         host,
		Name:         name,
		Port:         port,
		User:         user,
		Password:     password,
		AppUser:      appUser,
		AppPassword:  appPassword,
		RunBaseLine:  runBaseLine,
		StartupLocal: startupLocal,
	}
}
