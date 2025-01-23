package services

var validStates = []string{"open", "closed", "accepted", "investigating"}

func IsValidState(state string) bool {
	for _, s := range validStates {
		if s == state {
			return true
		}
	}
	return false
}
