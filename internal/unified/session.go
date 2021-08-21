package unified

import (
	"errors"
	"github.com/emersion/go-smtp"
	"github.com/shauncampbell/unified-smtp/pkg/authenticator"
	"io"
	"io/ioutil"
	"log"
)



// A Session is returned after EHLO.
type Session struct{
	auth  authenticator.Authenticator
}

func (s *Session) AuthPlain(username, password string) error {
	if !s.auth.Authenticate(username, password) {
		return errors.New("invalid username or password")
	}
	return nil
}

func (s *Session) Mail(from string, opts smtp.MailOptions) error {
	log.Println("Mail from:", from)
	return nil
}

func (s *Session) Rcpt(to string) error {
	log.Println("Rcpt to:", to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if b, err := ioutil.ReadAll(r); err != nil {
		return err
	} else {
		log.Println("Data:", string(b))
	}
	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}
