package Services

import (
	"context"
	"encoding/json"
	"log"
	"th3y3m/e-commerce-platform/Repositories"

	"github.com/hibiken/asynq"
)

const (
	TypeDeleteUser = "delete:user"
)

type DeleteUserPayload struct {
	UserID string
}

func NewDeleteUserTask(userID string) (*asynq.Task, error) {
	payload, err := json.Marshal(DeleteUserPayload{UserID: userID})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeDeleteUser, payload), nil
}

func HandleDeleteUserTask(ctx context.Context, t *asynq.Task) error {
	var p DeleteUserPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	log.Printf("Deleting user with ID: %s", p.UserID)

	if err := Repositories.DeleteUser(p.UserID); err != nil {
		return err
	}

	return nil
}
