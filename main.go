package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/urfave/cli/v2"
)

var (
	version = "dev"
)

func main() {
	app := &cli.App{
		Name:  "modifylt",
		Usage: "modify launch template",
		Action: func(c *cli.Context) error {
			return modify(c.String("launch-template-id"), c.String("default-version"))
		},
	}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "launch-template-id",
			Usage:    "Launch template ID",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "default-version",
			Usage:    "Version to set as default",
			Required: true,
		},
	}

	fmt.Fprintln(os.Stderr, "modifylt", version)
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
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
