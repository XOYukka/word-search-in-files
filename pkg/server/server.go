package server

import (
	"encoding/json"
	"net/http"
	"os"
	"word-search-in-files/pkg/searcher"
)

func ViewHandler(writer http.ResponseWriter, request *http.Request) {
	word := request.URL.Query().Get("word")
	if word == "" {
		http.Error(writer, "Missing 'word' parametr", http.StatusBadRequest)
		return
	}

	s := &searcher.Searcher{FS: os.DirFS(".")}

	files, err := s.Search(word)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(files); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func StartServer(port string) error {
	http.HandleFunc("/files/search", ViewHandler)
	err := http.ListenAndServe("localhost:"+port, nil)
	if err != nil {
		return err
	}
	return nil
}
