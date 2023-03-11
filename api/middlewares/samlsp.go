package middlewares

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"net/url"

	"github.com/crewjam/saml/samlsp"
)

func NewSamlSP() (*samlsp.Middleware, error) {
	keyPair, err := tls.LoadX509KeyPair("idp.example.com.crt", "idp.example.com.pem")
	if err != nil {
		log.Println("fail to load x509 key pair")
		return nil, err
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		log.Println("fail to parse x509 cert")
		return nil, err
	}

	idpMetadataURL, err := url.Parse("http://idp:80/simplesaml/saml2/idp/metadata.php")
	if err != nil {
		log.Println("fail to parse metadata url")
		return nil, err
	}

	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient, *idpMetadataURL)
	if err != nil {
		log.Println("fail to fetch metadata")
		return nil, err
	}

	rootURL, err := url.Parse("https://sp.example.org")
	if err != nil {
		log.Println("fail to parse root url")
		return nil, err
	}

	samlSP, err := samlsp.New(samlsp.Options{
		URL:         *rootURL,
		Key:         keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate: keyPair.Leaf,
		IDPMetadata: idpMetadata,
	})

	return samlSP, err
}
