package main

import (
	"fmt"
	"os/exec"
)

func main() {

	removeOldSiaga()
	gitCloneSiaga()
	storageLinkSiaga()
	setPermissionSiaga()
	fmt.Println("Successfully updated 'server-siaga-skripsi' directory!")

}

func removeOldSiaga() {

	fmt.Println("Removing the old 'server-siaga-skripsi' directory...")
	cmd := exec.Command("rm", "-rf", "/var/www/html/server-siaga-skripsi/")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout))
	fmt.Println("Successfully removed 'server-siaga-skripsi'.")

}

func gitCloneSiaga() {

	fmt.Println("Doing git clone for 'server-siaga-skripsi' directory...")
	cmd := exec.Command("git", "clone", "https://github.com/JoniJonnez/server-siaga-skripsi.git")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout))
	fmt.Println("Done!")

}

func storageLinkSiaga() {

	fmt.Println("Removing the old 'public/storage' directory...")
	cmd := exec.Command("rm", "-rf", "/var/www/html/server-siaga-skripsi/public/storage/")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout))
	fmt.Println("Done!")

	fmt.Println("Doing 'php artisan storage:link' for the new 'server-siaga-skripsi' directory...")
	cmd = exec.Command("php", "artisan", "storage:link")
	cmd.Dir = "./server-siaga-skripsi"
	stdout, err = cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout))
	fmt.Println("Done!")

}

func setPermissionSiaga() {

	fmt.Println("Setting up chmod for 'server-siaga-skripsi' directory...")
	cmd := exec.Command("chmod", "-R", "777", "/var/www/html/server-siaga-skripsi/")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout))
	fmt.Println("Done!")

}