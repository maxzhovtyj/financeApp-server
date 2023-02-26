package v1

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/magiconair/properties/assert"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"github.com/maxzhovtyj/financeApp-server/pkg/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_userIdentity(t *testing.T) {
	manager, err := auth.NewManager("test")
	if err != nil {
		t.Fatal(err)
	}

	userId := primitive.NewObjectID().Hex()

	accessToken, err := manager.NewJWT(userId, time.Minute)
	if err != nil {
		t.Fatal(err)
	}

	testTable := []struct {
		name                 string
		userId               string
		authorizationHeader  string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "OK",
			userId:               userId,
			authorizationHeader:  fmt.Sprintf("Bearer %s", accessToken),
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: "user authorized",
		},
		{
			name:                 "Empty auth header",
			userId:               userId,
			authorizationHeader:  "",
			expectedStatusCode:   http.StatusUnauthorized,
			expectedResponseBody: models.ErrInvalidAuthorizationHeader.Error(),
		},
		{
			name:                 "Empty access token in header",
			userId:               userId,
			authorizationHeader:  "Bearer ",
			expectedStatusCode:   http.StatusUnauthorized,
			expectedResponseBody: models.ErrInvalidAuthorizationHeader.Error(),
		},
		{
			name:                 "Missing bearer in auth header",
			userId:               userId,
			authorizationHeader:  " " + userId,
			expectedStatusCode:   http.StatusUnauthorized,
			expectedResponseBody: models.ErrInvalidAuthorizationHeader.Error(),
		},
		{
			name:                 "Invalid bearer in auth header",
			userId:               userId,
			authorizationHeader:  "Bear " + userId,
			expectedStatusCode:   http.StatusUnauthorized,
			expectedResponseBody: models.ErrInvalidAuthorizationHeader.Error(),
		},
		{
			name:                 "No space between bearer and token",
			userId:               userId,
			authorizationHeader:  "Bearer" + userId,
			expectedStatusCode:   http.StatusUnauthorized,
			expectedResponseBody: models.ErrInvalidAuthorizationHeader.Error(),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			handler := Handler{
				service:      nil,
				tokenManager: manager,
			}

			router := echo.New()

			router.GET("/test-handler", func(ctx echo.Context) error {
				userIdString, err := getUserIdFromContext(ctx)
				if err != nil {
					return newErrorResponse(ctx, http.StatusUnauthorized, err)
				}

				assert.Equal(t, userIdString, testCase.userId)

				return ctx.String(http.StatusOK, "user authorized")
			}, handler.userIdentity)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/test-handler", bytes.NewBufferString(""))
			req.Header.Set("Authorization", testCase.authorizationHeader)

			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}
}
