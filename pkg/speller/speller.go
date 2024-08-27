package speller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"tz-kode/internal/entity"
)

func CheckTexts(texts []string) (entity.Response, error) {
	res, err := send(texts)
	return res, err
}

func send(postData []string) (yr entity.Response, err error) {

	resp, err := http.PostForm("https://speller.yandex.net/services/spellservice.json/checkTexts", url.Values{
		"text": postData,
	})
	if err != nil {
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = resp.Body.Close()
	if err != nil {
		return
	}

	if err = json.Unmarshal(body, &yr); err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		return yr, errors.New(fmt.Sprint("Response status: ", resp.Status))
	}

	return
}
