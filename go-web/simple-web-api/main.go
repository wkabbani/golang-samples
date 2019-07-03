package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/gorilla/mux"
)

// Note represents a simple note
type Note struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

var notes []Note

func (note *Note) isValid() error {
	switch {
	case len(note.Text) == 0:
		return errors.New("text can't be empty")
	default:
		return nil
	}
}

func main() {
	// read port
	port := flag.Uint("p", 8080, "The port on which the server will listen")
	flag.Parse()

	// initialize notes
	initNotes()

	// setup http router
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/api/v1/notes", getNotes).Methods("GET")
	r.HandleFunc("/api/v1/notes/{id:[0-9]+}", getNote).Methods("GET")
	r.HandleFunc("/api/v1/notes", createNote).Methods("POST")
	r.HandleFunc("/api/v1/notes/{id:[0-9]+}", deleteNote).Methods("DELETE")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), r))
}

func initNotes() {
	notes = append(notes, Note{ID: 1, Text: "create get notes list endpoint", Done: false})
	notes = append(notes, Note{ID: 2, Text: "create get single note endpoint", Done: true})
	notes = append(notes, Note{ID: 3, Text: "create post note endpoint", Done: false})
	notes = append(notes, Note{ID: 4, Text: "create delete note endpoint", Done: true})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("simple notes api"))
}

func getNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func getNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	for _, note := range notes {
		if note.ID == id {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(note)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func createNote(w http.ResponseWriter, r *http.Request) {
	// JSON decode body
	var note Note
	_ = json.NewDecoder(r.Body).Decode(&note)

	// validate
	if err := note.isValid(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// find the id
	note.ID = 0
	for _, n := range notes {
		if n.ID > note.ID {
			note.ID = n.ID
		}
	}
	note.ID++

	// add note
	notes = append(notes, note)

	// return note
	w.Header().Add("Location", path.Join(r.URL.Path, strconv.Itoa(note.ID)))
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, note := range notes {
		if note.ID == id {
			notes = append(notes[0:index], notes[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
