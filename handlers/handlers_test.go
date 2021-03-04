package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockObject struct {
	mock.Mock
}

func (m *MockObject) GetCustomData(c *gin.Context) ([]byte, error) {
	args := m.Called(c)
	return args.Get(0).([]byte), args.Error(1)
}

func TestGetCustomData(t *testing.T) {
	//AAA: Arrange, Act Assert
	//Arrange
	mockClient := &MockObject{}
	gin.SetMode(gin.TestMode)
	err := errors.New("Error ficticio")
	//expectedResponse := `{"page":2,"per_page":6,"total":12,"total_pages":2,"data":[{"id":7,"email":"michael.lawson@reqres.in","first_name":"Michael","last_name":"Lawson","avatar":"https://reqres.in/img/faces/7-image.jpg"},{"id":8,"email":"lindsay.ferguson@reqres.in","first_name":"Lindsay","last_name":"Ferguson","avatar":"https://reqres.in/img/faces/8-image.jpg"},{"id":9,"email":"tobias.funke@reqres.in","first_name":"Tobias","last_name":"Funke","avatar":"https://reqres.in/img/faces/9-image.jpg"},{"id":10,"email":"byron.fields@reqres.in","first_name":"Byron","last_name":"Fields","avatar":"https://reqres.in/img/faces/10-image.jpg"},{"id":11,"email":"george.edwards@reqres.in","first_name":"George","last_name":"Edwards","avatar":"https://reqres.in/img/faces/11-image.jpg"},{"id":12,"email":"rachel.howell@reqres.in","first_name":"Rachel","last_name":"Howell","avatar":"https://reqres.in/img/faces/12-image.jpg"}],"support":{"url":"https://reqres.in/#support-heading","text":"To keep ReqRes free, contributions towards server costs are appreciated!"}}`
	//mockClient.On("GetCustomData", mock.AnythingOfType("*gin.Context")).Return([]byte(expectedResponse), nil)
	mockClient.On("GetCustomData", mock.AnythingOfType("*gin.Context")).Return([]byte{}, err)

	//Act
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	handler := gin.HandlerFunc(RequestWrapper(mockClient.GetCustomData))
	router := gin.Default()
	router.Use(handler)
	router.ServeHTTP(rec, req)

	//Assert
	assert.Equal(t, rec.Body.String(), "{\"status\":\"OK\"}")
	mockClient.AssertExpectations(t)
}
