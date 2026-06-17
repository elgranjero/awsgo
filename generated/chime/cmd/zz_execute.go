package cmd

func Execute(args []string) error {
	if p := _chimeCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_chimeCmd.Name()}, args...))
		return p.Execute()
	}
	_chimeCmd.SetArgs(args)
	return _chimeCmd.Execute()
}
