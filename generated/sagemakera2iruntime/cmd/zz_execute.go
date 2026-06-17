package cmd

func Execute(args []string) error {
	if p := _sagemakera2iruntimeCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_sagemakera2iruntimeCmd.Name()}, args...))
		return p.Execute()
	}
	_sagemakera2iruntimeCmd.SetArgs(args)
	return _sagemakera2iruntimeCmd.Execute()
}
