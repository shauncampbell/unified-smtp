package config

import (
	"github.com/shauncampbell/unified-smtp/internal/authenticator/ldap"
	"github.com/shauncampbell/unified-smtp/internal/authenticator/memory"
)

// Authenticator is a struct providing configuration of an authenticator.
type Authenticator struct {
	Type     string        `yaml:"type"`
	InMemory memory.Config `yaml:"in_memory"`
	LDAP     ldap.Config   `yaml:"ldap"`
}
