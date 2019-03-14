// Copyright 2019 Greg Lanthier
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build ignore

package main

import (
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filepath.Walk("client", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		name := info.Name()
		if strings.HasSuffix(name, ".generated.go") {
			return nil
		}
		if !strings.HasSuffix(name, ".go") {
			return nil
		}
		gen := name[0:len(name)-len(".go")] + ".generated.go"
		oldpath, _ := filepath.Abs(filepath.Join(filepath.Dir(path), name))
		newpath, _ := filepath.Abs(filepath.Join(filepath.Dir(path), gen))
		return os.Rename(oldpath, newpath)
	})
}
