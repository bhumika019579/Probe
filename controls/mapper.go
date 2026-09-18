package controls

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/bhumika019579/probe/internal/detector"
)

type ControlResult struct {
	ControlID   string     `json:"control_id"`
	Domain      string     `json:"domain"`
	Name        string     `json:"name"`
	Status      string     `json:"status"`
	Evidence    []detector.Evidence `json:"evidence"`
	Description string     `json:"description"`
}
type EvidenceItem struct {
	detector.Evidence
	EvidenceID string `json:"evidence_id"`
}

var detectorAlwaysPartial = map[string]bool{
	"hardening": true,
}

func MapEvidence(evidenceByDetector map[string][]detector.Evidence, defs []Control) []ControlResult {
	results := make([]ControlResult, 0, len(defs))
	for _, def := range defs {
		ev := evidenceByDetector[def.Detector]
		status := "gap"
		if len(ev) > 0 {
			if detectorAlwaysPartial[def.Detector] {
				status = "partial"
			} else {
				status = "found"
			}
		}
		results = append(results, ControlResult{
			ControlID:   def.ID,
			Domain:      def.Domain,
			Name:        def.Name,
			Status:      status,
			Evidence:    ev,
			Description: def.Description,
		})
	}
	return results
}
func EvidenceID(ControlID string, e detector.Evidence) string {
	raw := fmt.Sprintf("%s|%s|%d|%s", ControlID, e.File, e.Line, e.Description)
	sum:=sha256.Sum256([]byte(raw))
	return hex.EncodeToString(sum[:])[:12]
}