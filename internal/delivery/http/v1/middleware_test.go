package v1

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/magiconair/properties/assert"
	"github.com/maxzhovtyj/financeApp-server/pkg/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_userIdentity(t *testing.T) {
	testTable := []struct {
		name                 string
		userId               string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "OK",
			userId:               primitive.NewObjectID().Hex(),
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: "user authorized",
		},
	}

	manager, err := auth.NewManager("test")
	if err != nil {
		t.Fatal(err)
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			handler := Handler{
				service:      nil,
				tokenManager: manager,
			}

			accessToken, err := handler.tokenManager.NewJWT(testCase.userId, time.Minute)
			if err != nil {
				t.Fatal(err)
			}

			router := echo.New()

			router.GET("/test-handler", func(ctx echo.Context) error {
				userId := ctx.Get(userIdCtx)

				assert.Equal(t, userId, testCase.userId)

				return ctx.String(http.StatusOK, "user authorized")
			}, handler.userIdentity)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/test-handler", bytes.NewBufferString(""))
			req.Header.Set("Authorization", "Bearer "+accessToken)

			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}
}
