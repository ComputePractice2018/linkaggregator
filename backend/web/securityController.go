package web

import (
	"../../src/model"
	"../../src/dao/userRepository"
	"../../src/config"
	"../../src/utils"
	"../../src/dto"
	"github.com/labstack/echo"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
	"time"
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

func Reg(requestContext echo.Context) error {

	login := requestContext.FormValue("login")
	password := requestContext.FormValue("password")
	email := requestContext.FormValue("email")

	if isCredentialsInvalid(login, password) || utils.IsEmpty(email) {
		return requestContext.JSON(http.StatusOK, dto.Error("Invalid credentials"))
	}

	if userRepository.IsUserExist(email, login) {
		return requestContext.JSON(http.StatusOK, dto.Error("User already exist"))
	}

	var user model.User

	user.Id = bson.NewObjectId()
	user.Email = email
	user.Login = login
	user.Registration = time.Now()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.Password = string(hashedPassword)

	if err := userRepository.Insert(user); err != nil {
		panic(err)
	}

	requestContext = utils.SetXAuthCookieValue(requestContext, generateToken(user.Login))

	return requestContext.JSON(http.StatusOK, dto.Success())
}

func Login(requestContext echo.Context) error {

	login := requestContext.FormValue("login")
	password := requestContext.FormValue("password")

	if isCredentialsInvalid(login, password) {
		return requestContext.JSON(http.StatusOK, dto.Error("Invalid credentials"))
	}

	user, err := userRepository.FindByLogin(login)

	if err != nil {
		return requestContext.JSON(http.StatusOK, dto.Error("User not found"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return requestContext.JSON(http.StatusOK, dto.Error("Incorrect password"))
	}

	requestContext = utils.SetXAuthCookieValue(requestContext, generateToken(user.Login))

	return requestContext.JSON(http.StatusOK, dto.Success())
}

func Principal(c echo.Context) error {
	login := c.Get("login")
	textLogin := fmt.Sprintf("%v", login)

	user, err := userRepository.FindByLogin(textLogin)

	if err != nil {
		return c.JSON(http.StatusOK, dto.Error("Unexpected error"))
	}

	return c.JSON(http.StatusOK, dto.Item(dto.UserInfo(user)))
}

func Logout(context echo.Context) error {
	context = utils.RemoveCookie(context)
	return context.JSON(http.StatusOK, dto.Success())
}

func isCredentialsInvalid(login, password string) bool {
	return utils.IsEmpty(login) || utils.IsEmpty(password)
}

func generateToken(login string) string {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["login"] = login
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	signedXAuthToken, err := token.SignedString([]byte(config.JWT_SALT))

	if err != nil {
		panic(err)
	}

	return signedXAuthToken
}
