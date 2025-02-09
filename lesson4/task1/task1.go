package task1

import (
	"bufio"
	"context"
	"io"
)

func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {
	reader := bufio.NewReader(r)
	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		default:
			data, err := reader.ReadByte()
			if err != nil {
				if err == io.EOF {
					return false, nil
				}
				return false, err
			}
			if data == seq[0] {
				var flag bool = true
				for i := 1; i < len(seq); i++ {
					dataNext, err := reader.ReadByte()
					if err != nil {
						if err == io.EOF {
							return false, nil
						}
						return false, err
					}
					if dataNext != seq[i] {
						flag = false
						err = reader.UnreadByte()
						if err != nil {
							return false, err
						}
						break
					}
				}
				if flag {
					return true, nil
				}
			}
		}
	}
}
