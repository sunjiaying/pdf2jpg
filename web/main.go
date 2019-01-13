package main

import (
	"net/http"
	"strings"
	"log"
	"fmt"
	"runtime"
	"os/exec"
	"io/ioutil"
	"path"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println(r.Form)
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
		fmt.Fprintf(w, "Hello Wrold!")
		
		pdffilepath := "d:/1/test.pdf"
		outputpath := "d:/1/"

		t1 := time.Now()
		if runtime.GOOS == "windows" {
			cmd := exec.Command("magick.exe", "convert", "-density", "200", pdffilepath, outputpath + "image-%03d.jpg")

			out, err := cmd.Output()
			if err != nil {  
				fmt.Println(err)
			}
			fmt.Println(string(out))
			listJpgs(outputpath, w)
		} else {
			cmd := exec.Command("convert", "-v")

			out, err := cmd.Output()
			if err != nil {  
					fmt.Println(err)  
			}  
			fmt.Println(string(out))
		}

		t2 := time.Now()
		fmt.Println()
		fmt.Println(t2.Sub(t1))
}

func main() {
    http.HandleFunc("/", sayhelloName)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func listJpgs(pathname string, w http.ResponseWriter) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
			if path.Ext(fi.Name()) == ".jpg" {
					fmt.Fprintf(w, fi.Name())
			}
	}
	return err
}