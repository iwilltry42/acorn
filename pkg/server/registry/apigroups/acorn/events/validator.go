package events

import (
	"context"
	"fmt"

	apiv1 "github.com/acorn-io/runtime/pkg/apis/api.acorn.io/v1"
	"github.com/acorn-io/runtime/pkg/event"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

type validator struct{}

func (validator) ValidateName(_ context.Context, obj runtime.Object) (result field.ErrorList) {
	e := obj.(*apiv1.Event)

	id, err := event.ContentID(e)
	if err != nil {
		result = append(result, field.InternalError(
			field.NewPath("metadata", "name"),
			fmt.Errorf("failed to generate content ID for event: %w", err),
		))
		return
	}

	if e.Name != id {
		result = append(result, field.Invalid(
			field.NewPath("metadata", "name"),
			e.Name,
			fmt.Sprintf("name should be %q", id),
		))
	}

	return
}
