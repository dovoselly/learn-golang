package dao

import (
	"context"
	"echo-app/database"
	"echo-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AdminFindByUsername(username string) (models.Admin, error) {
	var (
		adminCol = database.AdminCol()
		ctx      = context.Background()
		admin    models.Admin
	)

	// find
	filter := bson.M{
		"username": username,
	}
	err := adminCol.FindOne(ctx, filter).Decode(&admin)

	if err != nil {
		return admin, err
	}

	return admin, nil
}

func InitAdminAccount() {
	var (
		adminCol = database.AdminCol()
		ctx      = context.Background()
	)

	count, _ := adminCol.CountDocuments(ctx, bson.D{})

	if count == 0 {
		admin := models.Admin{
			ID:             primitive.NewObjectID(),
			Email:          "admin123@gmail.com",
			Username:       "admin123",
			HashedPassword: "123456",
			FullName:       "admin",
			DateOfBirth:    "2022-05-10",
			Gender:         "Male",
			Phone:          "0779925153",
		}
		adminCol.InsertOne(ctx, admin)
	}
}

func AdminProfileFindByID(ID string) (models.Admin, error) {
	var (
		adminCol     = database.AdminCol()
		ctx          = context.Background()
		adminProfile models.Admin
	)

	// objectID
	objID, _ := primitive.ObjectIDFromHex(ID)

	// find profile
	filter := bson.M{"_id": objID}
	err := adminCol.FindOne(ctx, filter).Decode(&adminProfile)

	// if err
	if err != nil {
		return adminProfile, err
	}

	return adminProfile, nil

}

func UpdateMyProfileAdmin(ID string, newProfile models.Admin) error {
	var (
		adminCol = database.AdminCol()
		ctx      = context.Background()
	)

	objID, _ := primitive.ObjectIDFromHex(ID)
	update := models.Admin{
		FullName:    newProfile.FullName,
		DateOfBirth: newProfile.DateOfBirth,
		Gender:      newProfile.Gender,
		Phone:       newProfile.Phone,
	}

	// UpdateOne
	_, err := adminCol.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})

	if err != nil {
		return err
	}

	return nil
}
