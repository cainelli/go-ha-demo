package config

// Config is the service configuration
type Config struct {
	Healthy   bool `json:"healthy,omitempty"`
	Ready     bool `json:"ready,omitempty"`
	Delay     int  `json:"delay,omitempty"`
	WarmingUp bool `json:"warmingup,omitempty"`
}

// Load config from files and environment
func Load() *Config {
	c := Config{
		Healthy:   true,
		Ready:     false,
		WarmingUp: false,
		Delay:     0,
	}

	return &c
}
