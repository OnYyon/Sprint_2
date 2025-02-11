## Использование тайм-аута во время выполнения функции
>На предыдущем уроке мы познакомились с тайм-аутами. Они контролируют продолжительность операций и предотвращают их бесконечное выполнение. При работе с тайм-аутами обычно используются горутины и каналы для параллельных вычислений.
>На этом уроке мы узнаем о паттернах работы с тайм-аутами.
Как мы уже обсуждали, тайм-ауты реализует пакет time. Мы можем использовать оттуда функцию AfterFunc. Вот её сигнатуруа:
```go
func AfterFunc(d Duration, f func()) *Timer
```
AfterFunc ждёт, когда истечёт время d, а затем вызывает функцию f в собственной горутине. Она возвращает Timer, который можно использовать для отмены вызова с помощью его метода Stop.

Рассмотрим пример с подробными комментариями:
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//создание канала chan_c
	chan_c := make(chan string)
	//создание потока выполнения горутины
	go func() {
		/**
		  создание фиктивной операции, которая занимает 3 секунды

		  примечание: если вы хотите увидеть сообщение об успешном выполнении инструкций, сократите время до 1 секунды с миллисекундами или увеличьте время в параметре time.AfterFunc; также вместо Sleep() можно добавить операции, которые требуют больше времени
		  **/
		time.Sleep(3 * time.Second)
		chan_c <- "Инструкции выполнены успешно."
	}()
	//создание тайм-аута, который не даёт функции выполняться дольше 2 секунд
	timeout := time.AfterFunc(2*time.Second, func() {
		chan_c <- "Время выполнения истекло."
	})

	result := <-chan_c
	fmt.Println(result)
	timeout.Stop() // Отмена функции тайм-аута
}
```
## Контексты и отмена с тайм-аутом
На предыдущих уроках мы рассматривали пакет context. Вспомним функцию WithTimeout из него:
```go
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```
WithTimeout возвращает WithDeadline(parent, time.Now().Add(timeout)). Отмена этого контекста освобождает связанные с ним ресурсы, поэтому код должен вызвать отмену, как только операции, выполняемые в контексте, завершатся:
```go
func slowOperationWithTimeout(ctx context.Context) (Result, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel() // освобождает ресурсы, если slowOperation завершается до окончания тайм-аута
	return slowOperation(ctx)
}
```
Рассмотрим пример использования контекста с тайм-аутом:
```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//создание контекста WithTimeout, который ограничивает продолжительность в течение 2 секунд
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	//создание канала chan_c
	chan_c := make(chan string)
	//создание потока выполнения горутины
	go func() {
		// создание фиктивной операции, которая занимает 3 секунды с использованием Sleep()
		time.Sleep(3 * time.Second)
		chan_c <- "Инструкции успешно завершены."
	}()

	select {
	case result := <-chan_c:
		fmt.Println(result)
	case <-ctx.Done():
		fmt.Println("Время операции истекло или было отменено.")
	}
}
```
## Паттерн для параллельных запросов с тайм-аутом
Мы можем использовать комбинацию горутин, каналов и оператора select, чтобы обрабатывать тайм-ауты индивидуально для каждого запроса при выполнении нескольких параллельных запросов. Этот паттерн полезен, например, при выполнении параллельных HTTP-запросов с тайм-аутом.

Рассмотрим пример такого подхода:
```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetchURL(url string, c chan string) {
	//создание http-клиента с тайм-аутом
	client := http.Client{
		/**
		  установка тайм-аута для HTTP-запроса; вы можете уменьшить или увеличить это время
		  - при уменьшении вы можете часто сталкиваться с тайм-аутом
          - при увеличении данные будут получены за большее время
		  **/
		Timeout: 5 * time.Second,
	}
	//получение ответа по URL
	resp, err := client.Get(url)
	if err != nil {
		c <- fmt.Sprintf("Ошибка при получении %s: %s", url, err)
		return
	}
	defer resp.Body.Close()
	c <- fmt.Sprintf("Ответ от %s: Статус - %s", url, resp.Status)
}
func main() {
	//создание массива URL-адресов
	urls := []string{
		"https://yandex.ru",
		"https://lyceum.yandex.ru",
		"https://translate.yandex.com",
		// Симуляция несуществующего URL
		"https://ihumaunkabir.com",
	}
	//создание канала C
	c := make(chan string, len(urls))
	//итерация по массиву URL-адресов
	for _, url := range urls {
		go fetchURL(url, c)
	}
	//установка общего тайм-аута для всех запросов
	timeout := time.After(15 * time.Second)
	//итерация до конца массива URL-адресов
	for i := 0; i < len(urls); i++ {
		select {
		case result := <-c:
			fmt.Println(result)
		case <-timeout:
			fmt.Println("Произошел тайм-аут. Прерывание остальных запросов.")
			return
		}
	}
}
```
