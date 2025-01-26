package packagemanager

// TODO: Flesh out required methods

type Dnf struct {
}

func (a Dnf) String() string {
	return "dnf"
}

func (a Dnf) Install(packageName string, opts ...string) error {
	return nil
	// args := make([]string, 0, 2+len(opts))
	// args = append(args, "install")
	// args = append(args, packageName)
	// args = append(args, opts...)
	// args = append(args, "-y")
	//
	// return exec.Command("apt", args...).Run()
}
