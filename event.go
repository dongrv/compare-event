package event

import "errors"

// Arithmetic 算数比较逻辑
type Arithmetic int

const (
	Equal            Arithmetic = iota // 等于
	NotEqual                           // 不等于
	GreaterThan                        // 大于
	LessThan                           // 小于
	GreaterThanEqual                   // 大于等于
	LessThanEqual                      // 小于等于
)

// Argument 参数
type Argument struct {
	Compare   int64      // 被比较值
	Condition Arithmetic // 比较条件
	Standard  int64      // 标准值，目标值
}

// Logic 逻辑判断
type Logic int

const (
	LogicAnd Logic = iota // 逻辑与
	LogicOr               // 逻辑或
	LogicNot              // 逻辑非
)

// Condition 比较函数执行单元
type Condition func(compare int64, arithmetic Arithmetic, standard int64) (bool, error)

// Compare 数值比较
func Compare(standard int64, arithmetic Arithmetic, compare int64) (bool, error) {
	switch arithmetic {
	case Equal:
		return compare == standard, nil
	case NotEqual:
		return compare != standard, nil
	case GreaterThan:
		return compare > standard, nil
	case LessThan:
		return compare < standard, nil
	case GreaterThanEqual:
		return compare >= standard, nil
	case LessThanEqual:
		return compare <= standard, nil
	}
	return false, errors.New("invalid compare arithmetic")
}

type event interface {
	Id() int                                                             // 事件Id
	Type() int                                                           // 事件类型
	Do(handlers []Condition, args []Argument, logic Logic) (bool, error) // 运行条件判断，使用与或非衡量条件
}

type Event struct {
	event
	id    int // 事件Id
	type_ int // 类型
}

func NewEvent(id, type_ int) *Event {
	return &Event{id: id, type_: type_}
}

func (ev *Event) Id() int {
	return ev.id
}

func (ev *Event) Type() int {
	return ev.type_
}

func (ev *Event) Do(handlers []Condition, args []Argument, logic Logic) (bool, error) {
	if len(handlers) == 0 {
		return false, errors.New("functions is empty")
	}
	if len(handlers) != len(args) {
		return false, errors.New("functions and arguments does not match")
	}
	switch logic {
	case LogicAnd: // 都满足条件
		for k, v := range handlers {
			b, err := v(args[k].Standard, args[k].Condition, args[k].Compare)
			if b == false {
				return false, err
			}
		}
		return true, nil
	case LogicOr: // 至少一个满足条件
		for k, v := range handlers {
			b, err := v(args[k].Standard, args[k].Condition, args[k].Compare)
			if b == true {
				return true, err
			}
		}
		return false, nil
	case LogicNot: // 都不能满足条件
		for k, v := range handlers {
			b, err := v(args[k].Standard, args[k].Condition, args[k].Compare)
			if b == true {
				return false, err
			}
		}
		return true, nil
	}
	return false, errors.New("invalid event condition ")
}
