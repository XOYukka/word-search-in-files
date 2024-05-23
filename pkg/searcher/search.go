package searcher

import (
	"bufio"
	"fmt"
	"io/fs"
	"regexp"
	"strings"
	"sync"
	"word-search-in-files/pkg/internal/dir"
)

type Searcher struct {
	FS fs.FS
}

var wg sync.WaitGroup

func (s *Searcher) Search(word string) ([]string, error) {
	var foundFiles []string
	fileNames, err := dir.FilesFS(s.FS, ".") // Files list from examples
	if err != nil {
		return nil, err
	}

	for _, fileName := range fileNames {
		wg.Add(1)
		result := make(chan bool)
		go searchInFile(s, fileName, word, result)
		if <-result {
			foundFiles = append(foundFiles, strings.Split(fileName, ".")[0])
		}
		fmt.Println(fileName)
	}
	wg.Wait()
	return foundFiles, nil
}

func searchInFile(s *Searcher, fileName, word string, out chan<- bool) {
	f, err := s.FS.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	defer wg.Done()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if FindWord(scanner.Text(), word) {
			out <- true
			return
		}
	}
	out <- false
}

func FindWord(str, word string) bool {

	validWordPattern := regexp.MustCompile(`^[a-zA-Zа-яА-Я]+(-[a-zA-Zа-яА-Я]+)?$`)

	if !validWordPattern.MatchString(word) {
		return false
	}

	words := strings.FieldsFunc(str, func(r rune) bool {
		return r == ' ' || r == ',' || r == '.' || r == '!' || r == '?'
	})

	for _, w := range words {
		if strings.EqualFold(w, word) {
			return true
		}
	}

	return false
}
