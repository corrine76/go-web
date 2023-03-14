package funcs

import (
	"context"
	"encoding/json"
	openai "github.com/sashabaranov/go-openai"
	"html/template"
	"net/http"
	"net/url"
	"os"
)

type ChatGPTRequest struct {
	Messages []openai.ChatCompletionMessage `json:"messages"`
	Ask      string                         `json:"ask"`
}

type Config struct {
	APIKey string `json:"api_key"`
	Proxy  string `json:"proxy"`
}

type Page struct {
	Title    string
	APIKey   string
	Messages []openai.ChatCompletionMessage
	Reply    string
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

func sendText(req *openai.ChatCompletionRequest, apiKey, proxy string) (string, error) {
	config := openai.DefaultConfig(apiKey)
	proxyUrl, err := url.Parse(proxy)
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	config.HTTPClient = &http.Client{
		Transport: transport,
	}

	ctx := context.TODO()
	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(ctx, *req)
	if err != nil {
		return "", err
	}
	req.Messages = append(req.Messages, resp.Choices[0].Message)
	return resp.Choices[0].Message.Content, nil

	// resp, err := client.CreateChatCompletion(
	// 	context.TODO(),
	// 	openai.ChatCompletionRequest{
	// 		Model: openai.GPT3Dot5Turbo,
	// 		Messages: []openai.ChatCompletionMessage{
	// 			{
	// 				Role:    openai.ChatMessageRoleUser,
	// 				Content: content,
	// 			},
	// 		},
	// 	},
	// )
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

		req := r.FormValue("req")

		response, err := generateText(prompt, length, config.APIKey, config.Proxy)
		if err != nil {
			http.Error(w, "Failed to generate text: "+err.Error(), http.StatusInternalServerError)
			return
		}

		page := Page{
			Title:    "ChatGPT API Demo",
			APIKey:   config.APIKey,
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
		Title: "ChatGPT API Demo",
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
