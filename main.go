package main

import (
    "context"
    //"fmt"
    "log"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/ec2"
)

var foo = "foo"

func main() {
    cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
    if err != nil {
        log.Fatalf("unable to load SDK config, %v", err)
    }

	client := ec2.NewFromConfig(cfg)

	_, err = client.ModifyLaunchTemplate(context.TODO(), &ec2.ModifyLaunchTemplateInput{
		LaunchTemplateId: aws.String("foo"),
		DefaultVersion: &foo,
	})
	if err != nil {
        log.Fatalf("ModifyLaunchTemplate error: %v", err)
    }

    // // Using the Config value, create the DynamoDB client
    // svc := dynamodb.NewFromConfig(cfg)

    // // Build the request with its input parameters
    // resp, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{
    //     Limit: aws.Int32(5),
    // })
    // if err != nil {
    //     log.Fatalf("failed to list tables, %v", err)
    // }

    // fmt.Println("Tables:")
    // for _, tableName := range resp.TableNames {
    //     fmt.Println(tableName)
    // }
}
