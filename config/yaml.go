package config


var DefaultConfig =`
cpu:
    cpuRate: 10

`

type cpu struct {
	CpuRate int64
}
type YamlConfig struct {
	Cpu struct {
		    CpuRate int64 `yaml:"cpuRate"`
	    }`yaml:"cpu"`
}