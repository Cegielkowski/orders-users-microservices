package config

type Config struct {
	Databases struct {
		Mysql struct {
			Host     string `json:"host"`
			User     string `json:"user"`
			Port     int    `json:"port"`
			Password string `json:"password"`
			Database string `json:"database"`
		} `json:"mysql"`
		Use string `json:"use"`
	} `json:"databases"`
	Verbose bool `json:"verbose"`
}

// CONFIG Global configuration variable.
var CONFIG Config