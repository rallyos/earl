package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/shiftingphotons/earl/app"
)

/* SendShortenRequest sends a URL to be shortened through the REST API.
Main use is by calling the `shorten` subcommand using the executable. Host by default will be https://earl.sh but it can be
also passed as a string argument. Of course the url is coming through the url argument that is given to the executable for the shorten subcommand.

Keeping this here makes things a bit polluted, but it's small enough to not need its own package, and it doesn't fit the others, so...it stays here for now.
*/
func SendShortenRequest(url, host string) (string, error) {
	dest := fmt.Sprintf("%s/api/shorten", host)
	reqBody := fmt.Sprintf("{ \"url\": \"%s\" }", url)
	body := bytes.NewBufferString(reqBody)

	resp, err := http.Post(dest, "application/json", body)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	defer resp.Body.Close()

	shortURL := struct {
		Url string
	}{}

	err = json.NewDecoder(resp.Body).Decode(&shortURL)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	return shortURL.Url, nil
}

func main() {
	shortenCmdLifecycle := flag.String("lifecycle", "", "lifecycle")

	shortenCmd := flag.NewFlagSet("shorten", flag.ExitOnError)
	shortenCmdURL := shortenCmd.String("url", "", "url")
	shortenCmdHost := shortenCmd.String("host", "https://earl.sh", "host")

	if len(os.Args) > 1 && os.Args[1] == "shorten" {
		shortenCmd.Parse(os.Args[2:])
		shortenedURL, err := SendShortenRequest(*shortenCmdURL, *shortenCmdHost)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Printf("Your shortened URL is %s\n", shortenedURL)
	} else {
		flag.Parse()

		// Flag will have higher precedence if environment variable is set, so we override.
		if len(*shortenCmdLifecycle) > 0 {
			os.Setenv("EARL_LIFECYCLE", *shortenCmdLifecycle)
		}

		server := new(app.Server)
		fmt.Println("Starting server...")
		server.Initialize()
	}

}
