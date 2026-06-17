package cmd

func Execute(args []string) error {
	if p := _kendraCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_kendraCmd.Name()}, args...))
		return p.Execute()
	}
	_kendraCmd.SetArgs(args)
	return _kendraCmd.Execute()
}
