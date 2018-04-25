package opts

func main() {}

type RuleSet struct {
	dependencies map[string][]string
	conflicts    map[string][]string
}

func NewRuleSet() *RuleSet {
	rs := new(RuleSet)
	rs.conflicts = map[string][]string{}
	rs.dependencies = map[string][]string{}
	return rs
}

func (rs *RuleSet) AddDep(dep1, dep2 string) {
	rs.dependencies[dep1] = append(rs.dependencies[dep1], dep2)
	return
}

// IsCoherent Implement the algorithm that
// checks that an instance of that data structure
// is coherent, that is, that no option can depend,
// directly or indirectly, on another option and
// also be mutually exclusive with it.
func (rs *RuleSet) IsCoherent() (b bool) {
	for part, dependencies := range rs.dependencies {
		mx := isMutualExclusive(part, dependencies, rs.dependencies)
		if mx {
			return
		}
	}
	return true
}

// terrible, basically I want to find out if in deps slice
// has a the key of depsMap and if p (other map key)
// is equals to a element in depsP2 map
// this is where It gets to mutex
func isMutualExclusive(p string, deps []string, depsMap map[string][]string) (mx bool) {
	for nextPart, depsP2 := range depsMap {
		if nextPart == p {
			continue
		}
		for _, element := range deps {
			if element == nextPart {
				for _, el2 := range depsP2 {
					if el2 == p {
						return true
					}
				}
			}
		}
	}
	return
}

func (rs *RuleSet) AddConflict(part1, part2 string) {
	rs.conflicts[part1] = append(rs.conflicts[part1], part2)
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
