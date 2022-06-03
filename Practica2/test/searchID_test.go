package test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSearchID(t *testing.T) {
	_, err := http.NewRequest("GET", "/searchID", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	fmt.Println("All is working fine")
}
