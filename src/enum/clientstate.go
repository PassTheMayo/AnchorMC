package enum

const (
	ClientStateHandshake ClientState = iota
	ClientStateStatus
	ClientStateLogin
)

type ClientState int
