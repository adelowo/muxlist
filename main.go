package muxlist

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

type MuxLister struct {
	extractor Extractor
}

//Allocates and returns a new Multiplexer lister
func NewMuxLister(e Extractor) *MuxLister {

	return &MuxLister{extractor: e}
}

//Prints out a table containing data from a multiplexer
//This only works when in Development mode.
//In Production mode, this is a no-op
func (lister *MuxLister) Table() {

	fmt.Println(lister.extractor.List())
}

//Returns a human readable name for a http Handler
func GetHumanReadableNameForHandler(h http.Handler) string {
	return runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
}
