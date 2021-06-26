package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"flag"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var (
	version = "dev"
)

func main() {
	fmt.Fprintln(os.Stderr, "modifylt", version)

	id := flag.String("launch-template-id", "", "Launch template Id")
	defaultVersion := flag.String("default-version", "", "Version to set as default")

	flag.Parse()

	if *id == "" || *defaultVersion == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := modify(*id, *defaultVersion); err != nil {
		die(err)
	}
}

func die(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}

func modify(id string, defaultVersion string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	client := ec2.NewFromConfig(cfg)

	output, err := client.ModifyLaunchTemplate(context.TODO(), &ec2.ModifyLaunchTemplateInput{
		LaunchTemplateId: &id,
		DefaultVersion:   &defaultVersion,
	})
	if err != nil {
		return err
	}

	json, err := json.MarshalIndent(output.LaunchTemplate, "", "\t")
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", json)

	return nil
}
