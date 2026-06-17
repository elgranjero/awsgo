package cmd

func Execute(args []string) error {
	if p := _signerCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_signerCmd.Name()}, args...))
		return p.Execute()
	}
	_signerCmd.SetArgs(args)
	return _signerCmd.Execute()
}
