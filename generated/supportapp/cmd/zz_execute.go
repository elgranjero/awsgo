package cmd

func Execute(args []string) error {
	if p := _supportappCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_supportappCmd.Name()}, args...))
		return p.Execute()
	}
	_supportappCmd.SetArgs(args)
	return _supportappCmd.Execute()
}
