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

//func WriteArgs(f string) {
//d, err := yaml.Marshal(yf, a)
//Must(err)

//err = ioutil.WriteFile(f, d, 0755)
//Must(err)
//}
