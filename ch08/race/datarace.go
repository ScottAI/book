package singlegoroutine

var realNum = make(chan int)//设置数字值
var delta = make(chan int)//设置设置的增减额

func SetNumber(n int)  {
	realNum <- n
}
func ChangeByDelta(d int)  {
	delta <- d
}
func GetNumber() int{
	return <- realNum
}
func moitor()  {
	var i int //把数值限定在方法内，goroutine运行后仅在goroutine内
	for {
		select{
		case i = <- realNum:
		case d := <- delta:
			i += d
		case realNum <- i:
		}
	}
}
func init()  {
	go moitor()
}