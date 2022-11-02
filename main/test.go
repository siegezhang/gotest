package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

// qmpacs://DxStartW_64.exe 00700143599 00700143599 %E6%B5%8B%E8%AF%95%E4%B8%93%E7%94%A8 1963-2-31
func main() {
	fmt.Printf("%v\n", os.Args[1])
	var cmd_com string
	cmd_all := os.Args[1]
	fmt.Printf("\ninput: %v\n", cmd_all)
	cmd_cmd_args_arr := strings.Split(strings.Split(cmd_all, "//")[1], "%20")
	cmd_com = cmd_cmd_args_arr[0]
	patientId := cmd_cmd_args_arr[1]
	familyName := cmd_cmd_args_arr[2]
	firstName, _ := url.QueryUnescape(cmd_cmd_args_arr[3])
	birthDate := strings.Replace(cmd_cmd_args_arr[4], "/", "", -1)
	args := cmd_com + " " + strings.Join([]string{patientId, familyName, firstName, birthDate}, " ")
	args = strings.Replace(args, "/", "", -1)
	fmt.Printf("\ncommand: cmd /c %v\n", args)
	cmd := exec.Command("cmd", "/c", args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("err:%v", err)
	}
	time.Sleep(2 * time.Second)
}
