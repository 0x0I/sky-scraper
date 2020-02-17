package main

import (
  "fmt"
  "log"
  "os"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/ec2"
  "github.com/urfave/cli/v2"
)

func main() {
  var provider string
  var region string
  var resource string

  cli.VersionFlag = &cli.BoolFlag{
    Name: "print-version", Aliases: []string{"V"},
    Usage: "print only the version",
  }

  app := &cli.App{
    Name: "sky-scraper",
    Version: "v0.1.0",
    Usage: "scrape cloud infrastructure resource data from major cloud providers for developing provisioning insights",
    Flags: []cli.Flag {
      &cli.StringFlag{
        Name: "provider",
        Aliases: []string{"p"},
        Value: "aws",
        Usage: "cloud provider to target for infrastructure resource data",
        Destination: &provider,
      },
      &cli.StringFlag{
        Name: "region",
        Aliases: []string{"R"},
        Value: "us-east-1",
        Usage: "cloud provider region to focus scrape",
        Destination: &region,
      },
      &cli.StringFlag{
        Name: "resource",
        Aliases: []string{"r"},
        Value: "cpu",
        Usage: "cloud provider resource to scrape",
        Destination: &resource,
      },
    },
    Action: func(c *cli.Context) error {
      fmt.Println("The sky's the limit. Let's scrape It!")

      instance := "t2.micro"
      if c.NArg() > 0 {
        instance = c.Args().Get(0)
      }
      
      fmt.Println("Cloud Provider:", provider)
      fmt.Println("Cloud Region:", region)
      fmt.Println("Cloud Resource:", resource)
      fmt.Println("Resource Instance(s):", instance)

      switch provider {
	  case "aws":
		fmt.Println("Scraping AWS stats for:", resource)
        ec2svc := ec2.New(session.New(&aws.Config{Region: aws.String(region),}))
        params := &ec2.DescribeInstanceTypesInput{
          InstanceTypes: []*string{aws.String(instance)},
        }
        result, err := ec2svc.DescribeInstanceTypes(params)
        if err != nil {
          fmt.Println("There was an error listing instances!", err.Error())
          log.Fatal(err.Error())
        }
        fmt.Println(result)
	  default:
		fmt.Printf("Scraping unsupported for: %s.\n", provider)
	  }
      return nil
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
