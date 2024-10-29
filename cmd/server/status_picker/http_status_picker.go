package status_picker

import (
	"github.com/mroth/weightedrand"
)

type HttpStatusPicker struct {
	*weightedrand.Chooser
}

func NewHttpStatusPicker() (*HttpStatusPicker, error) {
	choices := GetHttpStatuses()
	chooser, err := weightedrand.NewChooser(*choices...)
	if err != nil {
		return nil, err
	}

	return &HttpStatusPicker{chooser}, nil
}

func GetHttpStatuses() *[]weightedrand.Choice {
	return &[]weightedrand.Choice{
		{Item: 200, Weight: 60},
		{Item: 400, Weight: 10},
		{Item: 401, Weight: 10},
		{Item: 404, Weight: 10},
		{Item: 500, Weight: 10},
	}
}

func (p *HttpStatusPicker) PickRandom() int {
	return p.Pick().(int)
}
