package models

import "time"

type Otp struct {
	Email     string    `gorm:"primaryKey;not null"`        // Định nghĩa khóa chính
	Code      string    `gorm:"primaryKey;not null"`        // Mã OTP
	ExpiredAt time.Time `gorm:"not null;column:expired_at"` // Thời gian hết hạn
}
