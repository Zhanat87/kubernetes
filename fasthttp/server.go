package main

import (
	"fmt"

	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"github.com/Sirupsen/logrus"
	"os"
	stdlog "log"
	kitlog "github.com/go-kit/kit/log"
	"github.com/Zhanat87/go/util"
)

var logger = logrus.New()

func logrusHandler(c *routing.Context) error {
	logger.Infof("logrusHandler path: %s. uri: %s", string(c.Path()), string(c.Request.RequestURI()))
	//if len(c.Param("test")) > 0 {
	//	logger.Infof("test param: %s", c.Param("test"))
	//}
	return nil
}

func goKitLogHandler(c *routing.Context) error {
	logger := kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stdout))
	stdlog.SetOutput(kitlog.NewStdlibAdapter(logger))
	stdlog.Printf("goKitLogHandler path: %s. uri: %s", string(c.Path()), string(c.Request.RequestURI()))
	return nil
}

func main() {
	router := routing.New()

	router.Use(logrusHandler)

	router.Use(goKitLogHandler)

	router.Get("/", func(c *routing.Context) error {
		fmt.Fprintf(c, "Hello, world!")
		return nil
	})

	router.Get("/test/<test>", func(c *routing.Context) error {
		fmt.Fprintf(c, "test action. test param value: %s", c.Param("test"))
		return nil
	})

	router.Get("/json", func(c *routing.Context) error {
		return c.Write(util.H{"a": "b", "c": 1})
	})

	panic(fasthttp.ListenAndServe(":8080", router.HandleRequest))
}