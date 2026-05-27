package defaults

// Setter is an interface for setting default values
type Setter interface {
	SetDefaults()
}

func callSetter(v interface{}) { _ = "STUB: not implemented"; return }
