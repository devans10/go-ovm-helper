package ovmHelper

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// RoundTripFunc is for returning a test response to the client
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip is the test http Transport
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func testAccPreChecks(t *testing.T) {

	if os.Getenv("OVM_ACC") == "" {
		t.Skip("set OVM_ACC to run purestorage acceptance tests (provider connection is required)")
	}

	entrypoint := os.Getenv("OVM_ENDPOINT")
	username := os.Getenv("OVM_USERNAME")
	password := os.Getenv("OVM_PASSWORD")
	if entrypoint == "" {
		t.Fatalf("OVM_ENDPOINT must be set for acceptance tests")
	}
	if (username != "") && (password == "") {
		t.Fatalf("OVM_PASSWORD must be set if OVM_USERNAME is set for acceptance tests")
	}
}

func testAccGenerateClient(t *testing.T) *Client {

	username := os.Getenv("OVM_USERNAME")
	password := os.Getenv("OVM_PASSWORD")
	entrypoint := os.Getenv("OVM_ENDPOINT")

	c := NewClient(username, password, entrypoint)

	return c
}

func TestAccClient(t *testing.T) {
	testAccPreChecks(t)
	testAccGenerateClient(t)
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
