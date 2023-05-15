package app

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	router  *mux.Router
	debug   bool
	testing bool
}

func NewServer() *Server {
	s := Server{router: mux.NewRouter()}
	return &s
}

func (s *Server) cors() {
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
		handlers.AllowedMethods([]string{"POST", "DELETE", "GET", "OPTIONS", "PUT", "PATCH"}),
	)
	s.router.Use(cors)
	s.router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func (s *Server) respond(w http.ResponseWriter, v interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Println(err)
	}
}

func (s *Server) LogDebug(v ...interface{}) {
	if s.debug {
		log.Println(v...)
	}
}

func (s *Server) validate(toValidate interface{}) error {
	v := validator.New()
	return v.Struct(toValidate)
}

func (s *Server) getParams(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func (s *Server) getQueryParams(r *http.Request) map[string][]string {
	return r.URL.Query()
}

func (s *Server) error(w http.ResponseWriter, err error, statusCode int) {
	log.Println(err)
	s.respond(w, struct {
		Error string `json:"error"`
	}{err.Error()}, statusCode)
}

func (s *Server) decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (s *Server) registerRoute(path string, handler http.HandlerFunc, methods []string, middleware ...func(handlerFunc http.HandlerFunc) http.HandlerFunc) {
	for _, mw := range middleware {
		handler = mw(handler)
	}

	handler = s.middlewareJson(handler)

	if s.debug {
		handler = s.middlewareAccessLog(handler)
	}

	s.router.HandleFunc(path, handler).Methods(methods...)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) LocalDev(port string) {
	s.debug = true
	s.Serve(port)
}

func (s *Server) Serve(port string) {
	s.routes()
	s.cors()
	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, s.router))
}
