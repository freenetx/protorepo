package inters

import "context"

type ChannelData interface {
	SelfKey() string
	DestinationKey() string
	UniqueKey() string

	Send(d []byte) (err error)
	Recv() ([]byte, error)

	Destroy() error
}

type ChannelClient interface {
	ChannelData

	Reconnect(ctx context.Context) error
}

type ChannelServerObserver interface {
	OnNewChannelData(ChannelData)
	OnDestroyChannelData(ChannelData)
}

type ChannelServer interface {
	Start(ob ChannelServerObserver) error
	Stop()
	Wait()
	StopAndWait()
}
