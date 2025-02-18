package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(filepath.Base(os.Args[0]))
	fmt.Println(strings.Join(os.Args[1:], " "))
}
