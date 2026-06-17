package cmd

func Execute(args []string) error {
	if p := _elementalinferenceCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_elementalinferenceCmd.Name()}, args...))
		return p.Execute()
	}
	_elementalinferenceCmd.SetArgs(args)
	return _elementalinferenceCmd.Execute()
}
