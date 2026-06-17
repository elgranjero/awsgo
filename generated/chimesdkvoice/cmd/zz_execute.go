package cmd

func Execute(args []string) error {
	if p := _chimesdkvoiceCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_chimesdkvoiceCmd.Name()}, args...))
		return p.Execute()
	}
	_chimesdkvoiceCmd.SetArgs(args)
	return _chimesdkvoiceCmd.Execute()
}
