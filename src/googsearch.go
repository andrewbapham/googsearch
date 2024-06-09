package main

import (
	"fmt"
	"log"
//	"net/http"
	"os"
//	"strings"
//	"time"

	"github.com/urfave/cli/v2"
//	"golang.org/x/net/html"
)

var websites string
var show_result_count int

func main() {
	app := &cli.App{
    Name: "googsearch",
    Usage: "Search for results on Google",
    Flags: []cli.Flag{
      &cli.StringFlag{
        Name: "websites",
        Aliases: []string{"w"},
        Usage: "List of websites to include",
        Destination: &websites,
      },
      &cli.IntFlag{
        Name: "number",
        Aliases: []string{"n"},
        Usage: "Number of results to display",
        Value: 5,
        Destination: &show_result_count,
      },
    },
    Action: main_func,
  }

  if err := app.Run(os.Args); err != nil {
    log.Fatal(err)
  }
}

func main_func(*cli.Context) error{
  fmt.Println("Flag website: ", websites)
  fmt.Println("Flag number:", show_result_count)
  return nil
}
