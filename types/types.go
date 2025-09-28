package types

import (
	"html/template"
	"net/http"
)

var Tmpl *template.Template
var LoadedPlugins = make(map[string]bool)
var MenuItems []NavigationConfig

type ReturnData struct {
	Title string
	Data  any
}

func (rd ReturnData) IsEmpty() bool {
	return rd.Title == "" && rd.Data == nil
}

type PluginHandler func(http.ResponseWriter, *http.Request) ReturnData
