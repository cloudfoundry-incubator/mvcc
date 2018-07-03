package mvcc

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	UUID        string
	AccessToken string
}

type Organization struct {
	Name string `json:"name"`
	UUID string `json:"guid"`
}

func RandomUUID(prefix string) (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%s", prefix, u), nil
}
