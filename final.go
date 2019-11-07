func AtoiBase(s string, t string) int {
	ans := 0
	ln := 0
	st := map[rune]bool{}
	for _, c := range t {
		if st[c] || c == '-' || c == '+' {
			return ans
		}
		st[c] = true
		ln++
	}
	if ln > 1 {
		for _, c := range s {
			cnt := 0
			if st[c] {
				for _, j := range t {
					if j == c {
						break
					}
					cnt++
				}
				ans = ans*ln + cnt
			}
		}
	}
	return ans
}

--------------------------
//#PrintNbrBase
import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(s int, t string) {
	ans := ""
	ln := 0
	for _, c := range t {
		if c == c {
			ln++
		}
	}
	mx_p := ln
	if s < 0 {
		ans = "-"
		mx_p *= -1
	}
	if ln > 1 {

		for s/mx_p >= ln {
			mx_p *= ln
		}
		for mx_p != 0 {
			ans = ans + string(t[s/mx_p])
			s = s - s/mx_p*mx_p
			mx_p /= ln
		}
		x := map[rune]bool{}
		for _, c := range t {
			if c == '+' || c == '-' {
				ans = "NV"
				break
			}
			if x[c] == false {
				x[c] = true
			} else {
				ans = "NV"
				break
			}
		}
	} else {
		ans = "NV"
	}
	for _, c := range ans {
		z01.PrintRune(c)
	}
}

----------
//#rotatevowels
import (
	"os"

	"github.com/01-edu/z01"
)

func check(x rune) bool {
	if x == 'a' || x == 'A' || x == 'e' || x == 'E' || x == 'o' || x == 'O' || x == 'u' || x == 'U' || x == 'i' || x == 'I' {
		return true
	}
	return false
}
func main() { 
	arg := os.Args[1:]
	rep := []rune{}
	ans := ""
	ln := 0
	IsF := true
	for _, c := range arg {
		for _, j := range c {
			if check(j) {
				rep = append(rep, j)
				ln++
			}
		}
		if IsF {
			ans = c
			IsF = false
			continue
		}
		ans = ans + " " + c
	}
	cur := 0
	for _, c := range ans {
		if check(c) {
			z01.PrintRune(rep[ln-cur-1])
			cur++
		} else {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}

----
//#ConvertBase
func ConvertBase(s, t, p string) string {
	ln := 0
	ln2 := 0
	ln3 := 0
	ans := ""
	st_t := map[rune]int{}
	for c := range s {
		ln = c + 1
	}
	for i, c := range t {
		st_t[c] = i
		ln2 = i + 1
	}
	for c := range p {
		ln3 = c + 1
	}
	pw := 1
	cnt := 0
	for i := ln - 1; i >= 0; i-- {
		cnt = cnt + st_t[[]rune(s)[i]]*pw
		pw *= ln2
	}
	for cnt != 0 {
		ans = string(p[cnt%ln3]) + ans
		cnt /= ln3
	}
	return ans
}

---
//#StrLen

func StrLen(str string) int {
	ln := 0
	for _, c := range str {
		if c == c {
			ln++
		}
	}
	return ln
}

---
//#FirstRune

func FirstRune(s string) rune {
	var ans rune
	for _, c := range s {
		return c
	}
	return ans
}
---
//#LastRune

func LastRune(s string) rune {
	var ans rune
	ln := 0
	for _, c := range s {
		if c == c {
			ln++
		}
	}
	cnt := 0
	for _, c := range s {
		if cnt == ln-1 {
			return c
		}
		cnt++
	}
	return ans
}

---
//#SortWordArr

func SortWordArr(arr []string) {
	ln := 0
	for i := range arr {
		ln = i + 1
	}
	for i := 0; i < ln; i++ {
		for j := i + 1; j < ln; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

---
//#GCD

private static ulong GCD(ulong a, ulong b)
{
    while (a != 0 && b != 0)
    {
        if (a > b)
            a %= b;
        else
            b %= a;
    }

    return a == 0 ? b : a; //tut  a == 0   togda b inache a 
}

---
//#Split

func Split(str, charset string) []string {
	ln := 0
	ok := true
	ln_s := 0
	for c := range str {
		ln_s = c + 1
	}
	for i := 0; i < ln_s; i++ {
		same := true
		ln2 := 0
		for j, t := range charset {
			if rune(str[i+j]) != t {
				same = false
				break
			}
			ln2 = j
		}

		if same {
			if !ok {
				ln++
			}
			i += (ln2)
			ok = true

			continue
		}
		ok = false
	}
	ans := make([]string, ln+1)
	x := 0
	for i := 0; i < ln_s; i++ {
		same := true
		ln2 := 0
		for j, t := range charset {
			if rune(str[i+j]) != t {
				same = false
				break
			}
			ln2 = j
		}
		if same {
			if !ok {
				x++
			}
			i += (ln2)
			ok = true
			continue
		}
		ans[x] = ans[x] + string(str[i])
		ok = false
	}

	return ans
}

---
//#ListRemoveIf

func ListRemoveIf(l *List, data_ref interface{}) {

	for l.Head != nil && CompStr(l.Head.Data, data_ref) {
		l.Head = l.Head.Next
	}

	it := l.Head

	for it != nil {

		cur := it.Next
		for cur != nil && CompStr(cur.Data, data_ref) {
			cur = cur.Next
		}
		it.Next = cur
		it = it.Next
	}
}

---
//#ListSize
func ListSize(l *List) int {
	ln := 0
	it := l.Head
	for it != nil {
		it = it.Next
		ln++
	}
	return ln
}