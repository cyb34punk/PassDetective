package extract

import (
	"os"
	"path/filepath"
)

func newBash() bash {
	return bash{}
}

type bash struct {
}

func (z bash) Process() (ShellDetections, error) {
	historyFile := filepath.Join(homeDir, ".bash_history")
	if _, err := os.Stat(historyFile); err != nil {
		return nil, ErrNotFound
	}
	return processFile(historyFile, Bash, newStdAnalyzer())
}
