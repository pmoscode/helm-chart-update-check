package utils

import (
	"github.com/Masterminds/semver/v3"
	"testing"
)

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
