package dict

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	Head struct {
	} `json:"head"`
	Def []struct {
		Text string `json:"text"`
		Pos  string `json:"pos"`
		Tr   []struct {
			Text string `json:"text"`
			Pos  string `json:"pos"`
			Syn  []struct {
				Text string `json:"text"`
			} `json:"syn"`
			Mean []struct {
				Text string `json:"text"`
			} `json:"mean"`
			Ex []struct {
				Text string `json:"text"`
				Tr   []struct {
					Text string `json:"text"`
				} `json:"tr"`
			} `json:"ex"`
		} `json:"tr"`
	} `json:"def"`
}

func Lookup(text []string) (string, error) {
	var msg string
	respBody := fmt.Sprintf("https://dictionary.yandex.net/api/v1/dicservice.json/lookup?key=%s&lang=ru-ru&text=%s", os.Getenv("DICT_KEY"), text[0])
	resp, err := http.Get(respBody)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response Response
	json.Unmarshal(body, &response)
	if len(response.Def) == 0 {
		return "", nil
	}
	for _, def := range response.Def {
		msg += fmt.Sprintf("%s (%s), значения:\n", strings.ToTitle(def.Text), def.Pos)
		for _, meaning := range def.Tr {
			msg += fmt.Sprintln()
			msg += fmt.Sprintf("* %s (%s), синонимы:\n", strings.ToTitle(meaning.Text), meaning.Pos)
			for count, synonym := range meaning.Syn {
				msg += synonym.Text
				if count != len(meaning.Syn)-1 {
					msg += ", "
				}
			}
			for count, meani := range meaning.Mean {
				msg += meani.Text
				if count != len(meaning.Syn)-1 {
					msg += ", "
				}
			}
		}
	}

	return msg, nil
}
