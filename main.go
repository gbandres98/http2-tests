package main

import (
	"context"
	"log"
	"net/http"

	"github.com/PaackEng/paackit"
	infraHTTPX "github.com/PaackEng/paackit/httpx"
	"github.com/PaackEng/paackit/paack"
	"github.com/PaackEng/test-server/endpoint"
)

const (
	serviceName = "RetailerService"
)

func main() {
	// if err := run(); err != nil {
	// 	fmt.Fprintf(os.Stderr, "%s\n", err)
	// 	os.Exit(1)
	// }

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
	})

	log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil))
}

func run() error {
	// Context
	ctx := context.Background()

	_httpCfg := infraHTTPX.NewDefaultConfig()

	// Infra
	http := infraHTTPX.NewHTTPX(*_httpCfg)

	privateChain := []infraHTTPX.Middleware{}

	// Endpoints
	// Country
	testHTTPX := endpoint.NewTestEndpoint(endpoint.TestEndpointDI{
		Middleware: privateChain,
	})

	// http register
	http.Register([]infraHTTPX.Service{
		// country
		testHTTPX,
	})

	//   Boot
	p := paackit.New(paackit.Config{
		Ctx:  ctx,
		Name: serviceName,
	})

	err := p.RegisterTransporter([]paack.Transporter{
		http,
	})
	if err != nil {
		return err
	}

	err = p.Start()
	if err != nil {
		return err
	}

	return nil
}
