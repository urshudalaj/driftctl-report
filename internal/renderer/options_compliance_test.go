package renderer

import "testing"

func TestWithCompliance_Enables(t *testing.T) {
	opts := DefaultOptions()
	WithCompliance(true)(&opts)
	if !opts.Compliance {
		t.Fatal("expected Compliance to be true")
	}
}

func TestWithCompliance_Disables(t *testing.T) {
	opts := DefaultOptions()
	opts.Compliance = true
	WithCompliance(false)(&opts)
	if opts.Compliance {
		t.Fatal("expected Compliance to be false")
	}
}

func TestWithCompliance_DefaultIsFalse(t *testing.T) {
	opts := DefaultOptions()
	if opts.Compliance {
		t.Fatal("expected Compliance default to be false")
	}
}

func TestWithComplianceTopN_PositiveValue(t *testing.T) {
	opts := DefaultOptions()
	WithComplianceTopN(5)(&opts)
	if opts.ComplianceTopN != 5 {
		t.Fatalf("expected ComplianceTopN=5, got %d", opts.ComplianceTopN)
	}
}

func TestWithComplianceTopN_ZeroMeansUnlimited(t *testing.T) {
	opts := DefaultOptions()
	WithComplianceTopN(10)(&opts)
	WithComplianceTopN(0)(&opts)
	if opts.ComplianceTopN != 0 {
		t.Fatalf("expected ComplianceTopN=0, got %d", opts.ComplianceTopN)
	}
}

func TestWithComplianceTopN_NegativeClampedToZero(t *testing.T) {
	opts := DefaultOptions()
	WithComplianceTopN(-3)(&opts)
	if opts.ComplianceTopN != 0 {
		t.Fatalf("expected ComplianceTopN=0 for negative input, got %d", opts.ComplianceTopN)
	}
}
