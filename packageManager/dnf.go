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

// TODO: Flesh out required methods

func (a Dnf) IsInstalled(packageName string) bool {
	return false
}

func (a Dnf) Remove(packageName string, opts ...string) error {

	return nil
}

func (a Dnf) Update(packageName string, opts ...string) error {

	return nil
}

func (a Dnf) Upgrade(packageName string, opts ...string) error {
	return nil
}

func (a Dnf) Version() string {
	return ""
}
