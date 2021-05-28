package stack

type stack struct {
	dataList []interface{}
}

func NewStack() *stack {
	return &stack{}
}

func (s *stack) Pop(ele interface{}) {
	s.dataList = append(s.dataList, ele)
}

func (s *stack) Push() interface{} {
	lastEleIdx := len(s.dataList) - 1
	ele := s.dataList[lastEleIdx]
	s.dataList = s.dataList[:lastEleIdx]
	return ele
}

func (s *stack) IsEmpty() bool {
	return len(s.dataList) == 0
}

func (s *stack) Size() int {
	return len(s.dataList)
}
