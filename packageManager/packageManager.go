package packagemanager

type PackageManager interface {
	Install(packageName string, opts ...string) error
	IsInstalled(packageName string) bool
	Remove(packageName string, opts ...string) error
	Update(packageName string, opts ...string) error
	Upgrade(packageName string, opts ...string) error
	Version() string
	String() string
}

func New(packagemanager string) PackageManager {
	switch packagemanager {
	case "apt":
		return Apt{}
	case "dnf":
		return Dnf{}
	case "pacman":
		return Pacman{}
	case "brew":
		return Brew{}
	default:
		panic("Unsupported package manager")
	}
}
