package renderer

// defaultTemplate is the built-in HTML template used to render drift reports.
const defaultTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>driftctl Report</title>
  <style>
    body { font-family: Arial, sans-serif; margin: 2rem; background: #f5f5f5; color: #333; }
    h1 { color: #2c3e50; }
    .meta { color: #666; font-size: 0.9rem; margin-bottom: 1.5rem; }
    .summary { display: flex; gap: 1rem; flex-wrap: wrap; margin-bottom: 2rem; }
    .card { background: #fff; border-radius: 6px; padding: 1rem 1.5rem; box-shadow: 0 1px 4px rgba(0,0,0,0.1); min-width: 140px; }
    .card .label { font-size: 0.8rem; color: #888; text-transform: uppercase; }
    .card .value { font-size: 2rem; font-weight: bold; }
    .managed .value { color: #27ae60; }
    .unmanaged .value { color: #e67e22; }
    .missing .value { color: #e74c3c; }
    .drifted .value { color: #8e44ad; }
    .coverage .value { color: #2980b9; }
    table { width: 100%; border-collapse: collapse; background: #fff; border-radius: 6px; overflow: hidden; box-shadow: 0 1px 4px rgba(0,0,0,0.1); margin-bottom: 2rem; }
    th { background: #2c3e50; color: #fff; padding: 0.6rem 1rem; text-align: left; font-size: 0.85rem; }
    td { padding: 0.55rem 1rem; border-bottom: 1px solid #eee; font-size: 0.85rem; }
    tr:last-child td { border-bottom: none; }
    h2 { color: #2c3e50; margin-top: 2rem; }
    .badge { display: inline-block; padding: 2px 8px; border-radius: 12px; font-size: 0.75rem; font-weight: bold; }
    .badge-ok { background: #d5f5e3; color: #1e8449; }
    .badge-warn { background: #fdebd0; color: #a04000; }
    .badge-err { background: #fadbd8; color: #922b21; }
  </style>
</head>
<body>
  <h1>&#128202; driftctl Drift Report</h1>
  <p class="meta">Generated at: {{.GeneratedAt}}</p>

  <div class="summary">
    <div class="card coverage"><div class="label">Coverage</div><div class="value">{{printf "%.1f" .Summary.CoveragePercent}}%</div></div>
    <div class="card managed"><div class="label">Managed</div><div class="value">{{.Summary.TotalManaged}}</div></div>
    <div class="card unmanaged"><div class="label">Unmanaged</div><div class="value">{{.Summary.TotalUnmanaged}}</div></div>
    <div class="card missing"><div class="label">Missing</div><div class="value">{{.Summary.TotalMissing}}</div></div>
    <div class="card drifted"><div class="label">Drifted</div><div class="value">{{.Summary.TotalDrifted}}</div></div>
  </div>

  {{if .Report.UnmanagedResources}}
  <h2>&#9888;&#65039; Unmanaged Resources</h2>
  <table>
    <thead><tr><th>Type</th><th>ID</th></tr></thead>
    <tbody>
    {{range .Report.UnmanagedResources}}
      <tr><td>{{.Type}}</td><td>{{.ID}}</td></tr>
    {{end}}
    </tbody>
  </table>
  {{end}}

  {{if .Report.MissingResources}}
  <h2>&#10060; Missing Resources</h2>
  <table>
    <thead><tr><th>Type</th><th>ID</th></tr></thead>
    <tbody>
    {{range .Report.MissingResources}}
      <tr><td>{{.Type}}</td><td>{{.ID}}</td></tr>
    {{end}}
    </tbody>
  </table>
  {{end}}

  {{if .Report.DifferentResources}}
  <h2>&#128260; Drifted Resources</h2>
  <table>
    <thead><tr><th>Type</th><th>ID</th></tr></thead>
    <tbody>
    {{range .Report.DifferentResources}}
      <tr><td>{{.Type}}</td><td>{{.ID}}</td></tr>
    {{end}}
    </tbody>
  </table>
  {{end}}

  <footer style="margin-top:3rem;font-size:0.8rem;color:#aaa;">driftctl-report &mdash; infra drift auditing</footer>
</body>
</html>
`
