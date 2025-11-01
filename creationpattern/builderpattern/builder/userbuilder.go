package builder

import (
	"builderpattern/model"
	"time"
)

type UserBuilder struct {
	ID        string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *UserBuilder) SetID(id string) *UserBuilder {
	b.ID = id
	return b
}

func (b *UserBuilder) SetUsername(username string) *UserBuilder {
	b.Username = username
	return b
}

func (b *UserBuilder) SetEmail(email string) *UserBuilder {
	b.Email = email
	return b
}

func (b *UserBuilder) SetPassword(password string) *UserBuilder {
	b.Password = password
	return b
}

func (b *UserBuilder) SetCreatedAt(createdAt time.Time) *UserBuilder {
	b.CreatedAt = createdAt
	return b
}

func (b *UserBuilder) SetUpdatedAt(updatedAt time.Time) *UserBuilder {
	b.UpdatedAt = updatedAt
	return b
}

func (b *UserBuilder) Build() *model.User {
	return &model.User{
		ID:        b.ID,
		Username:  b.Username,
		Email:     b.Email,
		Password:  b.Password,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}
