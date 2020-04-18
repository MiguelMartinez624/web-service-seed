package persistency

import (
	"context"
	"fmt"

	"github.com/miguelmartinez624/web-service-seed/domains/movies"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBMoviesStore struct {
	db *mongo.Collection
}

func NewMongoDBMovieStore(client *mongo.Database) *MongoDBMoviesStore {
	return &MongoDBMoviesStore{db: client.Collection("movies")}
}

func (s *MongoDBMoviesStore) GetAll(ctx context.Context) (movies []movies.Movie, err error) {
	cursor, err := s.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &movies)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (s *MongoDBMoviesStore) GetByID(ctx context.Context, ID string) (movie *movies.Movie, err error) {
	var m movies.Movie

	err = s.db.FindOne(ctx, bson.M{"ID": ID}).Decode(&m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *MongoDBMoviesStore) Save(ctx context.Context, movie *movies.Movie) (ID string, err error) {
	result, err := s.db.InsertOne(ctx, movie)
	if err != nil {
		return "", err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return fmt.Sprintf("%v", result.InsertedID), nil
}
