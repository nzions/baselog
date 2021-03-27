package baselog

import "fmt"

// Logger Interface
type Logger interface {
	Tracef(string, ...interface{})
	Bugf(string, ...interface{})
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Logf(string, ...interface{})
	Log(string)
	Errorf(string, ...interface{})
	Errore(error)
}

// a basic implementation of the logger, just prints to stdout, or a function specified
type STDOutLog struct {
	LogFunc func(string)
}

func (x *STDOutLog) logIt(m string) {
	if x.LogFunc != nil {
		x.LogFunc(m)
		return
	}
	fmt.Println(m)
}

func (x *STDOutLog) Tracef(f string, o ...interface{}) {
	m := fmt.Sprint("tce | ", fmt.Sprintf(f, o...))
	x.logIt(m)
}

func (x *STDOutLog) Bugf(f string, o ...interface{}) {
	m := fmt.Sprint("bug | ", fmt.Sprintf(f, o...))
	x.logIt(m)
}

func (x *STDOutLog) Debugf(f string, o ...interface{}) {
	m := fmt.Sprint("dbg | ", fmt.Sprintf(f, o...))
	x.logIt(m)
}

func (x *STDOutLog) Infof(f string, o ...interface{}) {
	m := fmt.Sprint("inf | ", fmt.Sprintf(f, o...))
	x.logIt(m)
}

func (x *STDOutLog) Warnf(f string, o ...interface{}) {
	m := fmt.Sprint("wrn | ", fmt.Sprintf(f, o...))
	x.logIt(m)
}

func (x *STDOutLog) Logf(f string, o ...interface{}) {
	m := fmt.Sprint("log | ", fmt.Sprintf(f, o...))
	x.logIt(m)
}

func (x *STDOutLog) Errorf(f string, o ...interface{}) {
	m := fmt.Sprint("err | ", fmt.Sprintf(f, o...))
	x.logIt(m)
}

func (x *STDOutLog) Errore(err error) {
	m := fmt.Sprintf("err | %T %s", err, err.Error())
	x.logIt(m)
}

func (x *STDOutLog) Log(f string) {
	m := fmt.Sprint("log | ", f)
	x.logIt(m)
}
