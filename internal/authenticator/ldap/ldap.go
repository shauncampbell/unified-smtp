package ldap

import (
	"crypto/tls"
	"fmt"

	ldap "github.com/go-ldap/ldap/v3"
	"github.com/rs/zerolog/log"
)

// Authenticator provides an ldap authenticator.
type Authenticator struct {
	cfg          *Config
	allowedUsers map[string]string
}

// dial connects to the ldap server.
func (a *Authenticator) dial() (*ldap.Conn, error) {
	if a.cfg.Protocol == "ldaps" {
		return ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port), a.getTLSConfig())
	}
	return ldap.Dial("tcp", fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port))
}

// Authenticate checks whether or not a user is authenticated.
func (a *Authenticator) Authenticate(username, password string) bool {
	var l *ldap.Conn
	var err error

	l, err = a.dial()

	if err != nil {
		log.Error().Msgf("failed to connect to ldap server: %s", err.Error())
		return false
	}

	defer l.Close()

	if a.cfg.UseStartTLS {
		// Reconnect with TLS
		err = l.StartTLS(a.getTLSConfig())
		if err != nil {
			log.Error().Msgf("failed to connect to ldap server with StartTLS: %s", err.Error())
			return false
		}
	}

	// First bind with a read only user
	err = l.Bind(a.cfg.BindUser, a.cfg.BindPass)
	if err != nil {
		log.Error().Msgf("failed to connect to ldap server with specified bind credentials: %s", err.Error())
		return false
	}

	if a.cfg.LoginNameAttribute == "" {
		a.cfg.LoginNameAttribute = "sAMAccountName"
	}

	if a.cfg.ObjectClass == "" {
		a.cfg.ObjectClass = "organizationalPerson"
	}

	searchRequest := ldap.NewSearchRequest(
		a.cfg.BaseDN, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=%s)(%s=%s))", a.cfg.ObjectClass, a.cfg.LoginNameAttribute, username),
		[]string{"dn", a.cfg.LoginNameAttribute}, // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Error().Msgf("failed to find any results: %s", err.Error())
		return false
	}

	for _, entry := range sr.Entries {
		if entry.GetAttributeValue(a.cfg.LoginNameAttribute) != username {
			continue
		}

		// Disconnect and then reconnect as the user to validate the password.
		l.Close()
		l, err = a.dial()
		if err != nil {
			log.Error().Msgf("failed to unbind: %s", err.Error())
			return false
		}

		// Try and bind as the other user.
		err = l.Bind(entry.DN, password)
		if err != nil {
			log.Error().Msgf("failed to bind with specified username and password: %s", err.Error())
			return false
		}
		return true
	}

	return a.allowedUsers[username] == password
}

func (a *Authenticator) getTLSConfig() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: !a.cfg.AllowUnverifiedTLS, //nolint:gosec
	}
}

// New creates a new in memory authenticator
func New(config *Config) (*Authenticator, error) {
	var l *ldap.Conn
	var err error

	auth := &Authenticator{cfg: config}
	l, err = auth.dial()

	if err != nil {
		log.Error().Msgf("failed to connect to ldap server: %s", err.Error())
		return nil, fmt.Errorf("failed to connect to ldap server: %w", err)
	}
	defer l.Close()

	if config.UseStartTLS {
		// Reconnect with TLS
		err = l.StartTLS(auth.getTLSConfig())
		if err != nil {
			log.Error().Msgf("failed to connect to ldap server with TLS: %s", err.Error())
			return nil, fmt.Errorf("failed to connect to ldap server with TLS: %w", err)
		}
	}

	// First bind with a read only user
	err = l.Bind(config.BindUser, config.BindPass)
	if err != nil {
		log.Error().Msgf("failed to connect to ldap server with specified bind credentials: %s", err.Error())
		return nil, fmt.Errorf("failed to connect to ldap server with specified credentials: %w", err)
	}

	return auth, nil
}
