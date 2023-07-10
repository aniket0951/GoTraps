package repository

import "github.com/aniket0951/models"

var (
	userCollection = make(map[string]interface{})
)

type UserRepository interface{
	CreateUser(user models.User) int
}

type userRepository struct {
	DBMap map[string]interface{}
}

func NewUserRepository() UserRepository {
	return &userRepository{
		DBMap:userCollection,
	}
}

func (db *userRepository) CreateUser(user models.User) int {
	db.DBMap[user.Name] = user
	return user.Id
}