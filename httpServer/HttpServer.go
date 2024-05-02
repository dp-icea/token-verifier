package httpServer

import (
	"log"
	"net/http"
	"token-signer-validator/config"
	"token-signer-validator/tokenValidator"
)

type HttpServer struct {
}

type RequestParser interface {
	ParseRequest(r *http.Request) (tokenValidator.ValidateTokenRequest, error)
}

func (h HttpServer) Serve(validator tokenValidator.TokenValidator) {
	conf := config.GetGlobalConfig()
	var parser RequestParser = HttpRequestParser{}

	http.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Token validate requested")
		log.Println(r)

		if r.Method != "GET" {
			http.Error(w, "Invalid request method", 405)
			return
		}

		request, err := parser.ParseRequest(r)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		tokenValid := validator.ValidateToken(request)

		if !tokenValid {
			http.Error(w, "Token signature not valid", 400)
		}

		return
	})
	log.Fatal(http.ListenAndServe(":"+conf.Port, nil))

}
