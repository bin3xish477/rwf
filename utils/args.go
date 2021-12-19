package utils

type LArg map[string]interface{}
type SArg map[string]interface{}
type Env map[string]string

type ToolArgs struct {
	Tool    string `yaml:"tool"`
	Long    []LArg `yaml:"long"`
	Short   []SArg `yaml:"short"`
	EnvVars []Env  `yaml:"env"`
}
