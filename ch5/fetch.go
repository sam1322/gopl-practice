package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {

	urls := os.Args[1:]
	for _, url := range urls {
		fileName, n, err := fetch(url)
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		fmt.Println("filename : ", fileName, ", bytes copied", n)
	}
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
			//err = fmt.Errorf("Hello this is a testing error ")
			fmt.Println("Sup")
		}
	}()
	if err != nil {
		return "", 0, err
	}
	fmt.Println("going to copy the html into the file")
	n, err = io.Copy(f, resp.Body)
	//if closeErr := f.Close(); err == nil {
	//	err = closeErr
	//}
	return local, n, err
}
