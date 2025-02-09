package task3

import "time"

func QuizRunner(questions, answers []string, answerCh chan string) int {
	correctAnsw := 0
	for i := 0; i < len(questions); i++ {
		select {
		case ans := <-answerCh:
			if ans == answers[i] {
				correctAnsw++
			}
		case <-time.After(1 * time.Second):
			continue
		}
	}
	return correctAnsw
}
