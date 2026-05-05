# driftctl-report

> Generates human-readable HTML reports from driftctl JSON output for infra drift auditing.

---

## Installation

```bash
go install github.com/yourusername/driftctl-report@latest
```

Or build from source:

```bash
git clone https://github.com/yourusername/driftctl-report.git
cd driftctl-report
go build -o driftctl-report .
```

---

## Usage

First, generate a JSON scan output using [driftctl](https://github.com/snyk/driftctl):

```bash
driftctl scan --output json://drift.json
```

Then pass the JSON file to `driftctl-report` to produce an HTML report:

```bash
driftctl-report --input drift.json --output report.html
```

Open `report.html` in your browser to review a formatted summary of drifted, missing, and unmanaged resources.

### Options

| Flag | Default | Description |
|------|---------|-------------|
| `--input` | `drift.json` | Path to driftctl JSON output file |
| `--output` | `report.html` | Path for the generated HTML report |
| `--title` | `Drift Report` | Custom title displayed in the report |

---

## Example Output

The generated report includes:

- **Coverage score** and scan summary
- **Unmanaged resources** — resources found in cloud but not in IaC
- **Missing resources** — resources defined in IaC but absent in cloud
- **Changed resources** — resources with configuration drift

---

## Requirements

- Go 1.21+
- driftctl JSON output (v0.39.0+)

---

## License

This project is licensed under the [MIT License](LICENSE).