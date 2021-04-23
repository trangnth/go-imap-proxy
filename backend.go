package proxy

import (
	"crypto/tls"
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/backend"
	"github.com/emersion/go-imap/client"
)

type Security int

const (
	SecurityNone Security = iota
	SecuritySTARTTLS
	SecurityTLS
)

type Backend struct {
	Addr      string
	Security  Security
	TLSConfig *tls.Config

	unexported struct{}
}

func New(addr string) *Backend {
	log.Printf("NEWWWWW")
	return &Backend{
		Addr:     addr,
		Security: SecuritySTARTTLS,
	}
}

func NewTLS(addr string, tlsConfig *tls.Config) *Backend {
	return &Backend{
		Addr:      addr,
		Security:  SecurityTLS,
		TLSConfig: tlsConfig,
	}
}

func (be *Backend) Klogin(username, password string) (*client.Client, error) {
	log.Printf("KKKKSSOOO2222222")
	var c *client.Client
	var err error
	if be.Security == SecurityTLS {
		if c, err = client.DialTLS(be.Addr, be.TLSConfig); err != nil {
			return nil, err
		}
	} else {
		log.Printf("KKKKSSOOO")
		if c, err = client.Dial(be.Addr); err != nil {
			return nil, err
		}

		if be.Security == SecuritySTARTTLS {
			if err := c.StartTLS(be.TLSConfig); err != nil {
				return nil, err
			}
		}
	}

	if err := c.Login(username, password); err != nil {
		return nil, err
	}

	return c, nil
}

func (be *Backend) Login(_ *imap.ConnInfo, username, password string) (backend.User, error) {
	log.Printf("KKKKSSOOO1111111111")
	c, err := be.Klogin(username, password)
	if err != nil {
		return nil, err
	}

	u := &user{
		be:       be,
		c:        c,
		username: username,
	}
	return u, nil
}
