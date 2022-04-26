package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var token string

func TestLogin(t *testing.T) {
	e := echo.New()
	requestBody, _ := json.Marshal(map[string]interface{}{
		"email":    "alka@gmail.com",
		"password": "1234",
	})
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/auth")

	handler := New(&mockRepoUser{})
	handler.Login(context)

	type ResponseStructure struct {
		Code    int
		Message string
		Status  bool
		Data    interface{}
	}

	var response ResponseStructure

	json.Unmarshal([]byte(res.Body.Bytes()), &response)
	log.Warn(response.Data)
	assert.True(t, response.Status)
	assert.NotNil(t, response.Data)
	data := response.Data.(map[string]interface{})
	log.Warn(data)
	token = data["Token"].(string)
}

func TestInsert(t *testing.T) {

}

func TestGetAllUser(t *testing.T) {
	t.Run("Get All User", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userHandler := New(&mockRepoUser{})
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("RH$SI4")})(userController.GetAllUser())(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    []entity.User
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, resp.Data[0].Nama, "Alka")

	})
}

type mockRepoUser struct{}

func (mru *mockRepoUser) Insert(newUser entity.User) (entity.User, error) {
	return newUser, nil
}
func (mru *mockRepoUser) GetAll() ([]entity.User, error) {
	return []entity.User{{Nama: "Alka", Email: "alka@gmail.com"}}, nil
}
func (mru *mockRepoUser) Login(email string, password string) (entity.User, error) {
	return entity.User{Model: gorm.Model{ID: uint(1)}, email: "alka@gmail.com", password: "1234"}, nil
}
