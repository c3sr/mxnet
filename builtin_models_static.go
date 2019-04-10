// Code generated by go-bindata.
// sources:
// builtin_models/AlexNet.yml
// builtin_models/ResNet18_v1.yml
// builtin_models/SSD_512_VGG16_Atrous_VOC.yml
// builtin_models/SSD_MobileNet_v1_COCO.yml
// DO NOT EDIT!

package mxnet

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _alexnetYml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\x4b\x8b\xe4\x36\x10\xbe\xfb\x57\x14\xf4\x61\x12\x98\xf1\xa3\xdd\xd3\xdb\x63\xc8\x42\xd2\x87\x4d\x20\x99\xc3\x42\x1e\xb0\x2c\x4d\x59\x2e\xb7\xb5\x2b\x4b\x46\x2a\x4f\x4f\xef\xaf\x0f\x7a\xf4\x8b\x5d\xb2\xe4\x62\x2c\xd5\x57\x55\x5f\x7d\x55\x92\x34\x8e\xd4\xc0\xcf\x8a\x5e\x9f\x89\x61\x01\x7e\x0d\xa6\x87\xa3\x99\x2d\x8c\xa6\x23\x95\xf5\x16\x47\x3a\x18\xfb\xb9\xc9\x00\x22\xfe\x8f\x7f\x22\xfa\x6c\x82\xde\x58\xe0\x81\x92\x0b\xc0\x0b\x59\x27\x8d\x6e\xe0\xee\xed\x4f\x55\x5e\xe5\xe5\xdd\x0d\x3c\x99\x41\x18\xcd\x16\xa5\xe6\xec\xec\x50\xe5\x25\x2c\xce\x00\xa9\x7b\x63\x47\xe4\xf8\x0f\x8e\x46\xd4\x2c\xc5\xd9\x1e\xad\x99\x8f\x83\x52\x93\x6d\x60\x01\xe7\x85\x83\xd9\x51\x07\x6c\x60\x22\xeb\x91\x91\x1e\xd0\x0b\xaa\x39\xc4\xcc\x00\x70\xec\xd6\x2b\x5f\x1a\xc0\x7e\x9a\x1b\xb0\x28\x27\x6b\x3e\x91\xe0\x42\xa0\x1d\xd5\x83\xc0\xbe\xa7\x26\xc0\x1e\xc4\x34\x07\xa4\xf8\x2e\x72\x1f\x90\xd3\x24\xd6\x2b\x45\xcd\x77\x9d\x12\x30\xb9\xfd\x37\x95\x6b\x6c\x47\x4e\x58\x39\x71\x90\xee\x6d\x06\xa9\x35\xbf\x8d\xb8\x27\xd8\x2a\x74\x4e\xf6\x52\x44\xfd\x42\xf1\xf7\x70\x18\xa4\x18\x40\x3a\x08\xca\x53\x07\x46\x87\xd6\x05\x1f\xef\xdc\x21\xa3\x23\xce\x33\x80\x3f\x1d\x9d\x87\xa3\xb7\x66\x84\x77\x6a\x36\x7a\xfb\x57\x12\xf2\x8b\x31\x79\x66\xa9\x27\x4b\x5a\x90\xf3\xe2\x5f\x56\x41\x77\x9c\x7c\x1b\x0a\x38\x50\xeb\x24\x93\xff\x25\x16\x79\x0e\x91\x78\x2b\xf5\xfe\x66\x6e\x1e\x60\x60\x9e\x5c\x53\x14\x7b\x9f\xe9\x41\xbc\xe4\xe3\xab\x26\xce\xa5\x29\x02\x66\xf7\xc5\x98\x42\xdc\x14\x96\x0f\x3c\xaa\x4c\x49\x41\xda\x51\x03\xb3\xb6\xe4\xd8\x4a\xc1\xd4\xc1\x02\xd2\xbe\x1f\xea\x4b\x22\xa9\xa7\x99\x03\xdf\x58\x48\x5c\x87\xfc\x7c\x9c\xa8\x01\x19\x04\x5c\x40\x2f\xad\xe3\x68\xf6\x50\x54\x92\x8f\xa1\x41\x37\xc2\xfb\xc0\x11\x73\xf2\xbb\x32\x9f\x32\x5f\x85\x0a\x11\x26\xf4\xc7\x81\xc9\xba\x38\x1e\x00\xa4\x68\x24\xcd\xbb\x48\x61\x96\x9a\x37\xc9\x12\xbc\x76\x0a\x8f\x7e\xc2\xcb\xb4\xa9\xf0\x68\x66\x6e\xe0\x6e\xfb\xeb\xdf\x77\x69\x4f\x18\x65\xec\xce\x17\xd5\xc0\xdd\xfb\x77\xbf\x9c\xf6\x3b\x39\x92\xf6\x27\xc6\x35\xf0\xa1\xbe\x87\xe5\x72\x15\x3e\x1f\x93\x7d\x24\xd4\x0d\x7c\xa8\x96\xf5\x3d\x54\xd5\x9b\x7b\xa8\xca\xd5\xc7\xcc\xcc\x3c\xcd\xec\xe9\x2d\x42\x09\x9e\xd8\xa9\x9c\x68\xcb\x20\x09\x76\xdb\x92\xe0\x81\xdf\x52\x21\xba\x5d\x8a\xcf\xbe\x21\x65\xc2\x28\x6c\xc3\x4c\x5c\x09\xe5\x79\x78\x0e\x97\xad\xec\x6b\xe1\x7a\x65\x90\xeb\x65\x14\xd9\x9a\x16\x5b\xa9\x24\x4b\x72\x67\x01\x61\x01\x52\x77\xf4\x7a\x22\x75\x83\x82\x80\xf2\xb7\xcd\x85\x4b\x4c\xd3\x13\xf2\x6c\xc9\xed\x66\xab\x9a\x30\xa8\x4d\x51\xb8\x3a\xc7\x11\xbf\x18\x8d\x07\x97\x0b\x33\x16\x8e\x8d\xa5\x3c\x1c\xd8\xdc\xd8\x7d\xe1\x8e\xda\x11\xbb\x22\x8c\x86\x26\x4e\x1b\x39\xbf\xf2\x6d\x54\x31\x90\xf8\xec\xe6\xb1\x81\x55\xb7\xac\x57\xed\xe3\xa6\xae\x51\xe0\x6a\xf5\xb4\xdc\x94\xeb\x47\xac\x36\x65\xd7\xd6\x65\xb5\xc6\x2c\x4c\xad\x57\xc3\x4d\x24\x64\xef\x59\xc7\x41\xde\x5b\x9c\x06\x40\xdd\xc1\x81\xe4\x7e\x60\x07\xce\xcc\x56\x90\x2f\xa0\x45\x47\xff\x8f\x7a\x88\xe9\x8a\x70\x02\xe3\x81\x14\x2f\x05\x2a\xf2\xeb\x0c\x62\xb2\xdd\x84\x3c\x34\x31\xfd\x83\x3b\x8e\xad\x51\xf9\x27\x17\x46\x20\x51\xb8\x41\x94\x65\x59\xe6\xa1\x7d\x9e\x92\x74\x3b\xb4\x62\x90\x2f\xe9\x8a\xec\x51\x39\x7f\x7c\x64\x0f\x8e\xf8\xde\x77\x20\xb6\xe1\xc4\xdd\xdf\x5b\x08\xfe\x87\x0d\xa0\x86\xe4\x1d\x9c\xe3\x8c\x5e\x48\x5d\xcb\x10\x37\x42\xb8\x8e\xb4\x61\xf2\xff\xc9\xab\x97\x8a\xc2\xc3\xe6\x4e\xf3\xf0\xb5\x8a\x07\xc9\x43\x9a\x88\x4b\xca\x98\xea\xaa\x6d\xd8\x76\x8f\x6f\x48\x6c\x36\xeb\xba\xef\x6b\xaa\x69\xf9\x44\xa2\x5b\x11\x76\xab\xba\xab\xfa\x2b\x49\x2e\x4e\x4f\xe5\xda\x37\x7b\x89\xeb\xb6\x5a\xb5\xd4\xa1\x58\x76\x42\x88\x16\x37\xd5\x9b\xcd\xe3\xf2\x29\x43\x66\x2b\xdb\x99\xe3\xfd\x4a\xaf\x6c\x31\x35\xfb\x62\xc9\x00\x3e\x4b\xdd\x35\xb0\x7d\x7e\x4e\x3a\xf8\xb5\xaf\x47\xd3\x6c\x51\x81\x26\x0e\x2f\xef\x0f\xdb\xe7\xe7\x7b\x78\xef\x3f\x79\x9e\xff\xe8\xcf\xaf\x7f\x06\xa4\xde\xef\xd2\xbd\xdf\x5c\x5e\x82\xc5\xe9\x2d\x38\x3f\xa4\xe1\x9d\x4f\x0e\x19\xc0\x88\x5a\xf6\xe4\x78\x87\x33\x0f\xc6\x36\xb0\x1d\x48\xef\xe1\x77\x99\xfd\x1b\x00\x00\xff\xff\x0f\xd0\x3d\x9c\x59\x08\x00\x00"

func alexnetYmlBytes() ([]byte, error) {
	return bindataRead(
		_alexnetYml,
		"AlexNet.yml",
	)
}

func alexnetYml() (*asset, error) {
	bytes, err := alexnetYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "AlexNet.yml", size: 2137, mode: os.FileMode(420), modTime: time.Unix(1554878324, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resnet18_v1Yml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x55\x4d\x8f\xe3\x36\x0c\xbd\xfb\x57\x10\xc8\x61\x5b\x60\xc7\x76\x26\x9e\xac\x63\xa0\x7b\x99\x43\x51\xa0\xcd\x61\x81\x16\xbd\x05\xb4\x4c\xd9\xda\xb1\x25\x43\xa2\xf3\x31\xbf\xbe\x90\xe4\xc4\x09\x76\xdb\xed\xc5\x90\xcc\xf7\xc8\x27\x92\xa2\x34\x0e\x54\xc1\x17\x72\x7b\xe2\x75\x79\x38\xae\x61\x05\xfe\x1f\x18\x09\x17\x33\x59\x18\x4c\x43\x7d\x22\x2d\x0e\x74\x32\xf6\xad\x4a\x00\x22\xe7\x8f\xbf\xf7\xc4\xb0\x82\x9b\x09\xa4\xb1\xc0\x1d\xcd\x14\x80\x23\x59\xa7\x8c\xae\xe0\xc3\xe7\x5f\xd6\xe9\x3a\xcd\x3f\x3c\xc0\x67\x33\x08\xa3\xd9\xa2\xd2\x9c\xdc\x08\xeb\x34\x87\xd5\x0d\xa0\xb4\x34\x76\x40\x8e\x6b\x70\x34\xa0\x66\x25\x6e\xf6\x68\x4d\xbc\x1f\x54\x9a\x6c\x05\x2b\xb8\x6d\x1c\x4c\x8e\x1a\x60\x03\x23\x59\x8f\x8c\xf2\x80\x8e\xd8\x4f\xc1\x67\x02\x80\x43\xb3\x2d\xfc\xd1\x00\xda\x71\xaa\xc0\xa2\x1a\xad\xf9\x4a\x82\x33\x81\x76\xe8\x9f\x04\x4a\x49\x55\x80\x3d\x89\x71\x0a\x48\xf1\x43\x64\x1b\x90\xe3\x28\xb6\x45\x4f\xd5\x0f\x49\x33\x70\xa6\xfd\xb7\x94\x7b\x6c\x43\x4e\x58\x35\x72\x48\xdd\xe7\x04\xe6\xd2\xfc\x36\x60\x4b\xf0\xda\xa3\x73\x4a\x2a\x11\xf3\x17\x0e\xff\x11\x4e\x9d\x12\x1d\x28\x07\x21\xf3\xd4\x80\xd1\xa1\x74\x81\xe3\xc9\x0d\x32\x3a\xe2\x34\x01\xf8\xd3\x11\x58\x72\xfa\xda\x20\xd2\x9a\x01\x7e\xed\x27\xa3\x5f\xff\x9a\x93\xf9\x6e\x4c\x9a\x58\x92\x64\x49\x0b\x72\xbe\x00\xcb\x2e\xe4\x1e\x47\x5f\x8a\x0c\x4e\x54\x3b\xc5\xe4\x97\xc4\x22\x4d\x21\x8a\xaf\x95\x6e\x1f\x7a\xe7\x09\x3a\xe6\xd1\x55\x59\xd6\xfa\x48\x4f\xe2\x98\x0e\x67\x4d\x9c\x2a\x93\x05\xcc\xe1\xdd\x98\x4c\x3c\x1c\x2e\xed\x78\xe8\x93\x5e\x09\xd2\x8e\x2a\x98\xb4\x25\xc7\x56\x09\xa6\x06\x56\x30\xff\xf7\x8d\xbd\x04\x52\x7a\x9c\x38\xe8\x8d\x07\x89\xfb\x10\x9f\x2f\x23\x55\xa0\x42\x12\x57\x20\x95\x75\x1c\xcd\x1e\x8a\xbd\xe2\x4b\x28\xd2\x43\xf2\xbd\xe3\x88\xb9\xf2\xee\xcc\xd7\xc8\x77\xae\x82\x87\x11\xfd\x95\x60\xb2\x2e\xb6\x08\x00\xf5\x34\x90\xe6\x43\x94\x20\x7b\x83\xbc\x79\x9e\x6d\x81\x77\xe8\xf1\xe2\xfb\x3c\x4f\xcc\xc4\xe3\xc4\x9e\x18\xc1\xb5\x99\x74\xa3\x74\x5b\x9b\x73\xf2\x1d\x71\x11\x7e\x43\x41\x6d\xce\xb0\x02\xfc\x9e\xcc\x19\x7a\x53\x97\x7c\xab\xf4\xdf\x75\x8e\xd6\xd4\x58\xab\x5e\xb1\x22\xb7\xa8\xf5\x26\x49\xc8\x93\x25\x77\x98\x6c\x5f\xdd\xaa\xec\x36\x29\x0e\xf8\x6e\x34\x9e\x5c\x2a\xcc\x90\x39\x36\x96\xd2\xd0\xf2\xa9\xb1\x6d\xe6\x2e\xda\x11\xbb\x4c\x18\x61\xb2\x1e\x6b\xea\x5d\xca\x67\x7e\x74\x29\x3a\x12\x6f\x6e\x1a\x2a\x28\x8a\x75\x29\xca\x97\x26\x5f\xd3\x56\x94\xcf\x52\x6e\xcb\x62\x57\x97\xcf\x9b\xdd\xa7\x97\x5d\xb1\x29\x93\x50\x6f\x5f\x78\x37\x92\x50\x52\x91\x9b\x5b\xa0\xb5\x38\x76\x80\xba\x81\x13\xa9\xb6\x63\x07\xce\x4c\x56\x90\xef\x8a\x1a\x1d\x2d\xba\xff\x8f\xec\xe0\xd3\x65\xa1\x77\x63\x2b\x8b\x63\x76\x77\x9b\x12\x88\x01\x0f\x23\x72\x57\x45\x09\x4f\xee\x32\xd4\xa6\x4f\xbf\xba\x30\x9c\x66\x19\x0f\x88\x3c\xcf\xf3\x34\x94\xc3\xcb\x52\xee\x80\x56\x74\xea\x38\x0f\x19\x89\xbd\xf3\xcd\xa7\x24\x38\xe2\x8f\xbe\x9c\xf1\x7a\x5f\xf5\xfb\x9b\x8f\xe0\x17\x6c\x00\x35\xcc\xec\x40\x5e\x05\xe4\x22\xea\x3e\x15\xf1\x47\x70\xd7\x90\x36\x4c\x7e\x3d\xb3\xa4\xea\x29\x3c\x0d\xee\xda\x41\xdf\x66\xf2\xa4\xb8\x53\x51\xca\x12\x32\x86\x5a\x4a\x57\xaf\xf3\x7a\xb7\x2e\xca\x26\xdf\xd6\xeb\xb2\xc0\x52\x60\x81\x72\x53\xe7\x5b\x49\x45\xbe\xa1\xbb\x94\x2c\x24\x89\x3b\xda\xbe\x94\x54\x96\x4d\xf1\xfc\xb2\xab\x37\xf9\x27\x7c\xc9\xeb\xcd\x16\xf3\x32\xaf\x1b\x91\x20\xb3\x55\xf5\xc4\x71\x3a\xd1\x99\x2d\xce\x05\x5f\x2c\x09\xc0\x9b\xd2\x4d\x05\xaf\xfb\xfd\x9c\x07\xbf\xf7\xe7\xd1\x34\x59\xec\x41\x13\x87\xb7\xeb\xa7\xd7\xfd\xfe\x23\x7c\xf1\x9f\x34\x4d\x7f\xf6\xd7\xcf\x0f\x52\xa5\xdb\xc3\x3c\x39\xab\x65\x96\xae\xae\xd3\xf4\xf6\x14\x85\x97\x72\x26\x24\x00\x03\x6a\x25\xc9\xf1\x01\x27\xee\x8c\xad\xe0\xb5\x23\xdd\xc2\xef\x2a\xf9\x27\x00\x00\xff\xff\x34\x16\x0b\xf4\x9f\x07\x00\x00"

func resnet18_v1YmlBytes() ([]byte, error) {
	return bindataRead(
		_resnet18_v1Yml,
		"ResNet18_v1.yml",
	)
}

func resnet18_v1Yml() (*asset, error) {
	bytes, err := resnet18_v1YmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ResNet18_v1.yml", size: 1951, mode: os.FileMode(420), modTime: time.Unix(1554868703, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ssd_512_vgg16_atrous_vocYml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\x5f\x6f\xdb\x36\x10\x7f\xd7\xa7\x38\xc0\x0f\x69\x81\x46\xb2\x1c\xd7\xb5\x05\xac\xc0\xe0\x02\x7d\xd9\x12\x60\xc5\x82\xbd\x09\x27\xea\x24\x71\x91\x78\x02\x79\xb2\x9d\x7c\xfa\x81\xa4\x6c\x27\x68\xb7\x62\x2f\x82\xc8\xfb\xdd\xff\xdf\x1d\x0d\x0e\x54\xc0\xb7\x6f\x5f\xca\x8f\xf9\xaa\x7c\xfc\xfa\x35\xdf\x94\xbf\x8a\xe5\xc9\x95\x8f\x0f\x7b\x58\x80\x07\x00\x37\xf0\xcc\x93\x85\x81\x6b\xea\x93\xa4\xb1\x38\xd0\x91\xed\x53\x91\x00\x44\x0b\xbf\xff\x75\x4f\x02\x0b\xb8\x88\xa0\x61\x0b\xd2\xd1\xac\x03\x70\x20\xeb\x34\x9b\x02\x6e\x3e\xff\x92\xa7\x79\xba\xbc\x79\x03\x9f\xc5\xa0\xd8\x88\x45\x6d\x24\xb9\x28\xe4\xe9\x12\x16\x17\x80\x36\x0d\xdb\x01\x25\xfe\x83\xa3\x01\x8d\x68\x75\x91\x47\x69\xe2\xed\xa0\x36\x64\x0b\x58\xc0\xe5\xe0\x60\x72\x54\x83\x30\x8c\x64\x3d\x32\x86\x07\x74\xc0\x7e\x0a\x36\x13\x00\x1c\xea\xcd\xda\xa7\x06\xd0\x8e\x53\x01\x16\xf5\x68\xf9\x6f\x52\x92\x29\xb4\x43\x7f\xab\xb0\x69\xa8\x08\xb0\x5b\x35\x4e\x01\xa9\x7e\x8a\x6c\x03\x72\x1c\xd5\x66\xdd\x53\xf1\x53\xa5\x19\x38\xab\xfd\x77\x28\xaf\xb1\x35\x39\x65\xf5\x28\xa1\x74\x9f\x13\x98\x5b\xf3\x50\x79\x2d\xf8\x42\x42\x2a\xd4\x2e\x24\xfe\x01\x8e\x9d\x56\x1d\x68\x07\xa1\xea\x54\x03\x9b\xd0\xb6\xfd\xc3\xfe\x01\xde\xed\x79\x18\xd8\xcc\xca\xce\xd7\x7b\xcf\x46\xe8\x24\xef\xa1\x46\x41\x47\x92\x26\x00\x7f\x3a\x02\xe7\xea\x40\xa1\x43\xdb\xe6\x9b\x12\x23\x85\x0e\xac\xa0\xb1\x3c\xc0\xd7\x7e\x62\xb3\x7f\x9c\xab\xfd\xc2\x9c\x26\x96\x1a\xb2\x64\x14\x39\xdf\xa1\xeb\x29\x34\x07\x47\xdf\xab\x0c\x8e\x54\x39\x2d\xe4\x7f\x49\x54\x9a\x42\xcc\xae\xd2\xa6\x7d\x43\xae\x5b\xe8\x44\x46\x57\x64\x59\xeb\x3d\xdd\xaa\x43\x3a\x9c\x0c\x49\xaa\x39\x0b\x98\xf2\x85\x39\xab\xcf\xc9\xa7\x9d\x0c\x7d\xd2\x6b\x45\xc6\x51\x01\x93\xb1\xe4\xc4\x6a\x25\x54\xc3\x02\xe6\x7b\xcf\xfa\xab\x0f\x6d\xc6\x49\x42\xa8\x31\x87\x78\x0e\xae\xe5\x79\xa4\x02\xf4\x80\x2d\x79\x4a\x6b\xeb\x24\x8a\x3d\x14\x7b\x2d\xcf\xa1\x81\x6f\x1a\xe3\x0d\x47\xcc\x59\xef\x95\xf8\xec\xf9\x95\xa9\x60\x61\x44\x3f\x2e\x42\xd6\x45\xfa\x00\x50\x4f\x03\x19\x29\x63\x08\x4d\xcf\x28\x77\xab\x59\x16\xf4\xca\x1e\x9f\xfd\x0c\xdc\xf8\x76\xdd\x24\x3c\xc9\x38\x89\xd7\x8e\x1a\x15\x4f\xa6\xd6\xa6\xad\xf8\x94\xfc\x20\xc2\x08\xbf\xa0\xa0\xe2\x13\x2c\x00\x7f\x14\xeb\x0c\xbd\x84\x98\x7c\x1f\xee\xbf\x07\x5b\xf1\x89\xdc\x39\xd4\x65\xcc\xd5\x72\x85\x95\xee\xb5\xe8\xab\x28\x8f\x53\xd3\xa3\x73\xd7\xcb\x68\xa2\x21\x94\xc9\x92\x2b\x27\xdb\x17\x17\x36\xb8\xbb\x14\x07\x7c\x61\x83\x47\x97\x2a\x1e\x32\x27\x6c\x29\x0d\xb3\x93\xb2\x6d\x33\xf7\x6c\x1c\x89\xcb\x14\x2b\xce\x7a\xac\xa8\x77\xa9\x9c\xe4\xad\x49\xd5\x91\x7a\x72\xd3\x50\xc0\x7a\x9d\x6f\xd5\xf6\x63\xbd\xcc\x69\xa3\xb6\xab\xa6\xd9\x6c\xd7\xbb\x6a\xbb\xba\xdb\x7d\xfa\xb8\x5b\xdf\x6d\x93\x40\x0e\xcf\x12\x37\x92\xd2\x8d\x26\x37\xf3\xa5\xb5\x38\x76\x80\xa6\x86\x23\xe9\xb6\x13\x07\x8e\x27\xab\xc8\x53\xa8\x42\x47\x21\xee\x04\x22\xae\x1c\x51\xba\xff\x91\x45\x70\xe1\x32\x21\xe3\xd8\x36\x3d\x1f\xcf\x37\x7e\x2a\x07\xae\x74\x4f\x86\xa4\x3c\xe4\xa5\xcf\xb3\x5c\x2d\xf3\x6d\xb9\xcc\xcb\xd5\x36\x6b\x2c\xbf\x90\x29\xb5\x99\xc7\xaf\x0c\xfe\xd3\xb1\x4a\x00\xb4\x2b\xd1\xaa\x4e\x1f\xe6\x6d\xd5\x60\xef\x3c\x53\x75\x03\x62\x27\xfa\xe0\xfb\x1e\x97\xc5\x39\x01\xbf\x47\x10\xfc\x8f\x30\xa0\x81\x59\x3d\x68\x2f\x02\xf2\x9a\xde\xeb\x5a\xc4\x8b\x60\xae\x26\xc3\x42\xfe\x7f\xd6\x6a\x74\x4f\xe1\x91\x71\x67\xaa\x7d\x5f\xca\xa3\x96\x4e\xc7\x50\xae\x2e\xa3\xab\x6b\xef\xd4\xae\x6e\x76\xcb\x2d\x7e\xfa\xb4\x5b\x6e\x56\x8a\x96\x4b\xac\x68\xdb\xec\xb6\xb4\xc6\x3b\xaa\xf2\x04\x45\xac\xae\x26\x89\x1b\x89\x4e\x62\x71\x6e\xde\x55\x92\x00\x3c\x69\x53\x17\xb0\xbf\xbf\x9f\x53\xf2\x67\x1f\x9a\xa1\xc9\x62\x0f\x86\x24\x3c\x68\xef\xf6\xf7\xf7\x1f\xe0\x0f\xff\x49\xd3\xf4\xbd\x1f\x39\xbf\x61\xb5\x69\xcb\x79\x75\x16\x71\xc9\x2e\xce\xab\xf4\xf2\x36\x85\xa7\x73\x06\x27\x00\x03\x1a\xdd\x90\x93\x12\x27\xe9\xd8\x16\xb0\xef\xc8\xb4\xf0\x9b\x4e\xfe\x09\x00\x00\xff\xff\xc7\x3f\x47\x9a\xbe\x07\x00\x00"

func ssd_512_vgg16_atrous_vocYmlBytes() ([]byte, error) {
	return bindataRead(
		_ssd_512_vgg16_atrous_vocYml,
		"SSD_512_VGG16_Atrous_VOC.yml",
	)
}

func ssd_512_vgg16_atrous_vocYml() (*asset, error) {
	bytes, err := ssd_512_vgg16_atrous_vocYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "SSD_512_VGG16_Atrous_VOC.yml", size: 1982, mode: os.FileMode(420), modTime: time.Unix(1554868546, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ssd_mobilenet_v1_cocoYml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\x4d\x6f\xe3\x46\x0c\xbd\xeb\x57\x10\xf0\x61\x77\x81\x8d\x64\x3b\xde\xc4\xd6\xa1\x17\x2f\xd0\x4b\xd7\x01\xba\x68\xd1\x9b\x40\x8d\x28\x69\x1a\x69\x28\xcc\x50\xb1\x93\x5f\x5f\xcc\x87\xed\x04\x9b\x76\xd1\x8b\xa0\x19\x3e\x92\x8f\xe4\x9b\x19\x83\x23\x95\xf0\xfd\xfb\xd7\xea\x1b\xd7\x7a\xa0\x03\x49\xf5\xb4\xaa\xf6\x0f\xfb\x07\x58\x80\xb7\x02\xb7\xf0\xcc\xb3\x85\x91\x1b\x1a\xb2\xd6\xe2\x48\x47\xb6\x8f\x65\x06\x10\xbd\xbf\xfd\x75\x20\x81\x05\x5c\x4c\xd0\xb2\x05\xe9\x29\xb9\x00\x3c\x91\x75\x9a\x4d\x09\xab\x7c\x93\x2f\xdf\x40\x93\x09\x14\x1b\xb1\xa8\x8d\x64\xaf\xc0\x1e\x7a\x06\x68\xd3\xb2\x1d\x51\xe2\x3f\x38\x1a\xd1\x88\x56\x17\x7b\xb4\x66\x3e\x0e\x6a\x43\xb6\x84\x05\x5c\x16\x0e\x66\x47\x0d\x08\xc3\x44\xd6\x23\x23\x35\xa0\x27\x1c\xe6\x10\x33\x03\xc0\xb1\xb9\xdb\xf8\xb2\x00\xba\x69\x2e\xc1\xa2\x9e\x2c\xff\x4d\x4a\x0a\x85\x76\x1c\x6e\x14\xb6\x2d\x95\x01\x76\xa3\xa6\x39\x20\xd5\x4f\x91\x5d\x40\x4e\x93\xba\xdb\x0c\x54\xfe\xd4\x29\x01\x93\xdb\x7f\x53\x79\x8d\x6d\xc8\x29\xab\x27\x09\xad\xfb\x25\x83\x34\x96\x87\xda\x7b\xc1\x57\x12\x52\xa1\x77\xa1\xf0\xcf\x70\xec\xb5\xea\x41\x3b\x08\x5d\xa7\x06\xd8\x84\x91\x85\xc9\x7f\xdc\xf3\x38\xb2\x49\xce\xce\xf7\x7b\xcf\x46\xe8\x24\x9f\xa0\x41\x41\x47\x92\x67\x00\x7f\x38\x02\xe7\x9a\xea\xcb\x6a\x5d\x8d\x41\x3e\x86\x64\x95\x2f\x2b\xc5\x8a\xa1\xb5\x3c\xc2\xaf\xc3\xcc\x66\xff\x67\x6a\xf7\x0b\x73\x9e\x59\x6a\xc9\x92\x51\xe4\xfc\x88\xae\xab\x30\x1d\x9c\xfc\xb0\x0a\x38\x52\xed\xb4\x90\xff\x25\x51\x79\x0e\xb1\xbc\x5a\x9b\xee\x8d\xb2\x6e\xa0\x17\x99\x5c\x59\x14\x9d\xcf\x74\xa3\x9e\xf2\xf1\x64\x48\x72\xcd\x45\xc0\x54\x2f\xcc\x45\x73\xae\x3e\xef\x65\x1c\xb2\x41\x2b\x32\x8e\x4a\x98\x8d\x25\x27\x56\x2b\xa1\x06\x16\x90\xf6\xbd\xe2\xaf\x39\xb4\x99\x66\x09\x54\x63\x0d\x71\x1d\x52\xcb\xf3\x44\x25\xe8\x11\x3b\xf2\x9a\xd6\xd6\x49\x34\x7b\x28\x0e\x5a\x9e\xc3\x04\xdf\x4c\xc6\x07\x8e\x98\xb3\xdf\x2b\xf3\x39\xf3\xab\x50\x21\xc2\x84\xfe\xbc\x08\x59\x17\xf5\x03\x40\x03\x8d\x64\xa4\x8a\x14\xda\x81\x51\x6e\xd7\xc9\x16\xfc\xaa\x01\x9f\xfd\x21\xf8\x10\xd2\x54\x42\xc6\xb1\xfd\x90\xf1\x2c\xd3\x2c\x3e\x4a\xf4\xac\x79\x36\x8d\x36\x5d\xcd\xa7\xec\x1d\xa6\x11\x7e\x41\x41\xcd\x27\x58\x00\xbe\xc7\x39\x41\x2f\x54\xb3\x1f\x69\xff\x3b\xe9\x9a\x4f\xe4\x2e\x94\x2f\xe3\xaa\xc2\xfe\x87\xd8\x03\xcb\x35\xd6\x7a\xd0\xa2\xdf\x85\x3a\xc5\xf6\x8c\x55\x03\x3a\xf7\x2e\x2a\x59\x22\xac\x25\x94\xd9\x92\xab\x66\x3b\x94\x17\x21\xb9\xdb\x1c\x47\x7c\x61\x83\x47\x97\x2b\x1e\x0b\x27\x6c\x29\x0f\xe7\x2e\x67\xdb\x15\xee\xd9\x38\x12\x57\x78\x95\x17\x03\xd6\x34\xb8\x5c\x4e\xf2\x36\xa4\xea\x49\x3d\xba\x79\x2c\x61\xb3\x59\x6d\xd5\xf6\x4b\xb3\x5c\xd1\x9d\xda\xae\xdb\xf6\x6e\xbb\xd9\xd5\xdb\xf5\xed\xee\xfe\xcb\x6e\x73\xbb\xcd\x82\xae\xbc\xc0\xdc\x44\x4a\xb7\x9a\x5c\x92\x5a\x67\x71\xea\x01\x4d\x03\x47\xd2\x5d\x2f\x0e\x1c\xcf\x56\x91\x57\x5f\x8d\x8e\x02\xef\x0c\x22\xae\x9a\x50\xfa\xff\x51\x45\x48\xe1\x8a\x28\x8c\x76\xe0\xe3\x79\xc7\x9f\xe8\xcb\x69\xf6\x8f\x81\xaf\xb3\x5a\x2f\x57\xdb\x6a\xb9\xaa\xd6\xdb\xa2\xb5\xfc\x42\xa6\xd2\x26\x9d\xdc\x2a\xe4\xcf\xa7\x3a\x03\xd0\xae\x42\xab\x7a\xfd\x94\x6e\xba\x16\x07\xe7\x45\xae\x5b\x10\x3b\xd3\x67\x2f\x95\x78\xd1\x9c\x0b\xf0\x77\x10\x82\xff\x11\x06\x34\x90\xdc\x83\xf7\x22\x20\xaf\xe5\xbd\xee\x45\xdc\x08\xe1\x1a\x32\x2c\xe4\xff\x93\x57\xab\x07\x0a\x8f\x93\x3b\xab\xf3\xc7\x56\x1e\xb5\xf4\x3a\x52\xb9\xa6\x8c\xa9\xae\xb3\x53\xbb\xa6\xdd\x2d\xb7\x78\x7f\xbf\x5b\xde\xad\x15\x2d\x97\x58\xd3\xb6\xdd\x6d\x69\x83\xb7\x54\xaf\x32\x14\xb1\xba\x9e\x25\x5e\x66\x74\x12\x8b\x69\x78\x57\x4b\x06\xf0\xa8\x4d\x53\xc2\xfe\x70\x48\x25\xf9\xb5\xa7\x66\x68\xb6\x38\x80\x21\x09\x8f\xe1\xc7\xfd\xe1\xf0\x19\x7e\xf7\x9f\x3c\xcf\x3f\xf9\x53\xea\x6f\x67\x6d\xba\x2a\x5d\xbb\x25\xa4\xa7\x39\xad\x2f\xef\x5a\x78\x72\x13\x38\x03\x18\xd1\xe8\x96\x9c\x54\x38\x4b\xcf\xb6\x84\x7d\x4f\xa6\x83\xdf\x74\xf6\x4f\x00\x00\x00\xff\xff\x78\x32\x06\xf5\xf2\x07\x00\x00"

func ssd_mobilenet_v1_cocoYmlBytes() ([]byte, error) {
	return bindataRead(
		_ssd_mobilenet_v1_cocoYml,
		"SSD_MobileNet_v1_COCO.yml",
	)
}

func ssd_mobilenet_v1_cocoYml() (*asset, error) {
	bytes, err := ssd_mobilenet_v1_cocoYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "SSD_MobileNet_v1_COCO.yml", size: 2034, mode: os.FileMode(420), modTime: time.Unix(1554868546, 0)}
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
	"AlexNet.yml": alexnetYml,
	"ResNet18_v1.yml": resnet18_v1Yml,
	"SSD_512_VGG16_Atrous_VOC.yml": ssd_512_vgg16_atrous_vocYml,
	"SSD_MobileNet_v1_COCO.yml": ssd_mobilenet_v1_cocoYml,
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
	"AlexNet.yml": &bintree{alexnetYml, map[string]*bintree{}},
	"ResNet18_v1.yml": &bintree{resnet18_v1Yml, map[string]*bintree{}},
	"SSD_512_VGG16_Atrous_VOC.yml": &bintree{ssd_512_vgg16_atrous_vocYml, map[string]*bintree{}},
	"SSD_MobileNet_v1_COCO.yml": &bintree{ssd_mobilenet_v1_cocoYml, map[string]*bintree{}},
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

