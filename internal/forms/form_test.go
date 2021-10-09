package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/test", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/test", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/test", nil)

	r.PostForm = postedData
	form = New(r.PostForm)

	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestField_IsEmail(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("notemail", "ilham")

	r := httptest.NewRequest("POST", "/test", nil)
	r.PostForm = postedData
	form := New(r.PostForm)

	form.IsEmail("notemail")
	if form.Valid() {
		t.Error("form shows valid when field is not email")
	}

	if form.Errors.Get("notemail") == "" {
		t.Error("should have an error but did not get one")
	}

	postedData.Add("email", "ilham@mail.com")

	r, _ = http.NewRequest("POST", "/test", nil)

	r.PostForm = postedData
	form = New(r.PostForm)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("shows is not email fields when it is")
	}

	if form.Errors.Get("email") != "" {
		t.Error("should not have an error but got one")
	}
}

func TestField_Has(t *testing.T) {
	r, _ := http.NewRequest("POST", "/test", nil)
	form := New(r.Form)

	form.Has("notexist")
	if form.Valid() {
		t.Error("form shows has fields when fields missing")
	}

	if form.Errors.Get("notexist") == "" {
		t.Error("should have an error but did not get one")
	}

	postedData := url.Values{}
	postedData.Add("exist", "ilham")

	r, _ = http.NewRequest("POST", "/test", nil)
	r.PostForm = postedData
	form = New(r.PostForm)

	form.Has("exist")
	if !form.Valid() {
		t.Error("shows does not has fields when it does")
	}

	if form.Errors.Get("exist") != "" {
		t.Error("should not have an error but got one")
	}
}

func TestField_MinLength(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("notlength", "ilham")

	r, _ := http.NewRequest("POST", "/test", nil)
	r.PostForm = postedData
	form := New(r.PostForm)

	form.MinLength("notlength", 8)
	if form.Valid() {
		t.Error("form shows valid when min length fields missing")
	}

	if form.Errors.Get("notlength") == "" {
		t.Error("should have an error but did not get one")
	}

	r, _ = http.NewRequest("POST", "/test", nil)
	r.PostForm = postedData
	form = New(r.PostForm)

	postedData.Add("minlength", "ilhamisthegod")

	form.MinLength("minlength", 8)
	if !form.Valid() {
		t.Error("shows is not min length fields when it is")
	}

	if form.Errors.Get("minlength") != "" {
		t.Error("should not have an error but got one")
	}
}
