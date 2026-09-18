package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootcmd =&cobra.Command{
	Use:"probe",
	Short: "Probe scans repositories for SOC2 compliance evidence",
	Long:  `Probe is a CLI tool that scans a local repository for security and
	compliance evidence — authentication, access control, secrets, CI/CD,
	testing, dependencies, and more — and maps what it finds to SOC 2 controls.`,
}
func Execute(){
	if err:=rootcmd.Execute();err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
}
