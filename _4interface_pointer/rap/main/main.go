package main
import "fmt"
//其实也没什么好总结的。只有两点需要记住，一是 interface 是有自己对应的实体数据结构的，
//尽量不要用指针去指向 interface，因为 golang 对指针自动解引用的处理会带来陷阱。
type Rapper interface {
	Rap() string
}

type Dean struct {}

func (_ Dean) Rap() string {
	return "Im a rapper"
}

func doRap(p *Rapper) {
	fmt.Println((*p).Rap())
}

func main(){
	i := new(Rapper)
	*i = Dean{}
	fmt.Println((*i).Rap())
	doRap(i)
}

