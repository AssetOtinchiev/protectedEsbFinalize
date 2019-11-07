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

---
//#BrainFuck

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println()
		return
	}

	arr := make([]byte, 2048)
	p := &arr[0]
	arg := []rune(os.Args[1])

	i := 0
	j := 0

	for j < 2048 {
		for i < len(arg) {
			switch arg[i] {
			case '>':
				j++
				p = &arr[j]
			case '<':
				j--
				p = &arr[j]
			case '+':
				*p++
			case '-':
				*p--
			case '.':
				fmt.Println(string(*p))
			case '[':
				cont := 0

				if *p == 0 {

					for cont >= 0 {
						i++
						if arg[i] == ']' {
							cont--
						} else if arg[i] == '[' {
							cont++
						}
					}
				}
			case ']':
				cont := 0
				if *p != 0 {
					for cont >= 0 {
						i--
						if arg[i] == '[' {
							cont--
						} else if arg[i] == ']' {
							cont++
						}
					}
				}
			}

			i++
		}
		j = 2048
	}
}


---
//#Brackets

import (
	"os"

	piscine ".."
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		z01.PrintRune('\n')
		return
	}

	for i := 1; i < len(os.Args); i++ {
		arr := []rune(os.Args[i])
		var stack []rune

		if os.Args[i] == "" {
			piscine.PrintStr("OK")
			z01.PrintRune('\n')
		} else {
			for j := 0; j < piscine.StrLen(os.Args[i]); j++ {
				if arr[j] == '(' || arr[j] == '[' || arr[j] == '{' {
					stack = append(stack, arr[j])
				} else {
					switch arr[j] {
					case ')':
						if stack[len(stack)-1] == '(' {
							stack = stack[0 : len(stack)-1]

						}
					case ']':
						if stack[len(stack)-1] == '[' {
							stack = stack[0 : len(stack)-1]

						}
					case '}':
						if stack[len(stack)-1] == '{' {
							stack = stack[0 : len(stack)-1]
						}
					}
				}
			}

			if len(stack) == 0 {
				piscine.PrintStr("OK")
				z01.PrintRune('\n')
			} else {
				piscine.PrintStr("Error")
				z01.PrintRune('\n')
			}

		}
	}
}

---
//#Options

import (
	"os"

	"github.com/01-edu/z01"

	piscine ".."
)

func main() {
	alpha := "abcdefghijklmnopqrstuvwxyz"

	if len(os.Args) == 1 {
		piscine.PrintStr("options: " + alpha + "\n")
		return
	}

	alpharev := piscine.StrRev(alpha)

	numbers := "00000000 00000000 00000000 00000000"
	nbrs := []rune(numbers)

	var flags []rune

	for i := 1; i < len(os.Args); i++ {
		flag := []rune(os.Args[i])
		if flag[0] == '-' {
			for j := 1; j < piscine.ArrayRuneLength(flag); j++ {
				if flag[j] >= 'a' && flag[j] <= 'z' {
					flags = append(flags, flag[j])
				}
			}
		} else {
			piscine.PrintStr("Invalid Option\n")
			return
		}
	}

	for i := 0; i < piscine.ArrayRuneLength(flags); i++ {
		helper := piscine.Index(alpharev, string(flags[i]))
		if flags[i] == 'z' || flags[i] == 'y' {
			helper += 6
		} else if flags[i] <= 'x' && flags[i] >= 'q' {
			helper += 7
		} else if flags[i] <= 'p' && flags[i] >= 'i' {
			helper += 8
		} else if flags[i] <= 'h' && flags[i] >= 'a' {
			helper += 9
		}
		nbrs[helper] = '1'
	}

	if nbrs[27] == '1' {
		piscine.PrintStr("options: " + alpha + "\n")
		return
	}
	for i := 0; i < piscine.ArrayRuneLength(nbrs); i++ {
		z01.PrintRune(nbrs[i])
	}
	z01.PrintRune('\n')
}


---
//#SwapBits

func SwapBits(octet byte) byte {
	return ((octet&0x0F)<<4 | (octet&0xF0)>>4)
  }


//#PrintHex
import( 
	piscine ".."
		"os"
		
		"github.com/01-edu/z01"
	)
	func main(){
	
		 if piscine.Lent3(os.Args) >1 && piscine.Lent3(os.Args) <=2 {
			v:= piscine.Atoi(os.Args[1])
			if v == 0{
				z01.PrintRune(0)
			}
			dec2hex(v)
		}else{
			z01.PrintRune('\n')
		}
	
	}
	
	
	func dec2hex(n int){
		var arr [100]rune
		i:=0
		for n!=0{
			temp:=n%16
			if temp<10{
				arr[i]= rune(temp + 48 )
				i++
			}else{
				arr[i]=rune(temp + 87)
				i++
			}
			n=n/16
		}
		for j:=i-1;j>=0;j--{
			z01.PrintRune(arr[j])
		}
		z01.PrintRune('\n')
	}

	---
	//#printmemory

	import (
		"fmt"
		"github.com/01-edu/z01"
	)
	
	func PrintMemory(arr [10]int) {
		index := 0
		for i := 0; i < len(arr); i++ {
			index++
			if index == 5 || index == 9 {
				fmt.Println()
			}
	
			if arr[i] != 0 {
				// PrintNbrBase(arr[i],"0123456789ABCDEF")
				dec2hexa(arr[i])
				fmt.Print("00 0000 ")
			} else {
				fmt.Println("0000 0000")
			}
	
		}
	
	}
	
	func dec2hexa(n int) {
		var arr [100]rune
		i := 0
		for n != 0 {
			temp := n % 16
			if temp < 10 {
				arr[i] = rune(temp + 48)
				i++
			} else {
				arr[i] = rune(temp + 87)
				i++
			}
			n = n / 16
		}
		for j := i - 1; j >= 0; j-- {
			z01.PrintRune(arr[j])
		}
	}

	--
	//#Itoa

	func Itoa(nbr int) string {
		result = ""
		t := 1
	
		if nbr < 0 {
			result += "-"
			t = -1
		}
		if nbr != 0 {
			q := (nbr / 10) * t
			if q != 0 {
				Itoa(q)
			}
			d := ((nbr % 10) * t) + '0'
			result += string(rune(d))
		} else {
			result += "0"
		}
	
		return result
	}

	---
	//#Addprimesum
	import (
		piscine ".."
		"fmt"
		"os"
	)
	
	func main() {
	
		if len(os.Args) == 2 {
			nbr := piscine.Atoi(os.Args[1])
			printPrime(nbr)
	
		} else {
			fmt.Println("0")
		}
	}
	
	func IsPrime(value int) bool {
		if value <= 1 {
			return false
		}
		for i := 2; i < value; i++ {
			if value%i == 0 {
				return false
			}
		}
		return true
	}
	
	func printPrime(n int) int{
		sum := 0
		for i:= 2; i <= n; i++ {
			if IsPrime(i) {
			sum+= i
			
				//fmt.Println(sum) 
			}
			
		}
		fmt.Println(sum) 
		return sum
	}
	---
	//#fprime


import (
	piscine ".."
	"fmt"
	"os"
)

func main() {

	if len(os.Args) > 1 && len(os.Args) <= 2 {
		count := 0
		l := append((PrimeFactors(piscine.Atoi(os.Args[1]))))
		for i := 0; i < len(l); i++ {
			count++
			if count >= i && i > 0 {
				fmt.Print("*")
			}
			fmt.Print(l[i])

		}
		fmt.Println()

	} else {
		fmt.Println()
	}
}

func PrimeFactors(x int) []rune {
	factors := []rune{}

	candidate := 2
	for candidate <= x {
		if x%candidate == 0 {
			factors = append(factors, rune(candidate))
			x = x / candidate
		} else {
			candidate++
		}
	}

	return factors
}

---
//#revWstr
import (
	piscine ".."
	"fmt"
	"os"
)

func main() {

	if len(os.Args) > 1 && len(os.Args) <= 2 {
		str := piscine.SplitWhiteSpaces(os.Args[1])
		for i := len(str) - 1; i >= 0; i-- {
			fmt.Print(string(str[i]) + " ")
		}
		fmt.Println()
	} else {
		fmt.Println()
	}
}

--
//#Split

func Split(str, charset string) []string {
	resp := []string{}
	new := ""
		for i := 0; i < len(str); i++ {
			if Charset(str, charset, i) {
				if new != "" {
					resp = append(resp, new)
					new = ""
					i = i + len(charset) - 1
				}
			} else {
				new = new + string(str[i])
			}
		}
		if new != "" {
			resp = append(resp, new)
		}
		return resp
}


func Charset(str, charset string, i int) bool {

	j := 0
		for j < len(charset) && i < len(str) {
			if str[i] != charset[j] {
				return false
			}
			i++
			j++
		}
		return true
}

--
//#hiddenp

import (
	piscine ".."
	"fmt"
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if piscine.Lent3(os.Args) != 3 {
		z01.PrintRune('\n')
		} else {
		count:=0
		in:=0
		arr1 := []rune(os.Args[1])
		arr2 := []rune(os.Args[2])
	
			arr2=unique(Reverse(arr2))
			arr2=Reverse(arr2)
		//	fmt.Println(string(arr2))
			if os.Args[1] == "" {
				fmt.Println("1")
				} else {
					
					for g := 0; g < len(arr1); g++ {
						for j := count; j < len(arr2); j++ {

						if arr1[g]==arr2[j]{							
							in ++
							count=j
							
						}
					}
				}
			}
		//	fmt.Println(in,len(arr1) )
			if in == len(arr1){
				fmt.Println("1")
				}else {
					
					fmt.Println("0")
					
				}
		
	}
}

func unique(Slice []rune) []rune {
    keys := make(map[rune]bool)
    list := []rune{} 
    for _, entry := range Slice  {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}


func Reverse(s []rune) []rune {
    var reverse []rune
    for i := len(s)-1; i >= 0; i-- {
	
        reverse =append(reverse,s[i])
    }
    return reverse 
}


--
//#sortwordarr


func SortWordArr(array []string) {

	quickSrot2(array, 0, len(array)-1)

}


func quickSrot2(table []string, beg int, end int) {

	if beg < end {
		lockPivo := mudaVariavel2(table, beg, end)
		quickSrot2(table, beg, lockPivo-1)
		quickSrot2(table, lockPivo+1, end)

	}
}

func mudaVariavel2(table []string, beg int, end int) int {
	pivo := table[end]
	i := beg - 1

	for j := beg; j < end; j++ {
		if table[j] <= pivo {
			i++
			aux := table[j]
			table[j] = table[i]
			table[i] = aux
		}
	}

	aux := table[end]
	table[end] = table[i+1]
	table[i+1] = aux

	return i + 1
}

---
//#repeatalpha
import (
	"github.com/01-edu/z01"
	piscine ".."
	"os"
)


func main() {

count:=0
if piscine.Lent3(os.Args) >1 {
	arr:=[]rune(os.Args[1])
	for i:=0; i<len(arr);i++{

		if arr[i] >= 'a' && arr[i] <= 'z' || arr[i] >= 'A' && arr[i] <= 'Z'{
		
			if arr[i] >= 'a' && arr[i] <= 'z'{
				count = int(arr[i] - 'a') + 1
			}
		
			if arr[i] >= 'A' && arr[i] <= 'Z'{
				 count = int(arr[i] - 'A') + 1 
			}
			for ; count>0; count--{
				z01.PrintRune(arr[i])
			}
			}	
		}
		z01.PrintRune('\n')
	}else{
		z01.PrintRune('\n')
	}
}
---
//#inter
import (
	"os"
	piscine ".."
	"fmt"
)

func main() {

	str1 := os.Args[1]
	str2 := os.Args[2]
	arr1 := []rune(str1)
	arr2 := []rune(str2)
	
	if piscine.Lent3(os.Args)==3 {
		
		
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {

			if arr2[j] == arr1[i] {

				for k := i+1; k < len(arr1); k++ {
					if arr1[i] == arr1[k] {
						arr1[k] = 0

					}

				}

			}

		}
	}
	fmt.Println(string(arr1))
	}else{
		fmt.Println()
	}

}

---
//#tabmult

import (
	piscine ".."
	"fmt"
	"os"
)

func main() {
strarg := os.Args[1]
arg := piscine.Atoi(strarg)
	
    if piscine.Lent3(os.Args) < 1 {
		fmt.Println()
		return
	}

	for i := 1; i <= 9; i++ {

		fmt.Println(i, "*", arg, "=", tab(i, arg))

	}

}

func tab(a, b int) int {

	r := a * b
	return r
}