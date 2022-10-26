package config

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hanafiadhi/MyGrams/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func ConnectDB() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return &DB{client: client}
}
func colHelper(db *DB, collectionName string) *mongo.Collection {
	return db.client.Database("myGrams").Collection(collectionName)
}
func (db *DB) Register(input *model.NewUser) (*model.User, error) {
	var err error
	collection := colHelper(db, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	CheckEmail := model.IsEmailValid(input.Email)
	if !CheckEmail {
		return nil, gqlerror.Errorf("Email was not valid")
	}

	filter := bson.M{"email": bson.M{"$eq": input.Email}}
	counts, err := collection.CountDocuments(ctx, filter)
	if counts != 0 {
		return nil, gqlerror.Errorf("Email was exist")
	}

	filter2 := bson.M{"username": bson.M{"$eq": input.Username}}
	counts, err = collection.CountDocuments(ctx, filter2)
	if counts != 0 {
		return nil, gqlerror.Errorf("Username was exist")
	}

	if input.Age <= 8 {
		return nil, gqlerror.Errorf("Age must equal or more than 8")
	}
	if len(input.Password) <= 6 {
		return nil, gqlerror.Errorf("password must equal or more than 6")
	}
	times := int(time.Now().Unix())

	hashedPassword, _ := model.HashPassword(input.Password)
	res, err := collection.InsertOne(ctx, bson.M{
		"username":   input.Username,
		"password":   hashedPassword,
		"email":      input.Email,
		"age":        input.Age,
		"created_at": times,
		"updated_at": times,
	})

	if err != nil {
		return nil, err
	}
	user := &model.User{
		ID:        res.InsertedID.(primitive.ObjectID).Hex(),
		Username:  input.Username,
		Email:     input.Email,
		Password:  input.Password,
		Age:       input.Age,
		CreatedAt: times,
		UpdatedAt: times,
	}
	return user, err
}
func (db *DB) Login(loginUSer *model.LoginUser) (*model.User, error) {
	// var err error
	collection := colHelper(db, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	CheckEmail := model.IsEmailValid(loginUSer.Email)
	if !CheckEmail {
		return nil, gqlerror.Errorf("Email was not valid")
	}
	filter := bson.M{"email": bson.M{"$eq": loginUSer.Email}}
	counts, err := collection.Find(ctx, filter)
	result := make([]model.User, 0)
	for counts.Next(ctx) {
		var row model.User
		err := counts.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}
		result = append(result, row)
	}
	if correct := model.CheckPasswordHash(loginUSer.Password, result[0].Password); correct == false {
		return nil, errors.New("Password Or Email is not Valid!")
	}
	return &result[0], err
}
func (db *DB) CreateUser(input *model.NewUser) (*model.User, error) {
	var err error
	collection := colHelper(db, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	CheckEmail := model.IsEmailValid(input.Email)
	if CheckEmail == false {
		return nil, gqlerror.Errorf("Email was not valid")
	}
	filter := bson.M{"email": bson.M{"$eq": input.Email}}
	counts, err := collection.CountDocuments(ctx, filter)
	if counts != 0 {
		return nil, gqlerror.Errorf("Email was exist")
	}
	filter2 := bson.M{"username": bson.M{"$eq": input.Username}}
	counts, err = collection.CountDocuments(ctx, filter2)
	if counts != 0 {
		return nil, gqlerror.Errorf("Username was exist")
	}
	if input.Age <= 8 {
		return nil, gqlerror.Errorf("Age must equal or more than 8")
	}
	if len(input.Password) <= 6 {
		return nil, gqlerror.Errorf("password must equal or more than 6")
	}
	times := int(time.Now().Unix())
	res, err := collection.InsertOne(ctx, bson.M{
		"username":   input.Username,
		"password":   input.Password,
		"email":      input.Email,
		"age":        input.Age,
		"created_at": times,
		"updated_at": times,
	})

	if err != nil {
		return nil, err
	}
	user := &model.User{
		ID:        res.InsertedID.(primitive.ObjectID).Hex(),
		Username:  input.Username,
		Email:     input.Email,
		Password:  input.Password,
		Age:       input.Age,
		CreatedAt: times,
		UpdatedAt: times,
	}
	// user.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return user, err
}
