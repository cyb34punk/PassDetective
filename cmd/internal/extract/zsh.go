package extract

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var zshHistoryRE = regexp.MustCompile(`^:\s\d+:\d+;.*`)

func newZsh() zsh {
	return zsh{}
}

type zsh struct {
}

func (z zsh) Process() (ShellDetections, error) {
	historyFile := filepath.Join(homeDir, ".zsh_history")
	if _, err := os.Stat(historyFile); err != nil {
		return nil, ErrNotFound
	}
	return processFile(historyFile, Zsh, newZshAnalyzer())
}

func newZshAnalyzer() *zshAnalyzer {
	return &zshAnalyzer{
		stdAnalyzer: newStdAnalyzer(),
	}
}

type zshAnalyzer struct {
	*stdAnalyzer
}

// analyze override stdAnalyzer
func (z *zshAnalyzer) analyze(num int, line string) error {
	if !zshHistoryRE.MatchString(line) {
		z.append(line)
		return nil
	}
	command := strings.SplitN(line, ";", 2)[1]
	return z.initLine(num, command)
}
