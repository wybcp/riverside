package gtls

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/pkg/errors"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

type ClientTLS struct {
	CertFile string
	KeyFile string
	ServerName string
	CaFile string
}

func (c *ClientTLS)GetTLSCredentials()(credentials.TransportCredentials,error ) {
	ct, err := credentials.NewClientTLSFromFile(c.CertFile, c.ServerName)
	if err != nil {
		return nil,err
	}
	return ct,err
}
func (c *ClientTLS)GetTLSCredentialsByCA()(credentials.TransportCredentials,error ) {
	cert, err := tls.LoadX509KeyPair(c.CertFile, c.KeyFile)
	if err != nil {
		return nil,err
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(c.CaFile)
	if err != nil {
		return nil,err
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil,errors.New("certPool.AppendCertsFromPEM is not ok!")
	}
	ct:= credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   c.ServerName,
		RootCAs:      certPool,
	})

	return ct,err
}