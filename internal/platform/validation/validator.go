package validation

type Validator struct {
	errs map[string]string
}

func New() *Validator {
	return &Validator{errs: make(map[string]string)}
}

func (v *Validator) Errors() map[string]string {
	return v.errs
}

func (v *Validator) HasErrors() bool {
	return len(v.errs) > 0
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.errs[key]; !exists {
		v.errs[key] = message
	}
}
