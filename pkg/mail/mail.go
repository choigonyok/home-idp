package mail

import (
	"fmt"
	"log"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/util"
)

type SmtpClient struct {
	Auth *smtp.Auth
	Addr string
}

const (
	defaultMailSender = "idp@choigonyok.com"
)

func (s *SmtpClient) SendMail(receivers []string) error {
	code := util.GenerateRandNum(6)
	message := `type this code at the home-idp ui dashboard to sign up. ` + strconv.Itoa(code)
	headers := []string{
		"From: " + defaultMailSender,
		"To: " + receivers[0],
		"Subject: Home-IDP Confirmation Code",
		"MIME-Version: 1.0;",
		"Content-Type: text/html; charset=\"UTF-8\";",
	}
	emailBody := fmt.Sprintf("<html><body>%s</body></html>", message)
	emailMessage := fmt.Sprintf("%s\r\n\r\n%s",
		strings.Join(headers, "\r\n"), emailBody)

	// err := smtp.SendMail(s.Addr, nil, defaultMailSender, receivers, []byte(emailMessage))
	// if err != nil {
	// 	fmt.Println("Failed to send email:", err)
	// 	return err
	// }
	client, _ := smtp.Dial(s.Addr)

	if err := client.Mail(defaultMailSender); err != nil {
		log.Fatalf("Failed to set mail sender: %v", err)
	}

	if err := client.Rcpt(receivers[0]); err != nil {
		log.Fatalf("Failed to set recipient: %v", err)
	}

	wc, err := client.Data()
	if err != nil {
		log.Fatalf("Failed to get write closer: %v", err)
	}

	_, err = wc.Write([]byte(emailMessage))
	if err != nil {
		log.Fatalf("Failed to write message: %v", err)
	}

	err = wc.Close()
	if err != nil {
		log.Fatalf("Failed to close write closer: %v", err)
	}

	client.Quit()

	return nil
}

func NewClient(component util.Components) *SmtpClient {
	user := env.Get(env.GetEnvPrefix(component) + "_MANAGER_SMTP_USER")
	password := env.Get(env.GetEnvPrefix(component) + "_MANAGER_SMTP_PASSWORD")
	host := env.Get(env.GetEnvPrefix(component) + "_MANAGER_SMTP_HOST")
	port := env.Get(env.GetEnvPrefix(component) + "_MANAGER_SMTP_PORT")

	auth := smtp.PlainAuth("", user, password, host)
	addr := host + ":" + port

	return &SmtpClient{
		Auth: &auth,
		Addr: addr,
	}
}
