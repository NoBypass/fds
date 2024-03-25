package pipeline

type Broadcaster[T any] struct {
	receivers []chan<- T
	in        <-chan T
}

func NewBroadcaster[T any](in <-chan T) *Broadcaster[T] {
	b := Broadcaster[T]{
		receivers: make([]chan<- T, 0),
		in:        in,
	}

	go func() {
		for {
			select {
			case v, ok := <-in:
				if !ok {
					for _, r := range b.receivers {
						close(r)
					}
					return
				}

				for _, r := range b.receivers {
					r <- v
				}
			}
		}
	}()

	return &b
}

func (b *Broadcaster[T]) Attach() <-chan T {
	out := make(chan T)
	b.receivers = append(b.receivers, out)
	return out
}
