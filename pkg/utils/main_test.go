package utils

import (
	"github.com/Masterminds/semver/v3"
	"testing"
)

func TestGetExcludedVersionsEmpty(t *testing.T) {
	result, err := GetExcludedVersions("")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != nil {
		t.Errorf("expected nil, got %v", result)
	}
}

func TestGetExcludedVersionsInvalidVersion(t *testing.T) {
	_, err := GetExcludedVersions("not-a-semver")
	if err == nil {
		t.Fatal("expected error for invalid version, got nil")
	}
}

func TestGetExcludedVersionsSimpleEmpty(t *testing.T) {
	result := GetExcludedVersionsSimple("")
	if result != nil {
		t.Errorf("expected nil, got %v", result)
	}
}

func TestGetExcludedVersionsSimpleSingle(t *testing.T) {
	result := GetExcludedVersionsSimple("1.2.3")
	if len(result) != 1 {
		t.Fatalf("expected 1 element, got %d", len(result))
	}
	if result[0] != "1.2.3" {
		t.Errorf("expected '1.2.3', got %q", result[0])
	}
}

func TestGetExcludedVersionsSimpleMultiple(t *testing.T) {
	result := GetExcludedVersionsSimple("1.2.3,^2.0.0,~3.0.0-0")
	if len(result) != 3 {
		t.Fatalf("expected 3 elements, got %d", len(result))
	}
	expected := []string{"1.2.3", "^2.0.0", "~3.0.0-0"}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("element %d: expected %q, got %q", i, v, result[i])
		}
	}
}

func TestGetExcludedVersionsSimple(t *testing.T) {
	versionStr := "1.2.3"
	versionSemver := semver.New(1, 2, 3, "", "")

	versions, err := GetExcludedVersions(versionStr)
	if err != nil {
		t.Errorf("GetExcludedVersions failed: %v", err)
	}

	versionsLen := len(versions)
	if versionsLen != 1 {
		t.Fatalf("Resulting slice must equal 1 but is %d", versionsLen)
	}

	if versions[0].Compare(versionSemver) != 0 {
		t.Fatalf("Resulting slice must equal VersionSemver (%s) but is %s", versionSemver.String(), versions[0].String())
	}
}

func TestGetExcludedVersionsMultiple(t *testing.T) {
	versionStr := "1.2.3,3.4.0,5.2"
	versionSemver1 := semver.New(1, 2, 3, "", "")
	versionSemver2 := semver.New(3, 4, 0, "", "")
	versionSemver3 := semver.New(5, 2, 0, "", "")

	versions, err := GetExcludedVersions(versionStr)
	if err != nil {
		t.Errorf("GetExcludedVersions failed: %v", err)
	}

	versionsLen := len(versions)
	if versionsLen != 3 {
		t.Fatalf("Resulting slice must equal 3 but is %d", versionsLen)
	}

	if versions[0].Compare(versionSemver1) != 0 {
		t.Fatalf("Resulting slice must equal VersionSemver (%s) but is %s", versionSemver1.String(), versions[0].String())
	}
	if versions[1].Compare(versionSemver2) != 0 {
		t.Fatalf("Resulting slice must equal VersionSemver (%s) but is %s", versionSemver2.String(), versions[1].String())
	}
	if versions[2].Compare(versionSemver3) != 0 {
		t.Fatalf("Resulting slice must equal VersionSemver (%s) but is %s", versionSemver3.String(), versions[2].String())
	}
}
