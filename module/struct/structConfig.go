package _struct

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Path struct {
			PathServer     string `yaml:"pathServer"`
			PathOutputFile string `yaml:"pathOutputFile"`
		} `yaml:"path"`
		Log struct {
			Path     string `yaml:"path"`
			FileName string `yaml:"fileName"`
		} `yaml:"log"`
	} `yaml:"server"`
	Link struct {
		Home struct {
			UrlHome      string `yaml:"urlHome"`
			PathHomeHtml string `yaml:"pathHomeHtml"`
		} `yaml:"home"`
		Print struct {
			UrlPrint     string `yaml:"urlPrint"`
			PathHomeHtml string `yaml:"pathPrintHtml"`
		} `yaml:"print"`
		Output struct {
			UrlOutput      string `yaml:"urlOutput"`
			PathOutputHtml string `yaml:"pathOutputHtml"`
		} `yaml:"output"`
		Reverse struct {
			UrlReverse      string `yaml:"urlReverse"`
			PathReverseHtml string `yaml:"pathReverseHtml"`
		} `yaml:"reverse"`
		UploadsFiles struct {
			UrlUploadsFiles  string `yaml:"urlUploadsFiles"`
			PathUploadsFiles string `yaml:"pathUploadsFiles"`
		} `yaml:"uploadsFiles"`
		DownloadsFiles struct {
			UrlDownloadsFiles  string `yaml:"urlDownloadsFiles"`
			PathDownloadsFiles string `yaml:"pathDownloadsFiles"`
		} `yaml:"downloadsFiles"`
		LoadsAssets struct {
			UrlLoadsAssets  string `yaml:"urlLoadsAssets"`
			PathLoadsAssets string `yaml:"pathLoadsAssets"`
		} `yaml:"loadsAssets"`
	} `yaml:"link"`
}

func newConfig(config *Config) (*Config, error) {

	file, err := os.Open("config.yml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}

func (config *Config) GetConfig() *Config {
	conf, err := newConfig(config)
	if err != nil {
		fmt.Println(err)
	}
	return conf
}
