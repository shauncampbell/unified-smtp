// Package memory provides an in memory authenticator.
package memory

// Authenticator provides an in memory authenticator.
type Authenticator struct {
	allowedUsers map[string]string
}

// Authenticate checks whether or not a user is authenticated.
func (a *Authenticator) Authenticate(username, password string) bool {
	return a.allowedUsers[username] == password
}

// New creates a new in memory authenticator
func New(config *Config) *Authenticator {
	return &Authenticator{allowedUsers: config.AllowedUsers}
}
