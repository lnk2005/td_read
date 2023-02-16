package model

type Config struct {
	Host string `json:"host" yaml:"host"`
	Port string `json:"port" yaml:"port"`
	User string `json:"user" yaml:"user"`
	Pass string `json:"pass" yaml:"pass"`
}
