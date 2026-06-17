package cmd

func Execute(args []string) error {
	if p := _mediastoredataCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_mediastoredataCmd.Name()}, args...))
		return p.Execute()
	}
	_mediastoredataCmd.SetArgs(args)
	return _mediastoredataCmd.Execute()
}
