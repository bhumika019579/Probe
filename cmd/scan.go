package cmd

import (
	"fmt"

	"github.com/bhumika019579/probe/controls"
	"github.com/bhumika019579/probe/internal/detector"
	"github.com/bhumika019579/probe/internal/walker"
	"github.com/bhumika019579/probe/reports"
	"github.com/spf13/cobra"
)

var jsonOutputPath string
var scancmd = &cobra.Command{
	Use:   "scan [path]",
	Short: "Scan a local repository for SOC2 compliance evidence",
	Long: `Scan walks the given local directory, runs all 8 detectors against
    its files, maps what they find to SOC2 controls, and prints a report.`,
	Args: cobra.ExactArgs(1),
	RunE: runScan,
}
func init(){
	scancmd.Flags().StringVar(&jsonOutputPath,"json","","write a report as json to this filepath")
	rootcmd.AddCommand(scancmd)
}
func runScan(cmd *cobra.Command,args []string)error{
	target:=args[0]
	files,err:=walker.Walk(target)
	if err!=nil{
		return fmt.Errorf("walking %q: %w", target, err)
	}
	evidenceByDetector:=make(map[string][]detector.Evidence)
	for _,d:=range detector.AllDetectors(){
		ev:=d.Detect(files)
		evidenceByDetector[d.Category()]=ev
	}
	defs,err:=controls.LoadControls()
	if err != nil {
		return fmt.Errorf("loading controls: %w", err)
	}
	results:=controls.MapEvidence(evidenceByDetector,defs)
	report:=reports.Build(target,results)
	reports.PrintTerminal(report)
	if jsonOutputPath!=""{
		if err:=reports.WriteJson(report,jsonOutputPath);err!=nil{
			return fmt.Errorf("writing json report: %w", err)
		}
		fmt.Printf("\nJSON report written to %s\n", jsonOutputPath)
	}
	return nil
}