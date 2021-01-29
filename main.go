package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type bearerMiddlewareMetadata struct {
	IssuerURL string `json:"issuerURL"`
	ClientID  string `json:"clientID"`
}

const (
	bearerPrefix       = "bearer "
	bearerPrefixLength = len(bearerPrefix)
)

var ren = render.New()

func main() {
	// serve
	addr := ":8080"
	log.Printf("Listening on %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, handlers()))
}

func handlers() http.Handler {

	r := mux.NewRouter()
	r.HandleFunc("/token", test).Methods("GET").Headers("Authorization", "")
	return r
}

func test(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	rawToken := authHeader[bearerPrefixLength:]
	// ctx := context.Background()
	// provider, err := oidc.NewProvider(context.Background(), keycloakURL)
	// if err != nil {
	// 	return
	// }

	// verifier := provider.Verifier(&oidc.Config{
	// 	ClientID: keycloakClientID,
	// })

	// prefix := strings.HasPrefix(strings.ToLower(authHeader), bearerPrefix)
	// if !prefix {
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }
	// log.Println(prefix)

	// _, err = verifier.Verify(ctx, rawToken)
	// if err != nil {
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)

	// 	return
	// }

	ren.JSON(w, http.StatusOK, rawToken)
}
