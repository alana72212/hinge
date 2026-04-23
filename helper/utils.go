package helper

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func UUID() string {
	id := uuid.New()
	return strings.ToUpper(id.String())
}

func Request(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
