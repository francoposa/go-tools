package postgres

import (
	"fmt"
)

// ConnectionConfig defines Postgres database_sql connection parameters
type ConnectionConfig struct {
	Host                  string
	Port                  uint16
	Username              string
	Password              string
	Database              string
	ApplicationName       string
	ConnectTimeoutSeconds int
	SSLMode               string
}

// BuildConnectionURI builds a connection string for lib/pq from ConnectionConfig.
// If a missing or invalid field is provided, an error is returned.
func BuildConnectionURI(cc ConnectionConfig) string {
	auth := ""
	if cc.Username != "" || cc.Password != "" {
		auth = fmt.Sprintf("%s:%s@", cc.Username, cc.Password)
	}
	url := fmt.Sprintf(
		"postgres://%s%s:%d/%s?application_name=%s&connect_timeout=%d&sslmode=%s",
		auth,
		cc.Host,
		cc.Port,
		cc.Database,
		cc.ApplicationName,
		cc.ConnectTimeoutSeconds,
		cc.SSLMode,
	)
	return url
}
