package for1

import "fmt"

type Tttt struct {
	age   int
	name  string
	class string
}

func (t1 Tttt) PrintT(age1 int, name1 string, class1 string) int {
	t1.age = age1
	t1.name = name1
	t1.class = class1
	fmt.Println(t1.age)
	fmt.Println(t1.name)
	fmt.Println(t1.class)
	return 1
}
