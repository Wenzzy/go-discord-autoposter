package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func (d *API) PostMessage(channelID string, content *string, filesPaths []string) error {
	var (
		isWithFiles = (filesPaths != nil) && (len(filesPaths) > 0)
		buf         = new(bytes.Buffer)
		multi       *multipart.Writer
	)

	client := &http.Client{}

	if isWithFiles {

		data := map[string]interface{}{
			"content": nil,
			"tts":     "false",
		}
		if content != nil {
			data["content"] = *content
		}
		multi = multipart.NewWriter(buf)

		for _, filePath := range filesPaths {
			file, err := os.Open(filepath.Clean(filePath))
			if err != nil {
				return err
			}
			defer func() {
				_ = file.Close()
			}()

			fileWriter, err := multi.CreateFormFile("files", filePath)
			if err != nil {
				return errors.Errorf("failed to create form file: %v", err)
			}

			if _, err = io.Copy(fileWriter, file); err != nil {
				return errors.Errorf("failed to write file: %v", err)
			}
		}

		for key, val := range data {
			_ = multi.WriteField(key, fmt.Sprintf("%s", val))
		}

		_ = multi.Close()

	} else {
		jsonData, err := json.Marshal(map[string]interface{}{
			"content": content,
			"tts":     false,
		})
		if err != nil {
			return errors.Errorf("failed to marshal data: %v", err)
		}
		buf.Write(jsonData)
	}

	req, err := http.NewRequest("POST", d.baseURL+fmt.Sprintf("/channels/%s/messages", channelID), buf)
	if err != nil {
		return errors.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("Authorization", d.accessToken)
	if isWithFiles {
		req.Header.Add("Content-Type", multi.FormDataContentType())

	} else {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Errorf("failed to send request: %v", err)
	}

	// var body map[string]interface{}

	// err = json.NewDecoder(resp.Body).Decode(&body)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("Response: %v\n", body)

	_ = resp.Body.Close()

	return nil

}
