package utils

type LArg map[string]interface{}
type SArg map[string]interface{}
type Env map[string]string

type ToolArgs struct {
	Tool    string `yaml:"tool,omitempty"`
	Long    []LArg `yaml:"long,omitempty""`
	Short   []SArg `yaml:"short,omitempty""`
	EnvVars []Env  `yaml:"env,omitempty""`
}
