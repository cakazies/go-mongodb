package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/local/go-mongo/models"
	"github.com/local/testify/assert"
)

type (
	testCase struct {
		name         string
		input        string
		expectedData string
		expectedCode int
		path         string
		handler      func(echo.Context) error
		query        string
	}
	Response struct {
		Response Rest                   `json:"response"`
		Data     map[string]interface{} `json:"data,omitempty"`
	}
	Rest struct {
		Message string `json:"message,omitempty"`
		Code    string `json:"code,omitempty"`
	}
)

var (
	userJSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`
)

func TestGetRoom(t *testing.T) {
	tasks := []testCase{
		{
			name:         "Testing with student id",
			input:        "5d283a781ce2e869fcf97d8b",
			expectedData: "5d283a781ce2e869fcf97d8b",
			expectedCode: http.StatusOK,
			path:         "api/student",
			handler:      models.GetStudent,
			query:        "",
		},
		{
			name:         "Testing with random id",
			input:        "9897",
			expectedData: "<nil>", // because not value
			expectedCode: http.StatusBadRequest,
			path:         "api/student",
			handler:      models.GetStudent,
			query:        "",
		},
	}

	for _, tc := range tasks {
		t.Run(tc.name, func(t *testing.T) {
			// url := fmt.Sprintf("http://%s/%s/%s", "127.0.0.1:8000", tc.path, tc.input)
			// url := fmt.Sprintf("/%s/%s", tc.path, tc.input)
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/api/student/:id")
			c.SetParamNames("id")
			c.SetParamValues(tc.input)
			// Assertions
			if assert.NoError(t, models.GetStudent(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)
				assert.Equal(t, userJSON, rec.Body.String())
			}

			buf := rec.Body.Bytes()
			var respData Response
			if err := json.Unmarshal(buf, &respData); err != nil {
				t.Error("Can not parsing response testing. Error :", err)
			}
			// getData := fmt.Sprintf("%v", respData.Data["rm_id"])
			// assert.Equal(t, getData, tc.expectedData, "Expedted Data is Wrong")
		})
	}
}
