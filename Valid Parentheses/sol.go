type Stack struct {
	val []rune
}

func (s *Stack) Push(v rune) {
	s.val = append(s.val, v)

}
func (s *Stack) Pop() (rune, bool) {
	if len(s.val) == 0 {
		return 0, false
	} else {
		item := s.val[len(s.val)-1]
		s.val = s.val[:len(s.val)-1]
		return item, true
	}
}
func (s *Stack) Peek() rune {
	if len(s.val) == 0 {
		return 0
	} else {
		return s.val[len(s.val)-1]
	}
}
func (s*Stack) Len() int {
	return len(s.val)
}


func isValid(s string) bool {
	stack := new(Stack)
	if len(s) <=1 || len(s)%2!=0{
		return false
	}
    	if (s[0] == ']' || s[0] == ')' || s[0] == '}') {
		return false
	}
	for i := range s {
		if s[i] == '[' || s[i] == '{' || s[i] == '(' {
			stack.Push(rune(s[i]))
		} else if s[i] == ']' && stack.Peek() == '[' {
			stack.Pop()
		} else if s[i] == '}' && stack.Peek() == '{' {
			stack.Pop()
		} else if s[i] == ')' && stack.Peek() == '(' {
			stack.Pop()
		} else {
			return false
		}
	}
	if stack.Len() == 0 {
		return true
	}else {
		return false
	}
	
}