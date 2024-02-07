package extract

type Extractor interface {
	Process() (ShellDetections, error)
}

// Kind represents the extractor kind.
type Kind string

// the supported extractors
const (
	Zsh  Kind = "zsh"
	Bash Kind = "bash"
	All  Kind = "all"
)

// Create the Extractor by Kind
func Create(kind Kind) Extractor {
	switch kind {
	case Zsh:
		return newZsh()
	case Bash:
		return newBash()
	case All:
		return newAll()
	default:
		panic("unknown kind " + kind)
	}
}
