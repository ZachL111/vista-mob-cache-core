package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 56, Capacity: 100, Latency: 15, Risk: 14, Weight: 5}, wantScore: 79, wantDecision: "review"},
		{name: "case_2", signal: Signal{Demand: 86, Capacity: 76, Latency: 12, Risk: 11, Weight: 12}, wantScore: 159, wantDecision: "review"},
		{name: "case_3", signal: Signal{Demand: 66, Capacity: 82, Latency: 20, Risk: 21, Weight: 10}, wantScore: 27, wantDecision: "review"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
