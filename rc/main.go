package opts

func main() {}

type RuleSet struct {
	dependencies map[string][]string
}

func NewRuleSet() *RuleSet {
	rs := new(RuleSet)
	rs.dependencies = map[string][]string{}
	return rs
}

func (rs *RuleSet) AddDep(dep1, dep2 string) {
	rs.dependencies[dep1] = append(rs.dependencies[dep1], dep2)
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
