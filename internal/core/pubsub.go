package core

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

// PubSub responsible for communication intermediary between packages.
// The core is using watermill GoChannel pub/sub implementation
// and could be replaced by any other type as long as
// it's still using watermill package, otherwise
// the router wouldn't work as expected.
type PubSub struct {
	core   *gochannel.GoChannel
	router *message.Router
}

func NewPubSub(l *Logger) *PubSub {
	defer l.GetCore().Info("PubSub initialized")

	var (
		err    error
		gochan *gochannel.GoChannel
		router *message.Router
		logger = watermill.NewSlogLogger(l.GetCore())
	)

	{
		config := gochannel.Config{}
		gochan = gochannel.NewGoChannel(config, logger)
	}

	{
		config := message.RouterConfig{}
		router, err = message.NewRouter(config, logger)
		if err != nil {
			panic(err)
		}
	}

	return &PubSub{core: gochan, router: router}
}

func (ps *PubSub) GetCore() *gochannel.GoChannel {
	return ps.core
}

func (ps *PubSub) GetRouter() *message.Router {
	return ps.router
}
