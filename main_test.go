package main 

import "testing"

func TestSayHello(t *testing.T){
	expectedStr := "Hello Github Actions"

	if SayHello()!=expectedStr{
		t.Errorf("Expected string %s but got %s", expectedStr, SayHello())
	}
}