package soprocli

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func TranscodeFile() {
	// create a new HTTP client
	client := &http.Client{}

	// open the audio file to be uploaded
	file, err := os.Open("audio.mp3")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// create a new multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// add the audio file to the form
	part, err := writer.CreateFormFile("audio", "audio.mp3")
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := io.Copy(part, file); err != nil {
		fmt.Println(err)
		return
	}

	// close the multipart form
	if err := writer.Close(); err != nil {
		fmt.Println(err)
		return
	}

	// create a new HTTP request
	req, err := http.NewRequest("POST", "http://localhost:3000/process-audio", body)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// send the HTTP request and get the response
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// save the processed audio file to disk
	file, err = os.Create("processed-audio.mp3")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// read the response body and write it to the processed audio file
	if _, err := io.Copy(file, resp.Body); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Audio file processed successfully")
}
