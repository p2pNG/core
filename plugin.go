package core

import "github.com/go-chi/chi"

var plugins map[string]RouterPlugin

func init() {
	plugins = make(map[string]RouterPlugin)
}

func RegisterRouterPlugin(plugin RouterPlugin) {
	x := plugin.PluginInfo()
	plugins[x.Name] = plugin
}
func GetRouterPluginRegistry() []RouterPlugin {
	x := make([]RouterPlugin, len(plugins))
	i := 0
	for name := range plugins {
		x[i] = plugins[name]
		i++
	}
	return x
}

func GetRouterPlugin(name string) (RouterPlugin, bool) {
	x, ok := plugins[name]
	return x, ok
}

type RouterPlugin interface {
	PluginInfo() *PluginInfo
	GetRouter() chi.Router
}

type PluginInfo struct {
	Name    string
	Version string
	Prefix  string
	Buckets []string
}
