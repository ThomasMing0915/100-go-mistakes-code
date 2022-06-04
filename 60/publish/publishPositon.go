package publish

import (
	"context"
	"fmt"
	"github.com/ThomasMing0915/100-go-mistakes-code/60/flight"
	"time"
)

type PublishHandler struct {
	pub publisher
}

func NewPublishHandler() PublishHandler {
	return PublishHandler{
		pub: publisher{},
	}

}

type publisher struct{}

func (p publisher) Publish(ctx context.Context, position flight.Position) (string, error) {
	ch := make(chan string)
	go func() {
		// 模拟超时工作
		time.Sleep(time.Second * 5)
		select {
		case ch <- "result":
		default:
			return
		}
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case result := <-ch:
		return result, nil
	}
}

func (h PublishHandler) PublishPosition(position flight.Position) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	fmt.Printf("[%v] begin start handle \n", time.Now())
	res, err := h.pub.Publish(ctx, position)
	if err != nil {
		fmt.Printf("[%v] return  error: %v \n", time.Now(), err)
		return
	}

	fmt.Printf("[%v] %s \n", time.Now(), res)
}
