//+build mage

package main

import (
	"os"
	"runtime"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/mholt/archiver/v3"
)

var gocmd = sh.RunCmd("go")

// Runs go mod download and then installs the binary.
func Build() error {
	mg.Deps(Resolve)
	println("Running build")
	if err := gocmd("build"); err != nil {
		return mg.Fatal(1, "build error : "+err.Error())
	}
	return Zip()
}

func Zip() error {
	println("Creating zip")
	fileExt := ""
	if runtime.GOOS == "windows" {
		fileExt = ".exe"
	}
	buildFile := "gocrudrest" + fileExt
	targetZip := "build.zip"
	files := []string{buildFile}
	if err := os.Remove(targetZip); err != nil {
		return mg.Fatal(1, "error deleting : "+err.Error())
	}

	if err := archiver.Archive(files, targetZip); err != nil {
		return mg.Fatal(1, "archive error : "+err.Error())
	}
	return nil
}

// Runs go mod download.
func Resolve() error {
	println("Resolving dependencies")
	err := gocmd("mod", "download")
	if err != nil {
		return mg.Fatal(1, "download error : "+err.Error())
	}
	return nil
}
