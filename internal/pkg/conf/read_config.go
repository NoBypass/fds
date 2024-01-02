package conf

import "github.com/magiconair/properties"

type Config struct {
	Server struct {
		Port int `properties:"port"`
	} `properties:"server"`
	Database struct {
		Host      string `properties:"host"`
		Port      int    `properties:"port"`
		User      string `properties:"user"`
		Password  string `properties:"password"`
		Namespace string `properties:"namespace"`
		Name      string `properties:"name"`
	} `properties:"db"`
	Authentication struct {
		Jwt struct {
			Secret string `properties:"secret"`
		} `properties:"jwt"`
		Bot struct {
			Password string `properties:"password"`
		} `properties:"bot"`
	} `properties:"authentication"`
}

func ReadConfig() *Config {
	var cfg Config
	p := properties.MustLoadFile("config/config.properties", properties.UTF8)
	if err := p.Decode(&cfg); err != nil {
		panic("failed to load config file " + err.Error())
	}
	return &cfg
}
