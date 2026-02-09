package server

import (
	"fmt"
	"net/http"
)

type Server struct {
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("hello %s", r.Header.Get("userid"))))

}
