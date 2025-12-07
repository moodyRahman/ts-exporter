package main

import (
	"fmt"
	"net/http"
	"time"
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
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Body)

}
