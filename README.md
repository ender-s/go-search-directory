# go-search-directory
A command-line tool written in Golang that performs file search under the given directory based on filenames using multiple goroutines.

# Usage
To see usage of the tool, run the following command:
```bash
$ go run main.go --help
```
Output:
```bash
  -case-sensitive
        case sensitive (look for perfect match in terms of letter case)
  -max-thread-count int
        number of maximum threads running concurrently (default 16)
  -path string
        path to the directory to be searched
  -words string
        words to be searched in path names (separated by commas)
```

# Record of a Sample Run
[![asciicast](https://asciinema.org/a/QBugLtetriyfM1zLQEhZftMQY.svg)](https://asciinema.org/a/QBugLtetriyfM1zLQEhZftMQY)