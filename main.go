package main

import (
	log "github.com/xushikuan/microlog"
	"flag"
)


var envConsulAddress string
var envModuleName string
var envPort int

func init(){
	//moduleName := flag.String("moduleName", "router", "Set module name.(string)")
	level := flag.String("level", "info", "Set log level.Including panic,fatal,error,warn,warning,info,debug(string)")
	flag.StringVar(&envModuleName, "module-name", "router", "Set module name.(string)")
	flag.IntVar(&envPort, "port", 80, "Set module port.(int)")
	flag.StringVar(&envConsulAddress, "consul-address", "", "Set consul address.(string)")

	flag.Parse()

	envLevel,err := log.ParseLevel(*level)
	if err != nil{
		log.Panic("Set log level error.",err)
	}
	if envConsulAddress == ""{
		log.Panic("consul-address cannot be nil or empty.")
	}

	log.SetFormatter(&log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	log.SetLevel(envLevel)
	log.SetModuleName(envModuleName)
	//log.SetFormatter(&log.JSONFormatter{TimestampFormat: format_time})
	////log.SetOutput(os.Stdout)
	//log.SetLevel(log.DebugLevel)
	//log.SetModuleName("test-module")
	//log.SetHookType(log.Elasticsearch)
	//log.SetWriteLogProperties(log.WriteLogProperties{URL:"http://172.28.2.25:9200"})
	//hook, err := hookSyslog.NewSyslogHook("", "", syslog.LOG_INFO, "")
	//if err != nil {
	//	panic(err)
	//}
	//log.AddHook(hook)
	log.Info("module-name :",envModuleName)
	log.Info("consul-address :",envConsulAddress)
	log.Info("port :",envPort)
}

func main() {
	log.Infof("-------------------- start server %v--------------------",envModuleName)
	log.Debug("-------------------- asdfasdfsafasd--------------------")
	log.Info("-------------------- asdfasdfsafasd--------------------")
}