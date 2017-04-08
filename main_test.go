package muxlist

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetHumanReadableNameForHandler(t *testing.T) {

	h := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Found"))
	}

	name := GetHumanReadableNameForHandler(http.HandlerFunc(h))
	same := strings.Compare(
		"github.com/adelowo/muxlist.TestGetHumanReadableNameForHandler.func1",
		name)

	if same != 0 {
		t.Fatalf("Not the same. Expected %s. Got %s",
			"github.com/adelowo/muxlist.TestGetHumanReadableNameForHandler.func1",
			name)
	}
}

func TestGetHumanReadableNameForHandler2(t *testing.T) {

	h := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Found"))
	}

	name := GetHumanReadableNameForHandler(http.HandlerFunc(h))

	same := strings.Compare(
		"github.com/adelowo/muxlist.TestGetHumanReadableNameForHandler.func1",
		name)

	if same == 0 {
		t.Fatalf("Not the same. Expected %s. Got %s",
			"github.com/adelowo/muxlist.TestGetHumanReadableNameForHandler.func1",
			name)
	}

}

type MockExtractor struct {
}

func (e MockExtractor) Extract() ResultSet {
	return make(ResultSet, 10)
}

func (e MockExtractor) List() string {
	return "Formatted table"
}

func TestCannotMakeUseOfAnInvalidEnvironmentModeToSetUpAMuxLister(t *testing.T) {

	defer func() {
		recover()
	}()

	NewMuxLister(&MockExtractor{})
}

func TestMuxLister_Table(t *testing.T) {

	extractor := MockExtractor{}

	same := strings.Compare(extractor.List(), "Formatted table")

	if same != 0 {
		t.Fatalf(
			"Invalid list.. Expected %s. Got %s",
			"Formatted table",
			extractor.List())
	}
}
