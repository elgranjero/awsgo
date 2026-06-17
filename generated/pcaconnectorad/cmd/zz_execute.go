package cmd

func Execute(args []string) error {
	if p := _pcaconnectoradCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_pcaconnectoradCmd.Name()}, args...))
		return p.Execute()
	}
	_pcaconnectoradCmd.SetArgs(args)
	return _pcaconnectoradCmd.Execute()
}
