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
	log.Printf("NEWWWWWW9999")
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

func (be *Backend) login(username, password string) (*client.Client, error) {
	log.Printf("KKKKSSOOO2222222")
	var c *client.Client
	var err error
	if be.Security == SecurityTLS {
		if c, err = client.DialTLS(be.Addr, be.TLSConfig); err != nil {
			return nil, err
		}
	} else {
		log.Printf("KKKKSSOOO: %s", be.Security)
		if c, err = client.Dial(be.Addr); err != nil {
			log.Printf("VAO DAY")
			return nil, err
		}
		log.Printf("%s", be.Addr)
		if be.Security == SecuritySTARTTLS {
			log.Printf("HAY VAO DAY")
			//if err := c.StartTLS(be.TLSConfig); err != nil {
			//	log.Printf("AAAAA: %s", err)
			//	return nil, err
			//}
		}
	}

	if err := c.Login(username, password); err != nil {
		return nil, err
	}

	return c, nil
}

func (be *Backend) Login(_ *imap.ConnInfo, username, password string) (backend.User, error) {
	log.Printf("KKKKSSOOO1111111111")
	c, err := be.login(username, password)
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
