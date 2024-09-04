package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}
	bool1 := &Boolean{Value: true}
	bool2 := &Boolean{Value: true}
	int1 := &Integer{Value: 10}
	int2 := &Integer{Value: 10}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with the same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with the same content have different hash keys")
	}
	
	if bool1.HashKey() != bool2.HashKey() {
		t.Errorf("bools with the same content have different hash keys")
	}

	if int1.HashKey() != int2.HashKey() {
		t.Errorf("bools with the same content have different hash keys")
	}
}