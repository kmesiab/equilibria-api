package nrclex

import (
	"time"

	"gorm.io/gorm"
)

// Repository provides an interface to perform CRUD operations on NrcLex entities.
type Repository struct {
	DB *gorm.DB
}

// NewRepository creates a new instance of NrcLexRepository.
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

// Create inserts a new NrcLex into the database.
func (r *Repository) Create(nrcLex *NrcLex) error {
	return r.DB.Create(nrcLex).Error
}

// FindByID finds a NrcLex by its ID.
func (r *Repository) FindByID(id int64) (*NrcLex, error) {
	var nrcLex NrcLex

	err := r.DB.First(&nrcLex, id).Error

	if err != nil {
		return nil, err
	}

	return &nrcLex, nil
}

// Update updates an existing NrcLex.
func (r *Repository) Update(nrcLex *NrcLex) error {
	// This starts a new DB transaction.
	tx := r.DB.Begin()

	// Attempt to update the record that matches the given ID.
	// This will only update existing records and won't insert a new one if the record doesn't exist.
	result := tx.Model(&NrcLex{}).Where("id = ?", nrcLex.ID).Updates(nrcLex)

	// Check for other possible errors during the update.
	if err := result.Error; err != nil {
		tx.Rollback()
		return err
	}

	// Check if the record was found and updated. If not, rollback and return an error.
	if result.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	// Commit the transaction if everything was okay.
	return tx.Commit().Error
}

// Delete removes a NrcLex from the database.
func (r *Repository) Delete(id int64) error {
	return r.DB.Delete(&NrcLex{}, id).Error
}

// FindByMessageID finds a NrcLex by its Message ID.
func (r *Repository) FindByMessageID(messageID int64) (*NrcLex, error) {
	var nrcLex NrcLex
	result := r.DB.Where("message_id = ?", messageID).First(&nrcLex)
	if result.Error != nil {
		return nil, result.Error
	}
	return &nrcLex, nil
}

// FindByUserID finds all NrcLex entries associated with a given user_id.
func (r *Repository) FindByUserID(userID int64, limit int) ([]NrcLex, error) {
	var nrcLexes []NrcLex
	result := r.DB.Where("user_id = ?", userID).Find(&nrcLexes).Limit(limit)
	if result.Error != nil {
		return nil, result.Error
	}
	return nrcLexes, nil
}

func (r *Repository) FindRangeByUserID(userID int64, limit, offset int, startDate, endDate time.Time) (*[]NrcLex, error) {
	var nrcLexes []NrcLex

	result := r.DB.Where("user_id = ? AND created_at >= ? AND created_at <= ?", userID, startDate, endDate).
		Order("created_at desc"). // Assuming you might want to order the results; adjust as necessary
		Limit(limit).
		Offset(offset).
		Find(&nrcLexes)

	if result.Error != nil {

		return nil, result.Error
	}
	return &nrcLexes, nil
}
