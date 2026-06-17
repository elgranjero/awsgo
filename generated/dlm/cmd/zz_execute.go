package cmd

func Execute(args []string) error {
	if p := _dlmCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_dlmCmd.Name()}, args...))
		return p.Execute()
	}
	_dlmCmd.SetArgs(args)
	return _dlmCmd.Execute()
}
