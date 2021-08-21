// Package ldap provides functionality for working with ldap as an authentication service.
package ldap

// Config represents the configuration options for the ldap authenticator
type Config struct {
	Protocol           string `yaml:"protocol"`
	Host               string `yaml:"host"`
	Port               int    `yaml:"port"`
	BindUser           string `yaml:"bind_user"`
	BindPass           string `yaml:"bind_pwd"`
	BaseDN             string `yaml:"base_dn"`
	UseStartTLS        bool   `yaml:"use_start_tls"`
	AllowUnverifiedTLS bool   `yaml:"allow_unverified_tls"`
	LoginNameAttribute string `yaml:"login_name_attribute"`
	ObjectClass        string `yaml:"object_class"`
}
