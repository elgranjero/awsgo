package cmd

func Execute(args []string) error {
	if p := _pcaconnectorscepCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_pcaconnectorscepCmd.Name()}, args...))
		return p.Execute()
	}
	_pcaconnectorscepCmd.SetArgs(args)
	return _pcaconnectorscepCmd.Execute()
}
