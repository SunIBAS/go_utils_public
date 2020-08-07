package MainTools

import "net/http"

type ActionAtom struct {
	DearFn func(writer http.ResponseWriter,s string,config Config)	`json:"-"`
	Method string		`json:"method"`
	Description string	`json:"description"`
	Params []string		`json:"params"`
}
