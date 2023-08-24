package main

import (
	"flag"
	"fmt"
	"go-search-directory/searcher"
	"strconv"
	"strings"
	"time"
)

func main() {
	rootPathPtr := flag.String("path", "", "path to the directory to be searched")
	wordsPtr := flag.String("words", "", "words to be searched in path names (separated by commas)")
	threadLimitPtr := flag.Int("max-thread-count", 16, "number of maximum threads running concurrently")
	caseSensitivePtr := flag.Bool("case-sensitive", false, "case sensitive (look for perfect match in terms of letter case)")

	flag.Parse()

	rootPath := *rootPathPtr
	words := *wordsPtr
	threadLimit := *threadLimitPtr
	caseSensitive := *caseSensitivePtr

	fmt.Println("max-thread-count is " + strconv.Itoa(threadLimit))
	time.Sleep(1 * time.Second)

	if strings.Trim(rootPath, " ") == "" {
		panic("Path argument cannot be empty!")
	}

	if strings.Trim(words, " ") == "" {
		panic("Word list cannot be empty!")
	}

	wordsSlice := make([]string, 0)
	for _, word := range strings.Split(words, ",") {
		wordsSlice = append(wordsSlice, strings.Trim(word, " "))
	}

	s := searcher.New(rootPath, threadLimit, wordsSlice, caseSensitive)

	start := time.Now()

	s.SearchDirectory()
	elapsed := time.Since(start)
	fmt.Printf("Hit count: %d\nFolders scanned: %d\nFiles scanned: %d\n", s.HitCount, s.FoldersScanned, s.FilesScanned)
	fmt.Printf("Elapsed time: %s\n", elapsed)
}
