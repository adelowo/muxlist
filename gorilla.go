//Provides a Multiplexer lister for Gorilla/mux
package muxlist

import (
	"fmt"

	"github.com/gorilla/mux"
)

type GorillaMuxLister struct {
	router *mux.Router
}

func (m *GorillaMuxLister) Extract() ResultSet {

	var result ResultSet

	m.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {

		r := make(Result, 10)
		uri, err := route.GetPathTemplate()
		if err != nil {
			return err
		}

		host, err := route.GetHostTemplate()

		if err == nil {
			r[REQUEST_URI] = host + uri
		} else {
			r[REQUEST_URI] = uri
		}

		r[ROUTE_NAME] = route.GetName()
		r[HTTP_METHODS] = "" //Currently unavailable for Gorilla Mux but there's a pending PR for this

		r[HANDLER_NAME] = GetHumanReadableNameForHandler(route.GetHandler())

		result = append(result, r)
		return nil
	})

	return result
}

func (m *GorillaMuxLister) List() string {
	fmt.Println("+--------+-----------+--------------` + " +
		"--------------------+-----------------------------------+" +
		"----------------------------------------------------------------------" +
		"-+---------------------+" + "\n" +
		`METHOD NAME | REQUEST URI | ROUTE NAME | HTTP HANDLER ` + "\n" +
		"+--------+-----------+--------------` + " +
		"--------------------+-----------------------------------+" +
		"----------------------------------------------------------------------" +
		"-+---------------------+" + "\n")

	var val string
	for _, v := range m.Extract() {
		val += fmt.Sprintf(`[%s]`+" | "+`[%s]`+" | "+`[%s]`+" | "+`[%s]`+"\n \n",
			v[HTTP_METHODS],
			v[REQUEST_URI],
			v[ROUTE_NAME],
			v[HANDLER_NAME])
	}

	return fmt.Sprint(val + "+--------+-----------+-------------------------" +
		"---------+-----------------------------------+-----------------------" +
		"------------------------------------------------+---------------------+")
}

//Allocates and returns a new GorillaMuxLister
func NewGorillaMuxLister(r *mux.Router) *GorillaMuxLister {
	return &GorillaMuxLister{router: r}
}
