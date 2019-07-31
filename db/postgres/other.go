package postgres

import "fmt"

type InstrumentedClientOther struct{}

func NewInstrumentedClientOther() *InstrumentedClientOther {
	return &InstrumentedClientOther{}
}

func (c *InstrumentedClientOther) Insert(o string) error {
	fmt.Printf("posted metrics, inserted - %s\n", o)
	return nil
}

func (c *InstrumentedClientOther) Delete(o string) error {
	fmt.Printf("posted metrics, deleted - %s\n", o)
	return nil
}
