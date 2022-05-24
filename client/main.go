package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	transport := http.Transport{
		TLSClientConfig: tlsConfig(),
	}

	client := http.Client{
		Transport: &transport,
	}

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}

	log.Println(resp.Proto)
}

func tlsConfig() *tls.Config {

	caBundle := x509.NewCertPool()

	certs, _ := ioutil.ReadFile("../server.crt")

	caBundle.AppendCertsFromPEM(certs)
	t := &tls.Config{
		RootCAs: caBundle,
	}

	return t

}
