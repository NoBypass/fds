package utils

import "github.com/NoBypass/surgo"

func Error(res []surgo.Result) error {
	for _, r := range res {
		if r.Error != nil {
			return r.Error
		}
	}
	return nil
}
