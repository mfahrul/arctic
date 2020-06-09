package app

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

//Execute func
func Execute(dir string, name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Dir = dir

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	// Execute the command
	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}

	log.Println(stdBuffer.String())

}

//Curdir func
func Curdir() string {
	path, err := os.Getwd()
	if err != nil {
		Amsyong(err)
	}
	return path
}

//GetInput func
func GetInput(message string, scanner bufio.Scanner, required bool) (input string) {
	for {
		fmt.Print(message)
		scanner.Scan()
		input = scanner.Text()
		if len(input) == 0 {
			if required {
				continue
			} else {
				break
			}
		} else {
			break
		}
	}
	return
}

// IsDirExist func
func IsDirExist(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

// CopyDir copies a whole directory recursively
func CopyDir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if fd.Name() == "src" {
				continue
			}
			if err = CopyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = copyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

//MoveFile func
func MoveFile(oldLocation string, newLocation string) (err error) {
	err = os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// copyFile copies a single file from src to dst
func copyFile(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

//Amsyong keluar dari program
func Amsyong(msg interface{}) {
	log.Fatalln(msg)
}

//Log func
func Log(msg interface{}) {
	log.Println("LOG ===>", msg)
}

//BackupConfig func
func BackupConfig() {
	if err := copyFile(WorkDir+"App"+string(filepath.Separator)+"config"+string(filepath.Separator)+"config.go", WorkDir+"App"+string(filepath.Separator)+"config"+string(filepath.Separator)+"config.go.temp"); err != nil {
		panic(err)
	}
}

//RestoreConfig func
func RestoreConfig() {
	if err := MoveFile(WorkDir+"App"+string(filepath.Separator)+"config"+string(filepath.Separator)+"config.go.temp", WorkDir+"App"+string(filepath.Separator)+"config"+string(filepath.Separator)+"config.go"); err != nil {
		panic(err)
	}
}
