package api

import (
	"encoding/json"
	"fmt"
	"github.com/spitsynv2/stf-monitor/config"
	"github.com/spitsynv2/stf-monitor/model"
	"io"
	"net/http"
)

func GetDevices(provider string) ([]model.Device, error) {
	token := config.Conf.StfProviders[provider]
	if token == "" {
		return nil, fmt.Errorf("no token for provider %s", provider)
	}

	req, err := http.NewRequest("GET", provider+"/api/v1/devices", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("bad response %d: %s", resp.StatusCode, string(body))
	}

	var deviceResponse model.DeviceResponse
	if err := json.NewDecoder(resp.Body).Decode(&deviceResponse); err != nil {
		return nil, err
	}

	return deviceResponse.Devices, nil
}
