package repository

import (
	"belajar-clean-arsitektur/internal/entity"
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userModel struct {
	ID        int       `gorm:"column:id;primaryKey"`
	Email     string    `gorm:"column:email;unique"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (u *userModel) TableName() string {
	return "users"
}

type UserRepository struct {
	Log *logrus.Logger
	DB  *gorm.DB
}

func NewUserRepository(log *logrus.Logger, db *gorm.DB) *UserRepository {
	return &UserRepository{
		Log: log,
		DB:  db,
	}
}

func (r *UserRepository) Create(ctx context.Context, userEntity *entity.User) error {
	user := toModel(userEntity)

	tx := r.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := tx.Create(&user).Error; err != nil {
		r.Log.WithError(err).Error("failed to create user")
		return entity.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		r.Log.WithError(err).Error("failed to commit transaction create user")
		return entity.ErrInternalServerError
	}

	return nil
}

func toModel(user *entity.User) *userModel {
	return &userModel{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
