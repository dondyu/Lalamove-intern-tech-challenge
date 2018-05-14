package main

import (
	"testing"

	"github.com/coreos/go-semver/semver"
)

func stringToVersionSlice(stringSlice []string) []*semver.Version {
	versionSlice := make([]*semver.Version, len(stringSlice))
	for i, versionString := range stringSlice {
		versionSlice[i] = semver.New(versionString)
	}
	return versionSlice
}

func versionToStringSlice(versionSlice []*semver.Version) []string {
	stringSlice := make([]string, len(versionSlice))
	for i, version := range versionSlice {
		stringSlice[i] = version.String()
	}
	return stringSlice
}

func TestLatestVersions(t *testing.T) {
	testCases := []struct {
		versionSlice   []string
		expectedResult []string
		minVersion     *semver.Version
	}{
		{
			versionSlice:   []string{"1.8.11", "1.9.6", "1.10.1", "1.9.5", "1.8.10", "1.10.0", "1.7.14", "1.8.9", "1.9.5"},
			expectedResult: []string{"1.10.1", "1.9.6", "1.8.11"},
			minVersion:     semver.New("1.8.0"),
		},
		{
			versionSlice:   []string{"1.8.11", "1.9.6", "1.10.1", "1.9.5", "1.8.10", "1.10.0", "1.7.14", "1.8.9", "1.9.5"},
			expectedResult: []string{"1.10.1", "1.9.6"},
			minVersion:     semver.New("1.8.12"),
		},
		{
			versionSlice:   []string{"1.10.1", "1.9.5", "1.8.10", "1.10.0", "1.7.14", "1.8.9", "1.9.5"},
			expectedResult: []string{"1.10.1"},
			minVersion:     semver.New("1.10.0"),
		},
		{
			versionSlice:   []string{"2.2.1", "2.2.0"},
			expectedResult: []string{"2.2.1"},
			minVersion:     semver.New("2.2.1"),
		},
		// Implement more relevant test cases here, if you can think of any
		//Change of major versions
		{
			versionSlice:   []string{"2.1.1", "1.7.4", "2.2.13", "5.6.2", "2.1.3", "1.8.3", "3.2.5", "1.8.7", "1.9.10", "3.2.2", "2.2.4", "2.7.6", "5.6.0"},
			expectedResult: []string{"5.6.2", "3.2.5", "2.7.6", "2.2.13", "2.1.3", "1.9.10", "1.8.7"},
			minVersion:     semver.New("1.8.3"),
		},
		//All versions below minVersion
		{
			versionSlice:   []string{"2.1.1", "1.7.4", "2.2.13", "2.1.3", "1.8.3", "1.8.7", "1.9.10", "2.1.3", "2.2.4", "2.7.6"},
			expectedResult: []string{},
			minVersion:     semver.New("4.8.3"),
		},
		//Repeated versions
		{
			versionSlice:   []string{"1.8.4", "1.9.7", "1.8.4", "1.12.3", "1.14.5", "1.9.6", "1.9.7", "1.14.5"},
			expectedResult: []string{"1.14.5", "1.12.3", "1.9.7", "1.8.4"},
			minVersion:     semver.New("1.8.3"),
		},
		//Patches with letters
		{
			versionSlice:   []string{"1.5.0-alpha.1", "1.5.4", "1.6.3-alpha.1", "1.6.3"},
			expectedResult: []string{"1.6.3-alpha.1", "1.5.4"},
			minVersion:     semver.New("1.5.0"),
		},
	}

	test := func(versionData []string, expectedResult []string, minVersion *semver.Version) {
		stringSlice := versionToStringSlice(LatestVersions(stringToVersionSlice(versionData), minVersion))
		for i, versionString := range stringSlice {
			if versionString != expectedResult[i] {
				t.Errorf("Received %s, expected %s", stringSlice, expectedResult)
				return
			}
		}
	}

	for _, testValues := range testCases {
		test(testValues.versionSlice, testValues.expectedResult, testValues.minVersion)
	}
}
