package cmd

func Execute(args []string) error {
	if p := _pricingCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_pricingCmd.Name()}, args...))
		return p.Execute()
	}
	_pricingCmd.SetArgs(args)
	return _pricingCmd.Execute()
}
