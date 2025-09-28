package main

import (
	"dpage/types"
	"fmt"
	"net/http"
	"strings"
)

func PluginInit() {
	fmt.Printf("[BasePlugin] Is Loaded...\n")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) types.ReturnData {
	return types.ReturnData{
		Title: "Главная",
		Data:  "",
	}
}

func RedirectToNonSlash(w http.ResponseWriter, r *http.Request) types.ReturnData {
	path := r.URL.Path
	if path != "" || !(strings.HasSuffix(path, "/") || strings.Contains(path, ".")) {
		path := path + "/"
		http.Redirect(w, r, path, http.StatusMovedPermanently)
	}

	return types.ReturnData{Data: "norender"}
}
