// Objective: Start the Jet Ski

package main

import "reflect"

func main() {
	println("Logging in...")
	authorized := startup(login())
	if reflect.ValueOf(authorized).Bool() {
		println("Starting the engine")
		return
	}
	println("Startup failed")
}

func validSequence(i int, el interface{}) bool {
	return reflect.TypeOf(el).String() == "*main.Sequence" &&
		!reflect.ValueOf(el).IsNil() &&
		reflect.ValueOf(el).Elem().NumField() == 2 &&
		reflect.TypeOf(reflect.ValueOf(el).Elem().Field(0).Interface()).String() == "int" &&
		int(reflect.ValueOf(el).Elem().Field(0).Int()) == i*i-i &&
		!reflect.ValueOf(reflect.ValueOf(el).Elem().Field(1).Interface()).IsNil()
}

func startup(seq interface{}) bool {
	for i := 0; i < 5; i++ {
		if !validSequence(i, seq) {
			return false
		}
		seq = reflect.ValueOf(seq).Elem().Field(1).Interface()
	}

	return true
}


func login() interface{} {
	seq := &Sequence{
	    Field1: 0,
	    Field2: &Sequence{
	        Field1: 0,
	        Field2: &Sequence{
	            Field1: 2,
	            Field2: &Sequence{
	                Field1: 6,
	                Field2: &Sequence{
	                    Field1: 12,
	                    Field2: &Sequence{
	                        Field1: -1,
	                        Field2: nil,
	                    },
	                },
	            },
	        },
	    },
	}
	return seq
}

type Sequence struct {
    Field1 int
    Field2 *Sequence
}
