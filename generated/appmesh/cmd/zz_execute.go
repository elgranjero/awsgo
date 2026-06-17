package cmd

func Execute(args []string) error {
	if p := _appmeshCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_appmeshCmd.Name()}, args...))
		return p.Execute()
	}
	_appmeshCmd.SetArgs(args)
	return _appmeshCmd.Execute()
}
