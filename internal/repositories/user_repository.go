package repositories

import (
	"github.com/SalawatJoldasbaev/chat-app-golang/internal/interfaces"
	"github.com/SalawatJoldasbaev/chat-app-golang/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepositoryInterface {
	return &UserRepository{db: db}
}

func (u *UserRepository) Insert(user *models.User) (*models.User, error) {
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) IsEmailExists(email string, ignoreId string) bool {
	query := u.db.Where("email = ?", email)
	if ignoreId != "" {
		query = query.Where("id != ?", ignoreId)
	}
	if err := query.First(&models.User{}).Error; err != nil {
		return false
	}
	return true
}

func (u *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) FindById(id string) (*models.User, error) {
	var user models.User
	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) Update(id string, user *models.User) (*models.User, error) {
	if err := u.db.Model(&models.User{}).Where("id = ?", id).Clauses(clause.Returning{}).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) Delete(id string) error {
	if err := u.db.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) AddSession(session *models.Session) (*models.Session, error) {
	if err := u.db.Create(session).Error; err != nil {
		return nil, err
	}
	return session, nil
}

func (u *UserRepository) FindSessionByToken(token string) (*models.Session, error) {
	var session models.Session
	if err := u.db.Where("data = ?", token).First(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (u *UserRepository) DeleteSession(token string) error {
	if err := u.db.Where("data = ?", token).Delete(&models.Session{}).Error; err != nil {
		return err
	}
	return nil
}
