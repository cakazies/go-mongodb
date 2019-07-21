package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"

	"github.com/local/go-mongo/models"
)

type testCase struct {
	name         string
	input        string
	expectedData string
	expectedCode int
	path         string
	handler      func(echo.Context) error
	query        string
}

func TestGetRoom(t *testing.T) {
	tasks := []testCase{
		{
			name:         "Testing with student id",
			input:        "5d33185d5dbd1d1f0c4f5d3a",
			expectedData: "5d33185d5dbd1d1f0c4f5d3a",
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
			log.Println(rec.Body.String())
			// Assertions
			// if assert.NoError(t, models.GetStudent) {
			// 	assert.Equal(t, http.StatusOK, rec.Code)
			// 	assert.Equal(t, userJSON, rec.Body.String())
			// }

			// buf := resp.Body.Bytes()
			// var respData Response
			// if err := json.Unmarshal(buf, &respData); err != nil {
			// 	t.Error("Can not parsing response testing. Error :", err)
			// }
			// getData := fmt.Sprintf("%v", respData.Data["rm_id"])
			// assert.Equal(t, getData, tc.expectedData, "Expedted Data is Wrong")
		})
	}
}
