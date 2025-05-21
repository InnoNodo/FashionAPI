package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	var content string
	fmt.Print("Write your message: ")
	fmt.Scanln(&content)

	url := "https://openrouter.ai/api/v1/chat/completions"

	payload := map[string]interface{}{
		"model": "google/gemma-3n-e4b-it:free",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": content,
			},
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	godotenv.Load()

	req.Header.Set("Authorization", os.Getenv("OPEN_ROUTER"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Title", "FashionAPI")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
json.Unmarshal(body, &data)

choices, ok := data["choices"].([]interface{})
if ok && len(choices) > 0 {
    firstChoice, ok := choices[0].(map[string]interface{})
    if ok {
        message, ok := firstChoice["message"].(map[string]interface{})
        if ok {
            content, ok := message["content"].(string)
            if ok {
                fmt.Println(content)
                return
            }
        }
    }
}	

	
}
