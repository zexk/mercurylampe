package cat

import (
	"io"
	"log"
	"net/http"
	"os"
)

const url = "https://cataas.com/cat"

func GetCat() []byte {
	res, err := http.Get(url)

	if err != nil {
		log.Fatalf("%v", err)
	}

	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func WriteCat(cat []byte) *os.File {
	f, err := os.CreateTemp("", "tmp-")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	defer os.Remove(f.Name())

	if size, err := f.Write(cat); err != nil {
		log.Fatal(err)

	} else {
		log.Println(size)
	}

	return f
}
