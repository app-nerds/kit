package email

import (
	"gopkg.in/gomail.v2"
)

/*
IMailService provides an interface describing a service for working with email
*/
type IMailService interface {
	Connect() error
	Send(mail ...*Mail) error
}

/*
MailService provides methods for working with email
*/
type MailService struct {
	Config *Config
	Dialer *gomail.Dialer
	Sender gomail.SendCloser
}

/*
NewMailService creates a new instance of MailService
*/
func NewMailService(config *Config) *MailService {
	result := &MailService{
		Config: config,
		Dialer: &gomail.Dialer{
			Host: config.Host,
			Port: config.Port,
		},
	}

	if config.UserName != "" {
		result.Dialer.Username = config.UserName
	}

	if config.Password != "" {
		result.Dialer.Password = config.Password
	}

	return result
}

/*
Connect establishes a connections to an SMTP server
*/
func (s *MailService) Connect() error {
	var err error

	s.Sender, err = s.Dialer.Dial()
	return err
}

/*
Send sends an email
*/
func (s *MailService) Send(mail ...*Mail) error {
	mailItems := make([]*gomail.Message, len(mail))

	for index := 0; index < len(mail); index++ {
		m := gomail.NewMessage()
		m.SetAddressHeader("From", mail[index].From.EmailAddress, mail[index].From.Name)
		m.SetHeader("Subject", mail[index].Subject)
		m.SetBody("text/html", mail[index].Body)

		for _, p := range mail[index].To {
			m.SetAddressHeader("To", p.EmailAddress, p.Name)
		}

		mailItems[index] = m
	}

	return gomail.Send(s.Sender, mailItems...)
}
