package prisma

import (
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"os"
)

func Export(apiKey string) error {

	url := "https://www.bridgecrew.cloud/api/v1/suppressions"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", apiKey)

	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(res.Body)
	body, _ := io.ReadAll(res.Body)
	outfile := "suppressions.json"
	err := os.WriteFile(outfile, body, 0644)
	log.Printf("File %s written", outfile)

	if err != nil {
		log.Fatal().Err(err)
	}
	return nil
}
