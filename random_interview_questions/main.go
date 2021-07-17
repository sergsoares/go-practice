package main

import  (
	"os"
	"strconv"
	"fmt"
	"math/rand"
)

func fetch_answers() []string {
	questions := []string{"1. Describe a time when your boss was wrong. How did you handle the situation?", "2. How would you feel about reporting to a person younger than you?", "3. Describe a time you went above and beyond at work.", "4. Tell me about the last mistake you made.", "5. What do you want to accomplish in the first 30 days of this job?", "6. Describe a time you got angry at work.", "7. Describe a time when you had to give a person difficult feedback.", "8. Describe a time when you disagreed with your boss.", "9. Would you ever lie for a company?", "10. Tell me about how you dealt with a difficult challenge in the workplace.", "11. What do you really think about your previous boss?", "12. What has been the most rewarding experience of your career thus far?", "13. How would you deal with an angry or irate customer?", "14. Describe a time you chose to not help a teammate.", "15. Describe a time you went out of your way to help somebody.", "16. Describe a time when your work was criticized?", "17. What do you want to accomplish in the first 90 days of this job?", "18. Do you think you could have done better in your last job?", "19. How would you fire someone?"}
	return questions
}

func main() {
	args := os.Args[1:]
	value, err  := strconv.Atoi(args[0])

	if err != nil {
		panic(err)
	}

	for i := 1; i <= value; i++ {
		pos := rand.Intn(18)
		fmt.Println(fetch_answers()[pos])
	}
}
