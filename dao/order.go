package dao

import (
	"context"
	"echo-app/database"
	"echo-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllOrdersByUserId(ID primitive.ObjectID) ([]models.OrderBSON, error) {

	var (
		orders   []models.OrderBSON
		orderCol = database.OrderCol()
		ctx      = context.Background()
	)

	pipeline := []bson.M{
		{
			"$match": bson.M{"userId": ID},
		},
		{"$lookup": bson.M{
			"from":         "orderItems",
			"localField":   "items",
			"foreignField": "_id",
			"as":           "items",
		}},
	}

	cursor, err := orderCol.Aggregate(ctx, pipeline)

	err = cursor.All(context.Background(), &orders)

	if err != nil {
		return orders, err
	}

	return orders, nil
}

func CreateOrder(body models.OrderCreateBSON) error {
	var (
		orderCol = database.OrderCol()
		ctx      = context.Background()
	)

	// fmt.Println("BODY IN ORDER_DAO", body)

	// create order
	_, err := orderCol.InsertOne(ctx, body)
	if err != nil {
		return err
	}

	return nil
}