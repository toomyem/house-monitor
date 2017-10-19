package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_1(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/temperature", strings.NewReader("{}"))
	temperatureHandler(recorder, request)
	if recorder.Code != 201 {
		t.Fatalf("Expected 201, but got: %d", recorder.Code)
	}
}
