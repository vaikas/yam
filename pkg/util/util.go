package util

import "path/filepath"

func IsYAML(path string) bool {
	switch ext := filepath.Ext(path); ext {
	case ".yaml", ".yml":
		return true
	}

	return false
}

const ConfigFileName = ".yam.yaml"
