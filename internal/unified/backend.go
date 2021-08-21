// Package unified provides a backend for go-smtp which allows multiple users from unified sources.
package unified

import (
	"errors"
	"github.com/emersion/go-smtp"
	"github.com/shauncampbell/unified-smtp/pkg/authenticator"
)

// Unified creates a unified backend for go-smtp.
type Unified struct {
	auth  authenticator.Authenticator
	smtp.Backend
}

func (u *Unified) NewSession(_ smtp.ConnectionState, _ string) (smtp.Session, error) {
	return &Session{auth: u.auth}, nil
}

func (u *Unified) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	if !u.auth.Authenticate(username, password) {
		return nil, errors.New("invalid username or password")
	}
	return &Session{auth: u.auth}, nil
}

func (u *Unified) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return nil, smtp.ErrAuthRequired
}

// New creates a new unified backend.
func New(auth authenticator.Authenticator) smtp.Backend {
	return &Unified{auth: auth}
}
