package muxlist

//Constants that map to a key in a Result.
//This is provided to be used as "helpers".
const (
	HTTP_METHODS = iota
	REQUEST_URI
	ROUTE_NAME
	HANDLER_NAME
)

//Represents information for a single route
type Result map[int]string

//ResultSet contains all available information for a multiplexer
type ResultSet []Result

type Extractor interface {
	//This method is exported as the ResultSet might want to be inspected manually
	Extract() ResultSet
	List() string
}
