package core

import "github.com/go-chi/chi"

var plugins map[string]RouterPlugin

func init() {
	plugins = make(map[string]RouterPlugin)
}

// RegisterRouterPlugin is used to register a plugin
// Notice: should called by package init()
func RegisterRouterPlugin(plugin RouterPlugin) {
	x := plugin.PluginInfo()
	plugins[x.Name] = plugin
}

// GetRouterPluginRegistry returns a list of all registered RouterPlugin
func GetRouterPluginRegistry() []RouterPlugin {
	x := make([]RouterPlugin, len(plugins))
	i := 0
	for name := range plugins {
		x[i] = plugins[name]
		i++
	}
	return x
}

// GetRouterPlugin return a RouterPlugin from registry that matches the key
func GetRouterPlugin(name string) (p RouterPlugin, ok bool) {
	p, ok = plugins[name]
	return
}

// RouterPlugin is a type that is used as a p2pNG plugin.
// p2pNG will call PluginInfo first, so PluginInfo should returns static info,
// then fill the Config and call Init,
// at last GetRouter() to mount and init database buckets.
type RouterPlugin interface {
	PluginInfo() *PluginInfo
	GetRouter() chi.Router
	Init() error
	Config() interface{}
}

// PluginInfo is a type that describe a p2pNG plugin.
// Name is should be unique, usually use module path for it.
// Prefix indicates the http service path prefix.
// Buckets indicates which database bucket will use.
type PluginInfo struct {
	Name    string
	Version string
	Prefix  string
	Buckets []string
}
