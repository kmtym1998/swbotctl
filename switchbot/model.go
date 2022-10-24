package switchbot

type ListDeviceResponse struct {
	StatusCode int                    `json:"statusCode"`
	Body       ListDeviceResponseBody `json:"body"`
	Message    string                 `json:"message"`
}
type Device struct {
	DeviceID           string `json:"deviceId"`
	DeviceName         string `json:"deviceName"`
	DeviceType         string `json:"deviceType"`
	EnableCloudService bool   `json:"enableCloudService,omitempty"`
	HubDeviceID        string `json:"hubDeviceId"`
}
type InfraredRemote struct {
	DeviceID    string `json:"deviceId"`
	DeviceName  string `json:"deviceName"`
	RemoteType  string `json:"remoteType"`
	HubDeviceID string `json:"hubDeviceId"`
}
type ListDeviceResponseBody struct {
	DeviceList         []Device         `json:"deviceList"`
	InfraredRemoteList []InfraredRemote `json:"infraredRemoteList"`
}

type SendDeviceControlCommandsRequest struct {
	Command     string  `json:"command"`
	Parameter   *string `json:"parameter,omitempty"`
	CommandType *string `json:"commandType,omitempty"`
}
