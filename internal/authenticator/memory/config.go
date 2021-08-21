package memory

// Config represents the configuration options for the in memory authenticator
type Config struct {
	AllowedUsers map[string]string `yaml:"allowed_users"`
}
