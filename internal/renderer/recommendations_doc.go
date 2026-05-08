// Package renderer provides the HTML rendering pipeline for driftctl reports.
//
// # Recommendations
//
// The recommendations module analyses the drift summary and produces a ranked
// list of actionable suggestions to help engineers reduce infrastructure drift.
//
// Three categories of recommendation are currently supported:
//
//   - Import unmanaged resources – resources present in live infrastructure
//     but absent from IaC definitions.
//   - Reconcile deleted resources – resources declared in IaC but no longer
//     present in live infrastructure.
//   - Fix configuration drift – resources whose live state diverges from the
//     desired state expressed in IaC.
//
// Each recommendation is assigned a severity badge (success / warning / danger)
// based on the count of affected resources, using the same thresholds as the
// badge and alert modules.
//
// Recommendations are sorted by descending resource count so that the most
// impactful actions appear first.
//
// Use [WithRecommendations] to enable the section and
// [WithRecommendationsMaxItems] to cap the number of items rendered.
package renderer
