package cat

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

const baseUrl = "https://cataas.com/cat"

func GetCat() io.Reader {
	res, err := http.Get(baseUrl)

	if err != nil {
		log.Fatalf("%v", err)
	}

	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	return bytes.NewReader(data)
}

func GetGifCat() io.Reader {
	res, err := http.Get(baseUrl + "/gif")

	if err != nil {
		log.Fatalf("%v", err)
	}

	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	return bytes.NewReader(data)
}

func GetTextCat(text string) io.Reader {
	res, err := http.Get(baseUrl + "/says/" + text)

	if err != nil {
		log.Fatalf("%v", err)
	}

	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	return bytes.NewReader(data)
}
