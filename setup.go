package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/lunny/html2md"
	"github.com/skratchdot/open-golang/open"
	cli "gopkg.in/urfave/cli.v2"

	"golang.org/x/net/html"
)

func main() {
	app := &cli.App{
		Usage:   "Bootstraps a daily folder for advent of code 2018",
		Version: "v1.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "auth",
				Aliases: []string{"a"},
				Value:   os.Getenv("AOC_SESSION"),
				Usage:   "<REQURIED> Auth key to use.",
			},
			&cli.IntFlag{
				Name:    "day",
				Aliases: []string{"d"},
				Usage:   "<REQURIED> Day to bootstrap challange for.",
			},
			&cli.BoolFlag{
				Name:    "replace",
				Aliases: []string{"r"},
				Value:   false,
				Usage:   "Replaces the items",
			},
			&cli.PathFlag{
				Name:    "path",
				Aliases: []string{"p"},
				Usage:   "Path to place files, current path if not provided",
			},
			&cli.BoolFlag{
				Name:    "template",
				Aliases: []string{"t"},
				Value:   false,
				Usage:   "Creates a main template file",
			},
			&cli.BoolFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Value:   false,
				Usage:   "Donwloads the input file",
			},
			&cli.BoolFlag{
				Name:    "markdown",
				Aliases: []string{"md"},
				Value:   false,
				Usage:   "Downloads the challange and saves it as markdown",
			},
			&cli.BoolFlag{
				Name:    "open",
				Aliases: []string{"o"},
				Value:   false,
				Usage:   "Opens current days challange in browser or calander.",
			},
		},
		EnableShellCompletion: true,
		CommandNotFound: func(c *cli.Context, command string) {
			fmt.Fprintf(c.App.Writer, "Thar be no %q here.\n", command)
		},
		Action: func(c *cli.Context) error {
			if c.IsSet("help") {
				cli.ShowAppHelpAndExit(c, 0)
			}

			if !c.IsSet("day") {
				return cli.Exit("\tFlag: --day is requried", 0)
			}

			auth := c.String("auth")
			if auth == "" {
				return cli.Exit("\tFlag: --auth is requried or 'AOC_SESSION' environment variable", 0)
			}

			day := strconv.Itoa(c.Int("day"))

			var path string
			if c.IsSet("path") {
				path = c.Path("path")
			} else if p, err := os.Getwd(); err == nil {
				path = p
			} else {
				return cli.Exit("Could not find working dir", 0)
			}

			if !strings.HasSuffix(path, string(os.PathSeparator)) {
				path = path + string(os.PathSeparator)
			}

			os.Mkdir(path, 0755)

			if c.Bool("input") {
				target := path + "input.txt"
				if _, err := os.Stat(target); os.IsNotExist(err) || c.Bool("replace") {
					if err := downloadInputFile(target, day, auth); err != nil {
						fmt.Println("Could not download input file: " + err.Error())
					}
				}
			}

			if c.Bool("markdown") {
				target := path + "README.md"
				if _, err := os.Stat(target); os.IsNotExist(err) || c.Bool("replace") {
					if err := downloadQuestionFile(target, day, auth); err != nil {
						fmt.Println("Could not download input file: " + err.Error())
					}
				}
			}

			if c.Bool("template") {
				target := path + "main.go"
				if _, err := os.Stat(target); os.IsNotExist(err) || c.Bool("replace") {
					if err := createMainTemplate(target); err != nil {
						fmt.Println("Could not download input file: " + err.Error())
					}
				}
			}

			if c.Bool("open") {
				open.Run("https://adventofcode.com/2018/day/" + day)
			}

			return nil
		},
	}
	app.Run(os.Args)
}

// DownloadInputFile will download the input for the day
func downloadInputFile(target, day string, auth string) error {

	url := "https://adventofcode.com/2018/day/" + day + "/input"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: auth})

	// Get the data
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(url +
			"\nresp.StatusCode: " + strconv.Itoa(resp.StatusCode))
	}

	// Create the file
	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func downloadQuestionFile(target, day string, auth string) error {

	url := "https://adventofcode.com/2018/day/" + day
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: auth})

	// Get the data
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(url +
			"\nresp.StatusCode: " + strconv.Itoa(resp.StatusCode))
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	articles, err := getArticle(root)
	if err != nil {
		return err
	}
	md := html2md.Convert(renderNodes(articles))

	// Create the file
	err = ioutil.WriteFile(target, []byte(md), 0644)
	if err != nil {
		return err
	}
	return nil

}

func getArticle(doc *html.Node) ([]*html.Node, error) {
	var b []*html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "article" {
			b = append(b, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	if b != nil {
		return b, nil
	}
	return nil, errors.New("Missing <article> in the node tree")

}

func renderNodes(nodes []*html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	for _, n := range nodes {
		html.Render(w, n)
	}
	return buf.String()
}

func createMainTemplate(target string) error {
	err := ioutil.WriteFile(target, []byte(template), 0644)
	if err != nil {
		return err
	}
	return nil
}

var template = `package main

import (
	"fmt"
	"os"

	"github.com/urbansson/advent-of-code/util"
)

func main() {
	args := os.Args[1:]
	file := args[0]
	fmt.Println("Using input:", file)
	fc := util.ReadFile(file)

	for _, f := range fc {
		fmt.Println(f)
	}
}`
