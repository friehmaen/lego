// Package inwx implements a DNS provider for solving the DNS-01 challenge
// using the XML_RPC provided by inwx.de DNS registar.
// https://www.inwx.de / https://www.inwx.com
// https://www.inwx.de/de/offer/api
package inwx

import (
	"fmt"
	"os"

	"github.com/xenolf/lego/acme"
)

// DNSProvider is an implementation of the acme.ChallengeProvider interface.
type DNSProvider struct {
	username   string
	password   string
}

// NewDNSProvider returns a DNSProvider instance configured for inwx.
// Credentials must be passed in the environment variables: 
// username: INWX_USERNAME
// password: INWX_PASSWORD
func NewDNSProvider() (*DNSProvider, error) {
	username := os.Getenv("INWX_USERNAME")
	password := os.Getenv("INWX_PASSWORD")
	return NewDNSProviderCredentials(username, password)
}

// NewDNSProviderCredentials uses the supplied credentials to return a
// DNSProvider instance configured for inwx.
func NewDNSProviderCredentials(username, password string) (*DNSProvider, error) {
	if username == "" || password == "" {
		return nil, fmt.Errorf("inwx credentials missing")
	}

	return &DNSProvider{
		username: username,
		password: password,
	}, nil
}

// Present creates a TXT record to fulfil the dns-01 challenge.
func (d *DNSProvider) Present(domain, token, keyAuth string) error {
	fqdn, value, ttl := acme.DNS01Record(domain, keyAuth)

	fmt.Println("inwx: fqdn: ", fqdn, " value: ", value, " ttl:", ttl)

	return nil
}

// CleanUp removes the TXT record matching the specified parameters.
func (d *DNSProvider) CleanUp(domain, token, keyAuth string) error {
	fqdn, _, _ := acme.DNS01Record(domain, keyAuth)

	fmt.Println("inwx: fqdn: ", fqdn)

	return nil
}
