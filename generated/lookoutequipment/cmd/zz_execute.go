package cmd

func Execute(args []string) error {
	if p := _lookoutequipmentCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_lookoutequipmentCmd.Name()}, args...))
		return p.Execute()
	}
	_lookoutequipmentCmd.SetArgs(args)
	return _lookoutequipmentCmd.Execute()
}
