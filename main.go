package main

import ("fmt"
		"os"
		"log"
		"github.com/joho/godotenv"
		"github.com/go-chi/chi"
		"github.com/go-chi/cors"
		"net/http"
)

func main()  {
	fmt.Println("Hello, World!")

	godotenv.Load(".env")

	portString:=os.Getenv("PORT")

	if portString==""{
		log.Fatal("Port not found")
	}
	fmt.Println("Port:",portString)

	router:=chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router:=chi.NewRouter()
	v1Router.Get("/healtz",headerReadiness)
	v1Router.Get("/err",handleErr)
	router.Mount("/v1",v1Router)


	server:=&http.Server{
		Handler:router,
		Addr:   ":"+portString,
	}

	log.Printf("Server started on port %s",portString)

	err:=server.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}

}