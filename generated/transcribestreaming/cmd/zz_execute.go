package cmd

func Execute(args []string) error {
	if p := _transcribestreamingCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_transcribestreamingCmd.Name()}, args...))
		return p.Execute()
	}
	_transcribestreamingCmd.SetArgs(args)
	return _transcribestreamingCmd.Execute()
}
