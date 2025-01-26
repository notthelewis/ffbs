package packagemanager

// TODO: Flesh out required methods

type Brew struct {
}

func (a Brew) String() string {
	return "brew"
}

func (a Brew) Install(packageName string, opts ...string) error {
	return nil
	// args := make([]string, 0, 2+len(opts))
	// args = append(args, "install")
	// args = append(args, packageName)
	// args = append(args, opts...)
	// args = append(args, "-y")
	//
	// return exec.Command("apt", args...).Run()
}
