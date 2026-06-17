package cmd

func Execute(args []string) error {
	if p := _paymentcryptographydataCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_paymentcryptographydataCmd.Name()}, args...))
		return p.Execute()
	}
	_paymentcryptographydataCmd.SetArgs(args)
	return _paymentcryptographydataCmd.Execute()
}
