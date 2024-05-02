package httpServer

import (
	"encoding/json"
	"net/http"
	"token-signer-validator/tokenValidator"
)

type HttpRequestParser struct {
}

func (p HttpRequestParser) ParseRequest(r *http.Request) (tokenValidator.ValidateTokenRequest, error) {
	var entity tokenValidator.ValidateTokenRequest

	err := json.NewDecoder(r.Body).Decode(&entity)
	return entity, err
}
