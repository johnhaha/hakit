package hadata

type Binder string

func (b Binder) With(t ...string) string {
	v := NewStringBinder().BindString(string(b)).BindString(t...).Value()
	return v
}
