// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var convertMap = [][2]string{
	[2]string{
		`"github.com/mikespook/gearman-go`,
		`"github.com/gopkg/gearman`,
	},
	[2]string{
		`"github.com/mikespook/golib/signal`,
		`"github.com/gopkg/gearman/example/worker/signal`,
	},
}

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal("filepath.Walk: ", err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, "gen_helper.go") {
			return nil
		}
		if strings.HasSuffix(path, "hello.go") {
			return nil
		}
		if strings.HasSuffix(path, ".go") {
			fixImportPath(path)
		}
		return nil
	})
}

func fixImportPath(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("ioutil.ReadFile: ", err)
	}

	for _, v := range convertMap {
		oldPath, newPath := v[0], v[1]
		data = bytes.Replace(data, []byte(oldPath), []byte(newPath), -1)
	}
	if err = ioutil.WriteFile(filename, data, 0666); err != nil {
		log.Fatal("ioutil.WriteFile: ", err)
	}
	fmt.Printf("convert %s ok\n", filename)
}
