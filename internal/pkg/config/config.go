package config

// Config is the service configuration
type Config struct {
	Healthy   bool `json:"healthy,omitempty"`
	Ready     bool `json:"ready,omitempty"`
	WarmingUp bool `json:"warmingup,omitempty"`
}

// New default config
func New() *Config {
	c := Config{
		Healthy:   true,
		Ready:     false,
		WarmingUp: false,
	}

	return &c
}
