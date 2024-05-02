package main

import (
	"token-signer-validator/httpServer"
	"token-signer-validator/tokenValidator"
)

type Server interface {
	Serve(validator tokenValidator.TokenValidator)
}

func main() {
	var server Server = httpServer.HttpServer{}
	server.Serve(tokenValidator.New())
}
