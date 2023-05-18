package inters

import "context"

type ChannelData interface {
	Key() string
	UniqueKey() string

	Send(d []byte, key string) (err error)
	Recv() (d []byte, key string, err error)

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
