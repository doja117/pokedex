package pokepai

import "net/http"

func GetURL(url string) *http.Response {
	r, err := http.Get(url)
	if err != nil {
		return r
	} else {
		return nil
	}

}
