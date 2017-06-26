package fakes

type FakeStage struct {
	PerformCalls []*PerformCall
	SubStages    []*FakeStage
}

type PerformCall struct {
	Name      string
	Error     error
	SkipError error
	Stage     *FakeStage
}

func NewFakeStage() *FakeStage {
	return &FakeStage{}
}
