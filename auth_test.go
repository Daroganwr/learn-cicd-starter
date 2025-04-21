package main

import ("testing"
"github.com/bootdotdev/learn-cicd-starter/internal/auth"
//"errors"
"net/http"
//"strings"
)

func TestApiKey( t *testing.T) {
	custom_header := http.Header{}
	custom_header.Add("Authorization", "ApiKey somerandomjunk")
	want := "somerandomjunk"
	key, err := auth.GetAPIKey(custom_header)
	if err != nil {
		t.Errorf("%v", err)
	}
	if key != want {
		t.Errorf("GetApiKey(custom_header) = %v, want match for %#v, nil", key, want)
	}
}

func TestFaultyApiKey( t *testing.T) {
	custom_header := http.Header{}
	custom_header.Add("Authorization", "Bearer somerandomjunk")
	_, err := auth.GetAPIKey(custom_header)
	if err != nil {
		t.Errorf("%v", err)
	}
}