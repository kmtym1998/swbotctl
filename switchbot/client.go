package switchbot

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Switchbot struct {
	APIBaseURL, APIVersion, Token, Secret string
}

func NewClient(token, secret string) Switchbot {
	return Switchbot{
		APIBaseURL: "https://api.switch-bot.com",
		APIVersion: "v1.0",
		Token:      token,
		Secret:     secret,
	}
}

func (sw *Switchbot) SendDeviceControlCommands(
	deviceID string,
	input SendDeviceControlCommandsRequest,
) error {
	endpoint := fmt.Sprintf(
		"%s/%s/devices/%s/commands",
		sw.APIBaseURL, sw.APIVersion, deviceID,
	)

	log.Println("sw.Token", sw.Token)

	b, err := json.Marshal(input)
	if err != nil {
		return err
	}

	resp, err := sw.SendRequest(
		http.MethodPost,
		endpoint,
		b,
		map[string]string{
			"content-type":  "application/json",
			"authorization": sw.Token,
		},
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// FIXME: ロガー使う
	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println(string(b))

	return nil
}

func (sw *Switchbot) ListDevices() (*ListDeviceResponse, error) {
	endpoint := fmt.Sprintf(
		"%s/%s/devices",
		sw.APIBaseURL, sw.APIVersion,
	)
	resp, err := sw.SendRequest(
		http.MethodGet,
		endpoint,
		nil,
		map[string]string{
			"content-type":  "application/json",
			"authorization": sw.Token,
		},
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response from switchbot API: %v", string(b))
	}

	var data ListDeviceResponse
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
