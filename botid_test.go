package botid

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestUserAgent(t *testing.T) {
	ua, err := UserAgent("alice-01")
	if err != nil {
		t.Fatal(err)
	}
	if ua != "Mozilla/5.0 (compatible; cl0q-node/1.0; operator=alice-01; +https://cl0q.com/bot)" {
		t.Fatalf("unexpected UA %q", ua)
	}
	for _, w := range []string{"grabber", "scanner"} { // words WAFs treat as attack tools
		if strings.Contains(strings.ToLower(ua), w) {
			t.Fatalf("UA must not contain %q", w)
		}
	}
}

func TestUserAgentRejectsBadIDs(t *testing.T) {
	for _, id := range []string{"", "a", "Alice", "has space", "x;y", "a)b", "-lead", "_lead", strings.Repeat("a", 33)} {
		if _, err := UserAgent(id); err == nil {
			t.Errorf("node id %q should be rejected", id)
		}
	}
}

func TestUserAgentAcceptsServerIDs(t *testing.T) {
	// The live server mints nd_-prefixed ids (see PROTOCOL.md); the client must accept them.
	for _, id := range []string{"nd_d84cd67b", "nd_7f3a9c", "alice-01", "a_b-c1"} {
		if _, err := UserAgent(id); err != nil {
			t.Errorf("node id %q should be accepted: %v", id, err)
		}
	}
}

func TestUserAgentAcceptsBoundaries(t *testing.T) {
	// Exact 2-char and 32-char ids are the acceptance boundaries of idRE.
	for _, id := range []string{"ab", "a1", strings.Repeat("a", 32), "a" + strings.Repeat("-", 31)} {
		if _, err := UserAgent(id); err != nil {
			t.Errorf("boundary node id %q should be accepted: %v", id, err)
		}
	}
}

func TestUserAgentErrorWraps(t *testing.T) {
	_, err := UserAgent("bad id!")
	if !errors.Is(err, ErrInvalidNodeID) {
		t.Fatalf("expected ErrInvalidNodeID, got %v", err)
	}
}

func FuzzUserAgent(f *testing.F) {
	for _, id := range []string{"alice-01", "ab", "", "a", "-lead", "_lead", strings.Repeat("a", 33), "nd_d84cd67b"} {
		f.Add(id)
	}
	f.Fuzz(func(t *testing.T, id string) {
		ua, err := UserAgent(id)
		if err != nil {
			if !errors.Is(err, ErrInvalidNodeID) {
				t.Fatalf("UserAgent(%q) error is not ErrInvalidNodeID: %v", id, err)
			}
			return
		}
		if !strings.Contains(ua, "operator="+id) {
			t.Fatalf("UserAgent(%q) = %q does not carry the id", id, ua)
		}
	})
}

func FuzzIsOurs(f *testing.F) {
	for _, a := range []string{"cl0q-node", "googlebot", "", "cl0q", "cl0q-favicon-grabber", "*"} {
		f.Add(a)
	}
	f.Fuzz(func(t *testing.T, agent string) {
		_ = IsOurs(agent) // must never panic, whatever the input
	})
}

func ExampleUserAgent() {
	ua, err := UserAgent("alice-01")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(ua)
	// Output: Mozilla/5.0 (compatible; cl0q-node/1.0; operator=alice-01; +https://cl0q.com/bot)
}

func TestIsOurs(t *testing.T) {
	for _, v := range []string{"cl0q-node", "CL0Q-Node", " cl0q ", "cl0q-crawler", "cl0q-discovery", "cl0q-favicon", "cl0q-favicon-grabber"} {
		if !IsOurs(v) {
			t.Errorf("%q should address us", v)
		}
	}
	for _, v := range []string{"*", "googlebot", "cl0q-favicon-x", "cl0qbot", ""} {
		if IsOurs(v) {
			t.Errorf("%q should not address us", v)
		}
	}
}
