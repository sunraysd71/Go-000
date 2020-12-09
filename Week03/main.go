package main
/**
* 测试方法：
*        （1）在linux机器上给进程发送任意信号，goroutine都会终止，并按照退出顺序输出对应的日志。
*        （2）发送给8080端口/shutdown路径，如curl 127.0.0.1/shutdown，goroutine都会终止，并按照退出顺序输出对应的日志。
* 存在问题：
*        （1）最明显的问题是使用errGroup因为是包装所以函数除了context一个入参之外什么都没有，以及在课程上说的都是匿名函数，需要解决如何让有参数的函数能使用上errGroup的问题。
*	 （2）总体感觉有些别扭，因为信号服务器里面不需要开另外的goroutine，而http服务器里面需要开另外的goroutine。需要评估提供给他人调用函数里面再开goroutine的风险或看有没有更好办法让http服务器里面没有goroutine。
*	 （3）由于第一点，errGroup感觉用起来不太灵活，该场景是否不使用errGroup会更好。
* 版本：v1
*
*
*
*/

import (
	"fmt"
	"os"
	"os/signal"
	"context"
	"github.com/go-kratos/kratos/pkg/sync/errgroup"
	"net/http"
	"errors"
)

var SignalError = errors.New("Error: Signal Interrupt")
var HttpError = errors.New("Error: Http server shutdown")

func handleSignal(ctx context.Context) error {
	c := make(chan os.Signal, 1)
        signal.Notify(c)
	fmt.Println("Signal handler starting..")
	select {
		case <-ctx.Done():
			fmt.Println("Signal server shutting down due to http request.")
		case s:= <-c:
			fmt.Println("Signal server shutting down due to signal interrupt.")
			fmt.Println("Got signal:", s)
	}
	return SignalError
}

type myHandler struct {
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello\n")
}

func startServer(ctx context.Context) error {
	var stat bool = false
	m := http.NewServeMux()
	s := http.Server{Addr: ":8080", Handler: m}
	fmt.Println("Http server starting..")
	m.Handle("/", new(myHandler))
	m.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Http server shutting down due to http request.")
	stat = true
	s.Shutdown(ctx)})
	go func(){
		<-ctx.Done()
		if stat == false{
			fmt.Println("Http server shutting down due to signal interrupt.")
			s.Shutdown(ctx)
		}
	}()
	if err := s.ListenAndServe(); err != nil{
		return HttpError
	}
	return nil
}

func main() {
	ctx := context.Background()
	g := errgroup.WithCancel(ctx)
	g.Go(startServer)
	g.Go(handleSignal)
	if err := g.Wait();err != nil{
	fmt.Println(err)
	}
	fmt.Println("End main!")
}
