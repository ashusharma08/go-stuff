package consumer

import (
	"fmt"
	"time"

	"github.com/esoptra/go-prac/consumerproducer/producer"
	"github.com/esoptra/go-prac/consumerproducer/utils"
)

type Consumer struct {
	Name             string
	Producer         []*producer.Producer
	Widgets          []producer.Widget
	consumptionCount int
}

func GetConsumers(count int, producers []*producer.Producer) []*Consumer {
	c := make([]*Consumer, count)
	i := 0
	for range count {
		c[i] = &Consumer{
			Producer:         producers,
			Name:             fmt.Sprintf("Consumer %d", i+1),
			Widgets:          make([]producer.Widget, 0, 10),
			consumptionCount: 0,
		}
		i++
	}
	return c
}

func (c *Consumer) Consume() {
	fmt.Println("consumer ", c.Name, " starting up")
	i := 0
	for {
		if c.consumptionCount == 10 {
			return
		}
		if i >= len(c.Producer) {
			i = 0
		}
		p := c.Producer[i]
		num := utils.GetRandomNumber(2)
		for range num {
			select {
			case val := <-p.Widget:
				if len(c.Widgets) == 10 {
					fmt.Println(c.Name, " discarded")
					c.Widgets = c.Widgets[:0]
					c.consumptionCount++
					if c.consumptionCount == 10 {
						return
					}
					time.Sleep(time.Second * time.Duration(utils.GetRandomNumber(4)))
					break
				}
				c.Widgets = append(c.Widgets, val)
				fmt.Println("read ", val, " from ", p.Name, " by consumer ", c.Name)
			default:
				fmt.Println("waking up ", p.Name)
				p.GetUp <- true
			}
		}
		i++
	}
}
