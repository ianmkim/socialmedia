package routes

import (
    "time"
    "os"
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "crypto/sha256"
    "github.com/Kamva/mgm/v2"
    "github.com/gofiber/fiber/v2"
    "github.com/parvusvox/socialmedia/models"
    jwt "github.com/form3tech-oss/jwt-go"
)

/*
returns a binary value of whether the password hashes match record
*/
func PasswordMatch (pass string, user models.LoginUser) bool {
    if(Hash(pass) == user.Password){
        return true
    }
    return false
}

/*
performs a SHA256 hash to store password, obviously not secure bc it's
not salted, but it'll do for this toy project
*/
func Hash(data string) string {
    bData := []byte(data)
    hash := sha256.Sum256(bData)
    return string(hash[:])
}

/*
returns a JWT token if the login is successful
*/
func Login(c *fiber.Ctx) error {
    email := c.FormValue("email")
    pass := c.FormValue("password")

    db, client := Database()
    cur, _ := db.Collection("login_users").Find(context.Background(), bson.M{
        "email": email,
    })
    client.Disconnect(context.TODO())

    var user models.LoginUser
    for cur.Next(context.Background()){
        // decode the first record
        err := cur.Decode(&user)
        if err != nil {
            // if there is error, terminate and send response
            return c.Status(500).JSON(fiber.Map{
                "ok": false,
                "error" : err.Error(),
            })
        }
        // were only interested in the first result
        break
    }

    status := false
    var sToken string

    // if the password hashes match
    if(PasswordMatch(pass, user)){
        status = true
        
        // create new JWT token
        token := jwt.New(jwt.SigningMethodHS256)
        claims := token.Claims.(jwt.MapClaims)
        claims["username"] = user.Username
        // set expiration date to 24 hours from now 
        claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

        // convert token to string
        t, err := token.SignedString([]byte(os.Getenv("SECRET")))
        if err != nil {
            return c.Status(500).JSON(fiber.Map{
                "ok": false,
                "error": err.Error(),
            })
        }
        sToken = t
    }

    // send back token
    return c.Status(200).JSON(fiber.Map{
        "ok": status,
        "token" : sToken,
    })
}

/*
registers new login user
*/
func Register(c *fiber.Ctx) error {
    // get the user JSON
    params := new(models.LoginUser)
    if err := c.BodyParser(params); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "ok": false,
            "error": err.Error(),
        })
    }

    db, client := Database()
    // check if email is unique
    cur, _ := db.Collection("login_users").Find(context.Background(), bson.M{
        "email" : params.Email,
    })

    for cur.Next(context.Background()){
        // if this runs even once, unique check fails
        return c.Status(400).JSON(fiber.Map{
            "ok": false,
            "error": "User already exists",
        })
    }

    // hash the password
    params.Password = Hash(params.Password)
    // create user
    err := mgm.Coll(params).Create(params)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"ok": false, "err":err.Error()})
    }
    client.Disconnect(context.TODO())

    return c.Status(200).JSON(fiber.Map{
        "ok": true,
    })

}
