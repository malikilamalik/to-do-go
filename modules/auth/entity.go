package auth

import (
	"to-do-go/config"

	"github.com/dgrijalva/jwt-go"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type JwtClaim struct {
	Username string
	jwt.StandardClaims
}

var Wrapper = &JwtWrapper{
	SecretKey:       config.StringEnvVariable("JWT_SECRETKEY"),
	Issuer:          config.StringEnvVariable("JWT_ISSUER"),
	ExpirationHours: 24,
}
