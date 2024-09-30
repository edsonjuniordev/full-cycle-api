package main

import (
	"net/http"

	"github.com/edsonjuniordev/full-cycle-api/configs"
	_ "github.com/edsonjuniordev/full-cycle-api/docs"
	"github.com/edsonjuniordev/full-cycle-api/internal/entities"
	"github.com/edsonjuniordev/full-cycle-api/internal/infra/database"
	"github.com/edsonjuniordev/full-cycle-api/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go Expert - full-cycle-api
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  -

// @contact.name   Edson Júnior
// @contact.url    github.com/edsonjuniordev
// @contact.email  dev.edsonjunior@gmail.com

// @license.name   -
// @license.url    -

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entities.User{}, &entities.Product{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JWTExpiresIn)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/products", func(productRouter chi.Router) {
		productRouter.Use(jwtauth.Verifier(configs.TokenAuth))
		productRouter.Use(jwtauth.Authenticator)
		productRouter.Post("/", productHandler.CreateProduct)
		productRouter.Get("/", productHandler.ListProducts)
		productRouter.Get("/{id}", productHandler.GetProduct)
		productRouter.Put("/{id}", productHandler.UpdateProduct)
		productRouter.Delete("/{id}", productHandler.DeleteProduct)
	})

	router.Post("/users", userHandler.CreateUser)
	router.Post("/users/generate-token", userHandler.GetJWT)

	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", router)
}
