package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	Id       int
	UserName string
}

type Server struct {
	db      map[int]*User
	dbCache map[int]*User
	dbHit   int32
}

func NewServer() *Server {
	db := make(map[int]*User)
	dbCache := make(map[int]*User)
	for i := 0; i < 100; i++ {
		db[i+1] = &User{
			Id:       i + 1,
			UserName: fmt.Sprintf("User%d", i+1),
		}
	}
	return &Server{
		db:      db,
		dbCache: dbCache,
	}
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idstr)

	// Check cache
	user, found := s.dbCache[id]
	if found {
		json.NewEncoder(w).Encode(user)
		return
	}

	user, ok := s.db[id]
	s.dbHit++
	if !ok {
		panic("User not found")
	}
	json.NewEncoder(w).Encode(user)
	s.dbCache[id] = user
}

func main() {

}
