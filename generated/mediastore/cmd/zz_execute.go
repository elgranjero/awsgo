package cmd

func Execute(args []string) error {
	if p := _mediastoreCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_mediastoreCmd.Name()}, args...))
		return p.Execute()
	}
	_mediastoreCmd.SetArgs(args)
	return _mediastoreCmd.Execute()
}
