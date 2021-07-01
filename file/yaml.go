package file

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"

	"github.com/h0n9/cert-inspector/types"
)

type HostFile struct {
	Hosts    []types.Host `json:"hosts" yaml:"hosts"`
	filename string
}

func NewHostFile(filename string) *HostFile {
	return &HostFile{
		Hosts:    make([]types.Host, 0),
		filename: filename,
	}
}

func (hf *HostFile) Read() error {
	data, err := ioutil.ReadFile(hf.filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, hf)
	if err != nil {
		return err
	}
	return nil
}

func (hf *HostFile) Save() error {
	data, err := yaml.Marshal(hf)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(hf.filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
