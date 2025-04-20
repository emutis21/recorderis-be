package repository

import (
	"context"
	"fmt"
	"recorderis/cmd/services/repository/migrations"
	"recorderis/cmd/services/repository/models"
	"recorderis/internals/errors"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() (*Repository, error) {
	dsn := "host=localhost user=recorderis_user password=recorderis_pass dbname=recorderis_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	if err := migrations.InitDB(db); err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Memory{},
		&models.Description{},
		&models.Photo{},
	)

	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

/* users */
func (r *Repository) GetUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User

	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r *Repository) CreateUser(ctx context.Context, user *models.User) error {
	if user.Email == "" || user.PasswordHash == "" {
		return errors.NewValidationError("Email and password are required", nil)
	}

	var existingUser models.User
	if result := r.db.Where("email = ?", user.Email).First(&existingUser); result.Error == nil {
		return errors.NewValidationError("Email already exists", nil)
	}

	if result := r.db.Create(user); result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			return errors.NewValidationError("Email already exists", result.Error)
		}

		return errors.NewError(errors.ErrDatabase, "Failed to create user", result.Error)
	}

	return nil
}

func (r *Repository) GetUserById(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("User not found", result.Error)
		}

		return nil, errors.NewError(errors.ErrDatabase, "Failed to fetch user", result.Error)
	}

	return &user, nil
}

func (r *Repository) FindUserByUserID(ctx context.Context, userID string) (*models.User, error) {
	var user models.User
	result := r.db.Where("user_id = ?", userID).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("User not found", result.Error)
		}
		return nil, errors.NewError(errors.ErrDatabase, "Failed to fetch user", result.Error)
	}

	return &user, nil
}

func (r *Repository) UpdateUser(ctx context.Context, user *models.User) error {
	var existingUser models.User
	if err := r.db.First(&existingUser, user.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("User not found", err)
		}

		return errors.NewError(errors.ErrDatabase, "Failed to fetch user", err)
	}

	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if user.PasswordHash != "" {
		existingUser.PasswordHash = user.PasswordHash
	}

	result := r.db.Save(&existingUser)
	if result.Error != nil {
		return errors.NewError(errors.ErrDatabase, "Failed to update user", result.Error)
	}

	return nil
}

func (r *Repository) DeleteUser(ctx context.Context, id int) error {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("User not found", err)
		}

		return errors.NewError(errors.ErrDatabase, "Failed to fetch user", err)
	}

	result := r.db.Delete(&user)
	if result.Error != nil {
		return errors.NewError(errors.ErrDatabase, "Failed to delete user", result.Error)
	}

	return nil
}

func (r *Repository) GetDB() *gorm.DB {
	return r.db
}

func (r *Repository) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("User not found", result.Error)
		}
		return nil, errors.NewError(errors.ErrDatabase, "Failed to fetch user", result.Error)
	}

	return &user, nil
}

/* memories */
func (r *Repository) CreateMemory(ctx context.Context, memory *models.Memory) error {
	fmt.Println("Creating memory:", memory)
	if result := r.db.Create(memory); result.Error != nil {
		return errors.NewError(errors.ErrDatabase, "Failed to create memory", result.Error)
	}

	return nil
}

func (r *Repository) GetMemories(ctx context.Context, userID string) ([]models.Memory, error) {
	var memories []models.Memory

	result := r.db.Where("user_id = ?", userID).Find(&memories)
	if result.Error != nil {
		return nil, errors.NewError(errors.ErrDatabase, "Failed to fetch memories", result.Error)
	}

	return memories, nil
}

func (r *Repository) GetMemoryByMemoryID(ctx context.Context, memoryID string) (*models.Memory, error) {
	var memory models.Memory

	result := r.db.Where("memory_id = ?", memoryID).First(&memory)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("Memory not found", result.Error)
		}
		return nil, errors.NewError(errors.ErrDatabase, "Failed to fetch memory", result.Error)
	}

	return &memory, nil
}

func (r *Repository) UpdateMemory(ctx context.Context, memory *models.Memory) error {
	var existingMemory models.Memory
	if err := r.db.First(&existingMemory, memory.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("Memory not found", err)
		}

		return errors.NewError(errors.ErrDatabase, "Failed to fetch memory", err)
	}

	if memory.Title != "" {
		existingMemory.Title = memory.Title
	}

	if !memory.Date.IsZero() {
		existingMemory.Date = memory.Date
	}

	if memory.IsPublic {
		existingMemory.IsPublic = memory.IsPublic
	}

	if memory.UserID != "" {
		existingMemory.UserID = memory.UserID
	}

	if memory.Index != 0 {
		existingMemory.Index = memory.Index
	}

	if memory.Descriptions != nil {
		existingMemory.Descriptions = memory.Descriptions
	}

	if memory.Photos != nil {
		existingMemory.Photos = memory.Photos
	}

	if result := r.db.Save(&existingMemory); result.Error != nil {
		return errors.NewError(errors.ErrDatabase, "Failed to update memory", result.Error)
	}

	return nil
}

func (r *Repository) DeleteMemory(ctx context.Context, memoryID string) error {
	var memory models.Memory
	if err := r.db.Where("memory_id = ?", memoryID).First(&memory).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("Memory not found", err)
		}

		return errors.NewError(errors.ErrDatabase, "Failed to fetch memory", err)
	}

	result := r.db.Delete(&memory)
	if result.Error != nil {
		return errors.NewError(errors.ErrDatabase, "Failed to delete memory", result.Error)
	}

	return nil
}
