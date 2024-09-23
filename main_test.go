package main

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestGetIssues(t *testing.T) {
	router := InitRouter()
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest("GET", "/issues", nil))
	t.Run("Returns 200 status code", func(t *testing.T) {
		if recorder.Code != 200 {
			t.Error("Expected 200, got ", recorder.Code)
		}
	})

	fmt.Println("recorder", recorder)
	t.Error("Expected 200, got ", recorder.Code)
}
