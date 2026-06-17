package cmd

func Execute(args []string) error {
	if p := _textractCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_textractCmd.Name()}, args...))
		return p.Execute()
	}
	_textractCmd.SetArgs(args)
	return _textractCmd.Execute()
}
