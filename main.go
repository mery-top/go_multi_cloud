package main

import (
	"fmt"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Name     string `yaml:"name"`
		Provider string `yaml:"provider"`
		Compute  struct {
			Type   string `yaml:"type"`
			CPU    string `yaml:"cpu"`
			Memory string `yaml:"memory"`
			Image  string `yaml:"image"`
		} `yaml:"compute"`
		Storage struct {
			Type       string `yaml:"type"`
			BucketName string `yaml:"bucket_name"`
		} `yaml:"storage"`
	} `yaml:"app"`
}

const terraformTemplate = `
provider "aws" {
  region = "us-east-1"
}

resource "aws_s3_bucket" "app_bucket" {
  bucket = "{{.App.Storage.BucketName}}"
}

# Placeholder ECS Task
resource "aws_ecs_cluster" "main" {
  name = "{{.App.Name}}-cluster"
}
`

func main() {
	// Step 1: Read deploy.yaml
	data, err := os.ReadFile("deploy.yaml")
	if err != nil {
		panic(err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	fmt.Println("✅ Parsed YAML config")

	// Step 2: Generate main.tf from template
	tmpl, err := template.New("tf").Parse(terraformTemplate)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("main.tf")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, config); err != nil {
		panic(err)
	}

	fmt.Println("✅ main.tf generated successfully!")
}
