package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/parjinderpannu")
	if err != nil {
		log.Fatal("error: %s", err)
		/*
			log.Printf("error: %s", err)
			os.Exit(1)
		*/
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error:%s", resp.Status)
	}
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatal("error: can't copy - %s", err)
	}
}

/*
JSON <--> GO
true/false <-->
string <-->
null <--> nil
number <--> float64, float32, int8..16..32..64, int, uint8, ...
array <--> []any ([]interface{})
object <--> map[string]any, struct
*/
