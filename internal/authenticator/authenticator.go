// Package authenticator is a package for looking up authenticator implementations based on configuration.
package authenticator

import (
	"fmt"

	"github.com/shauncampbell/unified-smtp/internal/authenticator/ldap"

	"github.com/shauncampbell/unified-smtp/internal/authenticator/memory"
	"github.com/shauncampbell/unified-smtp/internal/config"
	"github.com/shauncampbell/unified-smtp/pkg/authenticator"
)

// ForConfiguration loads an authenticator based on the configuration.
func ForConfiguration(cfg *config.Authenticator) (authenticator.Authenticator, error) {
	switch cfg.Type {
	case "in_memory":
		return memory.New(&cfg.InMemory), nil
	case "ldap":
		return ldap.New(&cfg.LDAP)
	}
	return nil, fmt.Errorf("no authenticator named '%s'", cfg.Type)
}
