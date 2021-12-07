package adapters

import (
	"better-console-backend/config"
	"encoding/json"
	"fmt"
	"github.com/bettercode-oss/rest"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type WebHookGoRestAdapter struct {
}

func (WebHookGoRestAdapter) GetIgnoredItems() (interface{}, error) {
	client := rest.Client{}
	result := map[string]interface{}{}

	err := client.
		Request().
		SetResult(&result).
		Get(fmt.Sprintf("%v/webhook/hook/ignores", config.Config.WebHookGo.Url))

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (WebHookGoRestAdapter) AddIgnoreItem(newIgnoredItem map[string]interface{}) error {
	data := url.Values{}
	data.Set("instance", newIgnoredItem["instance"].(string))

	if newIgnoredItem["alertName"] != nil {
		data.Set("alert_name", newIgnoredItem["alertName"].(string))
	}

	if newIgnoredItem["job"] != nil {
		data.Set("job", newIgnoredItem["job"].(string))
	}

	if newIgnoredItem["status"] != nil {
		data.Set("status", newIgnoredItem["status"].(string))
	}

	client := &http.Client{}
	r, err := http.NewRequest("POST", config.Config.WebHookGo.Url+"/webhook/hook/ignore", strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		return err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	_, err = client.Do(r)
	if err != nil {
		return err
	}

	return nil
}

func (WebHookGoRestAdapter) GetTemplate() (interface{}, error) {
	client := rest.Client{}
	result := map[string]interface{}{}

	err := client.
		Request().
		SetResult(&result).
		Get(fmt.Sprintf("%v/webhook/hook/template", config.Config.WebHookGo.Url))

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (WebHookGoRestAdapter) UpdateTemplate(template map[string]interface{}) error {
	data := url.Values{}
	data.Set("content", template["content"].(string))

	client := &http.Client{}
	r, err := http.NewRequest("POST", config.Config.WebHookGo.Url+"/webhook/hook/template", strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		return err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	_, err = client.Do(r)
	if err != nil {
		return err
	}

	return nil
}

func (WebHookGoRestAdapter) CheckTemplateSyntax(template map[string]interface{}) (interface{}, error) {
	data := url.Values{}
	data.Set("content", template["content"].(string))

	client := &http.Client{}
	r, err := http.NewRequest("POST", config.Config.WebHookGo.Url+"/webhook/hook/template/check", strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(r)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	bytes, _ := ioutil.ReadAll(resp.Body)
	var responseBody = map[string]interface{}{}
	json.Unmarshal(bytes, &responseBody)

	return responseBody, nil
}

func (WebHookGoRestAdapter) ReloadTemplate() error {
	client := rest.Client{}
	result := map[string]interface{}{}

	err := client.
		Request().
		SetResult(&result).
		Post(fmt.Sprintf("%v/webhook/hook/template/reload", config.Config.WebHookGo.Url))

	if err != nil {
		return err
	}

	return nil
}
