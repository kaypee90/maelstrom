package main

import (
	"encoding/json"

	"os/exec"
)

func generateUUIDUsingCommand() (string, error) {
	cmd := exec.Command("uuidgen")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func decodeJson(data []byte) (map[string]any, error) {
	var body map[string]any
	if err := json.Unmarshal(data, &body); err != nil {
		return nil, err
	}

	return body, nil
}
