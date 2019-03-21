package models

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	ErrorNotFound          = errors.New("Models: resource not found.")
	ErrorInvalidID         = errors.New("Models: Invalid ID.")
	ErrorIncorrectPassword = errors.New("Models: Incorrect password.")
	secretString           = "dirty-secret-string"
)

type UserService struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
}

func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return err
	}

	return err
}

func NewUserService(connectionString string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln(err.Error())
	}

	db.LogMode(true)
	return &UserService{
		db: db,
	}, nil
}

func (us *UserService) Close() error {
	return us.db.Close()
}

func (us *UserService) ByID(id uint) (*User, error) {
	var user User

	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *UserService) ByEmail(email string) (*User, error) {
	var user User

	db := us.db.Where("email = ?", email)
	err := first(db, &user)

	return &user, err
}

func (us *UserService) Create(user *User) error {
	pswBytes := []byte(user.Password + secretString)
	hashedBytes, err := bcrypt.GenerateFromPassword(pswBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedBytes)
	user.PasswordHash = ""

	return us.db.Create(user).Error
}

func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrorInvalidID
	}

	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}

func (us *UserService) AutoMigrate() error {
	if err := us.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}

	return nil
}

func (us *UserService) HardReset() error {
	err := us.db.DropTableIfExists(&User{}).Error
	if err != nil {
		return err
	}

	return us.AutoMigrate()
}

func (us *UserService) Authentication(email, password string) (*User, error) {
	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password+secretString))
	switch err {
	case nil:
		return foundUser, nil
	case bcrypt.ErrMismatchedHashAndPassword:
		return nil, ErrorIncorrectPassword
	default:
		return nil, err
	}
}
