package io

import "io/ioutil"
import "os"

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func read(path string) string {
  dat, err := ioutil.ReadFile(path)
  check(err)
  return string(dat)
}

func write(path string, data string) {
  dat := []byte(data)
  err := ioutil.WriteFile(path, dat, 0644)
  check(err)
}

func list(path string) []string {
  files, _ := ioutil.ReadDir(path)
  stringFiles := make([]string, len(files))
  for i, file := range files {
    stringFiles[i] = file.Name()
  }
  return stringFiles
}

func exists(path string) bool {
  _, err := os.Stat(path)
  return os.IsNotExist(err)
}