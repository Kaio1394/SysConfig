package utils

import (
	"SysConfig/config"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func RequestValidateToken(token string) bool {
	url := config.ConfigViper.ConfigAuth.Host + ":" + strconv.Itoa(config.ConfigViper.ConfigAuth.Port) + config.ConfigViper.ConfigAuth.Endpoint
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return false
	}
	req.Header.Add("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	if !strings.Contains(string(body), "true") {
		return false
	}
	return true
}
