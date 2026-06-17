package cmd

func Execute(args []string) error {
	if p := _cloudfrontkeyvaluestoreCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_cloudfrontkeyvaluestoreCmd.Name()}, args...))
		return p.Execute()
	}
	_cloudfrontkeyvaluestoreCmd.SetArgs(args)
	return _cloudfrontkeyvaluestoreCmd.Execute()
}
