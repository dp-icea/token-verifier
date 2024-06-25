package httpServer

import (
	"net/http"
	"token-verifier/tokenVerifier"
)

type HttpRequestParser struct {
}

func (p HttpRequestParser) ParseRequest(r *http.Request) tokenVerifier.VerifyTokenRequest {
	entity := tokenVerifier.VerifyTokenRequest{
		AccessToken:   r.URL.Query().Get("access_token"),
		RequiredScope: r.URL.Query().Get("required_scope"),
	}

	return entity
}
