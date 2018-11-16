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
var subExtractor *GorillaMuxLister

func TestMain(m *testing.M) {

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	}).Name("index").Methods("GET")

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

	}).Name("users").Methods("POST")

	router.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
	}).Methods("GET") //no name

	router.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
	}) //no name or method

	extractor = NewGorillaMuxLister(router)

	// test subrouter
	router2 := mux.NewRouter()

	subrouter := router2.PathPrefix("/test").Subrouter()

	subrouter.HandleFunc("/{key}", func(w http.ResponseWriter, r *http.Request) {})

	subExtractor = NewGorillaMuxLister(router2)

	flag.Parse()

	os.Exit(m.Run())
}

func TestGorillaMuxLister_Extract(t *testing.T) {

	result := extractor.Extract()

	expected := ResultSet{
		Result{REQUEST_URI: "/", ROUTE_NAME: "index", HTTP_METHODS: "GET",
			HANDLER_NAME: "github.com/adelowo/muxlist.TestMain.func1"},

		Result{REQUEST_URI: "/users", ROUTE_NAME: "users", HTTP_METHODS: "POST",
			HANDLER_NAME: "github.com/adelowo/muxlist.TestMain.func2"},

		Result{REQUEST_URI: "/about", ROUTE_NAME: "", HTTP_METHODS: "GET",
			HANDLER_NAME: "github.com/adelowo/muxlist.TestMain.func3"},

		Result{REQUEST_URI: "/contact", ROUTE_NAME: "", HTTP_METHODS: "",
			HANDLER_NAME: "github.com/adelowo/muxlist.TestMain.func4"},
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v. \n Got %v", expected, result)
	}
}

func TestGorillaMuxLister_Subrouter(t *testing.T) {

	result := subExtractor.Extract()

	expected := ResultSet{
		Result{REQUEST_URI: "/test", ROUTE_NAME: "", HTTP_METHODS: "",
			HANDLER_NAME: "SUBROUTER"},

		Result{REQUEST_URI: "/test/{key}", ROUTE_NAME: "", HTTP_METHODS: "",
			HANDLER_NAME: "github.com/adelowo/muxlist.TestMain.func5"},
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("\nExpected %v.\n\tGot %v", expected, result)
	}
}
