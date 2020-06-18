package name

import "testing"

func TestName(t *testing.T) {
	want := "Nurali Virani"
	if got := FullName("nurali ", " virani"); got != want {
		t.Errorf("FullName() = %q, want = %q", got, want)
	}
}
