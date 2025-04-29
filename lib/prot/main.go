package prot

func New() AppGenerate {
	return &App{}
}

type AppGenerate interface {
	Generate(g Generator) AppTransform
}

type AppTransform interface {
	Transform(t Transformer) AppTransformSpeak
}

type AppTransformSpeak interface {
	AppTransform
	AppSpeak
}

type AppSpeak interface {
	Speak(s Speaker)
}
