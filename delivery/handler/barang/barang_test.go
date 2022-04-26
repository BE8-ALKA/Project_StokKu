package barang

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

func TestGetAllBarang(t *testing.T) {
	t.Run("Get All Barang", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/barangs")

		barangHandler := New(&mockRepoBarang{})
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("RH$SI4")})(barangController.GetAllTransaksi())(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    []entity.barang
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, resp.Data[0].Nama, "Fanta")

	})
}

type mockRepoBarang struct{}

func (mrb *mockRepoBarang) Insert(newBarang entity.Barang) (entity.Barang, error) {
	return newBarang, nil
}
func (mrb *mockRepoBarang) GetAll() ([]entity.Barang, error) {
	return []entity.Barang{{Nama: "Alka", Barang: "Lemonilo"}}, nil
}
func (mrb *mockRepoBarang) Login(email string, password string) (entity.Barang, error) {
	return entity.Barang{Model: gorm.Model{ID: uint(1)}, email: "alka@gmail.com", password: "1234"}, nil
}
