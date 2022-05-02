package utility

import (
	"fmt"
	"github.com/rogpeppe/go-internal/semver"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Validate Validates string against the regex provided returns bool
// or exits
func Validate(input string, regex string) bool {
	matched, err := regexp.MatchString(regex, input)
	if err != nil {
		log.Fatalln(err)
	}
	return matched
}

// NormalizeBranchName removes prefix and seperator from the branch name and returns
// semver branch
func NormalizeBranchName(branchName string, prefix string, separator string) string {
	return strings.Replace(branchName, fmt.Sprintf("%s%s", prefix, separator), "", 1)
}

func GetMaxSemver(v string, w string) string {
	return strings.Replace(semver.Max(fmt.Sprintf("v%s", v), fmt.Sprintf("v%s", w)), "v", "", 1)
}

func IncrementPatch(tag string) string {
	semverArray := strings.Split(tag, ".")
	i, _ := strconv.Atoi(semverArray[len(semverArray)-1])
	i = i + 1
	semverArray[len(semverArray)-1] = strconv.Itoa(i)
	return strings.Join(semverArray, ".")
}
