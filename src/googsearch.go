package main

import (
	"fmt"
	"log"
//	"net/http"
  //"net/url"
	"os"
  "strings"
  "regexp"
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
  website_filter_list := get_website_list_from_string(websites)
  fmt.Println("Filter list: ", website_filter_list)
  return nil
}

func get_website_list_from_string(str string) []string{
  entries := strings.Split(strings.ToLower(str), " ")

  var alias_map = map[string]string{
    "so": "stackoverflow.com",
    "gh": "github.com",
  }

  for i, w := range entries{
    url, ok := alias_map[w]
    if ok {
      entries[i] = url
    }else if !is_valid_domain(w) {
      log.Fatalf("Input filter %s is not a valid alias or is an invalid domain", w)
    }
  }

  return entries
}

func is_valid_domain(str string) bool{
  pattern, _ := regexp.Compile(`^([A-Za-z0-9-]{1, 63}\\.)+[A-Za-z]{2, 6}$`)
  starts_with_dash, _ := regexp.Compile(`-.*`)
  domain_ends_with_dash, _ := regexp.Compile(`^.*-\.`)
  return pattern.MatchString(str) && !starts_with_dash.MatchString(str) && ! domain_ends_with_dash.MatchString(str)
}
