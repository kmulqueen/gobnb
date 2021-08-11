package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)


type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"room1", "/room1", "GET", []postData{}, http.StatusOK},
	{"room2", "/room2", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"reservation-summary", "/reservation-summary", "GET", []postData{}, http.StatusOK},
	{"post-search-availability", "/search-availability", "POST", []postData{
		{key: "start", value: "08-10-2021"},
		{key: "end", value: "08-11-2021"},
	}, http.StatusOK},
	{"post-search-availability-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "08-10-2021"},
		{key: "end", value: "08-11-2021"},
	}, http.StatusOK},
	{"make-reservation-post", "/make-reservation", "POST", []postData{
		{key:"first-name", value: "Sam"},
		{key:"last-name", value: "Smith"},
		{key:"email", value: "sam@test.com"},
		{key:"phone-number", value: "123-456-7890"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()

	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close()

	for _, test := range theTests {
		if test.method == "GET" {
			res, err := testServer.Client().Get(testServer.URL + test.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if res.StatusCode != test.expectedStatusCode {
				t.Errorf("%s expected %d but got %d", test.name, test.expectedStatusCode, res.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range test.params{
				values.Add(x.key, x.value)
			}

			res, err := testServer.Client().PostForm(testServer.URL + test.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if res.StatusCode != test.expectedStatusCode {
				t.Errorf("%s expected %d but got %d", test.name, test.expectedStatusCode, res.StatusCode)
			}
		}
	}
}