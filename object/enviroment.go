package object

type Enviroment struct {
	store map[string]Object
	outer *Enviroment
}

func NewEnviroment() *Enviroment {
	s := make(map[string]Object)

	return &Enviroment{store: s, outer: nil}
}

func NewEnclosedEnviroment(outer *Enviroment) *Enviroment {
	env := NewEnviroment()
	env.outer = outer

	return env
}

func (e *Enviroment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]

	if !ok && e.outer != nil {
		obj, ok = e.outer.store[name]
	}

	return obj, ok
}

func (e *Enviroment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}