package maps

import "github.com/cheekybits/genny/generic"

type SomeK generic.Type
type SomeV generic.Type

type _Prefix_SomeKNSomeVIf interface {
	Load(SomeK) (SomeV, bool)
}

type _Prefix_SomeKNSomeVMutIf interface {
	Store(SomeK, SomeV)
	Delete(SomeK)
}

type _Prefix_SomeKNSomeVIterIf interface {
	Range(func(SomeK, SomeV) bool)
}

type _Prefix_SomeKNSomeV map[SomeK]SomeV

func (__ _Prefix_SomeKNSomeV) Load(k SomeK) (SomeV, bool) {
	v, ok := __[k]
	return v, ok
}

func (__ _Prefix_SomeKNSomeV) Store(k SomeK, v SomeV) {
	__[k] = v
}
func (__ _Prefix_SomeKNSomeV) Delete(k SomeK, v SomeV) {
	delete(__, k)
}
func (__ _Prefix_SomeKNSomeV) Range(f func(k SomeK, v SomeV) bool) {
	for k, v := range __ {
		if !f(k, v) {
			break
		}
	}
}

type __Prefix_SomeKNSomeV = _Prefix_SomeKNSomeV
type _Prefix_SomeKNSomeVSt struct {
	// do not want to export but want to use embedding method
	__Prefix_SomeKNSomeV
}

func New_Prefix_SomeKNSomeVSt(somes _Prefix_SomeKNSomeV) *_Prefix_SomeKNSomeVSt {
	return &_Prefix_SomeKNSomeVSt{
		__Prefix_SomeKNSomeV: somes,
	}
}
