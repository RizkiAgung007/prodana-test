package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strings" 
)

type AIRequest struct {
	ProductName string `json:"product_name"`
}

func GenerateProductDescription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req AIRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"message": "Format request tidak valid"}`, http.StatusBadRequest)
		return
	}

	apiKey := os.Getenv("APIKEY")
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-flash-latest:generateContent"

	payload := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{
						"text": "Buatkan deskripsi singkat maksimal 2 kalimat untuk produk: " + req.ProductName + ". Langsung berikan teks deskripsinya saja tanpa kata pembuka dan tanpa tanda kutip.",
					},
				},
			},
		},
	}

	jsonValue, _ := json.Marshal(payload)
	
	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-goog-api-key", apiKey) 

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		http.Error(w, `{"message": "Gagal terhubung ke server AI"}`, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	var text string
	if candidates, ok := result["candidates"].([]interface{}); ok && len(candidates) > 0 {
		candidate := candidates[0].(map[string]interface{})
		if content, ok := candidate["content"].(map[string]interface{}); ok {
			if parts, ok := content["parts"].([]interface{}); ok && len(parts) > 0 {
				text = parts[0].(map[string]interface{})["text"].(string)
			}
		}
	}

	if text != "" {
		text = strings.Trim(text, "\" \n\r")
		text = strings.Replace(text, "Berikut adalah deskripsi produknya: ", "", 1)
	} else {
		http.Error(w, `{"message": "AI tidak memberikan jawaban"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"result": text,
	})
}