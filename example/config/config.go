package config

type Database struct {
	Host string `mapstructure:"host" description:"host" default:"localhost"`
	Port int    `mapstructure:"port" description:"port" default:"3306"`
	User string `mapstructure:"user" description:"user" default:"root"`
	Pass string `mapstructure:"pass" description:"pass"`
}

type Server struct {
	Port int `mapstructure:"port" description:"port" default:"8080"`
}

type AppConfig struct {
	Debug    bool   `mapstructure:"debug" short:"d" description:"debug mode" default:"false"`
	Timezone string `mapstructure:"timezone" description:"timezone" default:"UTC"`

	Server Server `mapstructure:"server" group:"server"`

	Database map[string]Database `mapstructure:"database" group:"database" mapkey:"<dbkey>"`
}
