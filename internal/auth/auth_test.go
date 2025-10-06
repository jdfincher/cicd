package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

var ErrShouldFail = errors.New("some error message")

func TestGetApiKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:8080", nil)
	req.Header.Set("Authorization", "ApiKey somehashedtokenstring")

	got, err := GetAPIKey(req.Header)
	want := "somehashedtokenstring"
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
	req2, _ := http.NewRequest("GET", "", nil)
	req2.Header.Set("Auth", "Authorization ApiKey")

	got, err = GetAPIKey(req2.Header)
	want = ""
	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Fatalf("expected: ' %v ', got: ' %v '", want, got)
	}
}
