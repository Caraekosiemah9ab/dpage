package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"dpage/pkg/load"
	"dpage/types"
	"dpage/utils"
)

func main() {
	var ServerHostAddr (string)
	serverConfig, err := load.ServerConfig()
	if err != nil {
		fmt.Printf("Fail load server config: %v\n", err)
		return
	} else {
		if serverConfig[0].Addr != "" && serverConfig[0].Port != "" {
			ServerHostAddr = serverConfig[0].Addr + ":" + serverConfig[0].Port
		} else {
			fmt.Printf("Fail load server config: Addr or Port")
			return
		}
	}

	configRoute, err := load.RouteConfig()
	if err != nil {
		fmt.Printf("Fail load page config: %v\n", err)
		return
	}

	configNav, err := load.NavConfig()
	if err != nil {
		fmt.Printf("Fail load navigation config: %v\n", err)
		return
	} else {
		var mItems []types.NavigationConfig
		for _, cfg := range configNav {
			mItems = append(mItems, types.NavigationConfig{
				Name:   cfg.Name,
				URL:    cfg.URL,
				Target: cfg.Target,
			})
		}

		types.MenuItems = mItems
	}

	if err := load.Templates(); err != nil {
		fmt.Printf("Fail load templates: %v\n", err)
		return
	}

	pluginsDir, err := utils.GetCurrentDir()
	if err != nil {
		fmt.Printf("Fail open plugins dir\n")
		return
	}
	load.Plugins(filepath.Join(pluginsDir, "plugins"), configRoute)

	fmt.Printf("Server Addr -> ['%s']\n", ServerHostAddr)
	if err := http.ListenAndServe(ServerHostAddr, nil); err != nil {
		fmt.Printf("Error! %v\n", err)
	}
}
