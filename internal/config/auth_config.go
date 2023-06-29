package config

type AuthConfig struct {
	DatabaseURL   string
	DebugDatabase bool
	Port          int
}

func New() *AuthConfig {
	return &AuthConfig{}
}
