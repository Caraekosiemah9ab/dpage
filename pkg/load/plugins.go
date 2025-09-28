package load

import (
	"dpage/pkg"
	"dpage/types"
	"fmt"
	"net/http"
	"path/filepath"
	"plugin"
	"strings"
	"time"
)

func Plugins(pDir string, cRoute []types.RouteConfig) {
	fmt.Printf("Add routes: %d\n", len(cRoute))
	for _, cfg := range cRoute {
		pluginPath := filepath.Join(pDir, cfg.Plugin)

		plug, err := plugin.Open(pluginPath)
		if err != nil {
			fmt.Printf("Fail load plugin -> ['%s']: %v\n", cfg.Plugin, err)
			continue
		}

		sym, err := plug.Lookup(cfg.Handler)
		if err != nil {
			fmt.Printf("Handler ['%s'] not found -> ['%s']: %v\n", cfg.Handler, cfg.Plugin, err)
			continue
		}

		plugHandler, ok := sym.(func(http.ResponseWriter, *http.Request) types.ReturnData)
		if !ok {
			fmt.Printf("Bad handler ['%s'] -> ['%s']\n", cfg.Handler, cfg.Plugin)
			continue
		}

		sym2, _ := plug.Lookup("PluginInit")
		if PluginInit, ok := sym2.(func()); ok {
			if !types.LoadedPlugins[cfg.Plugin] {
				PluginInit()
				types.LoadedPlugins[cfg.Plugin] = true
			}
		}

		http.HandleFunc(cfg.URL, makeHandler(plugHandler, cfg.Template))
		fmt.Printf("-> ['%s'] -> [T: '%s'] -> [P: '%s']\n", cfg.URL, cfg.Template, cfg.Plugin)
	}
}

func makeHandler(plugHandler types.PluginHandler, tempName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rawData := plugHandler(w, r)
		if rawData.IsEmpty() {
			http.Error(w, "Not found Data", http.StatusInternalServerError)
			return
		}

		if str, ok := rawData.Data.(string); ok && strings.ToLower(str) == "norender" {
			return //fmt.Print("Found norender! Skipping render.\n")
		}

		pkg.SetDataField(&rawData.Data, "NavigationdPage", pkg.ToAnySlice(types.MenuItems))
		pkg.SetDataField(&rawData.Data, "CurrentYear", time.Now().Year())

		err := types.Tmpl.ExecuteTemplate(w, tempName, rawData)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error rendering -> ['%s']: %v", tempName, err), http.StatusInternalServerError)
			return
		}
	}
}
