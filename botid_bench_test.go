package botid

import "testing"

// BenchmarkIsOurs measures the prefix-matching hot path.
// Called per robots.txt check to determine if a UA belongs to us.
func BenchmarkIsOurs(b *testing.B) {
	agents := []string{
		"cl0q-node",           // exact match
		"cl0q-node/1.0",       // prefix match
		"cl0q-crawler",        // different known agent
		"mozilla/5.0",         // non-match (common case)
		"googlebot",           // non-match
		"",                    // empty
		"  CLOQ-NODE  ",       // case + whitespace normalization
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsOurs(agents[i%len(agents)])
	}
}

// BenchmarkIsOurs_Match measures the best-case (early match).
func BenchmarkIsOurs_Match(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsOurs("cl0q-node")
	}
}

// BenchmarkIsOurs_NoMatch measures the worst-case (full scan, no match).
func BenchmarkIsOurs_NoMatch(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsOurs("some-random-bot/1.0")
	}
}

// BenchmarkUserAgent measures UA string building.
// Called per HTTP request.
func BenchmarkUserAgent(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = UserAgent("test-node-123")
	}
}

// BenchmarkUserAgent_Invalid measures the validation failure path.
func BenchmarkUserAgent_Invalid(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = UserAgent("INVALID!!!")
	}
}
