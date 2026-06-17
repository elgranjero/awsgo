package cmd

func Execute(args []string) error {
	if p := _iotfleetwiseCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_iotfleetwiseCmd.Name()}, args...))
		return p.Execute()
	}
	_iotfleetwiseCmd.SetArgs(args)
	return _iotfleetwiseCmd.Execute()
}
