package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/joho/godotenv"
)

type TsResponse struct {
	Devices []struct {
		// Addresses                 []string  `json:"addresses,omitempty"`
		// Authorized                bool      `json:"authorized,omitempty"`
		// BlocksIncomingConnections bool      `json:"blocksIncomingConnections,omitempty"`
		// ClientVersion             string    `json:"clientVersion,omitempty"`
		// ConnectedToControl        bool      `json:"connectedToControl,omitempty"`
		// Created                   time.Time `json:"created,omitempty"`
		Expires time.Time `json:"expires,omitempty"`
		// Hostname                  string    `json:"hostname,omitempty"`
		// ID                        string    `json:"id,omitempty"`
		// IsExternal                bool      `json:"isExternal,omitempty"`
		KeyExpiryDisabled bool      `json:"keyExpiryDisabled"`
		LastSeen          time.Time `json:"lastSeen,omitempty"`
		// MachineKey                string    `json:"machineKey,omitempty"`
		Name string `json:"name,omitempty"`
		// NodeID                    string    `json:"nodeId,omitempty"`
		// NodeKey                   string    `json:"nodeKey,omitempty"`
		// Os                        string    `json:"os,omitempty"`
		// TailnetLockError          string    `json:"tailnetLockError,omitempty"`
		// TailnetLockKey            string    `json:"tailnetLockKey,omitempty"`
		// UpdateAvailable           bool      `json:"updateAvailable,omitempty"`
		// User                      string    `json:"user,omitempty"`
	} `json:"devices,omitempty"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		fmt.Println("using injected env variables")
	}

	// fmt.Println("https://api.tailscale.com/api/v2/tailnet/"+os.Getenv("TS_NET")+"/devices")

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {

		req, _ := http.NewRequest("GET", "https://api.tailscale.com/api/v2/tailnet/"+os.Getenv("TS_NET")+"/devices", nil)
		req.Header.Set("Authorization", "Bearer "+os.Getenv("TS_ACCESSKEY"))

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "internal error")
			return
		}

		var devices TsResponse
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "internal error")
			return
		}

		// fmt.Println("here")
		// fmt.Println(string(body))

		err = json.Unmarshal(body, &devices)

		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "internal error")
			return
		}

		// debug_out, err := json.MarshalIndent(devices, "", "	")
		// fmt.Println(string(debug_out))

		// if err != nil {
		// 	fmt.Println(err)
		// 	fmt.Println(string(body))
		// 	fmt.Fprintf(w, "internal error")
		// 	return
		// }
		// w.Header().Set("Content-Type", "application/json")

		funcMap := template.FuncMap{
			"toUnix": func(t time.Time) int { return int(t.Unix()) },
		}
		templ_date := "ts_expiry_date {{`{`}}name=\"{{.Name}}\" {{`}`}} {{toUnix .Expires}} \n" +
			"ts_expiry_enabled {{`{`}}name=\"{{.Name}}\" {{`}`}} {{.KeyExpiryDisabled}} \n"

		t_date, err := template.New("t").Funcs(funcMap).Parse(templ_date)
		if err != nil {
			fmt.Println(err)
		}

		// fmt.Println(len(devices.Devices))

		for _, device := range devices.Devices {
			// fmt.Println(device.Name)
			// device_s, _ := json.MarshalIndent(device, "", "  ")

			// fmt.Println(string(device_s))

			err = t_date.Execute(w, device)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Fprint(w, "\n")
			fmt.Fprint(w, "\n")

			// fmt.Fprint(w, "{")
			// fmt.Fprint(w, device.Name)
			// fmt.Fprint(w, "}")
		}

		// fmt.Fprint(w, string(debug_out))
	})

	http.ListenAndServe("0.0.0.0:5000", nil)

	// fmt.Println(devices)

}
