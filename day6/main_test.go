package main

import "testing"

func TestGroupGetUniqueAnswers(t *testing.T) {
	group := Group{People: []Person{{"abc"}, {"abz"}}}
	ans := group.GetTotalUniqueYesAnswers()
	if ans != 4 {
		t.Errorf("GetTotalUniqueYesAnswers(\"abc\", \"abz\") = %d, want 4", ans)
	}


	group = Group{People: []Person{{"vs"}, {"lqn"}, {"ti"}, {"uvl"}}}
	ans = group.GetTotalUniqueYesAnswers()

	if ans != 8 {
		t.Errorf("GetTotalUniqueYesAnswers() = %d, want 8", ans)
	}
}
