package extract

import (
	"fmt"
	"regexp"
	"strings"
)

// An analyzer can inspect each history line to find leaked password or secrets.
type analyzer interface {
	// analyze the provided line and return an error if there are
	// issues during the detection phase.
	analyze(num int, line string) error
	// matches return all the
	matches() map[string][]Detection
	finalize() error
}

func newStdAnalyzer() *stdAnalyzer {
	return &stdAnalyzer{detections: make(map[string][]Detection)}
}

type stdAnalyzer struct {
	lineNum    int
	line       string
	detections map[string][]Detection
}

// matches implements analyzer
func (s stdAnalyzer) matches() map[string][]Detection {
	return s.detections
}

// finalize executes a detection on the last file history line.
func (s *stdAnalyzer) finalize() error {
	return s.detect()
}

// analyze implements analyzer
func (s *stdAnalyzer) analyze(num int, line string) error {
	return s.initLine(num, line)
}

// initLine `detect` a leak and sets `num` and `line` as the current values of `s`.
func (s *stdAnalyzer) initLine(num int, line string) error {
	if err := s.detect(); err != nil {
		return err
	}
	s.lineNum = num
	s.line = strings.TrimSuffix(line, "\\\\")
	return nil
}

// append `in` to the current line. This is useful for history multi-lines files (like zsh).
func (s *stdAnalyzer) append(in string) {
	in = strings.TrimSuffix(in, "\\\\")
	s.line = fmt.Sprintf("%s %s", s.line, in)
}

// detect executes all the knownRegexes on the current line, when the first matches it will exit.
func (s *stdAnalyzer) detect() error {
	if s.line == "" {
		return nil
	}
	for provider, re := range knownRegexes {
		match, err := regexp.MatchString(re, s.line)
		if err != nil {
			return err
		}
		if !match {
			continue
		}
		current, ok := s.detections[provider]
		if !ok {
			current = make([]Detection, 0)
		}
		current = append(current, Detection{
			LineNum: s.lineNum,
			Text:    s.line,
		})
		s.detections[provider] = current
		break
	}
	return nil
}
