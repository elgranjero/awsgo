package cmd

func Execute(args []string) error {
	if p := _transcribeCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_transcribeCmd.Name()}, args...))
		return p.Execute()
	}
	_transcribeCmd.SetArgs(args)
	return _transcribeCmd.Execute()
}
