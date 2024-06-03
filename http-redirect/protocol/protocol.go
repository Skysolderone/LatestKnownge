package protocol

import "net/http"

type HttpRedirectPlugin struct {
	PreRequestHook (*http.Request)
}
