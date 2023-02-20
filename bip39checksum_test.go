package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{"crowd flip sign buyer truly sail balcony amazing spy garbage cool", []string{"ability", "advance", "air", "alien", "anger", "april", "aspect", "attack", "bachelor", "because", "best", "blind", "bone", "brick", "broccoli", "buyer", "canvas", "captain", "cereal", "champion", "claim", "click", "coast", "conduct", "coral", "crop", "cross", "curtain", "decrease", "demise", "detect", "dinosaur", "donkey", "dry", "dust", "electric", "enable", "enhance", "ethics", "excess", "exotic", "fatigue", "fever", "fish", "flight", "forum", "foster", "garlic", "gasp", "glad", "govern", "gun", "harvest", "hello", "horror", "humor", "icon", "inhale", "insect", "involve", "joke", "know", "ladder", "leaf", "liberty", "lobster", "lounge", "maid", "mean", "merry", "mistake", "more", "motor", "nation", "nominee", "north", "often", "olympic", "organ", "panda", "peanut", "permit", "pig", "plunge", "poverty", "prevent", "protect", "purity", "ramp", "ready", "reflect", "resemble", "rigid", "roof", "rule", "say", "scissors", "second", "setup", "shove", "silent", "slogan", "slush", "solid", "speak", "squeeze", "sting", "strike", "submit", "swim", "tag", "task", "ticket", "today", "top", "tree", "trick", "turn", "uncover", "urban", "valley", "video", "vocal", "wage", "wear", "wink", "wolf", "zone"}},
		{"grit swim bronze output various fiscal cement normal warrior wild area hire refuse rely group culture phone cliff unfair curve claw secret reward", []string{"arch", "devote", "fix", "jewel", "lucky", "result", "sea", "wage"}},
	}

	for _, test := range tests {
		got := brute(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("brute(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}
