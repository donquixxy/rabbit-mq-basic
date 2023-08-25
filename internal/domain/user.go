package domain

import (
	"encoding/json"
	"time"
)

type User struct {
	ID        string    `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Age       string    `json:"age" gorm:"column:age"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (u *User) Unmarshall(data []byte) error {
	return json.Unmarshal(data, &u)
}

func (*User) TableName() string {
	return "user"
}
