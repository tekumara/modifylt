package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "modifylt",
		Usage:  "modify launch template",
		Action: modify,
	}

	app.Flags = []cli.Flag {
		&cli.StringFlag{
		  Name: "launch-template-id",
		  Usage: "Launch template ID",
		  Required: true,
		},
		&cli.StringFlag{
			Name: "default-version",
			Usage: "Version to set as default",
			Required: true,
		},
	  }

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func modify(c *cli.Context) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	client := ec2.NewFromConfig(cfg)

	_, err = client.ModifyLaunchTemplate(context.TODO(), &ec2.ModifyLaunchTemplateInput{
		LaunchTemplateId: aws.String(c.String("launch-template-id")),
		DefaultVersion:   aws.String(c.String("default-version")),
	})
	if err != nil {
		return err
	}

	return nil
}
