package prot

type Transformer interface {
	Start() error
	Transform(record *Record) error
	Close() error
}


type CustomTransformer struct {}

func (g *CustomTransformer) Start() error {
	return nil
}

func (g *CustomTransformer) Transform(record *Record) error {
	return nil
}

func (g *CustomTransformer) Close() error {
	return nil
}
