package request

type Parameters []Parameter

func (ps Parameters) String(name string) string {
	for _, p := range ps {
		if p.Name == name && p.Type == STRING {
			return p.Value.(string)
		}
	}
	return ""
}

func (ps Parameters) StringList(name string) []string {
	s := make([]string, 0)
	for _, p := range ps {
		if p.Name == name && p.Type == STRINGLIST {
			for _, val := range p.Value.([]string) {
				s = append(s, val)
			}
			return s
		}
	}
	return s
}

func (ps Parameters) Int(name string) int {
	for _, p := range ps {
		if p.Name == name && p.Type == INT {
			return p.Value.(int)
		}
	}
	return 0
}

func (ps Parameters) Uint64(name string) uint64 {
	return uint64(ps.Int(name))
}

type Parameter struct {
	Name string
	Type ParameterType

	DefaultValue  string
	AllowedValues []interface{}

	List          bool
	ListSeparator string
	Nullable      bool

	Value interface{}
}

func (p Parameter) IsAllowed(value interface{}) bool {
	if p.AllowedValues == nil || len(p.AllowedValues) == 0 {
		return true
	}

	for _, val := range p.AllowedValues {
		if value == val {
			return true
		}
	}

	return false
}
type ParameterType int
const (
	INT         ParameterType = iota
	INTLIST
	STRING
	STRINGLIST
)