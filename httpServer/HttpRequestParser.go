package httpServer

import (
	"encoding/json"
	"net/http"
	"token-verifier/tokenVerifier"
)

type HttpRequestParser struct {
}

func (p HttpRequestParser) ParseRequest(r *http.Request) (tokenVerifier.VerifyTokenRequest, error) {
	var entity tokenVerifier.VerifyTokenRequest

	err := json.NewDecoder(r.Body).Decode(&entity)
	return entity, err
}
