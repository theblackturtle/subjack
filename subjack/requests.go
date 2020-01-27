package subjack

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

func get(url string, ssl bool, followRedirects bool, userAgent string, timeout int) (body []byte) {
	if ssl {
		url = "https://" + url
	} else {
		url = "http://" + url
	}

	client := resty.New()
	if followRedirects {
		client.SetRedirectPolicy(resty.RedirectPolicyFunc(func(req *http.Request, via []*http.Request) error {
			return nil
		}))
	} else {
		client.SetRedirectPolicy(resty.RedirectPolicyFunc(func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}))
	}

	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetHeader("User-Agent", userAgent)
	client.SetHeader("Connection", "close")
	client.SetTimeout(time.Duration(timeout) * time.Second)
	client.SetCloseConnection(true)
	client.SetDisableWarn(true)

	resp, err := client.R().Get(url)
	if err != nil {
		return []byte{}
	}
	return resp.Body()
}
