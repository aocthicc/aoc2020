package main

import "testing"

func TestGroupGetUniqueAnswers(t *testing.T) {
	group := Group{People: []Person{{"abc"}, {"abz"}}}
	ans := group.GetUniqueYesAnswersCount()
	if ans != 4 {
		t.Errorf("GetUniqueYesAnswersCount() = %d, want 4", ans)
	}


	group = Group{People: []Person{{"vs"}, {"lqn"}, {"ti"}, {"uvl"}}}
	ans = group.GetUniqueYesAnswersCount()

	if ans != 8 {
		t.Errorf("GetUniqueYesAnswersCount() = %d, want 8", ans)
	}
}

func TestGroup_GetGroupUnanimousAnswers(t *testing.T) {
	group := Group{People: []Person{{"ab"}, {"ac"}}}

	ans := group.GetUnanimousAnswersCount()
	if ans != 1 {
		t.Errorf("GetUniqueYesAnswersCount() = %d, want 1", ans)
	}
}