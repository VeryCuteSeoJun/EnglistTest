package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func clearconsole() {
	for i := 0; i < 500; i++ {
		fmt.Printf("\n\n\n\n\n")
	}
}

func input() string {
	kbReader := bufio.NewReader(os.Stdin)
	str, err := kbReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	str = strings.TrimSpace(str)
	return str

}

func main() {
	test := 0
	fmt.Println("시험을 볼 종목을 선택해주세요. (ex: 영어1과 p15를 시험볼려면 0을 입력)")
	fmt.Println("\n영어1과 p14[0]\n영어1과 p15[1]\n영어1과 p16[2]\n영어1과 p17[3]")
	fmt.Println("\n영어2과 p32[4]\n영어2과 p33[5]\n영어2과 p34[6]\n영어2과 p35[7]")
	fmt.Println("\n영어3과 p50 (Intruoducing) [8]\n영어3과 p50 (Minji - Selfies in the Past) [9]\n영어3과 p51 (Yunho - Fun Places for Selfies) [10]\n영어3과 p52 (Jihun - Selfie Safety) [11]\n영어3과 p53 (Soyun - Selfies for a Better School Life) [12]")
	fmt.Scanln(&test)
	answers := []string{}
	switch test {
	case 0:
		answers = []string{"Seho and Jihun were talking in the hallway when Dami came over.",
			"\"Happy birthday!\" she said to Seho.",
			"\"Here. They're from my dad.\"",
			"\"Wow, two KBL tickets! Thanks!\"",
			"\"Who are you going to take with you?\" Dami asked.",
			"\"Minjun. He took me to a soccer game before. So, it's time to pay him back.\"",
			"\"You know what?\" Jihun cut in.",
			"\"Minjun isn't a fan of basketball. But I am!\"",
			"\"Well, I'll ask him first anyway,\" replied Seho.",
			"\"He won't go with you. Trust me,\" said Jihun.",
			"\"Who is this guy?\"",
			"Dami thought to herself, \"He wants Minjun's ticket.\""}
	case 1:
		answers = []string{"\"Oh! There's the bell. See you later,\" said Dami.",
			"She hurried to class.",
			"\"Come on, Jihun,\" said Seho, and he started to run.",
			"At the corner, Seho bumped into someone.",
			"\"Sorry!\" he said and continued to run.",
			"Just then, Jihun saw something on the floor.",
			"\"Wait, Seho!\" he said, but Seho was not there."}
	case 2:
		answers = []string{"After class, Seho went to Dami and said, \"I can't find one of my tickets. Did you happen to see it?\"",
			"\"No,\" she answered.",
			"\"Isn't it in your bag?\"",
			"\"No, it's not there. I think I lost it,\" said Seho.",
			"On her way home, Dami saw Jihun.",
			"He had the ticket in his hand.",
			"Dami got angry and said, \"Hey! Why do you...?\"",
			"Just then, Jihun saw Seho and shouted, \"Seho! I found a ticket in the hallway. I think it's yours.\"",
			"\"Thanks! I was looking for that!\" said Seho."}
	case 3:
		answers = []string{"\"He's not so bad,\" Dami thought.",
			"\"So, what were you saying? Do you have something to say, Dami?\" asked Jihun.",
			"\"Um, how about going to the school basketball game with me this Friday? It's the finals.\"",
			"Jihun looked really pleased.",
			"\"I'd love to!\""}
	case 4:
		answers = []string{"Bear was a black and brown cat with green eyes.",
			"He lived with a boy, Ryan.",
			"Ryan always thought that \"Bear\" was a perfect name for the cat because he had a black spot in the shape of a bear.",
			"Bear liked to go outside every morning and run after butterflies.",
			"He always came home just in time for dinner.",
			"Five blocks away, Max the cat lived with a girl, Sheila.",
			"When Sheila moved to this town last month, she was lonely.",
			"She had no friends there.",
			"But, after Max followed her home, he became a good friend to her."}
	case 5:
		answers = []string{"One day, Sheila saw Max sitting under the desk.",
			"He was making a strange sound.",
			"\"What's wrong?\" asked Sheila.",
			"She looked at him closely and found a bad cut on his leg.",
			"She took him to the animal hospital.",
			"The doctor said, \"He will get better if he gets enough rest. Keep him inside for a week.\"",
			"That night, at Ryan's house, there was no Bear.",
			"Ryan checked outside, but he couldn't find him.",
			"He made posters and put them up around town.",
			"A third night passed.",
			"Still no Bear."}
	case 6:
		answers = []string{"When Sheila was walking near her house, she saw a poster about the lost cat.",
			"She read it closely, and her eyes got big.",
			"\"This cat looks exactly like Max. It's so strange.\"",
			"She hurried home.",
			"\"Come on, Max! Let's go!\"",
			"She took him to the address on the poster.",
			"\"Ding-Dong.\"",
			"When Ryan heard the doorbell ring, he ran to the door and opened it.",
			"\"Bear, you're back!\"",
			"Ryan cried.",
			"Max jumped up into Ryan's arms."}
	case 7:
		answers = []string{"\"Let me guess,\" said Sheila.",
			"\"Your cat comes home only in the evenings, doesn't he?\"",
			"Ryan nodded.",
			"\"And you lost him last Friday, didn't you?\"",
			"Sheila said.",
			"\"Because this is my cat, too, and he usually comes to my home only during the day.\"",
			"\"Our cat has two families!\" said Ryan.",
			"\"Hey, if you have time, please come in and have some cookies.\"",
			"\"Sure,\" said Sheila.",
			"\"Thank you, Max,\" she thought.",
			"\"I met a good neighbor thanks to you!\""}
	case 8:
		answers = []string{"Have you ever heard of a \"selfie\"?",
			"When you take a photograph of yourself, it's a selfie.",
			"The students from Minji's photo club have searched for information about selfies for one month.",
			"Here are some of their presentations about selfies."}
	case 9:
		answers = []string{"Did people in the past take selfies?",
			"Though it wasn't easy at that time, the answer is yes.",
			"Look at this photo of Princess Anastasia.",
			"She used a mirror to take a picture of herself.",
			"She looks nervous.",
			"Can you guess why?",
			"Well, I think it was her first selfie.",
			"And it was probably the world's first teenage selfie ever."}
	case 10:
		answers = []string{"You can take selfies at world-famous places like Big Ben and the Leaning Tower of Pisa.",
			"To take great pictures, just do fun poses and use camera tricks.",
			"You can also visit special museums to take fun selfies.",
			"For example, there is a famous selfie museum in the Philippines.",
			"It has special spots to take selfies.",
			"You can touch the paintings and even step inside them.",
			"Look at the following pictures.",
			"Though the boys are not really riding horses, it looks like they are.",
			"Though the man is just holding a big brush, it looks like he is painting the Mona Lisa.",
			"Selfie museums exist in Korea, too.",
			"I have visited one in Chuncheon before.",
			"Why don't you go there yourself?"}
	case 11:
		answers = []string{"These selfies look great, but were they a good idea?",
			"I don't think so.",
			"They don't look safe.",
			"You should take special care when you take selfies in the wild or at high places like these.",
			"A monkey could bite you at any time, or you could fall.",
			"Here are some safety tips:",
			"1. Don't take selfies while you're walking.",
			"2. Do not pose with or near wild animals.",
			"3. Never take selfies in dangerous places."}
	case 12:
		answers = []string{"I think we can use selfies to make a better school life.",
			"We can do good things at school and take selfies.",
			"Then we can post the photos on our school website.",
			"I've watered the plants and flowers at school for one month.",
			"I've also helped the teacher at the school library many times.",
			"Look at my selfies of those things.",
			"How about joining me to create a better school life?"}
	default:
		fmt.Println("올바르지 않은 선택지입니다. 다시 선택해주세요")
		main()
		return
	}

	wrong := make(map[string]string)
	point := 0
	sepoint := 0
	skips := 0

	// 시험 시작
	clearconsole()
	for num, i := range answers {
		fmt.Printf("Question [%d/%d]\n", num+1, len(answers))
		fmt.Println("Write answer (press ENTER when done): ")
		inputstr := input()
		if inputstr == i {
			point += 1
			fmt.Println("CORRECT!\n\n")
		} else if inputstr == "S" || inputstr == "Skip" || inputstr == "skip" || inputstr == "SKIP" || inputstr == "s" || inputstr == "X" || inputstr == "x" {
			skips += 1
			fmt.Println("SKIP\n\n")
			continue
		} else if inputstr == "RESTART" || inputstr == "restart" || inputstr == "R" || inputstr == "r" || inputstr == "Restart" {
			main()
			return
		} else {
			wrong[strconv.Itoa(num+1)] = i
			fmt.Println("WRONG ANSWER")
			fmt.Print("Answer is : " + i + "\n\n")
		}
	}

	fmt.Printf("------------------------\nScore: %d/%d\n\n\n", point, len(answers)-skips)
	if len(wrong) == 0 {
		fmt.Println("틀린 문제가 없으므로 재시험을 치루지 않습니다.")
		return
	} else {
		clearconsole()
		fmt.Println("틀린 문제가 있으므로 재시험을 실시합니다.")
		// 재시험
		for num, i := range wrong {
			fmt.Printf("Question %v\n", num)
			fmt.Println("Write answer (press ENTER when done): ")
			inputstr := input()
			if inputstr == i {
				sepoint += 1
				fmt.Println("CORRECT!\n\n")
			} else {
				fmt.Println("WRONG ANSWER \n\n")
			}
		}
	}
	fmt.Printf("첫번째 시험에서 맞춘 문제의 수 : %d/%d \n", point, len(answers)-skips)
	fmt.Printf("재시험에서 맞춘 문제의 수 : %d/%d \n", sepoint, len(wrong))
	time.Sleep(time.Second * 5)
}
