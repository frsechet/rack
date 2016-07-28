// Code generated by go-bindata.
// sources:
// provider/aws/templates/service/fluentd.tmpl
// provider/aws/templates/service/syslog.tmpl
// DO NOT EDIT!

package aws

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

var _templatesServiceFluentdTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x58\x5b\x6f\xdb\x36\x14\x7e\xf7\xaf\x10\x84\x02\x05\x0a\x5b\x4e\xdd\x0e\xc3\x08\xec\xc1\x8b\xed\x2e\x9b\xbb\x1a\x76\xda\x3e\x0c\x79\x60\xe8\x63\x47\xb3\x2c\x0a\x14\x95\x2b\xfc\xdf\x77\x48\x5d\x22\x52\x94\xe2\x04\x89\x5a\x18\x90\xf9\xf1\xdc\xcf\x77\x8e\xf3\xf0\xe0\xad\x61\x13\xc6\xe0\xf9\x29\x88\xeb\x90\x81\xef\x1d\x0e\xbd\x87\x9e\xe7\xf9\xe3\x9f\xab\x73\xd8\x27\x11\x95\x30\xe3\x62\x4f\xe5\x0f\x10\x69\xc8\x63\x9f\x78\xfe\xe8\xe4\xe3\xc9\xe0\xe4\x37\xfc\xef\xf7\x15\xf6\x5b\x26\x93\x4c\xa6\x78\xa4\xae\x7a\x1e\xca\x15\x34\xde\x82\xf7\x6e\xd7\xf7\xde\xd1\x24\xf1\xc8\xef\x5e\x30\x4e\x92\x54\x89\xf7\xf4\xe3\x23\x28\x4b\x12\x10\x1a\x10\xfc\x43\xf7\x80\x87\xf3\x30\xde\x55\x62\x34\xec\x07\x8d\x32\x50\x4a\x11\xaf\x91\x85\xb2\x60\xce\xb7\x5f\x04\xcf\x12\xbc\xe5\x17\xf0\x43\xbf\x54\x0f\xf1\xba\x54\xe5\x7f\x17\x51\x4d\x64\x25\xb0\xa6\x63\x09\x1b\xa5\x41\x01\x4b\x49\xbd\xf2\x53\xcb\xf4\x17\x54\xa0\x81\x12\x43\x50\xdd\xb4\xe5\x4e\x20\x65\x22\x4c\x64\x11\xa3\x19\x6a\x89\xe5\xc4\xfb\xbe\x9c\xf7\x3d\x08\xb6\x81\xf7\x5e\xb2\x84\x0c\x87\x1b\x7d\xb0\x1e\x30\x1e\x45\xc0\x24\x17\x01\xdc\x52\x0c\x34\x04\x8c\xef\xc9\xe8\xf3\x68\xf4\xf9\xbd\xdf\x2f\xa5\x9e\xdf\x25\xda\xfb\x95\x14\x61\xbc\xf5\x2d\xb3\x96\x90\xf2\x4c\x30\x78\x8d\xd0\x2f\x40\xec\xc3\xb4\x48\x71\x2d\x38\x0b\xc1\x11\x2a\xc3\x9a\x92\xe2\x64\xcc\x4a\x67\x23\xba\xbf\x5c\x53\x72\x16\x5f\xf3\x1d\xcc\xb2\x38\x3f\xe8\xd7\xc1\xe5\xb7\x4a\x9b\x25\xe8\x31\x03\xd5\xd5\xda\xe9\xc1\x10\xb3\xc0\x30\xb0\x30\xa1\x51\x53\xc6\x2c\x26\xe4\x2f\x1e\x2a\x83\xfe\x35\x4e\xf0\x2c\x30\x8c\x51\x8f\x0d\x41\x50\xc4\xb7\x69\x03\xe7\x59\x6a\x4c\x83\xb1\x4b\x08\x59\xc2\xd6\xb2\xd9\x61\x79\x71\x8f\xee\xe9\x3d\x8f\xe9\x4d\xaa\xb2\x6d\x5f\xb9\xe8\xb5\xbd\x99\x41\x58\xe9\xa4\x8f\x19\xe3\x59\x2c\x5b\x83\xa9\x6d\x2b\x40\x67\xeb\xf6\x90\x16\xd2\x44\xfc\xbc\x90\x1e\x13\x51\x2a\x62\x82\xbe\x12\x15\x59\xf2\xd6\xa1\x7d\x91\x02\x67\x7c\x3a\x74\xa0\x27\x83\xad\xa2\x1d\xd2\xc1\x47\xe4\xc3\x33\x32\xdb\x73\xe8\xab\xda\x5e\xdb\x38\xcf\xbb\x8b\xd4\x3a\xb4\x67\x5d\x71\xb7\xf4\x2a\xbb\xac\x38\x69\x16\x46\x48\x5f\x66\x6b\x4f\x20\x41\xa2\x4c\xbf\xd9\xf9\x7d\x92\x20\x2a\xec\x45\xff\x18\xa2\x40\x6a\x94\x61\x4c\x95\x19\xad\x65\xf6\x05\xe4\x58\x4a\x57\xa1\x39\xf9\x44\x9f\x28\x61\x47\xb6\x4c\xee\xfe\x82\x4a\xfc\xd4\x94\x65\xb2\x53\x99\xbc\x82\x9d\x8e\x19\x36\x5d\x09\x53\xe5\x4e\x1c\xd1\xef\x9c\x53\x95\x9f\xb5\xa1\xd2\x12\x53\xff\x94\xaf\x6d\x16\xf5\x57\x9f\xfe\xc8\xd8\x0e\x1c\x84\xd0\xd5\xc6\x83\x63\xfa\x98\x71\x64\xf6\xdb\x57\x6d\xe0\x97\x52\xdf\xa7\xbf\xe1\x4e\x49\x17\x10\x01\x4d\x61\xa8\x83\x78\x8d\x4a\xa7\xf3\xe9\x78\x35\x55\x0b\x4c\x39\x62\x83\xfb\x30\x71\x67\xcb\x9c\xd5\x46\x14\x9b\xbb\x80\x79\xb5\x63\x8e\x19\x8e\xaf\x24\x65\x3b\x0d\x72\x8a\xf9\x93\xc6\xeb\x48\xf7\xa3\x1f\xc6\x6b\xb8\x0d\xae\x8a\x2f\x6a\x98\x25\x8f\x1a\x2a\x3a\x5a\x25\xc7\x9b\x39\xb2\x5b\xe4\xc2\x69\xcc\x12\x19\x30\xcc\x0b\x3f\xc6\xc2\xfa\xaf\x3e\x05\xfd\x73\x3c\xe1\x99\xd4\x7b\xdf\x2f\x4d\xe2\x71\x32\x95\x39\xc7\x0b\xb0\xed\x4f\x6b\x71\x8f\xd3\x34\xdb\x83\x42\x2f\x78\x14\xb2\xbb\x09\x67\xf8\xde\x98\x74\xb8\x13\xe1\x72\x5a\x1c\x98\xa1\xb0\xcb\xf2\x71\x55\x71\x14\x77\x2a\x53\xf2\xa8\xb2\xc1\xdc\x0d\xd6\x99\x6e\x36\xb8\xb7\x69\x9f\xa3\x88\xdf\x34\x69\xa9\x7d\x47\xc9\xcd\x2e\x56\x6d\x97\x31\x5e\xb9\x4b\x05\x9d\x8b\x82\xdd\x2f\x76\x3f\xd5\xdf\x0c\x07\x7c\x73\x8b\x1f\x0d\x70\x91\xff\xf8\xab\xbb\x44\x91\x2c\xaf\x14\x6e\x58\xaf\x06\x9d\x91\x3c\x63\x75\xeb\x2d\xc2\xe9\xcc\x5b\x1e\x84\xd6\xdc\x35\xc5\x15\x57\x3a\x72\xa8\xcf\x59\xc4\xb3\xf5\x0d\x95\xec\x8a\x2c\x32\xf9\x15\x70\x63\x66\x13\x2a\xa9\x83\xb6\x34\xde\xbd\xb4\x3a\xb0\x8d\x12\xd0\xd7\x9f\x2a\x03\x0d\x2a\x17\x74\x05\x6b\x2c\x05\xce\x15\xe3\x45\xae\xeb\xfd\xea\x54\x00\xc6\xb4\x9c\x54\xad\x5e\x1b\x50\xfc\x59\x01\x74\xdf\x89\xc5\x58\x22\x70\x7a\x8d\xb9\x4a\xdf\x30\x3a\xc6\xaa\xf8\x41\xfd\x7b\xbb\x70\xe9\x4a\xd9\xe8\x5f\xb5\x88\x23\xf9\x34\xb8\x04\x4d\xd9\x6f\xe8\xa3\xcb\xd6\xee\xe9\x5c\x61\x5a\x32\xe4\x1a\xd6\xd5\x9d\x32\xa2\x96\xb7\xad\xa2\xdc\xc1\xac\x39\xf2\xf4\x54\x2f\x1f\x47\xa2\x2a\x39\xaf\x62\x40\xc7\xda\x7e\x94\x0d\xa9\x4a\xf5\xf0\x35\x2c\x71\xcc\xf9\xe7\x58\x32\x74\x94\x79\xfe\xd8\x0c\xdf\xfe\xad\xbd\x4b\x39\xb6\xab\xc6\x8c\x7a\x62\x0e\x38\xcc\x2e\x48\xbd\xdc\x90\xf3\x51\x7f\xaa\x6a\x6b\x56\xd6\x96\x7e\xfb\xa9\x18\xd8\x77\xfe\xc2\xb9\xe8\xde\x1e\xce\xc6\x5f\xb1\xb4\xaa\x19\xac\xff\xcc\xd1\x3b\xf4\x1e\xd7\xe4\xff\x03\x00\x00\xff\xff\xb2\xbe\x00\x03\xb0\x12\x00\x00")

func templatesServiceFluentdTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesServiceFluentdTmpl,
		"templates/service/fluentd.tmpl",
	)
}

func templatesServiceFluentdTmpl() (*asset, error) {
	bytes, err := templatesServiceFluentdTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/service/fluentd.tmpl", size: 4784, mode: os.FileMode(420), modTime: time.Unix(1469711986, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesServiceSyslogTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x58\xdd\x6e\xdb\x36\x14\xbe\xf7\x53\x10\x44\x81\x02\x9d\x23\xc7\x09\x8a\x61\x04\x76\xe1\x25\x75\x97\x2d\x5d\x0d\x3b\x6d\x2f\x86\x5c\x30\x34\xed\x68\x96\x45\x81\xa4\x92\x26\x81\xdf\x7d\x87\xd4\x4f\x44\x8a\x52\x9c\x20\x51\x0b\x03\x32\x3f\x9e\xff\xf3\x9d\xe3\x3c\x3c\xa0\x25\x5f\xc5\x29\x47\x58\x71\x79\x13\x33\x8e\xd1\x6e\x37\x78\x18\x20\x84\x27\x3f\x16\x17\x7c\x9b\x25\x54\xf3\xa9\x90\x5b\xaa\xbf\x73\xa9\x62\x91\x62\x82\xf0\xd1\xe1\xf8\xf0\xe0\xf0\x37\xf8\x8f\x87\x06\xfb\x35\xd7\x59\xae\x15\x1c\x99\xab\x08\x81\x5c\x49\xd3\x35\x47\xef\x36\x43\xf4\x8e\x66\x19\x22\xbf\xa3\x68\x92\x65\xca\x88\x47\xf6\xc1\x00\xca\xb3\x8c\x4b\x0b\x88\xfe\xa1\x5b\x0e\x87\xe7\x71\xba\xa9\xc5\x58\xd8\x77\x9a\xe4\xdc\x28\x05\xbc\x45\x96\xca\xa2\x73\xb1\xfe\x2c\x45\x9e\xc1\x2d\x5c\xc2\x77\xc3\x4a\x3d\x4f\x97\x95\x2a\xfc\x4d\x26\x0d\x91\xb5\xc0\x86\x8e\x39\x5f\x19\x0d\x06\x58\x49\x1a\x54\x9f\x56\x26\x9e\x51\x09\x06\x6a\x08\x41\x7d\xd3\x97\x7b\xca\x15\x93\x71\xa6\xcb\x18\x2d\xee\x54\x22\xd6\xe8\xdb\xfc\x7c\x88\x78\xb4\x8e\xd0\x7b\xcd\xb2\x5f\x74\xa2\xc8\x68\x04\x07\x6a\x1c\x65\x14\xbc\xd7\x92\xc6\x89\x71\x8b\x89\x2d\x19\x8f\x8f\x8e\x3f\xbe\xc7\xc3\x4a\xe4\xc5\x5d\x66\x5d\x5f\x68\x19\xa7\x6b\xec\xd9\x34\xe7\x4a\xe4\x92\xf1\xd7\x88\xfb\x8c\xcb\x6d\xac\xca\xfc\x36\x22\x33\x93\xc2\x18\x19\x37\x94\x94\x27\x13\x56\x79\x9a\xd0\xed\xd5\x92\x92\xb3\xf4\x46\x6c\xf8\x34\x4f\x8b\x83\x61\x13\x5c\x7d\x6b\xb4\x79\x82\x1e\xc3\x5f\x5f\x6d\x9c\xee\x1c\x31\x33\x08\x03\x8b\x33\x9a\xb4\x65\x4c\x53\x42\xfe\x12\xb1\x31\xe8\x5f\xe7\x04\xce\x22\xc7\x18\xf3\xf8\x10\x00\x99\x9c\xb4\x70\xc8\x53\xe3\x1a\x0c\x2d\x42\xc8\x9c\xaf\x3d\x9b\x03\x96\x97\xf7\xe8\x96\xde\x8b\x94\xde\x2a\x93\x6d\xff\xca\xe5\xa0\xeb\xcd\x0d\xc2\xc2\x26\x7d\xc2\x98\xc8\x53\xdd\x19\x4c\x6b\x5b\x09\x3a\x5b\x76\x87\xb4\x94\x26\xd3\xe7\x85\x74\x9f\x88\x52\x99\x12\xf0\x95\x98\xc8\x92\xb7\x0e\xed\x8b\x14\x04\xe3\xd3\xa3\x03\x3c\x39\x58\x1b\xce\x21\x3d\x64\x44\x3e\x3c\x23\xb3\x83\x80\xbe\xba\xed\xad\x8d\xe7\x45\x77\x91\x46\x87\x0e\xbc\x2b\xe1\x96\x5e\xe4\x57\x35\x21\x4d\xe3\x04\xb8\xcb\x6d\xed\x53\x9e\x01\x4b\xaa\xaf\x7e\x7e\x9f\x24\x88\x1a\x7b\x39\xdc\x87\x28\x80\x17\x75\x9c\x52\x63\x46\x67\x99\x7d\xe6\x7a\xa2\x75\xa8\xd0\x82\x7c\x62\x4f\x8c\xb0\x3d\x5b\xa6\x70\x7f\x46\x35\x7c\x5a\xca\x72\xd9\xa9\x4a\x5e\xc9\x4e\xfb\x4c\x9a\xbe\x84\x99\x72\x27\x81\xe8\xf7\x0e\xa9\xda\xcf\xc6\x44\xe9\x88\x29\x3e\x11\x4b\x9f\x45\xf1\xe2\xf8\x8f\x9c\x6d\x78\x80\x10\xfa\xda\xf8\x60\x9f\x3e\x66\x02\x98\xfd\xe7\xab\x36\xf0\x4b\xa9\xef\xf8\x6f\x7e\xf7\x38\x72\x46\xca\xce\xd8\xe8\x3e\xce\xc2\x79\x71\x47\xb2\x13\xaf\xf6\xc8\x77\xaf\xf6\x4c\x2c\xc7\xc5\x85\xa6\x6c\x63\x41\x41\x31\x7f\xd2\x74\x99\xd8\xce\xc3\x71\xba\xe4\x3f\xa3\xeb\xf2\x8b\x06\x66\x2e\x92\x96\x8a\x9e\xa6\x28\xf0\x6e\x36\xfc\x66\xb8\x0c\x1a\x33\x07\xae\x8b\x8b\x12\x4f\xa1\x84\xfe\x6b\xce\x3b\x7c\x01\x27\x22\xd7\x76\xbd\xfb\xd8\xa6\x98\x20\x27\xb9\x13\xbb\x04\xfb\xfe\x74\x96\xf1\x44\xa9\x7c\xcb\x0d\x7a\x26\x92\x98\xdd\x9d\x0a\x06\xef\xad\x99\x06\xdb\x0f\xec\xa0\xe5\x81\x1b\x0a\xbf\x00\x1f\x97\x92\x40\x19\x2b\xad\xc8\xa3\xca\x16\x47\xb7\xf8\xe5\xd3\x6a\xc5\x99\x8d\xc7\x24\x49\xc4\x6d\x9b\x80\xba\xb7\x91\xc2\xec\x72\xa3\x0e\x19\x83\xaa\x12\x8e\x7a\x57\x02\xbf\x33\xfc\xce\x69\xbe\x39\x0e\x60\x77\x59\x3f\x3a\x80\x7d\x7d\xfc\x6b\xb8\x44\x81\x16\xaf\x0d\x6e\xd4\xac\x06\x9b\x91\x22\x63\x4d\xeb\x3d\x6a\xe9\xcd\x5b\x11\x84\xce\xdc\xb5\xc5\x95\x57\x7a\x72\x68\xcf\x59\x22\xf2\xe5\x2d\xd5\xec\x9a\xcc\x72\xfd\x85\xc3\x6e\xcc\x4e\xa9\xa6\x01\x82\xb2\xf8\xf0\x7a\x1a\xc0\xb6\x4a\xc0\x5e\x7f\xaa\x0c\x2c\xa8\x5a\xc5\x0d\xac\x35\xfe\x83\xcb\xc4\x8b\x5c\xb7\x9b\xd4\x89\xe4\x10\xd3\x6a\x26\x75\x7a\xed\x40\xe1\x07\x04\xa7\xdb\x5e\x2c\xc4\x12\x80\x9f\x6e\x20\x57\xea\x0d\xa3\xe3\x2c\x85\x1f\xcc\xbf\xb7\x0b\x97\xad\x94\x95\xfd\xf1\x0a\x38\x52\x4c\x83\x2b\x6e\x29\xfb\x0d\x7d\x0c\xd9\xda\x3f\x87\x6b\x4c\x47\x86\x42\x63\xb9\xbe\x53\x45\xd4\xf3\xb6\x53\x54\x38\x98\x0d\x47\x9e\x9e\xdf\xd5\x13\x48\x54\x2d\xe7\x55\x0c\xe8\x59\xd0\xf7\xb2\x41\x99\x54\x8f\x5e\xc3\x92\xc0\x9c\x7f\x8e\x25\xa3\x40\x99\x17\x8f\xcf\xf0\xdd\xdf\xfa\x5b\x53\x60\x8f\x6a\xcd\xa8\x27\xe6\x40\xc0\xec\x92\xd4\xab\x5d\xb8\x18\xf5\x27\xa6\xb6\xa6\x55\x6d\xd9\xb7\x1f\x86\x81\x71\xf0\xb7\xcc\x65\xff\xf6\x70\x36\xf9\x02\xa5\x55\xcf\x60\xfb\x07\x8d\xc1\x6e\x50\x2f\xc4\xff\x07\x00\x00\xff\xff\xf4\xf5\x28\xc2\x96\x12\x00\x00")

func templatesServiceSyslogTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesServiceSyslogTmpl,
		"templates/service/syslog.tmpl",
	)
}

func templatesServiceSyslogTmpl() (*asset, error) {
	bytes, err := templatesServiceSyslogTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/service/syslog.tmpl", size: 4758, mode: os.FileMode(420), modTime: time.Unix(1462490475, 0)}
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
	"templates/service/fluentd.tmpl": templatesServiceFluentdTmpl,
	"templates/service/syslog.tmpl": templatesServiceSyslogTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"service": &bintree{nil, map[string]*bintree{
			"fluentd.tmpl": &bintree{templatesServiceFluentdTmpl, map[string]*bintree{}},
			"syslog.tmpl": &bintree{templatesServiceSyslogTmpl, map[string]*bintree{}},
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

