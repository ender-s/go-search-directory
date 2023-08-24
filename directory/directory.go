package directory

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	AbsolutePath     string
	RightTrimmedPath string
}

func (d *Directory) SearchKeywords(keywords []string, caseSensitive bool) (bool, []string, []string, int, int) {
	files, err := ioutil.ReadDir(d.AbsolutePath)

	if err != nil {
		log.Printf("error while searching the directory %v:%v\n", d.AbsolutePath, err)
		return false, nil, nil, 0, 0
	}

	folderNames := make([]string, 0, 10)
	hits := make([]string, 0, 10)

	result := false
	fileCounter := 0
	dirCounter := 0
	for _, file := range files {
		fname := file.Name()
		if !caseSensitive {
			fname = strings.ToLower(fname)
		}

		for _, keyword := range keywords {
			if strings.Contains(fname, keyword) {
				result = true
				hits = append(hits, file.Name())
				break
			}
		}
		if file.IsDir() {
			folderNames = append(folderNames, file.Name())
			dirCounter++
		} else {
			fileCounter++
		}
	}
	return result, folderNames, hits, dirCounter, fileCounter
}

func New(path string) *Directory {
	rightTrimmedPath := strings.TrimRight(path, strconv.QuoteRune(os.PathSeparator))
	return &Directory{AbsolutePath: path, RightTrimmedPath: rightTrimmedPath}

}
