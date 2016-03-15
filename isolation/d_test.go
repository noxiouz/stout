package isolation

import (
	"errors"
	"testing"
	"time"

	"golang.org/x/net/context"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func init() {
	Suite(&initialDispatchSuite{})
}

type testDownstreamItem struct {
	code int
	args []interface{}
}

type testDownstream struct {
	ch chan testDownstreamItem
}

func (t *testDownstream) Reply(code int, args ...interface{}) error {
	t.ch <- testDownstreamItem{code, args}
	return nil
}

type testBox struct {
	err   error
	sleep time.Duration
}

func (b *testBox) Spool(ctx context.Context, name string, opts profile) error {
	select {
	case <-ctx.Done():
		return errors.New("canceled")
	case <-time.After(b.sleep):
		return b.err
	}
}

type initialDispatchSuite struct {
	ctx    context.Context
	cancel context.CancelFunc

	d       Dispatcher
	dw      *testDownstream
	session int
}

func (s *initialDispatchSuite) SetUpTest(c *C) {
	ctx, cancel := context.WithCancel(context.Background())

	boxes := isolationBoxes{
		"testError": &testBox{err: errors.New("dummy error from testBox")},
		"testSleep": &testBox{err: nil, sleep: time.Second * 2},
		"test":      &testBox{err: nil},
	}

	ctx = withArgsUnpacker(ctx, jsonArgsDecoder{})
	ctx = context.WithValue(ctx, isolationBoxesTag, boxes)

	s.ctx, s.cancel = ctx, cancel
	s.session = 100

	s.dw = &testDownstream{
		ch: make(chan testDownstreamItem),
	}

	ctx = withDownstream(s.ctx, s.dw)
	d := newInitialDispatch(ctx)
	s.d = d
}

func (s *initialDispatchSuite) TearDownTest(c *C) {
	s.cancel()
}

func (s *initialDispatchSuite) TestSpool(c *C) {
	var (
		args = profile{
			Isolate: Isolate{
				Type: "test",
			},
		}
		appName  = "application"
		spoolMsg = message{s.session, spool, []interface{}{args, appName}}
	)

	spoolDisp, err := s.d.Handle(&spoolMsg)
	c.Assert(err, IsNil)
	c.Assert(spoolDisp, FitsTypeOf, &spoolCancelationDispatch{})
	msg := <-s.dw.ch
	c.Assert(msg.code, Equals, replySpoolOk)
}

func (s *initialDispatchSuite) TestSpoolCancel(c *C) {
	var (
		args = profile{
			Isolate: Isolate{
				Type: "testSleep",
			},
		}
		appName   = "application"
		spoolMsg  = message{s.session, spool, []interface{}{args, appName}}
		cancelMsg = message{s.session, spoolCancel, []interface{}{}}
	)

	spoolDisp, err := s.d.Handle(&spoolMsg)
	c.Assert(err, IsNil)
	c.Assert(spoolDisp, FitsTypeOf, &spoolCancelationDispatch{})
	spoolDisp.Handle(&cancelMsg)
	msg := <-s.dw.ch
	c.Assert(msg.code, Equals, replySpoolError)
}

func (s *initialDispatchSuite) TestSpoolError(c *C) {
	var (
		args = profile{
			Isolate: Isolate{
				Type: "testError",
			},
		}
		appName  = "application"
		spoolMsg = message{s.session, spool, []interface{}{args, appName}}
	)

	spoolDisp, err := s.d.Handle(&spoolMsg)
	c.Assert(err, IsNil)
	c.Assert(spoolDisp, FitsTypeOf, &spoolCancelationDispatch{})
	msg := <-s.dw.ch
	c.Assert(msg.code, Equals, replySpoolError)
}
