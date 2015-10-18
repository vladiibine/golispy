package scopes

type Scope struct {
	Vals map[string]interface{}
	Parent *Scope
	Omni *Scope
}
func (scope *Scope) Contains(key string) bool{
	_, ok := scope.Vals[key]
	return ok
}
func (scope *Scope) Get(key string) interface{}{
	return scope.Vals[key]
}
func (scope *Scope) Set(key string, value interface{}){
	scope.Vals[key] = value
}

type IScope interface {
	Contains(key string) bool
	Get(key string) interface{}
	Set(key string, value interface{})
}