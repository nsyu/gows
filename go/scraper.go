package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://show.1688.com/pinlei/industry/pllist.html?spm=a260j.12536015.jr60bfo3.1.10f3700ecaYFQt&&sceneSetId=857&sceneId=2158&bizId=4404"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
