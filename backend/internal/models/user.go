package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Streak struct {
	Current      int       `json:"current" bson:"current"`
	Longest      int       `json:"longest" bson:"longest"`
	LastActivity time.Time `json:"lastActivity" bson:"last_activity"`
}

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email        string             `json:"email" bson:"email"`
	Name         string             `json:"name" bson:"name"`
	AvatarURL    string             `json:"avatarUrl" bson:"avatar_url"`
	AuthProvider string             `json:"authProvider" bson:"auth_provider"`
	ProviderID   string             `json:"providerId,omitempty" bson:"provider_id,omitempty"`
	PasswordHash string             `json:"-" bson:"password_hash,omitempty"`
	Role         string             `json:"role" bson:"role"`
	Streak       Streak             `json:"streak" bson:"streak"`
	CreatedAt    time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updated_at"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	User        UserResponse `json:"user"`
	AccessToken string       `json:"accessToken"`
	ExpiresIn   int          `json:"expiresIn"`
}

type UserResponse struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Role   string `json:"role"`
}

type UserProfileResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
	Role      string `json:"role"`
	Streak    Streak `json:"streak"`
}
