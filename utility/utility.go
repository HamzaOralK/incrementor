package utility

import (
	"fmt"
	"github.com/rogpeppe/go-internal/semver"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Validate(input string, regex string) bool {
	matched, err := regexp.MatchString(regex, input)
	if err != nil {
		log.Fatalln(err)
	}
	return matched
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
