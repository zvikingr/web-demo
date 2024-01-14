package dao

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	userTable = "user"

	// AdminUserName admin user's name
	AdminUserName = "admin"
	// AdminPassword admin user's password
	AdminPassword = "123456"
)

// User user table
type User struct {
	ID        uint      `gorm:"column:id;primary_key;autoIncrement;not null"`
	UserName  string    `gorm:"column:user_name;type:varchar(20);unique;not null"`
	Password  string    `gorm:"column:password;type:varchar(256);not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}

// TableName return user table name
func (User) TableName() string {
	return userTable
}

// UserDB user gorm db
var UserDB *gorm.DB

// InitUserDB init user gorm db
func InitUserDB() error {
	UserDB = db.Table(userTable).Debug()

	// init admin user
	user, err := GetUser(AdminUserName)
	if err != nil {
		return err
	}

	if user.UserName == AdminUserName {
		return nil
	}

	passwd, err := bcrypt.GenerateFromPassword([]byte(AdminPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = CreateUser(&User{UserName: AdminUserName, Password: string(passwd)})
	if err != nil {
		return err
	}

	return nil
}

// CreateUser create user in database
func CreateUser(u *User) error {
	tx := UserDB.Create(u)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected != 1 {
		return errors.New("rows affected != 1")
	}

	return nil
}

// GetUser get user from database
func GetUser(userName string) (*User, error) {
	user := new(User)

	tx := UserDB.Where("user_name = ?", userName).Find(user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}

// UpdateUser update user in database
func UpdateUser(username, password string) error {
	tx := UserDB.Where("user_name = ?", username).Updates(User{Password: password})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// DeleteUser delete user in database
func DeleteUser(userName string) error {
	tx := UserDB.Where("user_name = ?", userName).Unscoped().Delete(&User{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
