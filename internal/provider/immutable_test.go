package provider

import (
	"slices"
	"testing"
)

func TestGeneratedDescriptorsRecordImmutableFields(t *testing.T) {
	expected := map[string][]string{
		"mta_route_relay": {"name"},
		"allowed_ip":      {"address"},
	}

	descriptors := make(map[string]resourceDescriptor, len(generatedResources))
	for _, descriptor := range generatedResources {
		descriptors[descriptor.Name] = descriptor
	}

	for name, fields := range expected {
		descriptor, ok := descriptors[name]
		if !ok {
			t.Fatalf("descriptor %q was not generated", name)
		}

		for _, field := range fields {
			if !slices.Contains(descriptor.Immutable, field) {
				t.Errorf("descriptor %q does not record %q as immutable, got %v", name, field, descriptor.Immutable)
			}
		}
	}
}
