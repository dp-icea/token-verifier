package tokenVerifier

import (
	"crypto/rsa"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

	if !verifySigning(request.AccessToken, j.publicKey) {
		return false, "Token signature not valid"
	}

	if !verifyExpiration(request.AccessToken) {
		return false, "Token is expired"
	}

	if !verifyAudience(request.AccessToken, request.ExpectedAudience) {
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
		log.Println("Invalid token signature")
		return false
	}
	return true
}

func verifyAudience(tokenString string, expectedAudience string) bool {

	tokenAudience, err := extractUnverifiedClaims(tokenString, "aud")
	if err != nil {
		log.Println(err.Error())
		return false
	}
	if tokenAudience != expectedAudience {
		log.Printf("Invalid audience\nExpected: %s\nActual: %s", expectedAudience, tokenAudience)
		return false
	}
	return true
}

func verifyScope(tokenString string, expectedScope string) bool {
	tokenScope, err := extractUnverifiedClaims(tokenString, "scope")
	if err != nil {
		log.Println(err.Error())
		return false
	}

	scopeList := strings.Split(tokenScope, " ")

	for _, scope := range scopeList {
		if scope == expectedScope {
			return true
		}
	}
	log.Printf("Invalid scope\nExpected: %s\nActual: %s", expectedScope, tokenScope)
	return false
}

func verifyExpiration(tokenString string) bool {
	expiration, err := extractUnverifiedClaims(tokenString, "exp")
	if err != nil {
		log.Println(err)
		return false
	}
	exp, err := strconv.ParseInt(expiration, 10, 64)
	if err != nil {
		log.Println(err)
		return false
	}
	return exp >= time.Now().Unix()
}

func extractUnverifiedClaims(tokenString string, key string) (string, error) {
	var value string
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		switch claims[key].(type) {
		case float64, float32:
			value = fmt.Sprintf("%.0f", claims[key])
		default:
			value = fmt.Sprint(claims[key])
		}

	}

	if value == "" {
		return "", fmt.Errorf("invalid token payload")
	}
	return value, nil
}
