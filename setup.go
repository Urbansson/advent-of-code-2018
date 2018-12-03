package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/skratchdot/open-golang/open"
)

func main() {
	args := os.Args[1:]

	if len(os.Args) < 2 {
		fmt.Println("Missing argument for day")
		os.Exit(1)
	}
	day := args[0]

	auth, ok := os.LookupEnv("AOC_SESSION")
	if !ok {
		fmt.Println("Needs a session key.\nuse: export 'AOC_SESSION=<session_key>'")
		os.Exit(1)
	}

	if _, err := os.Stat("./day" + day); !os.IsNotExist(err) {
		fmt.Println("Folder for day" + day + " aldready exists.")
		os.Exit(1)
	}
	os.Mkdir("./day"+day, 0755)

	if err := downloadInputFile(day, auth); err != nil {
		os.Remove("./day" + day)
		fmt.Println(err)
		os.Exit(1)
	}

	if err := createMainTemplate(day); err != nil {
		os.Remove("./day" + day)
		fmt.Println(err)
		os.Exit(1)
	}

	open.Run("https://adventofcode.com/2018/day/" + day)

}

// DownloadInputFile will download the input for the day
func downloadInputFile(day string, auth string) error {

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
	out, err := os.Create("./day" + day + "/input.txt")
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

func createMainTemplate(day string) error {
	destFile := "./day" + day + "/main.go"
	err := ioutil.WriteFile(destFile, []byte(template), 0644)
	if err != nil {
		fmt.Println("Error creating", destFile)
		fmt.Println(err)
		return err
	}
	return nil
}

var template = `
package main
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
}
`
