package controller_test

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"net/http/httptest"
	"number-classifier/controller"
	"number-classifier/models"
	"testing"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/api/classify-number", controller.ClassifyNumber)
	return r
}

func TestClassifyNumberValid(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/api/classify-number?number=153", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status OK, got %v", w.Code)
	}

	var resp models.NumberResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse JSON response: %v", err)
	}

	if resp.Number != 153 {
		t.Errorf("expected number 153, got %v", resp.Number)
	}

	if resp.DigitSum != 9 {
		t.Errorf("expected digit sum 9, got %v", resp.DigitSum)
	}

	expectedProperties := []string{"armstrong", "odd"}
	if len(resp.Properties) != len(expectedProperties) {
		t.Errorf("expected properties %v, got %v", expectedProperties, resp.Properties)
	} else {
		for i := range expectedProperties {
			if resp.Properties[i] != expectedProperties[i] {
				t.Errorf("expected properties %v, got %v", expectedProperties, resp.Properties)
			}
		}
	}

	if resp.IsPrime {
		t.Errorf("expected not prime, got prime")
	}

	if resp.IsPerfect {
		t.Errorf("expected not perfect, got perfect")
	}

	if resp.FunFact == "" {
		t.Errorf("expected fun fact, got empty")
	}
}

func TestClassifyNumberInvalid(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/api/classify-number?number=alphabet", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status Bad Request, got %v", w.Code)
	}

	var resp models.ErrorResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse JSON error response: %v", err)
	}

	if resp.Number != "alphabet" {
		t.Errorf("expected number abc, got %v", resp.Number)
	}

	if !resp.Error {
		t.Errorf("expected error true, got false")
	}
}
