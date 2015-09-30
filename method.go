package momentum

import (
"reflect"
"sync"
)

type Method struct {
	Name  interface{}
	Value interface{}
	Statistics* Stat
	Sync *sync.Mutex

}

func (method *Method) Call(arguments interface{}) interface{} {
	method.Sync.Lock()
	defer method.Sync.Unlock()
	method.Statistics.Numberofcalls++
	result := reflect.ValueOf(method.Value).Call([]reflect.Value{
		reflect.ValueOf(arguments)})[0].Interface()
	return result

}