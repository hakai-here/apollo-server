package session

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/syamsv/apollo-server/pkg/models"
)

type SessionData struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	LastIp    string `json:"ip"`
}

func StoreData(sessionID string, userData string) error {
	if sessionID == "" || userData == "" {
		return errors.New("empty session ID or user data")
	}

	if err := manager.Verify.SetValue(sessionID, userData, time.Minute*15); err != nil {
		return fmt.Errorf("failed to store user data in session: %w", err)
	}
	return nil
}

func GetData(token string) (*models.Users, error) {
	if token == "" {
		return nil, errors.New("empty token")
	}

	data, err := manager.Verify.GetValue(token)
	if err != nil {
		return nil, fmt.Errorf("failed to get user data from session: %w", err)
	}

	user := new(models.Users)
	if err := json.Unmarshal([]byte(data), user); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user data: %w", err)
	}

	if err := manager.Verify.DeleteValue(token); err != nil {
		fmt.Println("failed to delete token from session:", err)
	}

	return user, nil
}

func GenerateSession(user *models.Users, Ip string) (string, error) {
	sessionId := uuid.NewString()
	sessionData := &SessionData{
		Id:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		LastIp:    Ip,
	}
	userData, err := json.Marshal(sessionData)
	if err != nil {
		return "", err
	}
	if err := manager.Auth.SetValue(sessionId, string(userData), time.Hour); err != nil {
		return "", err
	}

	return sessionId, nil

}

func GetSession(token string) (*SessionData, error) {
	sessionData := new(SessionData)
	data, err := manager.Auth.GetValue(token)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(data), sessionData); err != nil {
		return nil, err
	}
	return sessionData, nil
}
