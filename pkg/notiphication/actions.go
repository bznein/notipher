package notiphication

type Actions map[string]func()

func (a Actions) Keys() []string {
	keys := []string{}
	for k, _ := range a {
		keys = append(keys, k)
	}
	return keys
}
