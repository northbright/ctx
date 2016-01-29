package ctxcopy

import (
	"errors"
	"golang.org/x/net/context"
	"io"
)

const (
	// Zero buffer length error string.
	ZeroBufferLenError string = "Buffer length is 0."
	// Copy canceled string.
	CanceledStr string = "Copy operation is canceled."
)

// copyWithCancelation() reads bytes to buffer from source and write to destination from buffer with cancelation signal.
//
//   Params:
//     dst: Destination.
//     src: Source.
//     buf: Buffer(length should >= 0).
//     cancel: cancelation flag. It'll stop copy while the cancelation signal is set to true.
func copyWithCancelation(dst io.Writer, src io.Reader, buf []byte, cancel *bool) (err error) {
	if len(buf) == 0 {
		return errors.New(ZeroBufferLenError)
	}

	for {
		if cancel != nil && (*cancel) {
			return errors.New(CanceledStr)
		}

		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			return err
		} else {
			if n == 0 {
				break
			}
		}

		if _, err := dst.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}

// Copy() reads bytes to buffer from source and write to destination from buffer with cancelation signal.
//
//   Params:
//     ctx:
//       Google's Context type which carries deadlines, cacelation signals,
//       and other request-scoped values across API boundaries and between processes.
//       See https://godoc.org/golang.org/x/net/context for more.
//     dst: Destination.
//     src: Source.
//     buf: Buffer(length should >= 0).
func Copy(ctx context.Context, dst io.Writer, src io.Reader, buf []byte) (err error) {
	cancel := false
	c := make(chan error)

	go func() {
		c <- copyWithCancelation(dst, src, buf, &cancel)
	}()

	select {
	case err = <-c:
		return err

	case <-ctx.Done():
		// Set cancelation flag to true
		cancel = true
		// Wait for copyWithCancelation() return.
		<-c
		return ctx.Err()
	}
}
