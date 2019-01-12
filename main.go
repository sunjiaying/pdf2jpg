package main

import (
	"flag"
	"fmt"
	"runtime"
	"os/exec"
	"io/ioutil"
	"path"
	"time"
)

func main(){
	pdffilepath := flag.String("pdf", "", "Input PDF Filenpath")
	outputpath := flag.String("output", "", "Output path")

	flag.Parse()
	
	fmt.Println("pdffilepath:", *pdffilepath)
	fmt.Println("outputpath:", *outputpath)

	t1 := time.Now()
	if runtime.GOOS == "windows" {
		cmd := exec.Command("magick.exe", "convert", "-density", "200", *pdffilepath, *outputpath + "image-%03d.jpg")

		out, err := cmd.Output()
		if err != nil {  
			fmt.Println(err)
		}
		fmt.Println(string(out))
		listJpgs(*outputpath)
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

func listJpgs(pathname string) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
			if path.Ext(fi.Name()) == ".jpg" {
					fmt.Printf("%s\n", fi.Name())
			}
	}
	return err
}