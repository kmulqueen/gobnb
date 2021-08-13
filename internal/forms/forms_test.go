package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid instead of valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows doesn't have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it doesn't")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("shows form doesn't have field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("whatever", 10)
	if form.Valid() {
		t.Error("form shows min length for a field that doesn't exist")
	}

	isError := form.Errors.Get("whatever")
	if isError == "" {
		t.Error("should have an error but couldn't find it")
	}

	
	postedValues := url.Values{}
	postedValues.Add("some_field", "some value")
	form = New(postedValues)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows min length of 100 is met when it is actuallly shorter")
	}

	postedValues = url.Values{}
	postedValues.Add("another_field", "1234")
	form = New(postedValues)
	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("shows min length isn't met when it is")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("shouldn't have an error but got one")
	}


}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "test@test.com")
	form = New(postedValues)
	
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("shows form email isn't valid when it should be")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "123")
	form = New(postedValues)
	
	form.IsEmail("email")
	if form.Valid() {
		t.Error("shows form email is valid when it shouldn't be")
	}
}