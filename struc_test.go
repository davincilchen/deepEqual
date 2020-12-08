package main

import (
	

	"testing"
	"fmt"
	"reflect"
	"github.com/stretchr/testify/assert"
)

type Sample struct {
	ID int
	Name string
}
func TestStruct(t *testing.T) {
	fmt.Println("struct test....")
    got := Sample{
		ID: 1,
		Name: "haha",
	}
	want := Sample{
		ID: 1,
		Name: "haha",
	}
    if !reflect.DeepEqual(got, want) {
		 fmt.Println("!reflect.DeepEqual(got, want)", got,want)
	}else{
		fmt.Println("reflect.DeepEqual(got, want)", got,want)
	}
	if got == want { 
		fmt.Println("got == want", got,want)
   }else{
	   fmt.Println("got != want", got,want)
   }
   fmt.Printf("address of got %p, address of want %p \n",&got,&want)
   fmt.Println("got to string", fmt.Sprintf("%#v", got))
   assert.Equal(t, got, want)		
}