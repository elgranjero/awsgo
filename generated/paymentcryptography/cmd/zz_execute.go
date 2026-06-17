package cmd

func Execute(args []string) error {
	if p := _paymentcryptographyCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_paymentcryptographyCmd.Name()}, args...))
		return p.Execute()
	}
	_paymentcryptographyCmd.SetArgs(args)
	return _paymentcryptographyCmd.Execute()
}
