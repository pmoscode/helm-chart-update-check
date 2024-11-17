package utils

import (
	"github.com/Masterminds/semver/v3"
	"strings"
)

func GetExcludedVersions(versionString string) ([]*semver.Version, error) {
	if versionString == "" {
		return nil, nil
	}

	versionSplit := strings.Split(versionString, ",")
	versions := make([]*semver.Version, len(versionSplit))

	for idx, version := range versionSplit {
		semverVersion, err := semver.NewVersion(strings.TrimSpace(version))
		if err != nil {
			return nil, err
		}

		versions[idx] = semverVersion
	}

	return versions, nil
}

func GetExcludedVersionsSimple(versionString string) []string {
	if versionString == "" {
		return nil
	}

	return strings.Split(versionString, ",")
}
