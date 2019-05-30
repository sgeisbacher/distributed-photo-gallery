package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondJSON(resp http.ResponseWriter, data interface{}, err error) error {
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	resp.Header().Add("Content-Type", "application/json")
	jsonData, err := json.Marshal(data)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return err
	}
	fmt.Fprintln(resp, string(jsonData))
	return nil
}
