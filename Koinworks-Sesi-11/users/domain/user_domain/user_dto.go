package user_domain

import (
	"strings"
	"time"
	"users/utils/error_utils"

	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var secret_key = "SECRET_KEY"

type User struct {
	Id        int64     `json:"id"`
	Email     string    `json:"email" valid:"required~email is required"`
	Password  string    `json:"password" valid:"required~password is required"`
	Address   string    `json:"address" valid:"required~address is required"`
	Role      string    `json:"role" valid:"required~role is required"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) Validate() error_utils.MessageErr {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return error_utils.NewBadRequest(err.Error())
	}

	return nil
}

func (u *User) HashPass() {
	salt := 8
	password := []byte(u.Password)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	u.Password = string(hash)
}

func (u *User) ComparePassword(hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(u.Password))

	return err == nil
}

func (u *User) GenerateToken() string {
	claims := jwt.MapClaims{
		"id":    u.Id,
		"email": u.Email,
		"role":  u.Role,
		"exp":   time.Now().Add(time.Hour * 3).Unix(),
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	_ = parseToken

	signedToken, _ := parseToken.SignedString([]byte(secret_key))

	return signedToken
}

func (u *User) VerifyToken(tokenStr string) (jwt.MapClaims, error_utils.MessageErr) {
	if bearer := strings.HasPrefix(tokenStr, "Bearer"); !bearer {
		return nil, error_utils.NewNotAuthenticated("login to proceed")
	}

	stringToken := strings.Split(tokenStr, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, error_utils.NewNotAuthenticated("login to proceed")
		}
		return []byte(secret_key), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, error_utils.NewNotAuthenticated("login to proceed")
	}

	exp := int64(token.Claims.(jwt.MapClaims)["exp"].(float64))

	if (exp - time.Now().Unix()) <= 0 {
		return nil, error_utils.NewNotAuthenticated("login to proceed")
	}

	return token.Claims.(jwt.MapClaims), nil
}
