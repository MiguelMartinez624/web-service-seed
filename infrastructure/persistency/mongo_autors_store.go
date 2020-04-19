package persistency

import (
	"context"
	"fmt"

	"github.com/miguelmartinez624/web-service-seed/domains/autors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBAutorsStore struct {
	db *mongo.Collection
}

func NewMongoDBAutorsStore(client *mongo.Database) *MongoDBAutorsStore {
	return &MongoDBAutorsStore{db: client.Collection("autors")}
}

func (s *MongoDBAutorsStore) GetAll(ctx context.Context) (autors []autors.Autor, err error) {
	cursor, err := s.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &autors)
	if err != nil {
		return nil, err
	}

	return autors, nil
}

func (s *MongoDBAutorsStore) GetByID(ctx context.Context, ID string) (autor *autors.Autor, err error) {
	var m autors.Autor

	err = s.db.FindOne(ctx, bson.M{"ID": ID}).Decode(&m)
	if err != nil {

		switch err.Error() {
		case "mongo: no documents in result":
			return nil, autors.AutorDontExistError
		default:
			return nil, err
		}

	}

	return &m, nil
}

func (s *MongoDBAutorsStore) Save(ctx context.Context, autor *autors.Autor) (ID string, err error) {
	result, err := s.db.InsertOne(ctx, autor)
	if err != nil {
		return "", err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return fmt.Sprintf("%v", result.InsertedID), nil
}
