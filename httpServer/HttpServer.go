package httpServer

import (
	"log"
	"net/http"
	"token-verifier/config"
	"token-verifier/tokenVerifier"
)

type HttpServer struct {
}

type RequestParser interface {
	ParseRequest(r *http.Request) tokenVerifier.VerifyTokenRequest
}

func (h HttpServer) Serve(verifier tokenVerifier.TokenVerifier) {
	conf := config.GetGlobalConfig()
	var parser RequestParser = HttpRequestParser{}

	http.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Token validate requested")
		log.Println(r)

		if r.Method != "GET" {
			http.Error(w, "Invalid request method", 405)
			return
		}

		request := parser.ParseRequest(r)

		tokenValid, msg := verifier.VerifyToken(request)

		if !tokenValid {
			http.Error(w, msg, 400)
		}

		return
	})
	log.Fatal(http.ListenAndServe(":"+conf.Port, nil))

}
