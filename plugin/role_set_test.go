package gcpsecrets

import (
	"fmt"
	"testing"
)

func Test_roleSetServiceAccountDisplayName(t *testing.T) {
	res := roleSetServiceAccountDisplayName("display-name-that-is-not-truncated")
	checkDisplayNameLength(t, res)
	expected := fmt.Sprintf(serviceAccountDisplayNameTmpl, "display-name-that-is-not-truncated")
	checkEquals(t, expected, res)

	res = roleSetServiceAccountDisplayName("display-name-that-is-really-long-vault-plugin-secrets-gcp-role-name")
	checkDisplayNameLength(t, res)
	expected = fmt.Sprintf(serviceAccountDisplayNameTmpl, "display-name-that-is-really-long-vault-pl43b18db3")
	checkEquals(t, expected, res)
}

func checkDisplayNameLength(t *testing.T, displayName string) {
	if len(displayName) > serviceAccountDisplayNameMaxLen {
		t.Errorf("expected display name to be less than or equal to %v. actual name '%v'", serviceAccountDisplayNameMaxLen, displayName)
	}
}

func checkEquals(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Errorf("expected '%v' to equal actual '%v'", expected, actual)
	}
}
