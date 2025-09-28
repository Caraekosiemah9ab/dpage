package load

import (
	"encoding/json"
	"os"
	"path/filepath"

	"dpage/types"
	"dpage/utils"
)

func ServerConfig() ([]types.ServerConfig, error) {
	exeDir, err := utils.GetCurrentDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(exeDir, "configs", "server.json")

	fileData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var jsonData []types.ServerConfig
	err = json.Unmarshal(fileData, &jsonData)
	return jsonData, err
}

func RouteConfig() ([]types.RouteConfig, error) {
	exeDir, err := utils.GetCurrentDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(exeDir, "configs", "route.json")

	fileData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var jsonData []types.RouteConfig
	err = json.Unmarshal(fileData, &jsonData)
	return jsonData, err
}

func NavConfig() ([]types.NavigationConfig, error) {
	exeDir, err := utils.GetCurrentDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(exeDir, "configs", "navigation.json")

	fileData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var jsonData []types.NavigationConfig
	err = json.Unmarshal(fileData, &jsonData)
	return jsonData, err
}
