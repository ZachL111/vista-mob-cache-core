package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 56, Capacity: 100, Latency: 15, Risk: 14, Weight: 5}
	if got := Score(signal); got != 79 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 86, Capacity: 76, Latency: 12, Risk: 11, Weight: 12}
	if got := Score(signal); got != 159 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 66, Capacity: 82, Latency: 20, Risk: 21, Weight: 10}
	if got := Score(signal); got != 27 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
}
