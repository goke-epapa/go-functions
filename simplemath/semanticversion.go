package simplemath

import "fmt"

// SemanticVersion is a semvar type
type SemanticVersion struct {
	major, minor, patch int
}

// NewSemanticVersion creates a new SemanticVersion
func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

// String returns a string representation of the SemanticVersion struct
func (sv SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}

// IncrementMajor bumps the major number
func (sv *SemanticVersion) IncrementMajor(){
	sv.major++
}

// IncrementMinor bumps the minor number
func (sv *SemanticVersion) IncrementMinor(){
	sv.minor++
}

// IncrementPatch bumps the patch number
func (sv *SemanticVersion) IncrementPatch(){
	sv.patch++
}