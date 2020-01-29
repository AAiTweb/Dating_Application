package form


import "net/url"

type Input struct {
	Values  url.Values
	VErrors ValidationErrors

}



func (inVal *Input) Required(fields ...string) {
	for _, f := range fields {
		value := inVal.Values.Get(f)
		if value == "" {
			inVal.VErrors.Add(f, "This field is required field")
		}
	}
}
func (inVal *Input) Valid() bool {
	return len(inVal.VErrors) == 0
}