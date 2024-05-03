package main

import (
	"token-verifier/httpServer"
	"token-verifier/tokenVerifier"
)

type Server interface {
	Serve(verifier tokenVerifier.TokenVerifier)
}

func main() {
	var server Server = httpServer.HttpServer{}
	server.Serve(tokenVerifier.New())
}
