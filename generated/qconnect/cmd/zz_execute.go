package cmd

func Execute(args []string) error {
	if p := _qconnectCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_qconnectCmd.Name()}, args...))
		return p.Execute()
	}
	_qconnectCmd.SetArgs(args)
	return _qconnectCmd.Execute()
}
