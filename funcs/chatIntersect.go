package funcs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type ChatGPTRequest struct {
	Prompt string `json:"prompt"`
	Length int    `json:"length"`
}

type ChatGPTResponse struct {
	Completion string `json:"completion"`
}

type Config struct {
	APIKey string `json:"api_key"`
	Proxy  string `json:"proxy"`
}

type Page struct {
	Title    string
	APIKey   string
	Prompt   string
	Length   int
	Response string
}

func loadConfig(filename string) (Config, error) {
	var config Config
	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func generateText(prompt string, length int, apiKey, proxy string) (string, error) {
	requestBody, err := json.Marshal(ChatGPTRequest{prompt, length})
	if err != nil {
		return "", err
	}
	fmt.Println("发送请求", requestBody)

	apiEndpoint := "https://api.openai.com/v1/chat/completions"
	request, err := http.NewRequest(http.MethodPost, apiEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	request.Header.Set("Authorization", "Bearer "+apiKey)
	request.Header.Set("Content-Type", "application/json")

	proxyUrl, err := url.Parse(proxy)
	if err != nil {
		return "", err
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var chatGPTResponse ChatGPTResponse
	err = json.Unmarshal(responseBody, &chatGPTResponse)
	if err != nil {
		return "", err
	}

	fmt.Println("chatGPTResponse", chatGPTResponse)
	return chatGPTResponse.Completion, nil
}

func ChatDemoHandler(w http.ResponseWriter, r *http.Request) {
	config, err := loadConfig("config.json")
	if err != nil {
		http.Error(w, "Failed to load config: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
			return
		}

		prompt := r.FormValue("prompt")
		length := 50
		response, err := generateText(prompt, length, config.APIKey, config.Proxy)
		if err != nil {
			http.Error(w, "Failed to generate text: "+err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("接收响应", response)

		page := Page{
			Title:    "ChatGPT API Demo",
			APIKey:   config.APIKey,
			Prompt:   prompt,
			Length:   length,
			Response: response,
		}

		tmpl, err := template.ParseFiles("chat-intersect.html")
		if err != nil {
			http.Error(w, "Failed to parse template: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, page)
		if err != nil {
			http.Error(w, "Failed to execute template: "+err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	page := Page{
		Title:  "ChatGPT API Demo",
		APIKey: config.APIKey,
	}

	tmpl, err := template.ParseFiles("chat-intersect.html")
	if err != nil {
		http.Error(w, "Failed to parse template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		http.Error(w, "Failed to execute template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
