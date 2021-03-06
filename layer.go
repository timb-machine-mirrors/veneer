/*
Copyright 2020 The veneer Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package veneer

import (
	"archive/tar"
	"io"
	"io/ioutil"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/spf13/afero"
)

// LayerFs builds the filesystem for an image layer.
func LayerFs(layer v1.Layer, fs afero.Fs) error {
	return Walk([]v1.Layer{layer}, func(h *tar.Header, r io.Reader) error {
		data, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}
		return afero.WriteFile(fs, h.Name, data, h.FileInfo().Mode())
	})
}
