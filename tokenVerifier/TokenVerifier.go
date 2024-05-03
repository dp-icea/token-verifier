package tokenVerifier

type VerifyTokenRequest struct {
	AccessToken string `json:"access_token"`
}

type VerifyTokenResponse struct {
	Message string `json:"message"`
}

type TokenVerifier interface {
	VerifyToken(request VerifyTokenRequest) (bool, string)
}
