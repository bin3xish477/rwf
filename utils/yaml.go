package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ParseArgs(f string, a interface{}) {
	yf, err := ioutil.ReadFile(f)
	Must(err)

	err = yaml.Unmarshal(yf, a)
	Must(err)
}
