package auth

import (
	"2021_1_Noskool_team/internal/microservices/auth/models"
	"context"
)

type Repository interface {
	CreateSession(*models.Sessions) (*models.Sessions, error)
	CheckSession(*models.Sessions) (*models.Sessions, error)
	DeleteSession(*models.Sessions) error

	CreateAdminSession(context.Context, *models.Sessions) (*models.Sessions, error)
}
