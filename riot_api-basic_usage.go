package APIriot

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

type API struct {
	Url string
	Key string
	Region string
	Version string
}

func (api *API) getSummoner(username string) (string) {

	request, err := http.Get( fmt.Sprintf(api.Url, api.Region, api.Region, api.Version, username, api.Key) )
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}

	request.Body.Close()
	text := string(content[:])

	return text
}

func main() {

	api := API{
		Url: "https://%s.api.pvp.net/api/lol/%s/%s/summoner/by-name/%s?api_key=%s",
		Key: "<coloque aqui a sua API KEY>",
		Region: "<REGIAO-DO-SERVER>",
		Version: "<VERSÂO-DA-API>",
	}

	fmt.Printf("Response: %s \n", api.getSummoner("<SUMMONER-NAME-HERE>") )

}
