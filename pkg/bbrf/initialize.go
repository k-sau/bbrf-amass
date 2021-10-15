package bbrf

func Initialize(type_, program string, wildcard bool) {
	if type_ == "service" {
		serviceInitialize(program)
	} else if type_ == "urls" {
		urlsInitialize(program)
	} else if type_ == "unresolved" {
		unresolvedInitialize(program, wildcard)
	}
}
