// Package renderer — compliance module.
//
// # Compliance
//
// The compliance module computes per-resource-type and overall compliance
// scores based on the ratio of IaC-managed resources to the total number
// of discovered resources (managed + unmanaged + deleted).
//
// # Grading
//
// Grades are assigned as follows:
//
//	 A  ≥ 95 %
//	 B  ≥ 80 %
//	 C  ≥ 60 %
//	 D  ≥ 40 %
//	 F  < 40 %
//
// # Options
//
// Use [WithCompliance] to enable the section and [WithComplianceTopN] to
// restrict the table to the N worst-performing resource types (sorted by
// ascending compliance percentage).
package renderer
