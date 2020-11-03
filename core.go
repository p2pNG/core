package core

import "runtime/debug"

// GoModule returns the build info of this p2pNG
// build from debug.BuildInfo (requires Go modules).
// If no version information is available, a non-nil
// value will still be returned, but with an
// unknown version.
func GoModule() *debug.Module {
	var mod debug.Module
	return goModule(&mod)
}

// goModule holds the actual implementation of GoModule.
// Allocating debug.Module in GoModule() and passing a
// reference to goModule enables mid-stack inlining.
func goModule(mod *debug.Module) *debug.Module {
	mod.Version = "unknown"
	bi, ok := debug.ReadBuildInfo()
	if ok {
		mod.Path = bi.Main.Path
		// The recommended way to build p2pNG involves
		// creating a separate main module, which
		// TODO: track related Go issue: https://github.com/golang/go/issues/29228
		// once that issue is fixed, we should just be able to use bi.Main... hopefully.
		for _, dep := range bi.Deps {
			if dep.Path == ImportPath {
				return dep
			}
		}
		return &bi.Main
	}
	return mod
}

// ImportPath is the package import path for p2pNG core.
const ImportPath = "github.com/p2pNG/core"
