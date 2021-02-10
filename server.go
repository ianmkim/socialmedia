package main

import (
    "os"
    "log"
    "github.com/Kamva/mgm/v2"
    "github.com/gofiber/fiber/v2"

    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/template/django"
    "go.mongodb.org/mongo-driver/mongo/options"

    jwtware "github.com/gofiber/jwt/v2"
    "github.com/parvusvox/socialmedia/routes"
)

/**
initializes the mgm connection to mongo
this app uses both mgm AND the official mongo driver for max readability
**/
func init(){
    connectionString := os.Getenv("CONNSTRING")
    err := mgm.SetDefaultConfig(nil, "socialmedia", options.Client().ApplyURI(connectionString))
    if err != nil{
        log.Fatal(err) 
    }
}

func main(){
    // in case I decide I have more time and attempts a frontend
    engine := django.New("./views", ".html")

    // utility functions pretty self explanatory
    engine.AddFunc("index", IndexArr)
    engine.AddFunc("utsToTime", UtsToTime)
    engine.AddFunc("replaceSpaces",ReplaceSpaces)
    engine.AddFunc("findId", FindId)

    app := fiber.New(fiber.Config {
        Views: engine,
    })

    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
    }))

    app.Static("/static", "./static")

    app.Post("/login", routes.Login)
    app.Post("/register", routes.Register)

    // all routes beyond this line is protected through JWT
    app.Use(jwtware.New(jwtware.Config{
        SigningKey : []byte(os.Getenv("SECRET")),
    }))

    // gets all users
    // requires int page and int limit for pagination
    app.Get("/users", routes.GetUsers)
    // indexes the quotes for searching
    app.Get("/index", routes.IndexUsers)
    // gets user based on hex id
    // requires string id
    app.Get("/user/:id", routes.GetUser)
    // full text search on quotes
    // requires q string
    // optional int page and int limit
    app.Get("/searchQuotes", routes.GetSearchByQuote)
    // creates user
    app.Post("/user", routes.PostUser)

    // for deployment on heroku cause it uses dynamic ports
    port := os.Getenv("PORT")
    if port == ""{
        port = ":3000"
    } else {
        port = ":" + port
    }

    log.Fatal(app.Listen(port))
}
