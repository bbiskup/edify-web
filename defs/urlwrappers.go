package defs

import (
	"fmt"
	dsp "github.com/bbiskup/edify/edifact/spec/dataelement"
	msp "github.com/bbiskup/edify/edifact/spec/message"
	ssp "github.com/bbiskup/edify/edifact/spec/segment"
)

// Procides URL for message spec resource
func MsgSpecURL(msgSpec *msp.MsgSpec) string {
	return fmt.Sprintf("/specs/message/%s", msgSpec.Id)
}

// Provides URL for segment spec resource
func SegSpecURL(segSpec *ssp.SegSpec) string {
	return fmt.Sprintf("/specs/segment/%s", segSpec.Id)
}

// Provides URL for message spec part resource
func MsgSpecPartURL(msgSpecPart msp.MsgSpecPart) string {
	switch msgSpecPart := msgSpecPart.(type) {
	case *msp.MsgSpecSegPart:
		return SegSpecURL(msgSpecPart.SegSpec)
	case *msp.MsgSpecSegGrpPart:
		return fmt.Sprintf("TODO_seg_group_%s", msgSpecPart.Id())
	default:
		panic(fmt.Sprintf("Unsupported type %T", msgSpecPart))
	}
}

func DataElemSpecURL(dataElemSpec dsp.DataElemSpec) string {
	switch dataElemSpec := dataElemSpec.(type) {
	case *dsp.SimpleDataElemSpec:
		return fmt.Sprintf("/specs/simpledataelement/%s", dataElemSpec.Id())
	case *dsp.CompositeDataElemSpec:
		return fmt.Sprintf("/specs/compositedataelement/%s", dataElemSpec.Id())
	default:
		panic(fmt.Sprintf("Unexpected type: %T", dataElemSpec))
	}

}
