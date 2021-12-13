package adminServices

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/novanda1/sayurgo/models"
)

func Auth(body *models.Admin) (token string, message string, admin *models.Admin) {
	admin, err := GetAdminByPhone(*body.Phone)
	if err != nil {
		admin, err = CreateAdmin(*body)

		if err != nil {
			return "", "failed createadmin", admin
		}

		token, _ := SignToken(admin)

		return token, "successfully", admin
	}

	token, err = SignToken(admin)

	if err != nil {
		message = "failed to authenticated"
		return
	}

	message = "authenticated successfully"
	return
}

func SignToken(user *models.Admin) (string, error) {
	claims := jwt.MapClaims{
		"phone": user.Phone,
		"id":    user.ID,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("adminsecret"))

	return t, err
}
