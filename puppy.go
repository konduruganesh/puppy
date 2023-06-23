package puppy

import (
	dognew "github.com/konduruganesh/dognew"
)

func Bark() string {
	return "Woff!"
}

func Barks() string {
	return "Woof! woof! woof!"
}

func BigBark() {
	dognew.WhenGrownUp("Big Bark!!!!!")
}
