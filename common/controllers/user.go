package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/syamsv/apollo-server/common/database"
	"github.com/syamsv/apollo-server/common/session"
	"github.com/syamsv/apollo-server/pkg/mailer"
	"github.com/syamsv/apollo-server/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(email string, password string, Ip string) (string, error) {
	user, err := database.User.FetchProfileByEmail(email)
	if err != nil {
		return "", fmt.Errorf("failed to fetch user profile: %w", err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}
	sessionId, err := session.GenerateSession(user, Ip)
	if err != nil {
		return "", err
	}
	return sessionId, nil
}

func CreateUser(user *models.Users) error {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = string(passwordHash)

	userData, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user data: %w", err)
	}

	sessionID := uuid.NewString()

	if err := session.StoreData(sessionID, string(userData)); err != nil {
		return fmt.Errorf("failed to store user data in session: %w", err)
	}

	go func() {
		if err := mailer.SendActivationMail(user.Email, mailer.ReturnHtmlTemplate(sessionID)); err != nil {
			log.Println(err.Error())
		}
	}()

	return nil
}

func ActivateUserController(token string) error {
	if token == "" {
		return errors.New("empty token")
	}

	user, err := session.GetData(token)
	if err != nil {
		return fmt.Errorf("failed to get user data from session: %w", err)
	}

	if _, err := database.User.CreateUser(user); err != nil {
		return fmt.Errorf("failed to create user in the database: %w", err)
	}

	return nil
}
