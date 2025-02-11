package task1

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int) {
	isPrime := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p <= N; p++ {
		select {
		case <-stop:
			return
		default:
			if isPrime[p] == true {
				for i := p * p; i <= N; i += p {
					isPrime[i] = false
				}
			}
		}
	}

	for p := 2; p <= N; p++ {
		select {
		case <-stop:
			return
		default:
			if isPrime[p] {
				prime_nums <- p
			}
		}
	}
	close(prime_nums)
}
