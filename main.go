package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UTC().UnixNano())
	answers := [2]string{"Don't shout at me, Son! I'm not deaf.", "No, not since Woodstock!"}
	r, err := regexp.Compile("(?i).*bye.*")
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println("Hello honey, how are you?")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if r.MatchString(text) {
			fmt.Println("GOOD BYE!")
			break
		} else if text == strings.ToUpper(text) {
			fmt.Println(answers[rand.Intn(2)])
		} else {
			fmt.Println("Speak up Sonny, I can't hear you!")
		}
	}

}
