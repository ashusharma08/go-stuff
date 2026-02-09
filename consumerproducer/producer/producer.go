package producer

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/esoptra/go-prac/consumerproducer/utils"
)

type Widget string
type Producer struct {
	Widget chan Widget
	GetUp  chan bool
	Name   string
}

func GetProducers(count int) []*Producer {
	p := make([]*Producer, count)
	i := 0
	for range count {
		p[i] = &Producer{
			Widget: make(chan Widget, 10),
			GetUp:  make(chan bool, 1),
			Name:   fmt.Sprintf("Producer %d", i+1),
		}
		i++
	}
	return p
}

func (p *Producer) Produce(ctx context.Context) {
	fmt.Println(p.Name, " staring ")
	for {
		select {
		case <-ctx.Done():
			close(p.Widget)
			close(p.GetUp)
			return
		default:
			num := utils.GetRandomNumber(2)
			for range num {
				w := getRandomWidget()
				select {
				case p.Widget <- w:
					fmt.Println("sending ", w, " from ", p.Name)
				default:
					select {
					case <-p.GetUp:
						fmt.Println("wakig up ", p.Name)
					case <-ctx.Done():
						close(p.Widget)
						close(p.GetUp)
						return
					}
				}
			}
		}
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getRandomWidget() Widget {
	bts := make([]byte, 10)
	for i := range bts {
		bts[i] = letterBytes[rand.Intn(51)]
	}
	return Widget(string(bts))
}
