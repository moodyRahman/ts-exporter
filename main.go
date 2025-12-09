package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type TsResponse struct {
	Devices []struct {
		Addresses                 []string  `json:"addresses"`
		Authorized                bool      `json:"authorized"`
		BlocksIncomingConnections bool      `json:"blocksIncomingConnections"`
		ClientVersion             string    `json:"clientVersion"`
		ConnectedToControl        bool      `json:"connectedToControl"`
		Created                   time.Time `json:"created"`
		Expires                   time.Time `json:"expires"`
		Hostname                  string    `json:"hostname"`
		ID                        string    `json:"id"`
		IsExternal                bool      `json:"isExternal"`
		KeyExpiryDisabled         bool      `json:"keyExpiryDisabled"`
		LastSeen                  time.Time `json:"lastSeen"`
		MachineKey                string    `json:"machineKey"`
		Name                      string    `json:"name"`
		NodeID                    string    `json:"nodeId"`
		NodeKey                   string    `json:"nodeKey"`
		Os                        string    `json:"os"`
		TailnetLockError          string    `json:"tailnetLockError"`
		TailnetLockKey            string    `json:"tailnetLockKey"`
		UpdateAvailable           bool      `json:"updateAvailable"`
		User                      string    `json:"user"`
	} `json:"devices"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	req, _ := http.NewRequest("GET", "https://api.tailscale.com/api/v2/tailnet/"+os.Getenv("TS_NET")+"/devices", nil)
	req.Header.Set("Authorization", "Bearer "+os.Getenv("TS_AUTHKEY"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	var devices TsResponse
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	_ = json.Unmarshal(body, &devices)

	debug_out, err := json.MarshalIndent(devices.Devices[0], "", "	")
	fmt.Println(string(debug_out))

	if err != nil {
		fmt.Println(err)
		fmt.Println(string(body))
		return
	}

	// fmt.Println(devices)

}
