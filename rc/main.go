package opts

func main() {}

type RuleSet struct{}

func NewRuleSet() *RuleSet {
	return new(RuleSet)
}

func (rs *RuleSet) AddDep(dep1, dep2 string) {
	return
}

func (rs *RuleSet) IsCoherent() (b bool) {
	return
}

func (rs *RuleSet) AddConflict(a, b string) {
	return
}

type Opts struct{}

func New(rs *RuleSet) *Opts {
	return new(Opts)
}

func (o *Opts) Toggle(opt string) {
	return
}

func (o *Opts) StringSlice() (sl []string) {
	return
}
