package service

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{}

func (u User) ChangePassword(ID primitive.ObjectID, body model.UserChangePassword) error {
	// check currentPassword
	userBSON, _ := userDAO.GetById(ID)
	if util.CheckPasswordHash(body.CurrentPassword, userBSON.Password) != nil {
		return errors.New("CurrentPassword is incorrect")
	}

	// HashPassword truoc khi update
	newPassword, _ := util.HashPassword(body.NewPassword)

	// update password
	err := userDAO.UpdatePassword(ID, newPassword)

	if err != nil {
		return err
	}

	return nil
}

func (u User) GetInfo(ID primitive.ObjectID) (model.UserInfo, error) {
	var (
		info model.UserInfo
	)

	// get user
	user, err := userDAO.GetInfo(ID)
	if err != nil {
		return info, err
	}

	// convert to userInfo
	info = model.UserInfo{
		ID:          user.ID,
		FullName:    user.FullName,
		Email:       user.Email,
		Username:    user.Username,
		Avatar:      user.Avatar,
		Gender:      user.Gender,
		DateOfBirth: user.DateOfBirth,
		Phone:       user.Phone,
		Address:     user.Address,
	}

	return info, nil
}

func (u User) UpdateInfo(ID primitive.ObjectID, body model.UserUpdate) error {

	bodyBSON := model.UserInfoBSON{
		FullName:    body.FullName,
		Email:       body.Email,
		Phone:       body.Phone,
		DateOfBirth: body.DateOfBirth,
		Gender:      body.Gender,
		Address:     body.Address,
	}

	// update info
	if err := userDAO.UpdateInfo(ID, bodyBSON); err != nil {
		return err
	}

	return nil
}
