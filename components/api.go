package components

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Api() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	url := "https://quotes-by-api-ninjas.p.rapidapi.com/v1/quotes"

	req, _ := http.NewRequest("GET", url, nil)

	apiKey := os.Getenv("API_KEY")

	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", "quotes-by-api-ninjas.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	text = string(body)
	fmt.Println(text)

}
