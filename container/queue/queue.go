package queue

type queue []interface{}

func New() *queue {
	return &queue{}
}

func (s *queue) Offer(ele interface{}) {
	queue := *s
	queue = append(queue, ele)
	*s = queue
}

func (s *queue) Poll() interface{} {
	queue := *s
	*s = queue[1:]
	return queue[0]
}

func (s *queue) IsEmpty() bool {
	return len(*s) == 0
}

func (s *queue) Size() int {
	return len(*s)
}
