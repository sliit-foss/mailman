package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mailman/src/database"
)

type Repository[T any] struct {
	collection string
}

func NewRepository[T any](collection string) Repository[T] {
	return Repository[T]{collection: collection}
}

func (r Repository[T]) Create(payload T) primitive.ObjectID {
	result, err := database.UseDefault().Collection(r.collection).InsertOne(context.TODO(), payload)
	if err != nil {
		panic(err)
	}
	return result.InsertedID.(primitive.ObjectID)
}

func (r Repository[T]) FindByID(id primitive.ObjectID) T {
	user := new(T)
	doc := database.UseDefault().Collection(r.collection).FindOne(context.TODO(), primitive.M{"_id": id})
	doc.Decode(&user)
	return *user
}

func (r Repository[T]) FindAll() []T {
	var users []T
	cursor, err := database.UseDefault().Collection(r.collection).Find(context.Background(), primitive.M{})
	if err != nil {
		panic(err)
	}
	cursor.All(context.Background(), &users)
	return users
}

func (r Repository[T]) Update(id primitive.ObjectID, payload T) {
	_, err := database.UseDefault().Collection(r.collection).UpdateOne(context.Background(), primitive.M{"_id": id}, primitive.M{"$set": payload})
	if err != nil {
		panic(err)
	}
}

func (r Repository[T]) Delete(id primitive.ObjectID) {
	_, err := database.UseDefault().Collection(r.collection).DeleteOne(context.Background(), primitive.M{"_id": id})
	if err != nil {
		panic(err)
	}
}
