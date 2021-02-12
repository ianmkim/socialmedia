package routes

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Kamva/mgm/v2"
	. "github.com/gobeam/mongo-go-pagination"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/parvusvox/socialmedia/models"
)

func Handle(err error) {
	if err != nil {
		panic(err.Error())
	}
}

/**
returns a database instance and client to close connection later
**/
func Database() (*mongo.Database, *mongo.Client) {
	opt := options.Client().ApplyURI(os.Getenv("CONNSTRING"))
	// terminate if op takes more than 10 secs
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// connect to db
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		log.Fatal(err.Error())
	}
	// social media is the db we want
	return client.Database("socialmedia"), client
}

/**
Returns paginated user objects based on page number and item per page limit
**/
func GetUsers(c *fiber.Ctx) error {
	// retrieve arguments and set default values if none are provided
	page, pErr := strconv.ParseInt(c.Query("page"), 10, 32)
	limit, lErr := strconv.ParseInt(c.Query("limit"), 10, 32)
	if pErr != nil {
		page = 1
	}
	if lErr != nil {
		limit = 10
	}

	// open mongodb connection
	db, client := Database()
	coll := db.Collection("users")
	// filter empty for all objects
	paginated, err := New(coll).Limit(limit).Page(page).Filter(bson.M{}).Find()
	Handle(err)

	var userList []models.User
	for _, raw := range paginated.Data {
		var user *models.User
		// unmarshal each user into User struct
		if marshallErr := bson.Unmarshal(raw, &user); marshallErr == nil {
			userList = append(userList, *user)
		}
	}
	// close db connection
	client.Disconnect(context.TODO())

	return c.JSON(fiber.Map{
		"ok":    true,
		"users": userList,
	})
}

/**
retrieves user based on hex ID
**/
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Bad arguments",
		})
	}

	user := &models.User{}
	col := mgm.Coll(user)
	err := col.FindByID(id, user)
	Handle(err)

	return c.JSON(fiber.Map{
		"ok":   true,
		"user": user,
	})
}

/**
Full text search for quotes
**/
func GetSearchByQuote(c *fiber.Ctx) error {
	// retrieve arguments and set default values if none are provided
	page, pErr := strconv.ParseInt(c.Query("page"), 10, 32)
	limit, lErr := strconv.ParseInt(c.Query("limit"), 10, 32)
	if pErr != nil {
		page = 1
	}
	if lErr != nil {
		limit = 10
	}

	// get query string
	// no need to make safe, mongo driver handles that
	reqQuery := c.Query("q")
	// search query for full text search
	query := bson.M{
		"$text": bson.M{
			"$search": reqQuery,
		},
	}

	db, client := Database()
	coll := db.Collection("users")

	// filter empty for all objects
	paginated, err := New(coll).Limit(limit).Page(page).Filter(query).Find()
	Handle(err)

	var userList []models.User
	for _, raw := range paginated.Data {
		var user *models.User
		// unmarshal each user into User struct
		if marshallErr := bson.Unmarshal(raw, &user); marshallErr == nil {
			userList = append(userList, *user)
		}
	}
	// close db connection
	client.Disconnect(context.TODO())

	return c.JSON(fiber.Map{
		"ok":     true,
		"length": len(userList),
		"users":  userList,
	})

}

/**
Indexes users and the "quote" field for full text search
mongodb's built in full text search isn't great

something like elastic search would work better in production env
**/
func IndexUsers(c *fiber.Ctx) error {
	// open db and get column
	db, client := Database()
	coll := db.Collection("users")
	index := []mongo.IndexModel{
		{Keys: bsonx.Doc{{Key: "quote", Value: bsonx.String("text")}}},
	}
	opts := options.CreateIndexes()
	// creates index based on given keys
	_, err := coll.Indexes().CreateMany(context.Background(), index, opts)
	Handle(err)

	client.Disconnect(context.TODO())
	return c.JSON(fiber.Map{
		"ok": true,
	})
}

/**
Creates user document
**/
func PostUser(c *fiber.Ctx) error {
	// parse the incoming JSON into a User model
	params := new(models.User)
	if err := c.BodyParser(params); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error()})
	}

	// create the user
	err := mgm.Coll(params).Create(params)
	Handle(err)

	return c.JSON(fiber.Map{
		"ok": true,
		// retrieves the hex ID from the default mgm model
		"id": params.DefaultModel.IDField.ID.Hex(),
	})
}
