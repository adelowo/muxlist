package muxlist

import (
	"flag"
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

var extractor *GorillaMuxLister

func TestMain(m *testing.M) {

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	}).Name("index")

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

	}).Name("users")

	router.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {

	}) //no name

	extractor = NewGorillaMuxLister(router)

	flag.Parse()

	os.Exit(m.Run())
}

func TestGorillaMuxLister_Extract(t *testing.T) {

	result := extractor.Extract()

	expected := ResultSet{
		Result{REQUEST_URI: "/", ROUTE_NAME: "index", HTTP_METHODS: "",
			HANDLER_NAME: "github.com/adelowo/muxlist/gorillamux.TestMain.func1"},

		Result{REQUEST_URI: "/users", ROUTE_NAME: "users", HTTP_METHODS: "",
			HANDLER_NAME: "github.com/adelowo/muxlist/gorillamux.TestMain.func2"},

		Result{REQUEST_URI: "/about", ROUTE_NAME: "", HTTP_METHODS: "",
			HANDLER_NAME: "github.com/adelowo/muxlist/gorillamux.TestMain.func3"},
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v. \n Got %v", expected, result)
	}
}
