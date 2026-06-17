package cmd

func Execute(args []string) error {
	if p := _bcmpricingcalculatorCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_bcmpricingcalculatorCmd.Name()}, args...))
		return p.Execute()
	}
	_bcmpricingcalculatorCmd.SetArgs(args)
	return _bcmpricingcalculatorCmd.Execute()
}
