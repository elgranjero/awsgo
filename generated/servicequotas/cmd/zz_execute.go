package cmd

func Execute(args []string) error {
	if p := _servicequotasCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_servicequotasCmd.Name()}, args...))
		return p.Execute()
	}
	_servicequotasCmd.SetArgs(args)
	return _servicequotasCmd.Execute()
}
