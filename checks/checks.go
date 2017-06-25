package main

// This is incorrectly documented
type SelfTest struct {
}

type SelfTest1 struct {
}
type SelfTest2 struct {
}

func (test *SelfTest) MissingDocs(foo, bar string) {
}

// MissingParamDocs is here
func (test *SelfTest) MissingParamDocs(foo, bar string) {
}

// This is incorrectly documented
func (test *SelfTest) IncorrectDocs(foo, bar string) {
}
