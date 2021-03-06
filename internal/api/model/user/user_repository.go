package user

import (
	"fmt"
	mysql "golang-gin-api/pkg/db"
)

// Insert data into the user table
func (user *User) Insert() error {
	return mysql.DB.Model(&User{}).Create(&user).Error
}

// Realization of user registration function
func Register(username, pwd string, phone string, email string) error {
	fmt.Println(username, pwd, phone, email)

	if CheckUser(username) {
		return fmt.Errorf("%s", "used to already exist, please log in directly.")
	}

	// Construct user registration information
	user := User{
		Name:  username,
		Pwd:   pwd,
		Phone: phone,
		Email: email,
	}
	return user.Insert()
}

// Check user exist
func CheckUser(username string) bool {

	result := false
	var user User

	dbResult := mysql.DB.Where("name = ?", username).Find(&user)
	if dbResult.Error != nil {
		fmt.Printf("Failed to obtain user information:%v\n", dbResult.Error)
	} else {
		result = true
	}
	return result
}

// Check user information by username
func CheckUserByName(username string) (bool, User, error) {
	userData := User{}
	userExist := false

	var user User
	dbResult := mysql.DB.Where("name = ?", username).Find(&user)
	if dbResult.Error != nil {
		return userExist, userData, dbResult.Error
	} else {
		userExist = true
	}

	return userExist, user, nil
}
