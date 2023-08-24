package searcher

import (
	"fmt"
	"go-search-directory/directory"
	"path/filepath"
	"strings"
	"sync"
)

type WorkerStarted struct{}

type Searcher struct {
	rootPath         string
	goRoutineChannel chan WorkerStarted
	keywords         []string
	caseSensitive    bool
	wg               *sync.WaitGroup
	mutex            sync.Mutex
	FilesScanned     int
	FoldersScanned   int
	HitCount         int
}

func (s *Searcher) searchDirectory(dir *directory.Directory) {
	s.goRoutineChannel <- struct{}{}
	//fmt.Printf("Searching directory %s\n", dir.AbsolutePath)
	found, folderNames, hits, dirCount, fileCount := dir.SearchKeywords(s.keywords, s.caseSensitive)
	s.mutex.Lock()
	s.FoldersScanned += dirCount
	s.FilesScanned += fileCount
	if found {
		fmt.Printf("[+] %s\n", dir.AbsolutePath)
		for _, hit := range hits {
			s.HitCount += 1
			fmt.Printf("\t%s\n", hit)
		}
	}
	s.mutex.Unlock()

	for _, folderName := range folderNames {
		s.wg.Add(1)
		go s.searchDirectory(directory.New(
			filepath.FromSlash(
				fmt.Sprintf("%s/%s", dir.RightTrimmedPath, folderName))))
	}

	<-s.goRoutineChannel
	s.wg.Done()
}

func (s *Searcher) SearchDirectory() {
	s.wg.Add(1)
	go s.searchDirectory(directory.New(s.rootPath))

	s.wg.Wait()
}

func New(rootPath string, maxGoRoutineCount int, keywords []string, caseSensitive bool) *Searcher {
	if !caseSensitive {
		for i := 0; i < len(keywords); i++ {
			keywords[i] = strings.ToLower(keywords[i])
		}
	}

	return &Searcher{rootPath: rootPath,
		goRoutineChannel: make(chan WorkerStarted, maxGoRoutineCount),
		keywords:         keywords,
		caseSensitive:    caseSensitive,
		wg:               new(sync.WaitGroup),
	}
}
