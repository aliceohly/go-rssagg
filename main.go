package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/aliceohly/go-rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("$DB_URL must be set")
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	router := chi.NewRouter()

	router.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	v1Router.Get("/user", apiCfg.middlewareAuth(apiCfg.handlerUserGet))
	v1Router.Post("/user", apiCfg.handlerUserCreate)

	v1Router.Get("/feeds", apiCfg.handlerFeedsGet)
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerFeedCreate))

	router.Mount("/v1", v1Router)

	// why need to add & before http.Server?
	server := http.Server{Addr: ":" + port, Handler: router}

	log.Printf("Server started on port %s", port)
	// what does Fatal mean?
	log.Fatal(server.ListenAndServe())
}
