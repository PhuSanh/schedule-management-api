package handler

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"schedule-management-api/common"
	"schedule-management-api/model"
	"schedule-management-api/repository"
)

type AuthHandler struct {
	userRepo repository.UserRepo
}

func (a AuthHandler) Login(c echo.Context) (err error) {
	var body model.LoginForm
	if err = c.Bind(&body); err != nil {
		return RespondToClient(c, common.ERROR_REQUEST_DATA_INVALID, common.MSG_REQUEST_DATA_INVALID, err)
	}

	user := new(model.User)

	_ = a.userRepo.GetUserLogin(user, body)

	if user.ID == 0 {
		return RespondToClient(c, common.ERROR_GET_ROW_FROM_DB, common.MSG_GET_ROW_FROM_DB, err)
	}

	if ok := common.CheckPasswordHash(body.Password, user.Password); ok == false {
		return RespondToClient(c, common.ERROR_CHECK_PASSWORD, common.MSG_USERNAME_PASSWORD_NOT_MATCH, nil)
	}

	authClaims := &model.AuthClaims{
		Id: user.ID,
		Username: user.Username,
		Email: user.Email,
	}
	tokenString, err := a.createToken(authClaims)
	auth := model.AuthResponse{
		Token: tokenString,
		Username: body.Username,
	}
	return RespondToClient(c, common.ERROR_NO_ERORR, common.MSG_SUCEESS, auth)
}

func (a AuthHandler) Check(c echo.Context) (err error) {
	token := c.Request().Header.Get("Authorization")
	claims, _ := a.validateToken(token)
	return RespondToClient(c, 200, "Get token success", claims)
}

func (a AuthHandler) Register(c echo.Context) (err error) {
	body := new(model.RegisterForm)
	if err = c.Bind(&body); err != nil {
		return RespondToClient(c, common.ERROR_REQUEST_DATA_INVALID, common.MSG_REQUEST_DATA_INVALID, err)
	}

	hashedPass, _ := common.HashPassword(body.Password)
	user := &model.User{
		Username: body.Username,
		Password: hashedPass,
	}
	err = a.userRepo.CreateUser(user)
	if err != nil {
		return RespondToClient(c, common.ERROR_INSERT_ROW_TO_DB, common.MSG_INSERT_ROW_TO_DB, err)
	}
	return RespondToClient(c, common.ERROR_NO_ERORR, common.MSG_SUCEESS, user)
}

func (a AuthHandler) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if _, err := a.validateToken(token); err != "" {
			return RespondToClient(c, common.ERROR_UNAUTHORIZED, common.MSG_UNAUTHORIZED, err)
		}
		return next(c)
	}
}

func (a AuthHandler) createToken(data jwt.Claims) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), data)
	tokenString, err = token.SignedString([]byte("this_is_secret_key"))
	return tokenString, err
}

func (a AuthHandler) validateToken(tokenString string) (*model.AuthClaims, string) {
	token, err := jwt.ParseWithClaims(tokenString, &model.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("this_is_secret_key"), nil
	})

	if err != nil {
		return nil, err.Error()
	}

	if !token.Valid {
		return nil, "Token is not valid."
	}

	claims, _ := token.Claims.(*model.AuthClaims)
	return claims, ""
}