package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/marioTiara/todolistapp/internal/api/handlers"
	"github.com/marioTiara/todolistapp/mocks"
)

func TestLogin(t *testing.T) {
	testCases := []struct {
		name             string
		setRequest       func(e *echo.Echo) *http.Request
		expectedResponse int
		expectedError    error
	}{
		{
			name: "StatusOK login success",
			setRequest: func(e *echo.Echo) *http.Request {
				f := make(url.Values)
				f.Set("username", "mario")
				f.Set("password", "mario2023")

				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

				return req
			},
			expectedResponse: http.StatusOK,
			expectedError:    nil,
		},
		{
			name: "unauthorized",
			setRequest: func(e *echo.Echo) *http.Request {
				f := make(url.Values)
				f.Set("username", "wrongUsername")
				f.Set("password", "shhh!")

				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

				return req
			},
			expectedResponse: http.StatusOK,
			expectedError:    echo.ErrUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			defer ctr.Finish()

			e := echo.New()
			req := tc.setRequest(e)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			//Instantiate mock service
			mockService := mocks.NewMockService(ctr)

			//Instantiate hanlders
			handler := handlers.NewHandlers(mockService)

			//Act
			err := handler.Login(c)

			//Assert
			assert.Equal(t, tc.expectedResponse, rec.Code)
			if tc.expectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}

}
