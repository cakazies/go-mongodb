package test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/local/go-mongo/models"
	"github.com/local/go-mongo/utils"
	"github.com/local/testify/assert"
	"github.com/spf13/viper"
)

var (
	cfgFile string
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

func init() {
	initViper()
}
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
		fmt.Println(tc)
		t.Run(tc.name, func(t *testing.T) {
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
			getData := fmt.Sprintf("%v", respData.Data["rm_id"])
			assert.Equal(t, getData, tc.expectedData, "Expedted Data is Wrong")

		})
	}
}

func initViper() {
	viper.SetConfigFile("toml")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./../configs")
		viper.SetConfigName("config")
	}
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	utils.FailError(err, "Error Viper config")
	log.Println("Using Config File: ", viper.ConfigFileUsed())
}
