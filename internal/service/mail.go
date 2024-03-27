package service

import (
	"crypto/tls"
	"fmt"
	"github.com/AZRV17/Skylang/internal/config"
	"log"
	"net/smtp"
)

type EmailService struct {
	cfg config.Config
}

func NewEmailService(cfg config.Config) *EmailService {
	return &EmailService{
		cfg: cfg,
	}
}

func (m *EmailService) SendMailForPasswordReset(recipient string, resetCode int) error {
	headerMap := make(map[string]string)

	headerMap["To"] = recipient
	headerMap["Subject"] = "Password reset"
	headerMap["From"] = m.cfg.Email.Username

	mailMessage := ""

	for key, value := range headerMap {
		mailMessage += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	mailMessage += fmt.Sprintf("\r\n%s", "Ваш код для сброса пароля: "+fmt.Sprint(resetCode))

	log.Println(mailMessage)

	if err := m.sendEmail(mailMessage, recipient); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m *EmailService) sendEmail(message string, recipient string) error {
	auth := smtp.PlainAuth("", m.cfg.Email.Username, m.cfg.Email.Password, m.cfg.Email.Host)

	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         m.cfg.Email.Host,
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", m.cfg.Email.Host, m.cfg.Email.Port), tlsConf)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, m.cfg.Email.Host)
	if err != nil {
		return err
	}

	if err = client.Auth(auth); err != nil {
		return err
	}

	if err = client.Mail(m.cfg.Email.Username); err != nil {
		return err
	}

	if err = client.Rcpt(recipient); err != nil {
		return err
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	err = client.Quit()
	if err != nil {
		return err
	}

	//err := smtp.SendMail(fmt.Sprintf(
	//	"%s:%d", m.cfg.Email.Host, m.cfg.Email.Port), auth, m.cfg.Email.Username, []string{recipient}, []byte(message))
	//if err != nil {
	//	return err
	//}

	return nil
}

func (m *EmailService) SendMailWithTemplate() error {
	return nil
}
