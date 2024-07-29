package mail

import (
	"net/smtp"
)

type SmtpClient struct {
	Enabled bool        `json:"enabled"`
	Config  *SmtpConfig `json:"config"`
	Auth    *smtp.Auth
}

type SmtpConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
}
