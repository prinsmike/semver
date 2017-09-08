package semver

import "fmt"

type Version struct {
	Major      int
	Minor      int
	Patch      int
	PreRelease string
	BuildMeta  string
}

func NewSemVer(major, minor, patch int, pre, build string) *Version {
	return &Version{
		Major:      major,
		Minor:      minor,
		Patch:      patch,
		PreRelease: pre,
		BuildMeta:  build,
	}
}

func (v Version) String() string {

	if v.PreRelease != "" && v.PreRelease[:1] != "-" {
		v.PreRelease = "-" + v.PreRelease
	}

	if v.BuildMeta != "" && v.BuildMeta[:1] != "+" {
		v.BuildMeta = "+" + v.BuildMeta
	}

	return fmt.Sprintf("%d.%d.%d%s%s",
		v.Major, v.Minor, v.Patch, v.PreRelease, v.BuildMeta)
}
