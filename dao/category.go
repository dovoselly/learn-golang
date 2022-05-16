package dao

import (
	"context"
	"echo-app/database"
	"echo-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCategory(category models.CategoryBSON) error {
	var (
		categoryCol = database.CategoryCol()
		ctx         = context.Background()
	)

	// InsertOne
	_, err := categoryCol.InsertOne(ctx, category)

	return err
}

func GetListCategory() ([]models.CategoryBSON, error) {
	var (
		categoryCol = database.CategoryCol()
		ctx         = context.Background()
		categories  []models.CategoryBSON
	)

	cursor, err := categoryCol.Find(ctx, bson.D{})
	if err != nil {
		return categories, err
	}

	if err = cursor.All(ctx, &categories); err != nil {
		return categories, nil
	}

	return categories, nil
}

func GetCategoryByID(ID primitive.ObjectID) (models.CategoryBSON, error) {
	var (
		categoryCol = database.CategoryCol()
		ctx         = context.Background()
		category    models.CategoryBSON
	)

	// find category
	filter := bson.M{"_id": ID}
	if err := categoryCol.FindOne(ctx, filter).Decode(&category); err != nil {
		return category, err
	}

	return category, nil
}

func UpdateCategoryByID(ID primitive.ObjectID, body models.CategoryUpdateBody) error {
	var (
		categoryCol = database.CategoryCol()
		ctx         = context.Background()
	)

	update := bson.M{"name": body.Name, "description": body.Description}

	// UpdateOne
	_, err := categoryCol.UpdateOne(ctx, bson.M{"_id": ID}, bson.M{"$set": update})

	if err != nil {
		return err
	}

	return nil
}

//
func DeleteCategoryByID(ID primitive.ObjectID) error {
	var (
		categoryCol = database.CategoryCol()
		ctx         = context.Background()
	)

	// filter
	filter := bson.M{"_id": ID}

	// DeleteOne
	if _, err := categoryCol.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}
