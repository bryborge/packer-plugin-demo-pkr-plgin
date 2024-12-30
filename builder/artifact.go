package builder

type Artifact struct {
	// Fields representing your built artifact
}

func (a *Artifact) BuilderId() string {
	return "example"
}

func (a *Artifact) Files() []string {
	return []string{} // Return list of files created
}

func (a *Artifact) Id() string {
	return "example"
}

func (a *Artifact) String() string {
	return "Example artifact"
}

func (a *Artifact) State(name string) interface{} {
	return nil
}

func (a *Artifact) Destroy() error {
	return nil
}
