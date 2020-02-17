package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli/v2"
)

func main() {
  var provider string
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
        Name: "resource",
        Aliases: []string{"r"},
        Value: "cpu",
        Usage: "cloud provider resource to scrape",
        Destination: &resource,
      },
    },
    Action: func(c *cli.Context) error {
      fmt.Println("The sky's the limit. Let's scrape It!")

      instance := "all"
      if c.NArg() > 0 {
        instance = c.Args().Get(0)
      }
      
      fmt.Println("Cloud Provider:", provider)
      fmt.Println("Cloud Resource:", resource)
      fmt.Println("Resource Instance(s):", instance)

      return nil
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
