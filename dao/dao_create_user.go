package dao

import (
	"GoFiberAPI/dbConfig"
	"GoFiberAPI/dto"
	"context"
)

func DB_CreatePreparedness(object *dto.User) error {

	_, err := dbConfig.DATABASE.Collection("User").InsertOne(context.Background(), object)
	if err != nil {
		return err
	}
	return nil
}