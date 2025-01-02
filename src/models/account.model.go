package models

import (
	"database/sql"
	"time"
)

type Sex string

const (
	Nam  Sex = "Nam"
	Nữ   Sex = "Nữ"
	Khác Sex = "Khác"
)

// User đại diện cho bảng users trong cơ sở dữ liệu
type Account struct {
	ID             uint         `gorm:"primaryKey"`                                                             // Khóa chính
	FullName       string       `gorm:"column:fullName"`                                                        // Cột fullName
	Email          string       `gorm:"unique;column:email"`                                                    // Cột email và đảm bảo tính duy nhất
	Password       string       `gorm:"column:password"`                                                        // Cột password
	Status         string       `gorm:"column:status`                                                           // Cột status
	Adress         string       `gorm:"column:adress"`                                                          // Cột adress
	Phone          string       `gorm:"column:phone"`                                                           // Cột phone
	Avatar         string       `gorm:"column:avatar"`                                                          // Cột avatar
	Sex            string       `gorm:"column:sex;default:'Khác"`                                               // Cột giới tính
	Birthday       sql.NullTime `gorm:"column:birthday;type:date"`                                              // Cột birthday
	Deleted        bool         `gorm:"column:deleted;default:false"`                                           // Cột deleted
	CreatedAt      time.Time    `gorm:"column:createdAt;default:CURRENT_TIMESTAMP"`                             // Cột createdAt
	UpdatedAt      time.Time    `gorm:"column:updatedAt;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // Cột updatedAt
	RequestFriends []byte       `gorm:"column:requestFriends"`                                                  // Cột requestFriends (dạng JSON)
	AcceptFriends  []byte       `gorm:"column:acceptFriends"`                                                   // Cột acceptFriends (dạng JSON)
	StatusOnline   string       `gorm:"column:statusOnline"`                                                    // Cột statusOnline
	FriendList     []byte       `gorm:"column:friendList"`                                                      // Cột friendList (dạng JSON)
}
