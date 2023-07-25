package user

import "github.com/syamsv/apollo-server/pkg/models"

type Interface interface {
	FetchProfileByEmail(email string) (*models.Users, error)
	CreateUser(user *models.Users) (*models.Users, error)
}
