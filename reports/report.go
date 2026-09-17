package reports

import (
	"time"

	"github.com/bhumika019579/probe/controls"
)

type ControlOutput struct {
	ControlID   string            `json:"control_id"`
	Domain      string            `json:"domain"`
	Name        string            `json:"name"`
	Status      string            `json:"status"`
	Evidence    []controls.EvidenceItem `json:"evidence"`
	Description string            `json:"description"`
}
type Summary struct{
	Found   int `json:"found"`
	Partial int `json:"partial"`
	Gap     int `json:"gap"`
}
type Report struct{
	ScannedAt time.Time   `json:"scanned_at"`
	Target  string         `json:"target"`
	Controls []ControlOutput  `json:"controls"`
	Summary Summary            `json:"summary"`
}
func Build(target string,results []controls.ControlResult)Report{
	var out []ControlOutput
	summary:=Summary{}
	for _,r:=range results{
		var items []controls.EvidenceItem
		for _,e:=range r.Evidence{
			items = append(items, controls.EvidenceItem{
				Evidence: e,
				EvidenceID: controls.EvidenceID(r.ControlID,e),
			})
		}
		out = append(out, ControlOutput{
			ControlID: r.ControlID,
			Domain: r.Domain,
			Name: r.Name,
			Status: r.Status,
			Evidence: items,
			Description: r.Description,
		})
		switch r.Status{
		case "Found":
			summary.Found++
		case "partial":
			summary.Partial++
		case "Gap":
			summary.Gap++
		}
	    }
		return  Report{
			ScannedAt: time.Now(),
			Target: target,
			Controls: out,
			Summary: summary,
		
	}
}