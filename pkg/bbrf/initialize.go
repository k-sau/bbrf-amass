package bbrf

func Initialize(type_, program string) {
	if type_ == "service" {
		serviceInitialize(program)
	}
}
