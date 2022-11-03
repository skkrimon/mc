package util

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type ConfigYaml struct {
	Port       string   `yaml:"port"`
	UpdatePath string   `yaml:"update_path"`
	GinMode    string   `yaml:"gin_mode"`
	Servers    []Server `yaml:"servers"`
	Users      []User
}

type User struct {
	Username string `yaml:"username"`
	ApiKey   string `yaml:"api_key"`
}

type Server struct {
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
}

func (c *ConfigYaml) AddUser(user User) {
	c.Users = append(c.Users, user)
}

func (c *ConfigYaml) AddServer(server Server) {
	c.Servers = append(c.Servers, server)
}

func (c *ConfigYaml) LoadConfig() error {
	file, err := os.ReadFile("config.yml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		return err
	}

	return nil
}

func (c *ConfigYaml) WriteConfig() error {
	data, _ := yaml.Marshal(c)

	file, _ := os.OpenFile("config.yml", os.O_RDWR|os.O_CREATE, 0777)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file")
		}
	}(file)

	_, err := file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
