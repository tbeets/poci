package poci

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func RequireTrue(t *testing.T, b bool) {
	t.Helper()
	if !b {
		t.Fatalf("require true, but got false")
	}
}

func RequireFalse(t *testing.T, b bool) {
	t.Helper()
	if b {
		t.Fatalf("require false, but got true")
	}
}

func RequireNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("require no error, but got: %v", err)
	}
}

func RequireContains(t *testing.T, s string, subStrs ...string) {
	t.Helper()
	for _, subStr := range subStrs {
		if !strings.Contains(s, subStr) {
			t.Fatalf("require %q to be contained in %q", subStr, s)
		}
	}
}

func RequireError(t *testing.T, err error, expected ...error) {
	t.Helper()
	if err == nil {
		t.Fatalf("require error, but got none")
	}
	if len(expected) == 0 {
		return
	}
	// Try to strip nats prefix from Go library if present.
	const natsErrPre = "nats: "
	eStr := err.Error()
	if strings.HasPrefix(eStr, natsErrPre) {
		eStr = strings.Replace(eStr, natsErrPre, _EMPTY_, 1)
	}

	for _, e := range expected {
		if err == e || strings.Contains(eStr, e.Error()) || strings.Contains(e.Error(), eStr) {
			return
		}
	}
	t.Fatalf("Expected one of %v, got '%v'", expected, err)
}

func RequireEqual(t *testing.T, a, b string) {
	t.Helper()
	if strings.Compare(a, b) != 0 {
		t.Fatalf("require equal, but got: %v != %v", a, b)
	}
}

func RequireNotEqual(t *testing.T, a, b [32]byte) {
	t.Helper()
	if bytes.Equal(a[:], b[:]) {
		t.Fatalf("require not equal, but got: %v != %v", a, b)
	}
}

func RequireLen(t *testing.T, a, b int) {
	t.Helper()
	if a != b {
		t.Fatalf("require len, but got: %v != %v", a, b)
	}
}

func CheckForErr(totalWait, sleepDur time.Duration, f func() error) error {
	return checkForErr(totalWait, sleepDur, f)
}

func checkForErr(totalWait, sleepDur time.Duration, f func() error) error {
	timeout := time.Now().Add(totalWait)
	var err error
	for time.Now().Before(timeout) {
		err = f()
		if err == nil {
			return nil
		}
		time.Sleep(sleepDur)
	}
	return err
}

func CheckFor(t *testing.T, totalWait, sleepDur time.Duration, f func() error) {
	t.Helper()
	err := checkForErr(totalWait, sleepDur, f)
	if err != nil {
		t.Fatal(err.Error())
	}
}

var (
	_ = RequireTrue
	_ = RequireFalse
	_ = RequireNoError
	_ = RequireContains
	_ = RequireError
	_ = RequireEqual
	_ = RequireNotEqual
	_ = RequireLen
	_ = CheckForErr
	_ = CheckFor
)
