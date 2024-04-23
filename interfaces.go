package types

type Combiner[T any] interface {
	Combine(T, T)
}

type Identifier[T any] interface {
	Identify(T)
	Identified() T
}

type Setter[T any] interface {
	Set(T)
}

type Getter[T any] interface {
	Get() T
}

type SetterGetter[T any] interface {
	Setter[T]
	Getter[T]
}

type IDModifier[ID any] interface {
	Combiner[ID]
	Identifier[ID]
}

type Modifier[Data any, ID any] interface {
	SetterGetter[Data]
	IDModifier[ID]
}

type Creator[C Modifier[any, any]] interface {
	Create(C) error
}

type Reader[C Modifier[any, any]] interface {
	Read(C) error
}

type Updater[C Modifier[any, any]] interface {
	Update(C) error
}

type Deleter[C Modifier[any, any]] interface {
	Delete(C) error
}

type Actioner[Data any, ID any] interface {
	Creator[Modifier[any, any]]
	Reader[Modifier[any, any]]
	Updater[Modifier[any, any]]
	Deleter[Modifier[any, any]]
}

type Error interface {
	error
}
