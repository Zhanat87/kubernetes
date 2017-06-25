package main

import (
	"fmt"

	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"github.com/Sirupsen/logrus"
	"os"
	stdlog "log"
	kitlog "github.com/go-kit/kit/log"
)

var logger = logrus.New()

func logrusMiddleware(c *routing.Context) error {
	logger.Infof("logrusHandler path: %s. uri: %s", string(c.Path()), string(c.Request.RequestURI()))
	//if len(c.Param("test")) > 0 {
	//	logger.Infof("test param: %s", c.Param("test"))
	//}
	return nil
}

func goKitLogMiddleware(c *routing.Context) error {
	logger := kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stdout))
	stdlog.SetOutput(kitlog.NewStdlibAdapter(logger))
	stdlog.Printf("goKitLogHandler path: %s. uri: %s", string(c.Path()), string(c.Request.RequestURI()))
	return nil
}

func helloWorldHandler(c *routing.Context) error {
	fmt.Fprintf(c, "Hello, world!")
	return nil
}

func testHandler(c *routing.Context) error {
	fmt.Fprintf(c, "test action. test param value: %s", c.Param("test"))
	return nil
}

func pingHandler(c *routing.Context) error {
	fmt.Fprint(c, "pong")
	return nil
}

func main() {
	router := routing.New()

	router.Use(logrusMiddleware)

	router.Use(goKitLogMiddleware)

	router.Get("/", helloWorldHandler)

	router.Get("/test/<test>", testHandler)

	router.Get("/ping", pingHandler)

	panic(fasthttp.ListenAndServe(":8080", router.HandleRequest))
}