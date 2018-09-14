# cloud

## Introduction

Microlog project is Microservice Log Service.It is logging project.Its code override the [logrus](https://github.com/sirupsen/logrus).Including the basic feature of [logrus](https://github.com/sirupsen/logrus).
And asynchronously upload log to [Elasticsearch](https://www.elastic.co/products/elasticsearch) using [KAFKA](https://kafka.apache.org/).

### Technology
* [kafka](https://kafka.apache.org/) - kafka
* [elasticsearch](https://www.elastic.co/products/elasticsearch) - elasticsearch
* [logrus](https://github.com/sirupsen/logrus) - logging


## 1、Initial microlog
```
SetFormatter(&log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
SetLevel(log.DebugLevel)
SetModuleName("Test Module")
```
## 2 Used
```
Debug("Debug")
Println("Println")
Info("Info")
Warning("Warning")
Warn("Warn")
Error("Error")
Panic("Panic")
```

## 3 Expanding Log
```
tutorial := "tutorial"
underestimate := "underestimate"
Infof("I am a %v.Don't %v me.",tutorial,underestimate)
```

## 4 Upload Log To Elasticsearch
```
tutorial := "tutorial"
underestimate := "underestimate"
Infof("I am a %v.Don't %v me.",tutorial,underestimate)
```

## 5 Upload Log To Elasticsearch.
```
import (
	log "microlog"
	hooksLog "microlog/hooks/syslog"
	"log/syslog"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	log.SetLevel(log.DebugLevel)
	log.SetModuleName("test-module")

	log.SetHookType(log.Elasticsearch)
	log.SetWriteLogProperties(&log.WriteLogProperties{URL:"http://172.28.2.25:9200"})
	hook, err := hooksLog.NewSyslogHook("", "", syslog.LOG_INFO, "")
	if err != nil {
		panic(err)
	}
	log.AddHook(hook)
}

func main() {
	log.Debug("Debug")
	log.Println("Println")
	log.Info("Info")
	log.Warning("Warning")
	log.Warn("Warn")
	log.Error("Error")
}
```

## 6 Asynchronously Upload Log To Elasticsearch Using KAFKA.
```
Not working(Ongoing)
```
##About Me
```
My name is Shikuan Xu.I'm chinese.My English is not very good,but I'm working on it.
If you have any questions, please send email to me.
Of course.If you have a job opportunity in Canada, please let me know,too.

Email: heixiushamao@gmail.com

```

