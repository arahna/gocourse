package safeslice

type SafeSlice interface {
	Append(interface{})
	At(int) interface{}
	Close() []interface{}
	Delete(int)
	Len() int
	Update(int, UpdateFunc)
}

type UpdateFunc func(interface{}) interface{}

type safeslice chan command

type command struct {
	action  action
	key     int
	value   interface{}
	updater UpdateFunc
	result  chan interface{}
	data    chan []interface{}
}

type action int

const (
	push   action = iota
	remove
	at
	length
	update
	end
)

func (ss safeslice) Append(value interface{}) {
	ss <- command{action: push, value: value}
}

func (ss safeslice) At(key int) interface{} {
	reply := make(chan interface{})
	ss <- command{action: at, key: key, result: reply}
	return <-reply
}

func (ss safeslice) Close() []interface{} {
	reply := make(chan []interface{})
	ss <- command{action: end, data: reply}
	return <-reply
}

func (ss safeslice) Delete(key int) {
	ss <- command{action: remove, key: key}
}

func (ss safeslice) Len() int {
	reply := make(chan interface{})
	ss <- command{action: length, result: reply}
	return (<-reply).(int)
}

func (ss safeslice) Update(key int, updater UpdateFunc) {
	ss <- command{action: update, key: key, updater: updater}
}

func New() SafeSlice {
	ss := make(safeslice)
	ss.run()
	return ss
}

func (ss safeslice) run() {
	data := make([]interface{}, 0, 1000)
	go func() {
		for {
			command := <-ss
			switch command.action {
			case push:
				data = append(data, command.value)
			case remove:
				key := command.key
				if 0 <= key && key < len(data) {
					data = append(data[:key], data[key+1:]...)
				}
			case at:
				if 0 <= command.key && command.key < len(data) {
					result := data[command.key]
					command.result <- result
				} else {
					command.result <- nil
				}
			case length:
				command.result <- len(data)
			case update:
				if 0 <= command.key && command.key < len(data) {
					value := data[command.key]
					data[command.key] = command.updater(value)
				}
			case end:
				close(ss)
				command.data <- data
				return
			}
		}
	}()
}
