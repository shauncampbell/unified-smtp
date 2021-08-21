// Package authenticator provides an interface for authenticating a user
package authenticator

// Authenticator is an interface which allows a user to authenticate.
type Authenticator interface {
	Authenticate(username, password string) bool
}
