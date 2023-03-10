package v1

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/magiconair/properties/assert"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"github.com/maxzhovtyj/financeApp-server/internal/service"
	mock_service "github.com/maxzhovtyj/financeApp-server/internal/service/mocks"
	"github.com/maxzhovtyj/financeApp-server/pkg/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(r *mock_service.MockUsers, input models.User)

	testTable := []struct {
		name                 string
		requestBody          string
		serviceInput         models.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			requestBody: `{"firstName":"test","lastName":"test","email":"test@gmail.com","password":"qwerty123"}`,
			serviceInput: models.User{
				FirstName: "test",
				LastName:  "test",
				Email:     "test@gmail.com",
				Password:  "qwerty123",
			},
			mockBehavior: func(r *mock_service.MockUsers, input models.User) {
				r.EXPECT().SignUp(context.Background(), input).Return(nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:                 "Empty user firstName",
			requestBody:          `{"firstName":"","lastName":"test","email":"test@gmail.com","password":"qwerty123"}`,
			mockBehavior:         func(r *mock_service.MockUsers, input models.User) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: models.ErrInvalidInputBody.Error(),
		},
		{
			name:                 "Invalid user email",
			requestBody:          `{"firstName":"test","lastName":"test","email":"test@","password":"qwerty123"}`,
			mockBehavior:         func(r *mock_service.MockUsers, input models.User) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: models.ErrInvalidInputBody.Error(),
		},
		{
			name:                 "Invalid user password length",
			requestBody:          `{"firstName":"test","lastName":"test","email":"test@gmail.com","password":"123"}`,
			mockBehavior:         func(r *mock_service.MockUsers, input models.User) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: models.ErrInvalidInputBody.Error(),
		},
		{
			name:                 "Missing user password",
			requestBody:          `{"firstName":"test","lastName":"test","email":"test@gmail.com","password":""}`,
			mockBehavior:         func(r *mock_service.MockUsers, input models.User) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: models.ErrInvalidInputBody.Error(),
		},
		{
			name:                 "Missing user lastName",
			requestBody:          `{"firstName":"test","lastName":"","email":"test@test.com","password":"qwerty123"}`,
			mockBehavior:         func(r *mock_service.MockUsers, input models.User) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: models.ErrInvalidInputBody.Error(),
		},
		{
			name:                 "Missing request body",
			mockBehavior:         func(r *mock_service.MockUsers, input models.User) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: models.ErrInvalidInputBody.Error(),
		},
		{
			name:        "User already exists",
			requestBody: `{"firstName":"test","lastName":"test","email":"test@gmail.com","password":"qwerty123"}`,
			serviceInput: models.User{
				FirstName: "test",
				LastName:  "test",
				Email:     "test@gmail.com",
				Password:  "qwerty123",
			},
			mockBehavior: func(r *mock_service.MockUsers, input models.User) {
				r.EXPECT().SignUp(context.Background(), input).Return(models.ErrUserAlreadyExists)
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: models.ErrUserAlreadyExists.Error(),
		},
		{
			name:        "Some server error",
			requestBody: `{"firstName":"test","lastName":"test","email":"test@gmail.com","password":"qwerty123"}`,
			serviceInput: models.User{
				FirstName: "test",
				LastName:  "test",
				Email:     "test@gmail.com",
				Password:  "qwerty123",
			},
			mockBehavior: func(r *mock_service.MockUsers, input models.User) {
				r.EXPECT().SignUp(context.Background(), input).Return(fmt.Errorf("some error occurred"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: fmt.Errorf("some error occurred").Error(),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			usersService := mock_service.NewMockUsers(controller)

			testCase.mockBehavior(usersService, testCase.serviceInput)

			s := &service.Service{Users: usersService}
			handler := Handler{service: s}

			router := echo.New()

			router.GET(signUpUrl, handler.signUp)

			v := validator.New()
			router.Validator = &AppValidator{Validator: v}

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, signUpUrl, bytes.NewBufferString(testCase.requestBody))
			r.Header.Set("Content-Type", "application/json")
			r.Header.Set("Accept", "application/json")

			router.ServeHTTP(w, r)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}
}

func TestHandler_signIn(t *testing.T) {
	type signInServiceInput struct {
		Email    string
		Password string
	}
	type mockBehavior func(r *mock_service.MockUsers, input signInServiceInput)

	manager, err := auth.NewManager("test")
	if err != nil {
		t.Fatal(err)
	}

	refreshToken, err := manager.NewRefreshToken()
	if err != nil {
		t.Fatal(err)
	}

	accessToken, err := manager.NewJWT(primitive.NewObjectID().Hex(), time.Minute)
	if err != nil {
		t.Fatal(err)
	}

	testTable := []struct {
		name                 string
		requestBody          string
		mockBehavior         mockBehavior
		serviceInput         signInServiceInput
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			requestBody: `{"email":"test@gmail.com","password":"qwerty123"}`,
			mockBehavior: func(r *mock_service.MockUsers, input signInServiceInput) {
				r.EXPECT().SignIn(context.Background(), input.Email, input.Password).Return(accessToken, refreshToken, nil)
			},
			serviceInput: signInServiceInput{
				Email:    "test@gmail.com",
				Password: "qwerty123",
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: fmt.Sprintf(`{"accessToken":"%s","refreshToken":"%s"}%s`, accessToken, refreshToken, "\n"),
		},
		{
			name:                 "Missing email",
			requestBody:          `{"email":"","password":"qwerty123"}`,
			mockBehavior:         func(r *mock_service.MockUsers, input signInServiceInput) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: models.ErrInvalidInputBody.Error(),
		},
		{
			name:                 "Missing password",
			requestBody:          `{"email":"test@gmail.com","password":""}`,
			mockBehavior:         func(r *mock_service.MockUsers, input signInServiceInput) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: models.ErrInvalidInputBody.Error(),
		},
		{
			name:        "User was not found",
			requestBody: `{"email":"test@gmail.com","password":"qwerty123"}`,
			mockBehavior: func(r *mock_service.MockUsers, input signInServiceInput) {
				r.EXPECT().SignIn(context.Background(), input.Email, input.Password).Return("", "", models.ErrUserNotFound)
			},
			serviceInput: signInServiceInput{
				Email:    "test@gmail.com",
				Password: "qwerty123",
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: models.ErrUserNotFound.Error(),
		},
		{
			name:        "Some server error",
			requestBody: `{"email":"test@gmail.com","password":"qwerty123"}`,
			mockBehavior: func(r *mock_service.MockUsers, input signInServiceInput) {
				r.EXPECT().SignIn(context.Background(), input.Email, input.Password).Return("", "", fmt.Errorf("error occurred"))
			},
			serviceInput: signInServiceInput{
				Email:    "test@gmail.com",
				Password: "qwerty123",
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: fmt.Errorf("error occurred").Error(),
		},
		{
			name:                 "Empty request body",
			mockBehavior:         func(r *mock_service.MockUsers, input signInServiceInput) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: models.ErrInvalidInputBody.Error(),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)

			usersService := mock_service.NewMockUsers(controller)

			testCase.mockBehavior(usersService, testCase.serviceInput)

			s := &service.Service{Users: usersService}
			handler := Handler{service: s}

			router := echo.New()

			router.GET(signUpUrl, handler.signIn)

			v := validator.New()
			router.Validator = &AppValidator{Validator: v}

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, signUpUrl, bytes.NewBufferString(testCase.requestBody))
			r.Header.Set("Content-Type", "application/json")
			r.Header.Set("Accept", "application/json")

			router.ServeHTTP(w, r)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}
}
