package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	login, repo_no, err := githubInfo("parjinderpannu")
	if err != nil {
		log.Fatal("error: githubInfo - %s", err)
	}

	fmt.Printf("Login: %s \t repo_no: %d", login, repo_no)
	// resp, err := http.Get("https://api.github.com/users/parjinderpannu")
	// if err != nil {
	// 	log.Fatal("error: %s", err)
	// 	/*
	// 		log.Printf("error: %s", err)
	// 		os.Exit(1)
	// 	*/
	// }
	// if resp.StatusCode != http.StatusOK {
	// 	log.Fatalf("error:%s", resp.Status)
	// }
	// fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	// /*
	// 	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
	// 		log.Fatal("error: can't copy - %s", err)
	// 	}
	// */

	// var r Reply
	// dec := json.NewDecoder(resp.Body)
	// if err := dec.Decode(&r); err != nil {
	// 	log.Printf("error: can't decode - %s", err)
	// }
	// fmt.Printf("%#v\n", r)
}

func githubInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + login
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
		// log.Fatal("error: %s", err)
		/*
			log.Printf("error: %s", err)
			os.Exit(1)
		*/
	}
	var r Reply
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		log.Printf("error: can't decode - %s", err)
		return "", 0, err
	}

	return r.Login, r.Public_Repos, nil
}

type Reply struct {
	Login        string
	Public_Repos int
}

/*
JSON <--> GO [serialization or unmarshalling]
true/false <-->
string <-->
null <--> nil
number <--> float64, float32, int8..16..32..64, int, uint8, ...
array <--> []any ([]interface{}) slice of some type
object <--> map[string]any, struct

encoding/json API
JSON --> io.Reader --> Go: json.Decoder
JSON --> []byte --> Go: json.Unmarshal
Go --> io.Writer --> JSON: json.Encoder
Go --> []byte --> JSON: json:Marshal
*/
