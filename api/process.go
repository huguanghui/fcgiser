package api

import "net/http"

type EleMethod struct {
	Opt func(http.ResponseWriter, *http.Request) int
}

type ProcessURL struct {
	Process map[string]*EleMethod
}

func (p *ProcessURL) AddMethod(method string, opt *EleMethod) error {
	p.Process[method] = opt
	return nil
}

func (p *ProcessURL) Lookup(method string) *EleMethod {
	return p.Process[method]
}
