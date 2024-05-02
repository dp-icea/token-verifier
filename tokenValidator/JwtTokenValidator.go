package tokenValidator

import (
	"crypto/rsa"
	"log"
	"os"
	"strings"
	conf "token-signer-validator/config"

	"github.com/golang-jwt/jwt"
)

type JwtTokenValidator struct {
	publicKey *rsa.PublicKey
}

func New() JwtTokenValidator {
	config := conf.GetGlobalConfig()

	bytes, err := os.ReadFile(config.RsaPublicKeyFileName)
	if err != nil {
		log.Panic(err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(bytes)
	if err != nil {
		log.Panic(err)
	}

	return JwtTokenValidator{
		publicKey: publicKey,
	}
}

func (j JwtTokenValidator) ValidateToken(request ValidateTokenRequest) bool {
	parts := strings.Split(request.AccessToken, ".")
	err := jwt.SigningMethodRS256.Verify(strings.Join(parts[0:2], "."), parts[2], j.publicKey)
	if err != nil {
		return false
	}
	return true
}
