package application

import (
	"time"
)

const DefaultLocation = "Europe/Moscow"

type BuildVersion struct {
	Time   time.Time
	Commit string
}

// These values have to be set via LDFLAGS.
var (
	buildVersionTime   string
	buildVersionCommit string
)

// buildVersion - stores parsed value.
var buildVersion BuildVersion

func GetBuildVersion() (BuildVersion, error) {
	if !buildVersion.Time.IsZero() {
		return buildVersion, nil
	}

	buildVersion.Commit = buildVersionCommit

	var err error

	buildVersion.Time, err = time.Parse(time.RFC3339, buildVersionTime)
	if err != nil {
		buildVersion.Time = time.Now()
	}

	location, err := time.LoadLocation(DefaultLocation)
	if err == nil {
		buildVersion.Time = buildVersion.Time.In(location)
	}

	return buildVersion, nil
}
