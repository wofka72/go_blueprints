package main

import (
	"log"
	"net/http"
	"sync"
	"text/template"
	"path/filepath"
	"flag"
	//"os"
	//"go_blueprints/chapter1/trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
)


// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(
			filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	t.templ.Execute(w, data)
}


func main() {
	var addr = flag.String("host", ":8080", "The addr of the application.")
	flag.Parse() // parse the flags

	// set up gomniauth
	// you can use the hash or phrase that I want
	gomniauth.SetSecurityKey("some long key")
	gomniauth.WithProviders(
		facebook.New("149793185735298",
			"ilEAiYePbCPkls9tQcDW3llhyGY",
			"http://localhost:8080/auth/callback/facebook"),
		github.New("b4ec743073297d6f03e3",
			"fce9352f0d0c0e9d96cb050db4b5774de53e3f6a",
			"http://localhost:8080/auth/callback/github"),
		google.New("24677165843-i6gr6chmuq7ea6cnvb8ouch84q70e19p.apps.googleusercontent.com",
			"qWIdbjuC9kr4ZqzeJ3EuLgMW",
			"http://localhost:8080/auth/callback/google"),
	)

	r := newRoom()
	// r.tracer = trace.New(os.Stdout)

	http.Handle("/chat",
		MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)


	// get the room going
	go r.run()

	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
