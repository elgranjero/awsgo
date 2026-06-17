package cmd

func Execute(args []string) error {
	if p := _invoicingCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_invoicingCmd.Name()}, args...))
		return p.Execute()
	}
	_invoicingCmd.SetArgs(args)
	return _invoicingCmd.Execute()
}
