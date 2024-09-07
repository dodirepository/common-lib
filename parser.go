package pkg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// ParseBody :nodoc:
func ParseBody(r *http.Request, payload interface{}) error {

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Print(err.Error())
		return err
	}
	err = json.Unmarshal(body, payload)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}
