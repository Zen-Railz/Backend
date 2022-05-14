package errorr

func (e *entity) Is(target error) bool {
	t, ok := target.(*entity)
	if ok {
		return e.code == t.code
	} else {
		return false
	}
}
