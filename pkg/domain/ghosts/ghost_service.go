package ghosts

import "math/rand"

func GenerateGhosts() [10]*Ghost {
	var ghosts [10]*Ghost

	for i := 0; i < 10; i += 1 {
		ghosts[i] = makeGhost()
	}

	return ghosts
}

func makeGhost() *Ghost {
	ghostName := randomString(8)
	return NewGhost(ghostName)
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}
	return string(bytes)
}
