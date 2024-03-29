package middleware

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/TheLazarusNetwork/LazarusTunnel/model"
	"github.com/TheLazarusNetwork/LazarusTunnel/template"
	"github.com/TheLazarusNetwork/LazarusTunnel/util"
)

// IsValid check if model is valid
func IsValidSSH(name string, port int) (int, string, error) {
	// check if the name is empty
	if name == "" {
		return -1, "Tunnel Name is required", nil
	}

	// check the name field is between 3 to 40 chars
	if len(name) < 4 || len(name) > 12 {
		return -1, "Tunnel Name field must be between 4-12 chars", nil
	}

	// check if name or port is already in use
	tunnels, err := ReadSSHTunnels()
	if err != nil {
		return -1, "", err
	} else {
		for _, tunnel := range tunnels.Tunnels {
			if tunnel.Name == name {
				return -1, "Tunnel Already exists", err
			} else if tunnel.Port == strconv.Itoa(port) {
				return -1, "Port Already in use", err
			}
		}
	}

	// check the format of name
	if !util.IsLetter(name) {
		return -1, "Tunnel Name should be Aplhanumeric", nil
	}

	return 1, "", nil
}

// ReadSSHTunnels fetches all SSH Tunnels
func ReadSSHTunnels() (*model.Tunnels, error) {
	file, err := os.OpenFile(filepath.Join(os.Getenv("APP_CONF_DIR"), "nginx.json"), os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		util.LogError("File Open error: ", err)
		return nil, err
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		util.LogError("File Read error: ", err)
		return nil, err
	}

	var tunnels model.Tunnels
	err = json.Unmarshal(b, &tunnels.Tunnels)
	if err != nil {
		util.LogError("Unmarshal json error: ", err)
		return nil, err
	}

	return &tunnels, nil
}

// ReadSSHTunnel fetches a SSH Tunnel
func ReadSSHTunnel(tunnelName string) (*model.Tunnel, error) {
	tunnels, err := ReadSSHTunnels()
	if err != nil {
		return nil, err
	}

	var data model.Tunnel
	for _, tunnel := range tunnels.Tunnels {
		if tunnel.Name == tunnelName {
			data.Name = tunnel.Name
			data.Port = tunnel.Port
			data.CreatedAt = tunnel.CreatedAt
			data.Domain = tunnel.Domain
			data.Status = tunnel.Status
			break
		}
	}

	return &data, nil
}

// AddSSHTunnel creates a SSH Tunnel
func AddSSHTunnel(tunnel model.Tunnel) error {
	tunnels, err := ReadSSHTunnels()
	if err != nil {
		return err
	}

	var updatedTunnels []model.Tunnel
	updatedTunnels = append(updatedTunnels, tunnels.Tunnels...)
	updatedTunnels = append(updatedTunnels, tunnel)

	inter, err := json.MarshalIndent(updatedTunnels, "", "   ")
	if err != nil {
		util.LogError("JSON Marshal error: ", err)
		return err
	}

	err = util.WriteFile(filepath.Join(os.Getenv("APP_CONF_DIR"), "nginx.json"), inter)
	if err != nil {
		util.LogError("File write error: ", err)
		return err
	}

	err = UpdateNginxConfig()
	if err != nil {
		return err
	}

	return nil
}

// DeleteSSHTunnel deletes a SSH Tunnel
func DeleteSSHTunnel(tunnelName string) error {
	tunnels, err := ReadSSHTunnels()
	if err != nil {
		return err
	}

	var updatedTunnels []model.Tunnel
	for _, tunnel := range tunnels.Tunnels {
		if tunnel.Name == tunnelName {
			continue
		}
		updatedTunnels = append(updatedTunnels, tunnel)
	}

	inter, err := json.MarshalIndent(updatedTunnels, "", "   ")
	if err != nil {
		util.LogError("JSON Marshal error: ", err)
		return err
	}

	err = util.WriteFile(filepath.Join(os.Getenv("APP_CONF_DIR"), "nginx.json"), inter)
	if err != nil {
		util.LogError("File write error: ", err)
		return err
	}

	err = UpdateNginxConfig()
	if err != nil {
		return err
	}

	return nil
}

// UpdateNginxConfig updates the sites-available files
func UpdateNginxConfig() error {
	tunnels, err := ReadSSHTunnels()
	if err != nil {
		return err
	}

	path := filepath.Join(os.Getenv("NGINX_CONF_DIR"), os.Getenv("NGINX_INTERFACE_NAME"))
	if util.FileExists(path) {
		os.Remove(path)
	}

	for _, tunnel := range tunnels.Tunnels {
		_, err := template.NginxConfigTempl(tunnel)
		if err != nil {
			util.LogError("Nginx update error: ", err)
			return err
		}
	}

	return nil
}
