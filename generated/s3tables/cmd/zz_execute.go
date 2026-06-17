package cmd

func Execute(args []string) error {
	if p := _s3tablesCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_s3tablesCmd.Name()}, args...))
		return p.Execute()
	}
	_s3tablesCmd.SetArgs(args)
	return _s3tablesCmd.Execute()
}
