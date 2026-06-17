package cmd

func Execute(args []string) error {
	if p := _quicksightCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_quicksightCmd.Name()}, args...))
		return p.Execute()
	}
	_quicksightCmd.SetArgs(args)
	return _quicksightCmd.Execute()
}
