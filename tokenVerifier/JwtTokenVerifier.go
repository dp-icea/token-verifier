package tokenVerifier

import (
	"crypto/rsa"
	"fmt"
	"log"
	"os"
	"strings"
	conf "token-verifier/config"

	"github.com/golang-jwt/jwt"
)

type JwtTokenVerifier struct {
	publicKey *rsa.PublicKey
}

func New() JwtTokenVerifier {
	config := conf.GetGlobalConfig()

	bytes, err := os.ReadFile(config.RsaPublicKeyFileName)
	if err != nil {
		log.Panic(err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(bytes)
	if err != nil {
		log.Panic(err)
	}

	return JwtTokenVerifier{
		publicKey: publicKey,
	}
}

func (j JwtTokenVerifier) VerifyToken(request VerifyTokenRequest) (bool, string) {

	config := conf.GetGlobalConfig()

	if !verifySigning(request.AccessToken, j.publicKey) {
		return false, "Token signature not valid"
	}

	if !verifyAudience(request.AccessToken, config.ExpectedAudience) {
		return false, "Token audience not valid"
	}

	if !verifyScope(request.AccessToken, request.RequiredScope) {
		return false, "Token does not contain necessary scope: " + request.RequiredScope
	}

	return true, ""

}

func verifySigning(tokenString string, publicKey *rsa.PublicKey) bool {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return false
	}
	err := jwt.SigningMethodRS256.Verify(strings.Join(parts[0:2], "."), parts[2], publicKey)

	if err != nil {
		return false
	}
	return true
}

func verifyAudience(tokenString string, expectedAudience string) bool {

	tokenAudience, err := extractUnverifiedClaims(tokenString, "aud")
	if err != nil {
		return false
	}
	return tokenAudience == expectedAudience
}

func verifyScope(tokenString string, expectedScope string) bool {
	tokenScope, err := extractUnverifiedClaims(tokenString, "scope")
	if err != nil {
		return false
	}
	return tokenScope == expectedScope
}

func extractUnverifiedClaims(tokenString string, key string) (string, error) {
	var value string
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		value = fmt.Sprint(claims[key])
	}

	if value == "" {
		return "", fmt.Errorf("invalid token payload")
	}
	log.Println(value)
	return value, nil
}
