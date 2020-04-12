package metrics

import (
	"fmt"
	"regexp"
)

const (
	metricNameMaxLen = 100
	tagNameMaxLen    = 100
)

var (
	metricNameRgx = regexp.MustCompile("[a-z][a-z0-9/_]*")
	tagNameRgx    = regexp.MustCompile("[a-z0-9_]+")
)

func validateMetricName(name string) error {
	if len(name) > metricNameMaxLen {
		return fmt.Errorf("metric name: %s is too large (%d > %d)", name, len(name), metricNameMaxLen)
	}
	if !metricNameRgx.MatchString(name) {
		return fmt.Errorf("invalid metric name: %s, must match: %s", name, metricNameRgx)
	}
	return nil
}

func validateTagName(tag string) error {
	if len(tag) > tagNameMaxLen {
		return fmt.Errorf("tag name: %s is too large (%d > %d)", tag, len(tag), tagNameMaxLen)
	}
	if !tagNameRgx.MatchString(tag) {
		return fmt.Errorf("invalid tag name: %s, must match: %s", tag, tagNameRgx)
	}
	return nil
}

func sanitizeTagValue(value string) string {
	// TODO(zviadm): sanitize tag values to replace unsupported characters and reduce
	// its size if it is too large.
	return value
}
