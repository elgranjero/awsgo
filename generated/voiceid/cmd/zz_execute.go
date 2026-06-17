package cmd

func Execute(args []string) error {
	if p := _voiceidCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_voiceidCmd.Name()}, args...))
		return p.Execute()
	}
	_voiceidCmd.SetArgs(args)
	return _voiceidCmd.Execute()
}
