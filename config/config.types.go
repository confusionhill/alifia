package config

type Config struct {
	Port    string         `json:"port"`
	Type    string         `json:"type"`
	Servers []ServerConfig `json:"servers"`
}

type ServerConfig struct {
	Host string
}
