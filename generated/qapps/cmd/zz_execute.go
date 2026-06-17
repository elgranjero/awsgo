package cmd

func Execute(args []string) error {
	if p := _qappsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_qappsCmd.Name()}, args...))
		return p.Execute()
	}
	_qappsCmd.SetArgs(args)
	return _qappsCmd.Execute()
}
