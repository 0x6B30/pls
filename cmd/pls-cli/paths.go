package main

import (
	"os"
	"path/filepath"
)

const historyFileEnv = "PLS_HISTORY_FILE"

func historyFilePath() (string, error) {
	if custom := os.Getenv(historyFileEnv); custom != "" {
		return custom, nil
	}
	if dataHome := os.Getenv("XDG_DATA_HOME"); filepath.IsAbs(dataHome) {
		return filepath.Join(dataHome, "pls", "history.tsv"), nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".local", "share", "pls", "history.tsv"), nil
}
