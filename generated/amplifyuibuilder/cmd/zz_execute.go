package cmd

func Execute(args []string) error {
	if p := _amplifyuibuilderCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_amplifyuibuilderCmd.Name()}, args...))
		return p.Execute()
	}
	_amplifyuibuilderCmd.SetArgs(args)
	return _amplifyuibuilderCmd.Execute()
}
