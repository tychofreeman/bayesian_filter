package bayesian

import . "testing"
import . "big"

// Pulled from en.wikipedia.org/wiki/Bayes%27_theorem
func TestCalculateProbOfStudentIsGirlWhenStudentWearsPants(t *T) {
	probGirlsWearPants := NewRat(1, 2)
	probRandomStudentIsGirl := NewRat(2, 5)
	probRandomStudentIsBoy := NewRat(1, 1).Sub(NewRat(1, 1), probRandomStudentIsGirl)
	probBoysWearPants := NewRat(1, 1)

	probRandomStudentIsGirlWearingPants := NewRat(1, 1).Mul(probGirlsWearPants, probRandomStudentIsGirl)
	probRandomStudentIsBoyWearingPants := NewRat(1, 1).Mul(probRandomStudentIsBoy, probBoysWearPants)

	// This should be .8, but I'd like to calculate it instead of defining it...
	probRandomStudentWearsPants :=
		NewRat(1, 1).Add(probRandomStudentIsGirlWearingPants, probRandomStudentIsBoyWearingPants)

	probRandomStudentIsGirlIfStudentIsWearingPants :=
		(&BayesRule{probGirlsWearPants, probRandomStudentIsGirl, probRandomStudentWearsPants}).Calculate()

	expected := NewRat(1, 4)
	if probRandomStudentIsGirlIfStudentIsWearingPants.Cmp(expected) != 0 {
		t.Errorf("Expected %v, but got %v", expected, probRandomStudentIsGirlIfStudentIsWearingPants)
	}
}
