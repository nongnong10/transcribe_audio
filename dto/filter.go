package dto

type Filter interface {
	Process(in chan []byte) chan []byte
}
