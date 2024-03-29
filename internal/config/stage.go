package config

import (
	"cuelang.org/go/cue"
	"github.com/wabenet/dodo-config/pkg/cuetils"
	"github.com/wabenet/dodo-stage/pkg/box"
)

type Stage struct {
	Name      string
	Type      string
	Box       *box.Config
	Resources *Resources
	Options   *Options
}

func StagesFromValue(v cue.Value) (map[string]*Stage, error) {
	return StagesFromMap(v)
}

func StagesFromMap(v cue.Value) (map[string]*Stage, error) {
	out := map[string]*Stage{}

	err := cuetils.IterMap(v, func(name string, v cue.Value) error {
		r, err := StageFromStruct(name, v)
		if err == nil {
			out[name] = r
		}

		return err

	})

	return out, err
}

func StageFromStruct(name string, v cue.Value) (*Stage, error) {
	out := &Stage{Name: name}

	if p, ok := cuetils.Get(v, "name"); ok {
		if n, err := StringFromValue(p); err != nil {
			return nil, err
		} else {
			out.Name = n
		}
	}

	if p, ok := cuetils.Get(v, "type"); ok {
		if n, err := StringFromValue(p); err != nil {
			return nil, err
		} else {
			out.Type = n
		}
	}

	if p, ok := cuetils.Get(v, "box"); ok {
		if n, err := BoxFromValue(p); err != nil {
			return nil, err
		} else {
			out.Box = n
		}
	}

	if p, ok := cuetils.Get(v, "resources"); ok {
		if n, err := ResourcesFromValue(p); err != nil {
			return nil, err
		} else {
			out.Resources = n
		}
	}

	if p, ok := cuetils.Get(v, "options"); ok {
		if n, err := OptionsFromValue(p); err != nil {
			return nil, err
		} else {
			out.Options = n
		}
	}

	return out, nil
}
