package main // import "download"

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/dustin/go-humanize"
)

// WriteCounter : Progress counter
type WriteCounter struct {
	Total uint64
}

// Write : Use as io.Writer for io.TeeReader
func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

// PrintProgress : Print progress to console
func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

// DownloadFile : Download file
func DownloadFile(filepath string, url string) error {
	// Create filename + tmp
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}

	// Receive data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	// Create & pass progress
	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	// Below closes lock the file so, defer should not use
	resp.Body.Close()
	out.Close()

	fmt.Print("\n")

	// Remove tmp
	err = os.Rename(filepath+".tmp", filepath)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Download Started")

	// Original example
	// fileURL := "https://upload.wikimedia.org/wikipedia/commons/d/d6/Wp-w4-big.jpg"
	// err := DownloadFile("world-map-big.jpg", fileURL)

	// Below are for practices

	// MinGW64 - https://sourceforge.net/projects/mingw-w64/files/Toolchains%20targetting%20Win64/Personal%20Builds/mingw-builds/8.1.0/threads-posix/seh/
	fmt.Println("MinGW64")
	fileURL := "https://sourceforge.net/projects/mingw-w64/files/Toolchains%20targetting%20Win64/Personal%20Builds/mingw-builds/8.1.0/threads-posix/seh/x86_64-8.1.0-release-posix-seh-rt_v6-rev0.7z/download"
	err := DownloadFile("mingw64.7z", fileURL)
	if err != nil {
		panic(err)
	}

	// Go - https://golang.org/dl/
	fmt.Println("Go, Golang")
	fileURL = "https://dl.google.com/go/go1.11.2.windows-amd64.zip"
	err = DownloadFile("go.zip", fileURL)
	if err != nil {
		panic(err)
	}

	// Git - https://git-scm.com/download/win
	fmt.Println("Git for Windows")
	fileURL = "https://github.com/git-for-windows/git/releases/download/v2.19.2.windows.1/PortableGit-2.19.2-64-bit.7z.exe"
	err = DownloadFile("git.7z", fileURL)
	if err != nil {
		panic(err)
	}

	// FileZilla - https://filezilla-project.org/download.php?show_all=1
	fmt.Println("FileZilla")
	fileURL = "https://dl3.cdn.filezilla-project.org/client/FileZilla_3.38.1_win64.zip?h=2s_xL1javWLHC767beBiEQ&x=1543484750"
	err = DownloadFile("filezilla.zip", fileURL)
	if err != nil {
		panic(err)
	}

	// PuTTY - https://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html
	fmt.Println("PuTTY")
	fileURL = "https://the.earth.li/~sgtatham/putty/latest/w64/putty.zip"
	err = DownloadFile("filezilla.zip", fileURL)
	if err != nil {
		panic(err)
	}

	// VSCode - https://code.visualstudio.com/docs/?dv=winzip
	fmt.Println("Visual Studio Code")
	fileURL = "https://go.microsoft.com/fwlink/?Linkid=850641"
	err = DownloadFile("vscode.zip", fileURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Download Finished")
}
