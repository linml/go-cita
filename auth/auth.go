// Copyright (C) 2018 yejiayu

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package auth

import (
	"github.com/golang/glog"
	"github.com/streadway/amqp"

	"github.com/yejiayu/go-cita/auth/service"
	"github.com/yejiayu/go-cita/database"
	"github.com/yejiayu/go-cita/mq"
)

type Interface interface {
	Run(quit chan<- error)
}

func New(queue mq.Queue, dbFactory database.Factory) (Interface, error) {
	srv, err := service.NewService(dbFactory)
	if err != nil {
		return nil, err
	}

	return &auth{
		queue:      queue,
		mqHandler:  newMQ(queue, srv),
		rpcHandler: newRPC(queue, srv),
	}, nil
}

type auth struct {
	queue mq.Queue

	mqHandler  *mqHandler
	rpcHandler *rpcHandler
}

func (a *auth) Run(quit chan<- error) {
	go a.loopMQ(quit)
}

func (a *auth) loopMQ(quit chan<- error) {
	delivery, err := a.queue.Sub()
	if err != nil {
		quit <- err
		return
	}

	for msg := range delivery {
		go func(msg *amqp.Delivery) {
			key := mq.RoutingKey(msg.RoutingKey)
			data := msg.Body

			if err := a.mqHandler.Call(key, data); err != nil {
				glog.Error(err)
				msg.Ack(false)
			} else {
				msg.Ack(true)
			}
		}(&msg)
	}
}
