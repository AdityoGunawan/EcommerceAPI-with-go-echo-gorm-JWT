package databases

import (
	"ecommerce-api/config"
	"ecommerce-api/middlewares"
	"ecommerce-api/models"

	"golang.org/x/crypto/bcrypt"
)

var user models.Users

// function database untuk menampilkan user by id
func GetUser(id int) (interface{}, error) {
	users := models.Users{}
	type get_user struct {
		Nama  string
		Email string
	}
	err := config.DB.Find(&users, id)
	rows_affected := config.DB.Find(&users, id).RowsAffected
	if err.Error != nil || rows_affected < 1 {
		return nil, err.Error
	}
	return get_user{users.Name, users.Email}, nil
}

// function database untuk menambahkan user baru (registrasi)
func CreateUser(user *models.Users) (interface{}, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// function database untuk menghapus user by id
func DeleteUser(id int) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// function database untuk memperbarui data user by id
func UpdateUser(id int, user *models.Users) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Updates(&user).Error; err != nil {
		return nil, err
	}
	config.DB.First(&user, id)
	return user, nil
}

// function login database untuk mendapatkan token
func LoginUser(plan_pass string, user *models.Users) (interface{}, error) {
	err := config.DB.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return nil, err
	}

	// cek plan password dengan hash password
	match := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plan_pass))
	if match != nil {
		return nil, match
	}
	user.Token, err = middlewares.CreateToken(int(user.ID)) // generate token
	if err != nil {
		return nil, err
	}
	if err = config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user.Token, nil
}
