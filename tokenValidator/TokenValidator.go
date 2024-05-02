package tokenValidator

type ValidateTokenRequest struct {
	AccessToken string `json:"access_token"`
}

type ValidateTokenResponse struct {
	Message string `json:"message"`
}

type TokenValidator interface {
	ValidateToken(request ValidateTokenRequest) bool
}
