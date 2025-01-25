package packagemanager

import (
	"fmt"
	"os/exec"
)

type Apt struct {
}

func (a Apt) String() string {
	return "apt"
}

func (a Apt) Install(packageName string, opts ...string) error {
	args := make([]string, 0, 4+len(opts))
	args = append(args, "apt")
	args = append(args, "install")
	args = append(args, packageName)
	args = append(args, opts...)
	args = append(args, "-y")

	aptResult := exec.Command("sudo", args...).Run()
	if aptResult != nil {
		fmt.Println("Error installing", packageName, "with apt. trying with snap...")
		snapResult := exec.Command("sudo", "snap", "install", packageName).Run()
		if snapResult != nil {
			return snapResult
		}
	}

	fmt.Println("installed", packageName)
	return nil
}

// TODO: Flesh out required methods

func (a Apt) IsInstalled(packageName string) bool {
	return false
	//
	// // dpkg-query -s packageName | grep 'Status' | awk '{print $4}'
	// res, err := exec.Command(
	// 	"dpkg-query",
	// 	"-s",
	// 	packageName,
	// 	"|",
	// 	"grep",
	// 	"'Status'",
	// 	"|",
	// 	"awk",
	// 	"'{print $4}'",
	// ).Output()
	//
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return false
	// }
	//
	// return string(res) != ""
}

func (a Apt) Remove(packageName string, opts ...string) error {

	return nil
}

func (a Apt) Update(packageName string, opts ...string) error {

	return nil
}

func (a Apt) Upgrade(packageName string, opts ...string) error {
	return nil
}

func (a Apt) Version() string {
	return ""
}
