package external

import "golang-aws-lambda/src/external/kkj"

type External struct {
	Kkj *kkj.KKJExternal
}

func NewExternal() (*External, error) {
	return &External{
		Kkj: kkj.NewKKJExternal(),
	}, nil
}
