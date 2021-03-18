package coingecko

import (
	"fmt"
	"net/http"
	"testing"
)

func TestUtilService_Ping(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/ping")

		fmt.Fprint(w, `{"gecko_says": "(V3) To the Moon!"}`)
	})
	if ping, _, err := testClient.Util.Ping(); err != nil {
		t.Errorf("Error given: %s", err)
	} else if ping == nil {
		t.Error("Expected ping. Util.Ping is nil")
	}
}
