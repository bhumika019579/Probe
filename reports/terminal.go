package reports

import (
	"fmt"
	"github.com/fatih/color"
)

func PrintTerminal(r Report) {
	fmt.Printf("Scan target: %s\n", r.Target)
	fmt.Printf("Scanned at:  %s\n\n", r.ScannedAt.Format("2006-01-02 15:04:05"))
	for _,c:=range r.Controls{
		statusText:=statusLabel(c.Status)
		fmt.Printf("[%s] %s — %s\n", c.ControlID, c.Name, statusText)
		if len(c.Evidence)==0{
			fmt.Println("No Evidence Found")
		}else{
			for _,e:=range c.Evidence{
				fmt.Printf("    - %s:%d — %s (confidence: %s)\n", e.File, e.Line, e.Description, e.Confidence)
			}
		}
		fmt.Println()
	}
	fmt.Printf(
		"Summary: %s found, %s partial, %s gap\n",
		color.GreenString("%d", r.Summary.Found),
		color.YellowString("%d", r.Summary.Partial),
		color.RedString("%d", r.Summary.Gap),
	)
}
func statusLabel(status string)string{
	switch status{
	case "found":
		return color.GreenString("FOUND")
	case "partial":
		return color.YellowString("PARTIAL")
	case "gap":
		return color.RedString("GAP")
	default:
		return status

	}
}