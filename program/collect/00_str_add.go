package program

import "errors"

var ErrNaN = errors.New("not a number")

// StrAdd 两个数字字符串相加
func StrAdd(s1, s2 string) (string, error) {
	if s1 == "" || s2 == "" {
		return "", ErrNaN
	}

	if len(s1) < len(s2) {
		s1, s2 = s2, s1
	}

	var carry bool
	var result []byte
	for i := 0; i < len(s1); i++ {
		b1 := s1[len(s1)-i-1]
		b2 := byte('0')
		if i < len(s2) {
			b2 = s2[len(s2)-i-1]
		}

		n1 := b1 - '0'
		n2 := b2 - '0'
		if n1 > 9 || n2 > 9 {
			return "", ErrNaN
		}

		n := n1 + n2
		if carry {
			n++
		}

		if n > 9 {
			carry = true
			n -= 10
		} else {
			carry = false
		}

		result = append(result, n+'0')
	}

	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}

	return string(result), nil
}
