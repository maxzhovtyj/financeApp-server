package v1

import (
	"bytes"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/magiconair/properties/assert"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"github.com/maxzhovtyj/financeApp-server/internal/service"
	mock_service "github.com/maxzhovtyj/financeApp-server/internal/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO invalid test responses
func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUsers, user models.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           models.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"firstName":"test","lastName":"test","email":"test5@gmail.com","password":"qwerty"}`,
			inputUser: models.User{
				FirstName: "test",
				LastName:  "test",
				Email:     "test5@gmail.com",
				Password:  "qwerty",
			},
			mockBehavior: func(mockUsers *mock_service.MockUsers, user models.User) {
				mockUsers.EXPECT().SignUp(context.Background(), user).Return(nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			usersService := mock_service.NewMockUsers(controller)
			testCase.mockBehavior(usersService, testCase.inputUser)

			services := &service.Service{Users: usersService}
			handler := Handler{service: services}

			router := echo.New()
			router.POST(signUpUrl, handler.signUp)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, signUpUrl, bytes.NewBufferString(testCase.inputBody))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Accept", "application/json")

			// Make Request
			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}

func TestHandler_signIn(t *testing.T) {

}
