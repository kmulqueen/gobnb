package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/kmulqueen/gobnb/internal/config"
	"github.com/kmulqueen/gobnb/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	// What to store in session
	gob.Register(models.Reservation{})

	// Change to true when in production
	testApp.InProduction = false

	// Initialize a new session & set the lifetime to 24 hours
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true // Should session persist after browser window closed?
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	// Set session in config to session we just initialized
	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct {}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int) {

}

func (tw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}