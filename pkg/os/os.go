package os

import (
    "io/ioutil"
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
    "strings"
)

func RemoveContents(dir string) error {
    d, err := os.Open(dir)
    if err != nil {
        return err
    }
    defer d.Close()
    names, err := d.Readdirnames(-1)
    if err != nil {
        return err
    }
    for _, name := range names {
        err = os.RemoveAll(filepath.Join(dir, name))
        if err != nil {
            return err
        }
    }
    return nil
}


func ReplaceFileContent(filename string, searchFor string, replaceWith string) error {
    visit := func(path string, fi os.FileInfo, err error) error {
        read, err := ioutil.ReadFile(path)
        if err != nil {
            panic(err)
        }
        newContents := strings.Replace(string(read), searchFor, replaceWith, -1)
        err = ioutil.WriteFile(path, []byte(newContents), 0)
        if err != nil {
            panic(err)
        }
        return nil
    }
    return filepath.Walk(filename, visit)
}

func GetBasePath() string {
    _, b, _, _ := runtime.Caller(0)
    return filepath.Dir(b)
}

func RunCmd(command string, path string, stdout *os.File) error {
    cmd := exec.Command("/bin/sh", "-c", command)
    cmd.Stdout = stdout
    cmd.Dir = path
    return cmd.Run()
}
