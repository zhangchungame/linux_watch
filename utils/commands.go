package utils

import (
	"github.com/urfave/cli"
	"linux_watch/cpu"
	"fmt"
	"strings"
	"io/ioutil"
	"linux_watch/config"
	"gopkg.in/yaml.v2"
)

var (
	Yamlfile string
)

var (
	Cpucommand=cli.Command{
		Action:    WatchCpu,
		Name:      "cpu",
		Usage:     "watch cpu usage",
		ArgsUsage: "<genesisPath>",
		Category:  "system",
		Description: `watch cpu usage Description.`,
		Flags:[]cli.Flag{
			cli.StringFlag{Name: "cpu_rate", Value: "10", Usage: "cpu报警上线"},
		},
	}
)

func loadConfig()  *config.YamlConfig{
	var err error
	yamlConfig:=config.YamlConfig{}
	configByte:=[]byte(config.DefaultConfig)
	if !strings.EqualFold(Yamlfile,""){
		configByte,err=ioutil.ReadFile(Yamlfile)
		if err!=nil{
			panic(err)
		}
	}
	err=yaml.Unmarshal(configByte,&yamlConfig)
	if err!=nil{
		panic(err)
	}
	return &yamlConfig
}
func WatchCpu(cxt *cli.Context)  {
	yamlConfig:=loadConfig()
	watchCpu(yamlConfig)
}
func watchCpu(yamlConfig *config.YamlConfig)  {
	cpuwatch:=&cpu.CpuWatch{}
	cpuwatch.Refresh()
	columns:=make(map[string]string)
	columns["bogomips"]=""
	columns["model"]=""
	columns["model name"]=""

	aaa,err:=cpuwatch.ParamStateBatch(columns)
	fmt.Println(columns)
	fmt.Println(aaa,err)
}
func AllWatch(cxt *cli.Context)  {
	yamlConfig:=loadConfig()
	if yamlConfig.Cpu.CpuRate!=int64(0){
		watchCpu(yamlConfig)
	}
}
