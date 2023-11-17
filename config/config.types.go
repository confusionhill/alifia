package config

type Config struct {
	Servers []ServerConfig `json:"servers"`
	Type    string         `json:"type"`
}

type ServerConfig struct {
	Host string
}
