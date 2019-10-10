// Code generated by go-bindata.
// sources:
// ../../../../../../../../../tmp/vizier_yamls/etcd.yaml
// ../../../../../../../../../tmp/vizier_yamls/nats.yaml
// ../../../../../../../../../tmp/vizier_yamls/vizier.yaml
// DO NOT EDIT!

package cmd

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _tmpVizier_yamlsEtcdYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\x4d\x6f\xdb\x38\x10\xbd\xeb\x57\x10\xb9\x04\x58\x40\x4a\x9c\xec\x61\x97\x37\xd7\x71\x8b\x02\x49\x6c\xc4\x6e\x81\x9e\x8c\x11\x35\x4e\xd8\xf0\xab\xe4\x48\x68\xf2\xeb\x0b\x9a\x92\x62\x5b\x76\x1a\x34\xe5\x49\x1a\xbc\x99\x37\x7c\xf3\x38\xe0\xe4\x57\xf4\x41\x5a\xc3\x99\x2f\x41\x14\x50\xd3\x83\xf5\xf2\x19\x48\x5a\x53\x3c\xfe\x17\x0a\x69\xcf\x9a\x51\x89\x04\xa3\xec\x51\x9a\x8a\xb3\x89\xaa\x03\xa1\xbf\xb3\x0a\x33\x8d\x04\x15\x10\xf0\x8c\x31\x05\x25\xaa\x10\xbf\x18\x03\xe7\x38\x73\x2a\xd7\xd6\x48\xb2\x5e\x9a\xfb\x8c\x31\x03\x1a\x39\x43\x12\x55\x6e\x1d\x7a\x20\xeb\x33\x5f\x2b\x0c\x3c\xcb\x19\x38\xf9\xc9\xdb\xda\x6d\x0a\xe4\x1b\x54\x11\x2b\x97\x10\xb0\x10\xd6\xa3\x0d\x85\xb0\x3a\x63\xcc\x63\xb0\xb5\x17\xb8\x85\x14\xa9\xa5\xd0\x07\x4a\x10\x8f\xb5\x7b\xf9\xf7\x18\xc8\x7a\x8c\x81\x06\x7d\xd9\xa6\x9e\xfe\x73\x3a\x64\x06\x27\xf1\x27\xa1\x89\xa2\x84\x56\x81\x21\xab\xa8\x03\x59\xdd\x05\x2b\x5c\x4b\x23\xa3\x64\x6f\x61\x38\x39\x19\xd6\x73\xb6\x4a\xdd\x06\xf4\x8d\x14\xd8\xb6\x6e\x2a\x67\xa5\xa1\xf4\xe7\xe2\xa4\x02\xa1\xa1\xc6\xaa\x5a\xa3\x50\x20\x75\x0b\x6c\x30\xa1\x7e\x7f\xbb\x8d\x2a\x7b\xec\x15\x3a\x65\x9f\xf4\x1b\x6b\x1c\xea\x3f\xa0\xf0\xb8\x9f\x7d\x8f\x94\xe5\x79\x9e\xbd\xd3\x65\x1f\xa4\xa9\xa2\x83\xde\x6f\x36\xab\xf0\x0e\xd7\x31\xab\xbb\xd2\x2b\x0d\x65\x8c\x0d\x0d\x7f\xb8\x70\xa8\xcb\xef\x28\x68\x63\xe4\x94\xb3\x48\x63\x1c\x0b\x61\x6b\x43\x7d\x5a\x85\x6b\xa8\x55\xf7\x1f\x1c\x08\x8c\xad\x0f\x54\x8a\x63\x3a\x6b\x3a\x2d\xae\xfa\xf1\xbc\x5f\x83\x7d\xe6\xe0\x50\xf0\xcd\x3c\x9d\x92\x02\x02\x67\xa3\x8c\xb1\x80\x0a\x05\x59\x9f\x8a\x6b\x20\xf1\x70\xbd\xc5\x76\x84\x8f\x50\x3b\x05\x84\x6d\xd6\x56\xab\xf1\xa8\x9d\x02\x47\x4a\xa4\x73\xb8\x71\xc6\xba\x66\xe3\x11\xd6\x10\x48\x83\xbe\x2f\x99\x33\x61\xb5\x06\x53\xbd\x70\xe4\x07\x8a\xc4\x83\xa6\xd9\x06\x25\xbe\x9b\x6f\xab\xf9\xec\x6a\x75\x3b\xbe\x99\x2e\xe6\xe3\xc9\xb4\x07\x30\xd6\x80\xaa\xf1\xa3\xb7\x9a\x6f\x05\x19\x5b\x4b\x54\x55\xeb\xa8\x41\x7c\x0e\xf4\xc0\x7b\x15\x8a\x5e\xf6\xd7\x78\xff\x3e\x65\x0f\x93\x1a\xee\x91\xb3\x1f\x35\x3c\xc5\xd7\x96\x56\xea\xd9\x8e\x3c\xbc\x39\x2f\xfe\x2f\xfe\x7d\x75\x0c\xfb\x4e\x3d\xb6\xa8\xdb\xe7\x7c\xd1\x5a\x78\x4a\xa2\x6a\xdf\xd1\x9f\x78\xd8\xa9\x3c\x12\x1d\x75\xef\xf2\x7a\x91\x6a\x04\x02\x92\xbd\x47\x34\xea\x12\xfd\x8b\x54\x0e\xd1\x2f\x36\x9b\xaa\xbd\x56\x0c\xe4\xa4\x42\x2e\xd0\x6f\xb6\x57\x3a\x71\x09\xef\x21\x53\x68\x80\xed\x74\xd9\xc1\x0a\x25\xd1\xd0\x0e\xd6\xd9\xd6\x95\x11\x30\xed\xdc\xd7\x39\x60\xba\x9c\x5c\xad\xc6\x5f\x96\xb3\xd5\x64\x76\x33\x1f\x4f\x96\x9f\x67\xb7\xab\xbb\xe9\x72\x7a\x1b\xbf\xb2\x2d\x43\x70\x76\x32\x3a\x8f\xfb\x37\xc8\x67\xe4\xec\x32\x2d\xdc\x34\x8a\xcb\xe2\xa2\x18\x5d\x66\xbf\x02\x00\x00\xff\xff\x09\x0c\x05\x30\xd4\x07\x00\x00")

func tmpVizier_yamlsEtcdYamlBytes() ([]byte, error) {
	return bindataRead(
		_tmpVizier_yamlsEtcdYaml,
		"tmp/vizier_yamls/etcd.yaml",
	)
}

func tmpVizier_yamlsEtcdYaml() (*asset, error) {
	bytes, err := tmpVizier_yamlsEtcdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmp/vizier_yamls/etcd.yaml", size: 2004, mode: os.FileMode(436), modTime: time.Unix(1570684389, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmpVizier_yamlsNatsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\x4d\x6f\xe3\x38\x0c\xbd\xfb\x57\x08\xbd\x14\x58\xc0\x4e\xb2\x45\x81\xc2\xb7\x6c\xbf\xb0\xd8\x6d\xd6\x68\x80\x05\xe6\x54\x30\x32\x93\x68\x22\x4b\x82\x44\x1b\x4d\x7f\xfd\x40\x96\xe3\xd8\x8d\x93\xe9\xb4\xd3\x19\x9d\x2c\xea\xf9\x91\x7c\xa4\x68\x83\x11\xff\xa3\x75\x42\xab\x94\x55\x93\x68\x23\x54\x9e\xb2\x39\xda\x4a\x70\x9c\x72\xae\x4b\x45\x51\x81\x04\x39\x10\xa4\x11\x63\x12\x16\x28\x9d\x7f\x62\x0c\x8c\x49\x99\x91\x71\xa1\x95\x20\x6d\x85\x5a\x45\x8c\x29\x28\x30\x65\x0a\xc8\xc5\xda\xa0\x05\xd2\xb6\xb1\x3a\x03\x1c\xfd\x0b\x51\x1c\xc7\xd1\x67\x3a\x76\x68\x2b\xfc\xbe\x5b\xbb\x00\x9e\x40\x49\x6b\x6d\xc5\x0b\x90\xd0\x2a\xd9\x5c\xb9\x44\xe8\x51\x1b\xd0\xb5\x2c\x1d\xa1\x7d\xd4\x12\x3f\x2e\x83\x2d\x25\xba\x34\x8a\x19\x18\x71\x6f\x75\x69\x6a\x82\x7a\x8b\xcf\x84\xca\x47\xe5\x9a\x10\x22\xc6\x2c\x3a\x5d\x5a\x8e\x0d\x8a\x97\x8e\x74\xb1\x33\xe6\xb8\x14\x4a\xf8\x98\x5d\xc4\x58\x85\x76\xd1\xc0\xce\xff\x38\x3f\xf4\xe0\xe3\x18\x24\xf5\x07\x3c\xe4\xe8\x5a\x83\x0b\x55\xb0\x5a\xe2\x5b\xc8\xcf\xce\x06\x82\xd5\x6a\x29\x56\x05\x98\xc0\xea\x90\x5b\xa4\xf0\x6c\x74\xbe\x7f\x18\xe1\x33\xf2\xfd\x4e\xea\x55\x83\xaf\x23\x70\xdd\x0d\x84\xa6\x18\xb4\x8d\x48\x6f\x50\xd5\x27\xa8\x72\xa3\xc5\x0e\x87\x15\x86\xc7\xf7\x24\xd1\x76\xcf\x2b\x02\x29\x1c\xfd\xae\x66\x6a\x5a\xfb\x58\x2b\x0d\xa6\xa1\xf3\xc1\x3a\x7e\x28\x81\xbf\x84\xca\x7d\x70\x1f\xbd\x14\xf1\xa2\x21\xf2\xdd\xf6\x88\x4b\xff\xf6\x2e\xa7\x13\x31\x45\x8c\x1d\x6a\x3a\x7c\xeb\x5c\xb9\xf8\x8a\x9c\x6a\xb5\x06\xa7\xcc\xfb\x66\xd6\xaf\x95\x2b\x94\xfd\xd3\xc4\x6a\xba\xea\xc7\xa4\x7a\xe3\x94\x05\x63\xdc\x5e\x93\x1b\x34\x52\x6f\x0b\xfc\x8c\xcf\x8a\x33\xc8\xd3\xba\xff\x8d\x14\x1c\x5c\xca\x26\x11\x63\x0e\x25\x72\xd2\x36\x90\x17\x40\x7c\xfd\x6f\xc7\xdb\x11\x7f\x7e\x0d\xfb\x24\x2c\x8c\x04\xc2\x86\xaf\x93\x84\x5f\xb2\x47\x7d\x82\xfc\x18\x3d\x63\xbb\x34\xfc\xe2\x5a\x11\x08\x85\xb6\xa5\x8c\x19\xd8\x55\xc7\x41\x3c\xc0\xe0\x17\xaa\xaa\x0f\xf2\xce\x1e\xbe\x3c\x65\xff\xdd\x3c\xcd\xa6\x0f\xb7\xf3\x6c\x7a\x7d\xdb\x02\x18\xab\x40\x96\x78\x67\x75\x91\x76\x8c\x8c\x2d\x05\xca\xbc\x69\xb5\x03\x7b\x06\xb4\x4e\x5b\x09\x92\xb6\x1a\xa7\xfc\xfe\x7c\x97\x2d\x4c\x14\xb0\xc2\xd4\x6b\xa6\x90\x13\x56\x68\xb7\xb4\x16\x6a\x35\xea\x29\x94\x8e\x93\xcb\x64\x1c\x57\x13\x90\x66\x0d\x7f\xf6\xdf\xce\x4a\x29\x33\x2d\x05\xdf\xa6\xec\xef\xe5\x4c\x53\x66\xd1\x61\xdd\xf9\xa7\x8b\xe6\x97\xd1\x96\x7a\x95\x69\xab\x97\x69\x4b\x29\xbb\x1a\x5f\x8d\x3b\x39\x05\x2a\x8b\x90\x6f\x5f\x5a\xb3\xdf\x0a\x85\xce\x65\x56\x2f\xb0\x2b\xc1\x9a\xc8\xdc\x23\xf5\x55\x31\xb5\x1c\xa3\x57\x24\xbb\x60\x0e\xd8\x19\xab\x7f\x1a\x40\xde\xa0\x84\xed\x1c\xb9\x56\xb9\xbf\x26\x97\x1d\x04\x89\x02\x75\x49\xed\xe1\x45\x73\xe6\x7a\xa3\x60\x36\xa0\xc3\xeb\x9b\xdf\xfc\x75\x8c\x5a\xa9\xc3\x08\x98\x01\xb9\x66\x12\xbd\x67\x06\x18\x19\x7b\xe2\xa3\xb7\xdf\x89\x17\x0c\x51\xd3\x8e\x2d\x4c\xaa\x79\xfd\x13\x92\xee\x12\x89\x49\xba\x98\xa3\xad\xa9\xfa\x98\xeb\xe9\x9d\x90\x18\x52\xe4\x90\x70\x4b\x87\x10\xb4\xb4\x07\x85\x93\x41\xe0\x3f\xb8\x3d\xc0\x6d\x70\x1b\xbe\xc7\x41\xa7\x49\x72\x91\x8c\xa3\x6f\x01\x00\x00\xff\xff\x4f\x82\x99\xca\x8a\x0b\x00\x00")

func tmpVizier_yamlsNatsYamlBytes() ([]byte, error) {
	return bindataRead(
		_tmpVizier_yamlsNatsYaml,
		"tmp/vizier_yamls/nats.yaml",
	)
}

func tmpVizier_yamlsNatsYaml() (*asset, error) {
	bytes, err := tmpVizier_yamlsNatsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmp/vizier_yamls/nats.yaml", size: 2954, mode: os.FileMode(436), modTime: time.Unix(1570684389, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmpVizier_yamlsVizierYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x9a\xdb\x73\xda\x3e\x16\xc7\xdf\xf9\x2b\x34\x79\xe9\xcc\xce\x18\xc8\x6d\x37\xf5\x4c\x1f\x28\xa1\x29\xdb\x84\x7a\x30\x6d\xa7\x4f\x1e\x21\x1f\x40\x45\x96\x54\x49\xa6\x25\x7f\xfd\x8e\xaf\xd8\x5c\x1d\x08\x69\x37\x3f\xfc\x84\x75\x39\x3a\x3a\xfa\xea\xa3\x83\x6d\x2c\xe9\x57\x50\x9a\x0a\x6e\x23\x35\xc4\xa4\x8e\x43\x33\x11\x8a\x3e\x62\x43\x05\xaf\x4f\x6f\x74\x9d\x8a\xc6\xec\x7c\x08\x06\x9f\xd7\xa6\x94\xfb\x36\x6a\xb3\x50\x1b\x50\x7d\xc1\xa0\x16\x80\xc1\x3e\x36\xd8\xae\x21\xc4\xf0\x10\x98\x8e\x7e\x21\x84\xa5\xb4\x91\x64\x56\x20\x38\x35\x42\x51\x3e\xae\x21\xc4\x71\x00\x36\x9a\xd1\x47\x0a\xca\x22\xa0\x4c\x30\x56\x35\x15\x32\xd0\x76\xcd\x42\x58\xd2\x3b\x25\x42\x19\x5b\xb0\xd0\xd9\x59\x0d\x21\x05\x5a\x84\x8a\x40\x5a\x26\x85\xaf\xe3\x1f\x1a\x88\x02\x13\xfd\x9e\x81\x1a\xa6\xb5\x6f\xfe\xf5\xa6\x66\x59\x56\xed\xcf\xcd\x29\xeb\xb9\xef\xa4\xd4\x8c\x12\x48\x6e\x80\xfb\x52\x50\x7e\xa4\x39\xbe\xa7\xdc\x8f\xfc\x7f\x86\xe5\x13\x0c\xfa\x30\x8a\xba\x65\x73\xdd\xe2\x51\x0d\xa1\xd5\x78\x6f\xb0\xac\xc3\xe1\x0f\x20\x26\x0e\x63\xd2\xc9\x4d\x02\xd4\x22\x44\x84\xdc\xe4\xfd\x7c\x18\xe1\x90\x65\xf7\x5a\x62\x02\x91\xf3\x7f\x47\x9c\x16\x92\x78\xf6\x40\xe5\xa6\x9f\x3d\x52\xb3\x2c\x14\xa9\x9d\x03\xe6\x8f\x25\xb5\x52\x61\xaf\x0c\xab\x25\x90\xc8\x8e\x14\xca\xa4\xfa\x4e\xfa\x1a\x22\xad\x89\x31\xf2\x22\x1e\x24\xaa\xb6\xd1\x75\xf3\xa2\xd9\x4c\xee\x95\x30\x82\x08\x66\xa3\x41\xdb\x89\x4b\x0c\x56\x63\x30\x4e\xa9\x9d\x06\x06\xc4\x08\xb5\xc5\xd1\x55\x57\x6b\x08\x99\xb9\x84\x3c\xec\x5d\xe7\x98\xb1\xc9\x94\x7e\x60\x5c\xde\x56\x8c\xcb\xdb\x3d\xe3\xb2\x70\xf3\x05\x63\x93\x8b\xfb\xc0\xe0\x5c\x55\x0c\xce\xd5\x9e\xc1\x29\xf8\xf9\x82\xd1\x91\x4a\xfc\x9e\xef\xbf\xaf\x74\x21\x44\x57\x57\x97\x55\x02\x74\xdd\x8c\x03\xb4\x6c\xa9\x60\xe8\xa6\x52\xa0\xaf\x9b\xcd\xf3\xbd\x02\x1d\x4f\x39\x8f\x72\x4f\xf8\x10\x59\x3c\x66\x90\x7f\x86\xa0\xe6\xd6\x50\x89\x29\x1c\xbc\x47\x2f\x2b\xca\xf0\x72\x4f\x19\x2e\xf9\xba\x4b\x8a\x58\x4a\xdd\xc8\x43\x75\x0b\x92\x89\x79\x00\xdc\x1c\x06\xfa\x8d\x41\x52\x20\x19\x25\x58\xdb\x68\x75\xed\x03\x6c\xc8\xe4\xbe\x30\xd4\xc6\x09\xaf\xc7\x35\x04\x92\x61\x03\xa9\xb1\x82\xfb\xd1\xc5\x4a\x76\xb7\x58\x5e\x6b\x1b\xa1\x6c\x02\xd1\x45\x04\x37\x98\x72\x50\xb9\xbd\x28\x4b\x9b\x2d\x8c\x67\x0a\x70\xee\xbd\xff\x7e\x1b\x78\x6e\xf7\xae\xd7\xed\xdd\x79\x9f\x3a\xdf\xf3\x26\x08\xcd\x30\x0b\xe1\x83\x12\x81\x5d\x28\x44\x69\x42\xfb\x09\xe6\x69\x8a\x50\xbc\xa6\x30\xb7\xd1\x8f\x5f\xc6\xd2\x74\xcc\x29\x1f\x5b\x53\x98\x2f\x35\x49\x06\x96\xcc\x22\xc9\x92\x5b\x8b\x04\x79\xc5\x37\xb7\xd3\xff\xda\xe9\x7b\x83\x7b\xd7\x6b\x77\xfa\x83\x65\xdf\x6c\xd4\x88\x80\xaf\x1b\x11\x5d\x40\xd5\x89\x32\x3b\xac\xac\x99\xe0\xb2\x91\xa2\xc7\x05\x23\xed\xfb\x6e\xa7\x37\xd8\xed\x0a\x61\x14\xb8\xd9\xe4\x4a\xc1\xca\x36\x57\x52\x23\x1b\x5c\x89\x7d\x68\xed\x70\x03\x97\x5c\xa0\x01\x1e\x83\x8d\xc6\x44\x45\x39\xa4\x64\x96\x0f\x33\x8b\xf2\x91\xc2\x8d\x44\x46\x0d\x2c\xa9\x97\x44\xc0\x4b\x1a\x47\x4a\xd5\x66\xa3\xe6\xac\xa4\x71\x5e\x9f\xc3\x25\x73\x37\xd7\xa0\xb3\x94\x13\x25\xd7\x44\x68\xb3\xae\x66\x26\x58\x18\xc0\x43\x94\x0b\x96\xcc\x05\x51\x89\x83\xcd\x24\x9b\x62\x6d\x59\x52\xc5\xd2\x78\x0a\x4e\xc8\x98\x9b\x88\x6b\xb1\x0d\x72\xf9\xc5\x4d\x52\xf1\xd5\x8a\x63\xaf\x34\x2e\x0f\x97\xf4\x28\x0a\x3f\x29\xe9\xc5\x6d\xd3\x73\xce\x32\x4c\x5b\x49\xbf\x17\xe0\xd9\xae\xe4\xec\x28\x4c\x2b\xa4\x5a\x47\xe1\xda\xc2\xfe\x89\x6d\x27\xb6\x6d\x61\x5b\xa1\xf7\x6d\xd7\x6d\xbd\xbf\xef\x78\x77\x7d\xa7\xed\xb5\xbe\x0c\x3e\xae\xda\x38\x33\x2a\x84\xb3\x75\x9d\x7b\xad\x87\x8e\xeb\xb4\xda\x9d\xdd\x32\x19\x51\x60\xfe\x1a\x89\xc4\xe5\x09\xa4\xb2\x6d\x50\xcf\xb7\xe4\x13\x60\x9c\x6a\xbf\x3a\x90\x8b\x9b\x05\x55\xa2\xf1\xdb\x13\x73\x0f\x63\x2e\x13\xa1\x6f\x11\xc1\x79\x0c\xd3\x17\x66\xef\xca\xe0\x47\x62\xf0\xca\x38\xaf\x93\xc5\xed\xfb\x2f\xee\xa0\xd3\xf7\xba\xb7\xcf\xe0\x56\x36\x1e\xf5\x4f\xa7\xc3\x9f\x3f\x1d\x80\xcf\xca\x6b\x18\xb3\x70\x44\xc7\x0f\x58\xae\xac\x60\x61\x85\x52\xe1\x8f\xe8\xf8\x29\xdc\x8e\xba\x79\xf9\x7e\x79\x02\xbf\xcb\x1b\xed\xe9\xc9\xf5\xcd\xc6\xe4\xfa\xe6\x04\xfa\x83\x40\xbf\xf3\xe9\xde\x51\x08\x5f\x7c\x56\x77\x14\xb4\x17\x06\x78\x9d\x4c\x3f\x11\xf4\x6f\xc8\xaf\x77\x33\x33\x13\x62\x75\x58\x96\xa4\x8b\x4e\x58\xdb\xff\xb1\xfc\xcb\x32\x2d\x7f\x2c\x7e\x14\xa0\x65\xd6\x77\xd3\x0c\xab\x71\x49\x26\x96\xc5\xa8\x36\xc0\x0b\x45\xe7\x17\xff\xa9\x37\xeb\xcd\xfa\xb9\x7d\x7d\x59\x6a\x99\xbe\x96\xb4\x14\x68\xc1\x8a\x47\x74\x54\x89\xa5\x04\xee\x5b\x1a\xb0\x22\x13\xcb\x17\x01\xa6\x5c\x97\x5a\x44\x07\xb3\x1e\x51\x06\xef\x1a\x60\x48\x23\xbe\x5d\xde\x2b\x3f\x30\x07\xf2\x38\x0d\x1b\x63\x61\xf9\x5c\x07\x58\xff\xb4\x15\x30\xc0\x1a\xac\xf3\x7a\xb3\x7e\xbd\x34\xff\xb4\xcd\xff\x35\xac\x77\xa2\xd1\x30\xbd\x09\x4a\xbb\x80\x18\x75\x2d\x7a\xb7\x1b\x49\xb1\x94\xaa\xf3\x68\xf1\x86\xeb\x29\x29\x5b\xfa\x96\x2a\xb9\x14\x60\x9f\x72\xd0\xda\x51\x62\x08\xc5\x60\x4f\x8c\x91\x77\x65\x76\x20\x24\x13\xa6\x4d\x00\x33\x33\x79\x2c\x57\xad\xb1\x1e\x2f\x26\x99\x40\xe4\xf3\xc7\xc1\xc0\x71\x5f\x0b\x39\x93\xc8\xbf\x28\x37\x2b\xbd\x69\x3b\x0a\x3e\x97\xdf\x9b\x1d\x85\xa2\x4b\x83\x9c\x52\xc3\x53\x6a\xf8\xe7\x52\xc3\x58\x8c\x5e\x22\xc6\xea\x38\x5e\x91\x30\xaa\xf4\x0f\xfa\xf2\x1f\xfc\x3f\x19\x43\x20\xb8\x0b\x7b\x71\x11\x8f\x81\xaf\x7e\xa6\x95\x61\xe3\x50\xf6\x65\xd6\x9f\x99\x76\x99\xd9\x7d\x92\x45\x6d\xa8\x62\x94\x8f\x3d\xe0\x78\xc8\xc0\x93\x58\xe9\xe8\x36\xfe\x5e\x62\x18\x8e\xf4\xbb\x48\xe7\x79\x97\x4d\x80\x7c\x75\x3b\xdf\x75\xef\xd7\xec\xf9\x11\x66\x7a\xfd\xfb\x16\xa7\xff\xb9\xed\x39\xad\x75\xa4\x88\x13\xe3\x28\x0d\x23\x4f\xa0\x45\xbc\xa6\xdb\xf8\xb0\x58\x74\x84\x96\x3e\x6d\xcd\x2e\x46\x03\x6a\x74\xf9\xf8\x09\x20\x10\x6a\x6e\xa3\x8b\xeb\x7f\x3f\xd0\x42\x8d\x82\x9f\x21\xe8\xe5\xd6\x44\x86\x36\x3a\x6f\x36\x83\x0a\x36\x34\x90\x50\x51\x33\x6f\x0b\x6e\xe0\x77\x69\x3b\x13\x2c\xf1\x90\x32\x6a\x28\x2c\x0d\x80\x7d\xbf\x5c\x60\x21\xf7\xbb\xeb\x39\x83\x7e\xf9\xd5\x55\x56\xd1\xba\x7d\xe8\xf6\x0a\xe5\x52\xd1\x19\x65\x30\x06\xdf\x46\x25\xa1\x56\x81\x5e\xb4\x2e\x2b\xcc\x8b\x0a\x2d\x25\x84\x29\x05\x07\xfb\x9f\x39\x9b\x2f\x8d\x51\xb6\xa6\xe7\xab\x00\x2d\x97\x55\x32\xb3\x9b\xc4\x3e\xd7\x8e\x60\x94\xcc\xf3\x0f\x8d\x3e\x50\xa5\xcd\x37\x6a\x26\x1f\x85\x36\xbd\x1c\xbe\x93\xe4\xee\x97\x50\xd3\xd2\x90\xf1\x03\xd5\xee\x6d\xa9\x6c\x3f\xbc\x1b\x50\x01\xe5\xf1\xf7\xb3\x77\x0a\x13\x70\x40\x51\xe1\xbb\x40\x04\xf7\xb5\x8d\x2e\xb3\x43\xc8\x08\x06\x2a\x6e\x56\x4c\xb6\x46\x23\x20\xc6\x46\x3d\xe1\x92\x09\xf8\x21\x5b\x84\x24\x4e\x91\xb8\xf0\xc1\x52\x82\x41\x7d\x1a\x0e\x41\x71\x30\x10\x7f\x34\x1c\xe0\x68\xce\x9b\x0e\x98\x78\x72\x51\x28\x8b\x2a\x49\x42\x5b\x28\x49\xbe\xd3\xba\xa5\x2a\xe6\xf9\xbc\xb6\x5d\x03\x5b\xad\x96\xd7\x78\xbb\xe1\x45\xdb\xe7\x39\x09\xff\x17\x00\x00\xff\xff\x7e\xb2\x09\x95\x46\x30\x00\x00")

func tmpVizier_yamlsVizierYamlBytes() ([]byte, error) {
	return bindataRead(
		_tmpVizier_yamlsVizierYaml,
		"tmp/vizier_yamls/vizier.yaml",
	)
}

func tmpVizier_yamlsVizierYaml() (*asset, error) {
	bytes, err := tmpVizier_yamlsVizierYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmp/vizier_yamls/vizier.yaml", size: 12358, mode: os.FileMode(436), modTime: time.Unix(1570684389, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"tmp/vizier_yamls/etcd.yaml": tmpVizier_yamlsEtcdYaml,
	"tmp/vizier_yamls/nats.yaml": tmpVizier_yamlsNatsYaml,
	"tmp/vizier_yamls/vizier.yaml": tmpVizier_yamlsVizierYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"tmp": &bintree{nil, map[string]*bintree{
		"vizier_yamls": &bintree{nil, map[string]*bintree{
			"etcd.yaml": &bintree{tmpVizier_yamlsEtcdYaml, map[string]*bintree{}},
			"nats.yaml": &bintree{tmpVizier_yamlsNatsYaml, map[string]*bintree{}},
			"vizier.yaml": &bintree{tmpVizier_yamlsVizierYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

