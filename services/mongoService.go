package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"sbytes_v3/entities"
)

type MongoService struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoService(uri string) *MongoService {
	ms := &MongoService{}
	ms.Connect(uri)
	ms.ping()

	return ms
}

func (ms *MongoService) Connect(uri string) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatalln(err)
	}

	ms.client = client
	ms.collection = ms.client.Database("sbytes").Collection("tickets")
}

func (ms *MongoService) InsertTicket(ticket entities.Ticket) (interface{}, error) {
	_id, err := ms.collection.InsertOne(context.TODO(), ticket)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("Inserted ticket with id: ", _id)
	return _id.InsertedID, nil
}

func (ms *MongoService) FindTicketAsBsonDocument(id string) (bson.M, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	filter := bson.M{"_id": objectId}

	var ticket bson.M
	err = ms.collection.FindOne(context.TODO(), filter).Decode(&ticket)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return ticket, nil
}

func (ms *MongoService) ping() {
	err := ms.client.Ping(context.TODO(), readpref.Primary())

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Connected to MongoDB.")
}

func (ms *MongoService) Disconnect() {
	if err := ms.client.Disconnect(context.TODO()); err != nil {
		log.Println(err)
	}

	log.Println("Disconnected from MongoDB.")
}

func (ms *MongoService) UpdateTicket(uuid string, ticket entities.Ticket) (interface{}, interface{}) {
	objectId, err := primitive.ObjectIDFromHex(uuid)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	filter := bson.M{"_id": objectId}

	_, err = ms.collection.UpdateOne(context.TODO(), filter, ticket)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return objectId, nil
}
