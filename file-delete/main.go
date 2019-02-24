package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// DeleteParentDirectory is delete Parent Directory
func DeleteParentDirectory(path string) error {
	// RemoveAllはディレクトリが存在しなくても削除する
	path = filepath.Dir(path)
	fmt.Println(path)
	return os.RemoveAll(path)
}

func main() {
	fmt.Println("init")
	if err := DeleteParentDirectory("hoge/fuga"); err != nil {
		fmt.Println(err)
	}
}
