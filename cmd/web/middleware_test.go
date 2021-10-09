package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	Test(t)
}

func TestSessionLoad(t *testing.T) {
	Test(t)
}

func Test(t *testing.T) {
	var myH myHandler

	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		//
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}
