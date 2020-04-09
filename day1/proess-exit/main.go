package main

func main()  {
	//panic("call panic")
	//死锁
	//死锁(Deadlock)就是一个进程拿着资源A请求资源B，
	//另一个进程拿着资源B请求资源A，双方都不释放自己的资源，导致两个进程都进行不下去。
	ch := make(chan int)
	<-ch
}
