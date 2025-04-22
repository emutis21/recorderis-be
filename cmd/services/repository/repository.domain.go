package repository

import (
	"context"
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
		&models.Location{},
		&models.MemoryLocation{},
		&models.Tag{},
		&models.MemoryTag{},
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

	result := r.db.Where("memory_id = ?", memoryID).First(&memory).Preload("Descriptions").Preload("Photos")
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
	if err := r.db.Where("memory_id = ?", memory.MemoryID).First(&existingMemory).Error; err != nil {
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

	if memory.IsPublicPtr != nil {
		existingMemory.IsPublic = *memory.IsPublicPtr
	}

	if memory.UserID != "" {
		existingMemory.UserID = memory.UserID
	}

	if memory.IndexPtr != nil {
		existingMemory.Index = *memory.IndexPtr
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

/* descriptions */
func (r *Repository) GetDescriptions(ctx context.Context, memoryID string) ([]models.Description, error) {
	var descriptions []models.Description

	result := r.db.Where("memory_id = ?", memoryID).Find(&descriptions)
	if result.Error != nil {
		return nil, errors.NewError(errors.ErrDatabase, "Failed to fetch descriptions", result.Error)
	}

	return descriptions, nil
}

func (r *Repository) GetDescriptionByID(ctx context.Context, memoryID string, descriptionID string) (*models.Description, error) {
	var description models.Description

	result := r.db.Where("memory_id = ? AND description_id = ?", memoryID, descriptionID).First(&description)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("Description not found", result.Error)
		}
		return nil, errors.NewError(errors.ErrDatabase, "Failed to fetch description", result.Error)
	}

	return &description, nil
}

func (r *Repository) CreateDescription(ctx context.Context, memoryID string, description *models.Description) error {
	if result := r.db.Create(description); result.Error != nil {
		return errors.NewError(errors.ErrDatabase, "Failed to create description", result.Error)
	}

	return nil
}

func (r *Repository) UpdateDescription(ctx context.Context, memoryID string, description *models.Description) error {
	var existingDescription models.Description
	if err := r.db.Where("memory_id = ? AND description_id = ?", memoryID, description.DescriptionID).First(&existingDescription).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("Description not found", err)
		}

		return errors.NewError(errors.ErrDatabase, "Failed to fetch description", err)
	}

	if description.Text != "" {
		existingDescription.Text = description.Text
	}

	existingDescription.Index = description.Index

	existingDescription.Version++

	if result := r.db.Save(&existingDescription); result.Error != nil {
		return errors.NewError(errors.ErrDatabase, "Failed to update description", result.Error)
	}

	return nil
}

func (r *Repository) DeleteDescription(ctx context.Context, memoryID string, descriptionID string) error {
	var description models.Description
	if err := r.db.Where("memory_id = ? AND description_id = ?", memoryID, descriptionID).First(&description).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("Description not found", err)
		}

		return errors.NewError(errors.ErrDatabase, "Failed to fetch description", err)
	}

	result := r.db.Delete(&description)
	if result.Error != nil {
		return errors.NewError(errors.ErrDatabase, "Failed to delete description", result.Error)
	}

	return nil
}

/* locations & memory_locations */
func (r *Repository) GetLocations(ctx context.Context) ([]models.Location, error) {
	var locations []models.Location

	result := r.db.Find(&locations)
	if result.Error != nil {
		return nil, errors.NewError(errors.ErrDatabase, "Failed to fetch locations", result.Error)
	}

	return locations, nil
}

func (r *Repository) GetLocationByID(ctx context.Context, locationID string) (*models.Location, error) {
	var location models.Location

	result := r.db.Where("location_id = ?", locationID).First(&location)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("Location not found", result.Error)
		}
		return nil, errors.NewError(errors.ErrDatabase, "Failed to fetch location", result.Error)
	}

	return &location, nil
}

func (r *Repository) CreateLocation(ctx context.Context, location *models.Location) error {
	if result := r.db.Create(location); result.Error != nil {
		return errors.NewError(errors.ErrDatabase, "Failed to create location", result.Error)
	}

	return nil
}

func (r *Repository) UpdateLocation(ctx context.Context, location *models.Location) error {
	var existingLocation models.Location
	if err := r.db.Where("location_id = ?", location.LocationID).First(&existingLocation).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("Location not found", err)
		}

		return errors.NewError(errors.ErrDatabase, "Failed to fetch location", err)
	}

	if location.Location != "" {
		existingLocation.Location = location.Location
	}
	if location.Longitude != 0.0 {
		existingLocation.Longitude = location.Longitude
	}
	if location.Latitude != 0.0 {
		existingLocation.Latitude = location.Latitude
	}
	if location.City != "" {
		existingLocation.City = location.City
	}
	if location.Country != "" {
		existingLocation.Country = location.Country
	}

	if result := r.db.Save(&existingLocation); result.Error != nil {
		return errors.NewError(errors.ErrDatabase, "Failed to update location", result.Error)
	}

	return nil
}

func (r *Repository) DeleteLocation(ctx context.Context, locationID string) error {
	var location models.Location
	if err := r.db.Where("location_id = ?", locationID).First(&location).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("Location not found", err)
		}

		return errors.NewError(errors.ErrDatabase, "Failed to fetch location", err)
	}

	result := r.db.Delete(&location)
	if result.Error != nil {
		return errors.NewError(errors.ErrDatabase, "Failed to delete location", result.Error)
	}

	return nil
}

func (r *Repository) AssociateMemoryWithLocation(ctx context.Context, memoryID string, locationID string) error {
	var memory models.Memory
	if err := r.db.Where("memory_id = ?", memoryID).First(&memory).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("Memory not found", err)
		}

		return errors.NewError(errors.ErrDatabase, "Failed to fetch memory", err)
	}

	var location models.Location
	if err := r.db.Where("location_id = ?", locationID).First(&location).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("Location not found", err)
		}

		return errors.NewError(errors.ErrDatabase, "Failed to fetch location", err)
	}

	memoryLocation := models.MemoryLocation{
		MemoryID:   memory.MemoryID,
		LocationID: location.LocationID,
	}

	if result := r.db.Create(&memoryLocation); result.Error != nil {
		return errors.NewError(errors.ErrDatabase, "Failed to associate memory with location", result.Error)
	}

	return nil
}

func (r *Repository) DisassociateMemoryFromLocation(ctx context.Context, memoryID string, locationID string) error {
	var memory models.Memory
	if err := r.db.Where("memory_id = ?", memoryID).First(&memory).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("Memory not found", err)
		}

		return errors.NewError(errors.ErrDatabase, "Failed to fetch memory", err)
	}

	var location models.Location
	if err := r.db.Where("location_id = ?", locationID).First(&location).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("Location not found", err)
		}

		return errors.NewError(errors.ErrDatabase, "Failed to fetch location", err)
	}

	result := r.db.Where("memory_id = ? AND location_id = ?", memory.MemoryID, location.LocationID).Delete(&models.MemoryLocation{})
	if result.Error != nil {
		return errors.NewError(errors.ErrDatabase, "Failed to disassociate memory from location", result.Error)
	}

	return nil
}

func (r *Repository) GetLocationsByMemoryID(ctx context.Context, memoryID string) ([]models.Location, error) {
	var locations []models.Location

	result := r.db.Table("locations").
		Joins("JOIN memory_locations ON memory_locations.location_id = locations.location_id").
		Where("memory_locations.memory_id = ?", memoryID).
		Find(&locations)

	if result.Error != nil {
		return nil, errors.NewError(errors.ErrDatabase, "Failed to fetch locations", result.Error)
	}

	return locations, nil
}

func (r *Repository) GetMemoriesByLocationID(ctx context.Context, locationID string) ([]string, error) {
	var memoryLocations []models.MemoryLocation
	var memoryIDs []string

	result := r.db.Where("location_id = ?", locationID).Find(&memoryLocations)
	if result.Error != nil {
		return nil, errors.NewError(errors.ErrDatabase, "Failed to fetch memories", result.Error)
	}

	for _, memoryLocation := range memoryLocations {
		memoryIDs = append(memoryIDs, memoryLocation.MemoryID)
	}

	return memoryIDs, nil
}

/* tags & memory_tags */
func (r *Repository) GetTags(ctx context.Context) ([]models.Tag, error) {
	var tags []models.Tag

	result := r.db.Find(&tags)
	if result.Error != nil {
		return nil, errors.NewError(errors.ErrDatabase, "Failed to fetch tags", result.Error)
	}

	return tags, nil
}

func (r *Repository) GetTagByID(ctx context.Context, tagID string) (*models.Tag, error) {
	var tag models.Tag

	result := r.db.Where("tag_id = ?", tagID).First(&tag)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("Tag not found", result.Error)
		}
		return nil, errors.NewError(errors.ErrDatabase, "Failed to fetch tag", result.Error)
	}

	return &tag, nil
}

func (r *Repository) CreateTag(ctx context.Context, tag *models.Tag) error {
	if result := r.db.Create(tag); result.Error != nil {
		return errors.NewError(errors.ErrDatabase, "Failed to create tag", result.Error)
	}

	return nil
}

func (r *Repository) UpdateTag(ctx context.Context, tag *models.Tag) error {
	var existingTag models.Tag
	if err := r.db.Where("tag_id = ?", tag.TagID).First(&existingTag).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("Tag not found", err)
		}

		return errors.NewError(errors.ErrDatabase, "Failed to fetch tag", err)
	}

	if tag.Name != "" {
		existingTag.Name = tag.Name
	}

	if result := r.db.Save(&existingTag); result.Error != nil {
		return errors.NewError(errors.ErrDatabase, "Failed to update tag", result.Error)
	}

	return nil
}

func (r *Repository) DeleteTag(ctx context.Context, tagID string) error {
	var tag models.Tag
	if err := r.db.Where("tag_id = ?", tagID).First(&tag).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("Tag not found", err)
		}

		return errors.NewError(errors.ErrDatabase, "Failed to fetch tag", err)
	}

	result := r.db.Delete(&tag)
	if result.Error != nil {
		return errors.NewError(errors.ErrDatabase, "Failed to delete tag", result.Error)
	}

	return nil
}

func (r *Repository) AssociateMemoryWithTag(ctx context.Context, memoryID string, tagID string) error {
	var memory models.Memory
	if err := r.db.Where("memory_id = ?", memoryID).First(&memory).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("Memory not found", err)
		}

		return errors.NewError(errors.ErrDatabase, "Failed to fetch memory", err)
	}

	var tag models.Tag
	if err := r.db.Where("tag_id = ?", tagID).First(&tag).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("Tag not found", err)
		}

		return errors.NewError(errors.ErrDatabase, "Failed to fetch tag", err)
	}

	memoryTag := models.MemoryTag{
		MemoryID: memory.MemoryID,
		TagID:    tag.TagID,
	}

	if result := r.db.Create(&memoryTag); result.Error != nil {
		return errors.NewError(errors.ErrDatabase, "Failed to associate memory with tag", result.Error)
	}

	return nil
}

func (r *Repository) DisassociateMemoryFromTag(ctx context.Context, memoryID string, tagID string) error {
	var memory models.Memory
	if err := r.db.Where("memory_id = ?", memoryID).First(&memory).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("Memory not found", err)
		}

		return errors.NewError(errors.ErrDatabase, "Failed to fetch memory", err)
	}

	var tag models.Tag
	if err := r.db.Where("tag_id = ?", tagID).First(&tag).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("Tag not found", err)
		}

		return errors.NewError(errors.ErrDatabase, "Failed to fetch tag", err)
	}

	result := r.db.Where("memory_id = ? AND tag_id = ?", memory.MemoryID, tag.TagID).Delete(&models.MemoryTag{})
	if result.Error != nil {
		return errors.NewError(errors.ErrDatabase, "Failed to disassociate memory from tag", result.Error)
	}

	return nil
}

func (r *Repository) GetTagsByMemoryID(ctx context.Context, memoryID string) ([]models.Tag, error) {
	var tags []models.Tag

	result := r.db.Table("tags").
		Joins("JOIN memory_tags ON memory_tags.tag_id = tags.tag_id").
		Where("memory_tags.memory_id = ?", memoryID).
		Find(&tags)
	if result.Error != nil {
		return nil, errors.NewError(errors.ErrDatabase, "Failed to fetch tags", result.Error)
	}
	return tags, nil
}

func (r *Repository) GetMemoriesByTagID(ctx context.Context, tagID string) ([]string, error) {
	var memoryTags []models.MemoryTag
	var memoryIDs []string

	result := r.db.Where("tag_id = ?", tagID).Find(&memoryTags)
	if result.Error != nil {
		return nil, errors.NewError(errors.ErrDatabase, "Failed to fetch memories", result.Error)
	}

	for _, memoryTag := range memoryTags {
		memoryIDs = append(memoryIDs, memoryTag.MemoryID)
	}

	return memoryIDs, nil
}
