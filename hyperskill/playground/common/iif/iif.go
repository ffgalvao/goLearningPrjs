package iif

// https://stackoverflow.com/questions/19979178/what-is-the-idiomatic-go-equivalent-of-cs-ternary-operator
// https://github.com/icza/gox/blob/master/gox/if.go

type Iif bool

func (i Iif) Then(a, b interface{}) interface{} {
	if i {
		return a
	} else {
		return b
	}
}

func (i Iif) then(a, b interface{}) interface{} {
	if i {
		return a
	} else {
		return b
	}
}

func (i Iif) Int(iFTrue, ifFalse int) int {
	return i.then(iFTrue, ifFalse).(int)
}

func (i Iif) Str(a, b string) string {
	return i.String(a, b)
}

func (i Iif) String(a, b string) string {
	return i.then(a, b).(string)
}

type iif2s struct {
	Cond   bool
	ITrue  interface{}
	IFalse interface{}
}

type Iif2 iif2s

func (i Iif2) Then() interface{} {
	if i.Cond {
		return i.ITrue
	} else {
		return i.IFalse
	}
}

func IIF[T any](cond bool, then T, elsi T) T {
	if cond {
		return then
	} else {
		return elsi
	}
}
