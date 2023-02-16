package model

var GlobalConfig = &Config{}

type Config struct {
	Postgres DbConfig `yaml:"postgres"`
}

type DbConfig struct {
	Host string `json:"host" yaml:"host"`
	Port string `json:"port" yaml:"port"`
	User string `json:"user" yaml:"user"`
	Pass string `json:"pass" yaml:"pass"`
}
