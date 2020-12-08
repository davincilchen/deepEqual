package main

import (
	

	"testing"
	"strings"
	"fmt"
	"reflect"
	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	fmt.Println("string array test....")
    got := strings.Split("a:b:c", ":")
    want := []string{"a", "b", "c"};
    if !reflect.DeepEqual(got, want) {
		 fmt.Println("!reflect.DeepEqual(got, want)", got,want)
	}else{
		fmt.Println("reflect.DeepEqual(got, want)", got,want)
	}
	assert.Equal(t, got, want)		
}