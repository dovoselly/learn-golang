package service

import (
	"echo-app/dao"
	"echo-app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

func ListReview(productId primitive.ObjectID, query model.ReviewQuery) ([]model.Review, error) {
	filter := bson.M{"productId": productId}

	if query.Rating != "" {
		rating, _ := strconv.ParseInt(query.Rating, 10, 64)
		filter["rating"] = rating
	}

	optionsQuery := new(options.FindOptions)
	optionsQuery.SetSkip(query.Page * limit)
	optionsQuery.SetLimit(limit)
	if query.Sort != "" {
		var value int
		if string([]rune(query.Sort)[0]) != "-" {
			value = -1
		} else {
			value = 1
		}
		sortMap := map[string]interface{}{
			"price": value,
		}
		optionsQuery.SetSort(sortMap)
	}

	results, err := dao.ListReview(filter, optionsQuery)
	return results, err
}

func CreateReview(userId primitive.ObjectID, productId primitive.ObjectID, body model.CreateReview) error {
	//init insert data
	insertData := model.Review{
		ID:        primitive.NewObjectID(),
		UserId:    userId,
		ProductId: productId,
		Rating:    body.Rating,
		Content:   body.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := dao.CreateReview(insertData)
	return err
}