package tokenVerifier

type VerifyTokenRequest struct {
	AccessToken      string `json:"access_token"`
	RequiredScope    string `json:"required_scope"`
	ExpectedAudience string `json:"expected_audience"`
}

type VerifyTokenResponse struct {
	Message string `json:"message"`
}

type TokenVerifier interface {
	VerifyToken(request VerifyTokenRequest) (bool, string)
}
