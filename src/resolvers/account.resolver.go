package resolvers

import (
	"ThaiLy/configs"
	helper "ThaiLy/helpers"
	"ThaiLy/models"
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
)

func GetAccountResolver(p graphql.ResolveParams) (interface{}, error) {
	accountInterface := p.Context.Value("account")
	if accountInterface == nil {
		return nil, fmt.Errorf("Vui lòng đăng nhập")
	}
	account, ok := accountInterface.(models.Account)
	if !ok {
		return nil, fmt.Errorf("Tạo tài khoảng không thành côngcông")
	}

	fmt.Println(account)
	return map[string]interface{}{
		"id":       account.ID,
		"fullName": account.FullName,
		"email":    account.Email,
		"adress":   account.Adress,
		"phone":    account.Phone,
		"avatar":   account.Avatar,
		"sex":      account.Sex,
		"birthday": account.Birthday,
		"token":    "",
		"code":     200,
		"msg":      "",
	}, nil
}

func RegisterAccountResolver(p graphql.ResolveParams) (interface{}, error) {
	accountArgs, ok := p.Args["account"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Giá trị đầu vàovào không hợp lệ")
	}
	otp, otpOK := accountArgs["otp"].(string)
	if !otpOK {
		return nil, fmt.Errorf("Mã OTP không cócó lệ")
	}

	fullName, fullNameOK := accountArgs["fullName"].(string)
	email, emailOK := accountArgs["email"].(string)
	password, passwordOK := accountArgs["password"].(string)
	// Kiểm tra mã OTP
	var otpModel = models.Otp{}
	resultOTP := configs.GetDB().Where("code = ? AND email = ?", otp, email).First(&otpModel)
	if resultOTP.Error != nil {
		return nil, fmt.Errorf("Mã OTP không hợp lệ")
	}
	if !fullNameOK || !emailOK || !passwordOK {
		return nil, fmt.Errorf("Tên, email hoặc mật khẩu không hợp lệ")
	}
	account := models.Account{
		FullName: fullName,
		Email:    email,
		Password: password,
		Status:   "active",
	}
	fmt.Println(account)
	resultAccount := configs.GetDB().Create(&account)
	if resultAccount.Error != nil {
		return nil, fmt.Errorf("Đăng ký không thành công")
	}
	token := helper.CreateJWT(account.ID)
	resultOTP = configs.GetDB().Delete(&otpModel)
	if resultOTP.Error != nil {
		return nil, fmt.Errorf("Xóa mã OTP không thành công")
	}
	return map[string]interface{}{
		"id":       account.ID,
		"fullName": account.FullName,
		"email":    account.Email,
		"adress":   account.Adress,
		"phone":    account.Phone,
		"avatar":   account.Avatar,
		"sex":      account.Sex,
		"birthday": account.Birthday,
		"token":    token,
		"code":     200,
		"msg":      "Register success",
	}, nil
}

func LoginAccountResolver(p graphql.ResolveParams) (interface{}, error) {
	accountArgs, ok := p.Args["account"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Chưa nhập email hoặc mật khẩu")
	}
	email, emailOK := accountArgs["email"].(string)
	password, passwordOK := accountArgs["password"].(string)
	if !emailOK || !passwordOK {
		return nil, fmt.Errorf("email or password is invalid")
	}
	fmt.Println(email, password)
	var account = models.Account{}
	result := configs.GetDB().Where(
		"email = ? AND password = ?", email, password,
	).First(&account)

	if result.Error != nil {
		return nil, fmt.Errorf("invalid email or password")
	}
	token := helper.CreateJWT(account.ID)
	return map[string]interface{}{
		"id":       account.ID,
		"fullName": account.FullName,
		"email":    account.Email,
		"adress":   account.Adress,
		"phone":    account.Phone,
		"avatar":   account.Avatar,
		"sex":      account.Sex,
		"birthday": account.Birthday,
		"token":    token,
		"code":     200,
		"msg":      "Login success",
	}, nil
}

func UpdateAccountResolver(p graphql.ResolveParams) (interface{}, error) {
	accountInterface := p.Context.Value("account")
	if accountInterface == nil {
		return nil, fmt.Errorf("Vui lòng đăng nhậpnhập")
	}
	account, ok := accountInterface.(models.Account)
	if !ok {
		return nil, fmt.Errorf("Lỗi không xác định")
	}

	accountArgs, ok := p.Args["account"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Giá trị đầu vào không hợp lệ")
	}
	accountArgs["id"] = account.ID
	result := configs.GetDB().Model(&models.Account{}).Where("id = ?", account.ID).Updates(accountArgs)
	if result.Error != nil {
		return nil, fmt.Errorf("Cập nhật không thành công")
	}
	return map[string]interface{}{
		"id":       account.ID,
		"fullName": account.FullName,
		"adress":   account.Adress,
		"avatar":   account.Avatar,
		"sex":      account.Sex,
		"birthday": account.Birthday,
		"token":    "",
		"code":     200,
		"msg":      "Update success",
	}, nil
}

func CreateOtpResolver(p graphql.ResolveParams) (interface{}, error) {
	email, ok := p.Args["email"].(string)
	if !ok {
		return nil, fmt.Errorf("Giá trị đầu vào không hợp lệ")
	}
	resultAccount := configs.GetDB().Where("email = ?", email).First(&models.Account{})
	if resultAccount.Error == nil {
		return nil, fmt.Errorf("Email đã tồn tại")
	}
	otp := helper.RandomNumber(6)
	otpModel := models.Otp{
		Email:     email,
		Code:      otp,
		ExpiredAt: time.Now().Add(time.Minute * 5),
	}
	resultOTP := configs.GetDB().Create(&otpModel)
	if resultOTP.Error != nil {
		return nil, fmt.Errorf("Tạo mã OTP không thành công")
	}
	err := helper.SendMail(email, "Mã OTP", otp)
	if err != nil {
		return nil, fmt.Errorf("Gửi mail không thành công")
	}
	return map[string]interface{}{
		"code": 200,
		"msg":  "Create OTP success",
	}, nil

}
