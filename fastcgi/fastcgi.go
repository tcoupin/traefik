package fastcgi

import (
	"net/http"

	"github.com/yookoala/gofast"
)

type Transport struct {
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "http"
	client, err := gofast.SimpleClientFactory(gofast.SimpleConnFactory("tcp", req.URL.Host), 0)()
	// TODO check err

	// defer closing with error reporting
	defer func() {
		if client == nil {
			return
		}

		// signal to close the client
		// or the pool to return the client
		if err = client.Close(); err != nil {
			// log
		}
	}()

	// TODO, for now, we just use Do method. Do we need middleware https://godoc.org/github.com/yookoala/gofast#Middleware ?
	resp, err := client.Do(gofast.NewRequest(req))
	// TODO check err
	// TODO convert Resp to http.Response

	// Just to avoid complation error for now
	if resp != nil {
		err = nil
	}

	return http.Get("http://google.fr")
}
