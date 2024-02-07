package extract

// newAll return an Extractor which combines zsh and bash Extractors.
func newAll() all {
	return all{}
}

type all struct {
}

func (all) Process() (ShellDetections, error) {
	m1, err := newZsh().Process()
	if err != nil {
		return nil, err
	}
	m2, err := newBash().Process()
	if err != nil {
		return nil, err
	}
	m1[Bash] = m2[Bash]
	return m1, nil
}
