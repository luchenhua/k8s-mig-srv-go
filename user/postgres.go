package user

import "gorm.io/gorm"

func GetUser(db *gorm.DB, id int) (user User, err error) {
	db.Table("tbl_user").First(&user, id)
	return
}

func GetUsers(db *gorm.DB, query User) (users []User, err error) {
	db.Table("tbl_user").Where(&query).Find(&users)
	return
}

func CreateUser(db *gorm.DB, query User) (success bool, err error) {
	result := db.Table("tbl_user").Create(&query)
	switch result.RowsAffected {
	case 1:
		return true, nil
	default:
		return false, result.Error
	}
}

func UpdateUser(db *gorm.DB, query User, id int) (success bool, err error) {
	query.ID = id
	result := db.Table("tbl_user").Save(&query)
	switch result.RowsAffected {
	case 1:
		return true, nil
	default:
		return false, result.Error
	}
}
