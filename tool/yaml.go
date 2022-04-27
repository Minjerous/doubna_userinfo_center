package tool

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Role struct{
	//id决定了Role的优先级，从0到3，优先级增高
	//当角色拥有多个优先级，采用最高优先级
	Id         int     `yaml:"id"`
	Resource   Resource `yaml:",inline"`
	Permission Permission `yaml:",inline"`
}
type  Resource struct{
	Host string `json:"host"yaml:"host"`
	//资源的path
	Path string`json:"path"yaml:"path"`
	//资源的nethod
	Method string `json:"method"yaml:"method"`
}
//资源的host

type Permission struct{
	//允许访问资源的角色
	AuthorizedRoles []string `json:"authorized_roles"yaml:"authorized_roles"`
	//不允许访问资源的角色
	ForbiddenRoles []string `json:"forbidden_roles"yaml:"forbidden_roles"`
	//是否允许所有人访问，在Permission中优先级最高
	ALlowAnyone   bool ` json:"allow_anyone"yaml:"allow_anyone"`
}



func PareYaml() error {
	conf := new(Role)
	yamlFile, err := ioutil.ReadFile("./config/app.yaml")

	// conf := new(module.Yaml1)
	// yamlFile, err := ioutil.ReadFile("test.yaml")

	// conf := new(module.Yaml2)
	//  yamlFile, err := ioutil.ReadFile("test1.yaml")

	log.Println("yamlFile:", yamlFile)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	// err = yaml.Unmarshal(yamlFile, &resultMap)
	log.Println("conf", conf)
	return err

}
