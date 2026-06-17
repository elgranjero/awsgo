package cmd

func Execute(args []string) error {
	if p := _ssoadminCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ssoadminCmd.Name()}, args...))
		return p.Execute()
	}
	_ssoadminCmd.SetArgs(args)
	return _ssoadminCmd.Execute()
}
