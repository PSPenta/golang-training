package main

import (
	"fmt"
	"log"
	"net/http"

	"main/src/router"

	"github.com/rs/cors"
)

func main() {
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowedHeaders: []string{"*"}, //or you can your header key values which you are using in your application
	})
	router.Routes()
	fmt.Println("Go server started on port 9000")
	log.Fatal(http.ListenAndServe(":9000", corsOpts.Handler(router.Router)))
}
