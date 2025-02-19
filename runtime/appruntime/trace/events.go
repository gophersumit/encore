package trace

import (
	"errors"
	"fmt"

	"encore.dev/appruntime/model"
	"encore.dev/beta/errs"
	"encore.dev/internal/stack"
)

type EventType byte

const (
	RequestStart       EventType = 0x01
	RequestEnd         EventType = 0x02
	GoStart            EventType = 0x03
	GoEnd              EventType = 0x04
	GoClear            EventType = 0x05
	TxStart            EventType = 0x06
	TxEnd              EventType = 0x07
	QueryStart         EventType = 0x08
	QueryEnd           EventType = 0x09
	CallStart          EventType = 0x0A
	CallEnd            EventType = 0x0B
	AuthStart          EventType = 0x0C
	AuthEnd            EventType = 0x0D
	HTTPCallStart      EventType = 0x0E
	HTTPCallEnd        EventType = 0x0F
	HTTPCallBodyClosed EventType = 0x10
	LogMessage         EventType = 0x11
	PublishStart       EventType = 0x12
	PublishEnd         EventType = 0x13
	ServiceInitStart   EventType = 0x14
	ServiceInitEnd     EventType = 0x15
	CacheOpStart       EventType = 0x16
	CacheOpEnd         EventType = 0x17
)

func (te EventType) String() string {
	switch te {
	case RequestStart:
		return "RequestStart"
	case RequestEnd:
		return "RequestEnd"
	case GoStart:
		return "GoStart"
	case GoEnd:
		return "GoEnd"
	case GoClear:
		return "GoClear"
	case TxStart:
		return "TxStart"
	case TxEnd:
		return "TxEnd"
	case QueryStart:
		return "QueryStart"
	case QueryEnd:
		return "QueryEnd"
	case CallStart:
		return "CallStart"
	case CallEnd:
		return "CallEnd"
	case AuthStart:
		return "AuthStart"
	case AuthEnd:
		return "AuthEnd"
	case HTTPCallStart:
		return "HTTPCallStart"
	case HTTPCallEnd:
		return "HTTPCallEnd"
	case HTTPCallBodyClosed:
		return "HTTPCallBodyClosed"
	case LogMessage:
		return "LogMessage"
	case PublishStart:
		return "PublishStart"
	case PublishEnd:
		return "PublishEnd"
	case ServiceInitStart:
		return "ServiceInitStart"
	case ServiceInitEnd:
		return "ServiceInitEnd"
	case CacheOpStart:
		return "CacheOpStart"
	case CacheOpEnd:
		return "CacheOpEnd"
	default:
		return fmt.Sprintf("Unknown(%x)", byte(te))
	}
}

func (l *Log) BeginRequest(req *model.Request, goid uint32) {
	tb := NewBuffer(1 + 8 + 8 + 8 + 8 + 8 + 8 + 64)
	tb.Byte(byte(req.Type))
	tb.Now()
	tb.Bytes(req.SpanID[:])
	tb.Bytes(req.ParentID[:])
	tb.String(req.Service)
	tb.String(req.Endpoint)
	tb.UVarint(uint64(goid))
	tb.UVarint(0)                  // call expr idx; unused
	tb.UVarint(uint64(req.DefLoc)) // endpoint expr idx
	tb.String(string(req.UID))
	tb.UVarint(uint64(len(req.Inputs)))
	for _, input := range req.Inputs {
		tb.UVarint(uint64(len(input)))
		tb.Bytes(input)
	}

	if req.Type == model.PubSubMessage {
		tb.String(req.MsgData.Topic)
		tb.String(req.MsgData.Subscription)
		tb.String(req.MsgData.MessageID)
		tb.Uint32(uint32(req.MsgData.Attempt))
		tb.Time(req.MsgData.Published)
	}

	l.Add(RequestStart, tb.Buf())
}

func (l *Log) FinishRequest(req *model.Request, output [][]byte, err error) {
	tb := NewBuffer(64)
	tb.Bytes(req.SpanID[:])
	if err == nil {
		tb.Byte(0) // no error
		tb.UVarint(uint64(len(output)))
		for _, output := range output {
			tb.UVarint(uint64(len(output)))
			tb.Bytes(output)
		}
	} else {
		tb.Byte(1) // error
		tb.String(err.Error())
		tb.Stack(errs.Stack(err))
	}
	l.Add(RequestEnd, tb.Buf())
}

func (l *Log) BeginCall(call *model.APICall, goid uint32) {
	tb := NewBuffer(8 + 4 + 4 + 4)
	tb.UVarint(call.ID)
	tb.Bytes(call.Source.SpanID[:])
	tb.Bytes(call.SpanID[:])
	tb.UVarint(uint64(goid))
	tb.UVarint(0)                   // call expr idx; no longer used
	tb.UVarint(uint64(call.DefLoc)) // endpoint expr idx
	tb.Stack(stack.Build(3))
	l.Add(CallStart, tb.Buf())
}

func (l *Log) FinishCall(call *model.APICall, err error) {
	tb := NewBuffer(8 + 4 + 4 + 4)
	tb.UVarint(call.ID)
	if err != nil {
		msg := err.Error()
		if msg == "" {
			msg = "unknown error"
		}
		tb.String(msg)
	} else {
		tb.String("")
	}
	l.Add(CallEnd, tb.Buf())
}

func (l *Log) BeginAuth(call *model.AuthCall, goid uint32) {
	tb := NewBuffer(8 + 4 + 4 + 4)
	tb.UVarint(call.ID)
	tb.Bytes(call.SpanID[:])
	tb.UVarint(uint64(goid))
	tb.UVarint(uint64(call.DefLoc)) // auth handler expr idx
	l.Add(AuthStart, tb.Buf())
}

func (l *Log) FinishAuth(call *model.AuthCall, uid model.UID, err error) {
	tb := NewBuffer(64)
	tb.UVarint(call.ID)
	tb.String(string(uid))
	if err != nil {
		msg := err.Error()
		if msg == "" {
			msg = "unknown error"
		}
		tb.String(msg)
		tb.Stack(errs.Stack(err))
	} else {
		tb.String("")
		tb.Stack(stack.Stack{}) // no stack
	}
	l.Add(AuthEnd, tb.Buf())
}

type DBQueryStartParams struct {
	Query   string
	SpanID  model.SpanID
	Goid    uint32
	QueryID uint64
	TxID    uint64
	Stack   stack.Stack
}

func (l *Log) DBQueryStart(p DBQueryStartParams) {
	var tb Buffer
	tb.UVarint(p.QueryID)
	tb.Bytes(p.SpanID[:])
	tb.UVarint(p.TxID)
	tb.UVarint(uint64(p.Goid))
	tb.String(p.Query)
	tb.Stack(p.Stack)
	l.Add(QueryStart, tb.Buf())
}

func (l *Log) DBQueryEnd(queryID uint64, err error) {
	var tb Buffer
	tb.UVarint(queryID)
	if err != nil {
		tb.String(err.Error())
	} else {
		tb.String("")
	}
	l.Add(QueryEnd, tb.Buf())
}

type DBTxStartParams struct {
	SpanID model.SpanID
	Goid   uint32
	TxID   uint64
	Stack  stack.Stack
}

func (l *Log) DBTxStart(p DBTxStartParams) {
	var tb Buffer
	tb.UVarint(p.TxID)
	tb.Bytes(p.SpanID[:])
	tb.UVarint(uint64(p.Goid))
	tb.Stack(p.Stack)
	l.Add(TxStart, tb.Buf())
}

type DBTxEndParams struct {
	SpanID model.SpanID
	Goid   uint32
	TxID   uint64
	Commit bool
	Err    error
	Stack  stack.Stack
}

func (l *Log) DBTxEnd(p DBTxEndParams) {
	var tb Buffer
	tb.UVarint(p.TxID)
	tb.Bytes(p.SpanID[:])
	tb.UVarint(uint64(p.Goid))
	if p.Commit {
		tb.Byte(1)
	} else {
		tb.Byte(0)
	}
	if p.Err != nil {
		tb.String(p.Err.Error())
	} else {
		tb.String("")
	}
	tb.Stack(p.Stack)
	l.Add(TxEnd, tb.Buf())
}

func (l *Log) PublishStart(topic string, msg []byte, spanID model.SpanID, goid uint32, publishID uint64, skipFrames int) {
	var tb Buffer
	tb.UVarint(publishID)
	tb.Bytes(spanID[:])
	tb.UVarint(uint64(goid))
	tb.String(topic)
	tb.ByteString(msg)
	tb.Stack(stack.Build(skipFrames))
	l.Add(PublishStart, tb.Buf())
}

func (l *Log) PublishEnd(publishID uint64, messageID string, err error) {
	var tb Buffer
	tb.UVarint(publishID)
	tb.String(messageID)
	tb.Err(err)
	l.Add(PublishEnd, tb.Buf())
}

func (l *Log) GoStart(spanID model.SpanID, goctr uint32) {
	l.Add(GoStart, []byte{
		spanID[0],
		spanID[1],
		spanID[2],
		spanID[3],
		spanID[4],
		spanID[5],
		spanID[6],
		spanID[7],
		byte(goctr),
		byte(goctr >> 8),
		byte(goctr >> 16),
		byte(goctr >> 24),
	})
}

func (l *Log) GoClear(spanID model.SpanID, goctr uint32) {
	l.Add(GoClear, []byte{
		spanID[0],
		spanID[1],
		spanID[2],
		spanID[3],
		spanID[4],
		spanID[5],
		spanID[6],
		spanID[7],
		byte(goctr),
		byte(goctr >> 8),
		byte(goctr >> 16),
		byte(goctr >> 24),
	})
}

func (l *Log) GoEnd(spanID model.SpanID, goctr uint32) {
	l.Add(GoEnd, []byte{
		spanID[0],
		spanID[1],
		spanID[2],
		spanID[3],
		spanID[4],
		spanID[5],
		spanID[6],
		spanID[7],
		byte(goctr),
		byte(goctr >> 8),
		byte(goctr >> 16),
		byte(goctr >> 24),
	})
}

type ServiceInitStartParams struct {
	InitCtr uint64
	SpanID  model.SpanID
	Goctr   uint32
	DefLoc  int32
	Service string
}

func (l *Log) ServiceInitStart(p ServiceInitStartParams) {
	var tb Buffer
	tb.Bytes(p.SpanID[:])
	tb.UVarint(p.InitCtr)
	tb.UVarint(uint64(p.Goctr))
	tb.UVarint(uint64(p.DefLoc))
	tb.String(p.Service)
	l.Add(ServiceInitStart, tb.Buf())
}

func (l *Log) ServiceInitEnd(initCtr uint64, err error) {
	var tb Buffer
	tb.UVarint(initCtr)
	tb.Err(err)
	if err != nil {
		tb.Stack(errs.Stack(err))
	}
	l.Add(ServiceInitEnd, tb.Buf())
}

type CacheOpStartParams struct {
	DefLoc    int32
	Operation string
	IsWrite   bool
	Keys      []string
	Inputs    [][]byte
	SpanID    model.SpanID
	Goid      uint32
	OpID      uint64
	Stack     stack.Stack
}

func (l *Log) CacheOpStart(p CacheOpStartParams) {
	var tb Buffer
	tb.UVarint(p.OpID)
	tb.Bytes(p.SpanID[:])
	tb.UVarint(uint64(p.Goid))
	tb.UVarint(uint64(p.DefLoc))
	tb.String(p.Operation)
	tb.Bool(p.IsWrite)
	tb.Stack(p.Stack)

	tb.UVarint(uint64(len(p.Keys)))
	for _, k := range p.Keys {
		tb.String(k)
	}

	tb.UVarint(uint64(len(p.Inputs)))
	suffix := []byte("...")
	for _, val := range p.Inputs {
		const maxLen = 4 * 1024 // 4KiB
		tb.TruncatedByteString(val, maxLen, suffix)
	}

	l.Add(CacheOpStart, tb.Buf())
}

type CacheOpEndParams struct {
	OpID    uint64
	Res     CacheOpResult
	Err     error // must be set iff Res == CacheErr
	Outputs [][]byte
}

func (l *Log) CacheOpEnd(p CacheOpEndParams) {
	var tb Buffer
	tb.UVarint(p.OpID)
	tb.Byte(byte(p.Res))
	if p.Res == CacheErr {
		err := p.Err
		if err == nil {
			err = errors.New("unknown error")
		}
		tb.Err(err)
	}

	tb.UVarint(uint64(len(p.Outputs)))
	suffix := []byte("...")
	for _, val := range p.Outputs {
		const maxLen = 4 * 1024 // 4KiB
		tb.TruncatedByteString(val, maxLen, suffix)
	}
	l.Add(CacheOpEnd, tb.Buf())
}

type CacheOpResult uint8

const (
	CacheOK        CacheOpResult = 1
	CacheNoSuchKey CacheOpResult = 2
	CacheConflict  CacheOpResult = 3
	CacheErr       CacheOpResult = 4
)
