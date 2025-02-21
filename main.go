package main

import ("fmt"
		"os"
		"log"
		"github.com/joho/godotenv"
		"github.com/go-chi/chi"
		"github.com/go-chi/cors"
		_ "github.com/lib/pq"
		"github.com/Suyash0803/rss-agg/internal/database"

		"net/http"
		"database/sql"
		
		
)
type apiConfig struct {
	DB *database.Queries
}

func main()  {
	fmt.Println("Hello, World!")

	godotenv.Load(".env")

	portString:=os.Getenv("PORT")

	if portString==""{
		log.Fatal("Port not found")
	}
	fmt.Println("Port:",portString)

	dbURL:=os.Getenv("DATABASE_URL")
	if dbURL==""{
		log.Fatal("Database URL not found")
	}
	conn,err:=sql.Open("postgres",dbURL)
	if err!=nil{
		log.Fatal("Cant connect ")
	}


	apiCfg:=apiConfig{
		DB:database.New(conn),
	}
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
	v1Router.Post("/users",apiCfg.handlerCreateUser)
	router.Mount("/v1",v1Router)


	server:=&http.Server{
		Handler:router,
		Addr:   ":"+portString,
	}

	log.Printf("Server started on port %s",portString)

	err := server.ListenAndServe()
if err != nil {
    log.Fatal(err)
}

}