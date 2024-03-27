package x14nfile

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	addr       string
	FileSystem *FileSystem
	Logger     *log.Logger
}

//new server

func (s *Server) Start() {
	server := &Server{
		addr: "12345",
	}
	router := mux.NewRouter()
	router.Use(LoggingMiddle)

	router.HandleFunc("/search", server.SearchFile)
	router.HandleFunc("/getFile", s.GetFile)
	router.HandleFunc("/DeleteFile", s.DeleteFile)
	log.Printf("Server listening on %s\n", server.addr)
	log.Fatal(http.ListenAndServe(server.addr, router))
}

func LoggingMiddle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the client's IP address
		clientIP := r.RemoteAddr

		// Log the request
		log.Printf("Request from IP: %s\n", clientIP)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// display allfile path
func (s *Server) GetAllFile(w http.ResponseWriter, r *http.Request) {

}

// delete file
func (s *Server) DeleteFile(w http.ResponseWriter, r *http.Request) {

}

// search file
func (s *Server) SearchFile(w http.ResponseWriter, r *http.Request) {

}

// search and get file
func (S *Server) GetFile(w http.ResponseWriter, r *http.Request) {

}
