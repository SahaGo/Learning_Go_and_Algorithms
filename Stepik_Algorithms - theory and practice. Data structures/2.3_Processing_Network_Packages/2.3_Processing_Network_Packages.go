package main

import (
	"bufio"
	"fmt"
	"os"
)

type Packet struct {
	arrival  int
	duration int
	end      int
}

type QueueImpl struct {
	queue []Packet
}

type Queue interface {
	Push(pack Packet)
	PeekBegin() Packet
	PeekEnd() Packet
	Pop() Packet
	IsEmpty() bool
	Size() int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var size, count int
	fmt.Fscan(in, &size, &count)

	if count == 0 {
		return
	}

	input := make([]Packet, count)

	for i := 0; i < count; i++ {
		var arr, dur int
		fmt.Fscan(in, &arr, &dur)
		input[i].arrival = arr
		input[i].duration = dur
	}

	fmt.Fprintln(out, input[0].arrival)

	q := &QueueImpl{}
	input[0].end = input[0].arrival + input[0].duration
	q.Push(input[0])

	for i := 1; i < count; i++ {

		curr := input[i]

		for !q.IsEmpty() && curr.arrival >= q.PeekBegin().end {
			q.Pop()
		}

		if q.Size() >= size {
			fmt.Fprintln(out, -1)
			continue
		}

		if q.IsEmpty() {
			fmt.Fprintln(out, curr.arrival)
			curr.end = curr.arrival + curr.duration
		} else {
			prev := q.PeekEnd()
			fmt.Fprintln(out, prev.end)
			curr.end = prev.end + curr.duration
		}

		q.Push(curr)
	}
}

// Push добавить в конец очереди
func (q *QueueImpl) Push(pack Packet) {
	q.queue = append(q.queue, pack)
}

// PeekBegin взять из очереди
func (q *QueueImpl) PeekBegin() (result Packet) {
	if q.IsEmpty() {
		return
	} else {
		result = q.queue[0]
	}
	return
}

func (q *QueueImpl) PeekEnd() (result Packet) {
	if q.IsEmpty() {
		return
	} else {
		result = q.queue[len(q.queue)-1]
	}
	return
}

// Pop удалить из начала очереди
func (q *QueueImpl) Pop() (result Packet) {
	if q.IsEmpty() {
		return
	} else {
		result = q.queue[0]
		q.queue = q.queue[1:]
	}
	return
}

// IsEmpty проверить пуста ли очередь
func (q *QueueImpl) IsEmpty() bool {
	return len(q.queue) == 0
}

// Size вернуть размер очереди
func (q *QueueImpl) Size() int {
	return len(q.queue)
}
