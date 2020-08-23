package api

type Runner struct {
	commands map[string]*ProcessURL
}

func NewRunner() *Runner {
	return &Runner{
		commands: make(map[string]*ProcessURL),
	}
}

func (r *Runner) Use(k string, v *ProcessURL) {
	r.commands[k] = v
}

func (r *Runner) Lookup(name string) *ProcessURL {
	return r.commands[name]
}
