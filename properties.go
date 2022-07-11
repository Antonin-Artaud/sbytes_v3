package main

type Properties struct {
	Server struct {
		Host           string   `yaml:"host" env-default:"localhost"`
		Port           string   `yaml:"port" env-default:"8080"`
		TrustedProxies []string `yaml:"trusted_proxies" env-default:""`
		ContextTicket  string   `yaml:"context-ticket" env-default:"/api/v1/ticket"`
	} `yaml:"server"`
	Database struct {
		URI          string `yaml:"uri" env-default:"mongodb://localhost:27017"`
		DbName       string `yaml:"db-name" env-default:"sbytes"`
		DbCollection string `yaml:"db-collection" env-default:"tickets"`
	} `yaml:"database"`
	Ticket struct {
		ExpirationTime string `yaml:"expiration-time"`
	} `yaml:"ticket"`
}
