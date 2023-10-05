package log

import (
	"bytes"
	"fmt"
	"io"

	"strconv"
	"time"

	"github.com/x14n/x14nfile/http"
)

//Log 结构体表示一个日志记录器
//EntryNum 用于指定一个日志条目数
//Writer 用于指定日志的写入目标
//Interval 用于写入日志的最小间隔
//i 相当于日志的条目索引
//ch 日志条目传输通道
//t一个定时器，用于定期刷新日志。

type Log struct {
	EntryNum int
	Writer   io.Writer
	Interval time.Duration

	i  int
	ch chan entry
	t  *time.Ticker

	entries []entry
}

type entry struct {
	req  http.Request
	code string
	sent int
	t    time.Time
}

func (l *Log) Init() {
	l.t = time.NewTicker(l.Interval)
	l.ch = make(chan entry, 2*l.EntryNum)
	l.entries = make([]entry, l.EntryNum)
}

func (l *Log) Log(req http.Request, code string, sent int) {
	l.ch <- entry{req, code, sent, time.Now()}
}

func (l *Log) Loop() {
	for {
		select {
		case e := <-l.ch:
			l.entries[l.i] = e
			l.i++
			if l.i == l.EntryNum {
				l.flush()
			}
		}
	}
}

func (l *Log) flush() {
	if l.i == 0 {
		return
	}

	var w bytes.Buffer
	for i := 0; i < l.i; i++ {
		e := l.entries[i]
		sent := strconv.Itoa(e.sent)
		fmt.Fprintf(&w, "%s %s %s %s %s %s %s \"%s\"\n",
			e.t.Format(time.RFC3339),
			or(e.req.Method, "-"),
			or(e.req.Path, "-"),
			or(e.req.Version, "-"),
			or(e.code, "-"),
			or(sent, "-"),
			or(e.req.Headers.Get("Referer"), "-"),
			or(e.req.Headers.Get("User-Agent"), "-"),
		)
	}
	l.Writer.Write(w.Bytes())
	l.i = 0
}

func or(in, or string) string {
	if in == "" {
		return or
	} else {
		return in
	}
}
