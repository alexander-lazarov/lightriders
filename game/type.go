package game

const (
	// Hotseat - 2 players playing locally on the same machine
	Hotseat = iota

	// Server - hosting a game
	Server

	// Client - joining a game
	Client

	// Unknown - cannot determine
	Unknown
)

// Type of the game
type Type byte

// GetType the game type from a string
func GetType(args []string) (Type, string) {
	if len(args) == 1 {
		return Hotseat, ""
	}

	switch args[1] {
	case "hotseat", "":
		return Hotseat, ""
	case "server":
		return Server, ""
	case "client":
		if len(args) < 3 {
			return Unknown, ""
		}

		return Client, args[2]
	default:
		return Unknown, ""
	}
}
