package cmd

func Execute(args []string) error {
	if p := _swfCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_swfCmd.Name()}, args...))
		return p.Execute()
	}
	_swfCmd.SetArgs(args)
	return _swfCmd.Execute()
}
