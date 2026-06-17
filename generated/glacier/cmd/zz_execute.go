package cmd

func Execute(args []string) error {
	if p := _glacierCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_glacierCmd.Name()}, args...))
		return p.Execute()
	}
	_glacierCmd.SetArgs(args)
	return _glacierCmd.Execute()
}
