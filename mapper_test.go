package confmap_test

import (
	"testing"

	"github.com/memlib/confmap"
)

func equals(t *testing.T, expect, actual any) {
	if expect != actual {
		t.Errorf(`expects "%+v"(%T), actual "%+v"(%T)`, expect, expect, actual, actual)
	}
}

func Test_MapEnvs(t *testing.T) {
	type mappingObj struct {
		Field0 string `mapstructure:"field_0"`
		Field1 string `mapstructure:"field_1"`
	}

	t.Run("success", func(t *testing.T) {
		var obj mappingObj
		t.Setenv("PRFX_FIELD_0", "val0")
		t.Setenv("PRFX_FIELD_1", "val1")

		if err := confmap.MapEnvs("prfx", &obj); err != nil {
			t.Errorf("error = %v", err)
		}

		equals(t, "val0", obj.Field0)
		equals(t, "val1", obj.Field1)
	})

	t.Run("success generics", func(t *testing.T) {
		var obj *mappingObj
		t.Setenv("PRFX_FIELD_0", "val0")
		t.Setenv("PRFX_FIELD_1", "val1")

		obj, err := confmap.MapEnvsTyped[mappingObj]("prfx")
		if err != nil {
			t.Errorf("error = %v", err)
		}

		equals(t, "val0", obj.Field0)
		equals(t, "val1", obj.Field1)
	})

	t.Run("success_single", func(t *testing.T) {
		var obj mappingObj
		t.Setenv("PRFX_FIELD_0", "val0")

		if err := confmap.MapEnvs("prfx", &obj); err != nil {
			t.Errorf("error = %v", err)
		}

		equals(t, "val0", obj.Field0)
		equals(t, "", obj.Field1)
	})
}

type testObjWithDefaults struct {
	Field0 string `mapstructure:"field_0"`
	Field1 string `mapstructure:"field_1"`
}

func (d testObjWithDefaults) DefaultEnvs() map[string]any {
	return map[string]any{
		"field_0": "def_val0",
		"field_1": "def_val1",
	}
}

func TestDefaults(t *testing.T) {
	var obj testObjWithDefaults
	t.Setenv("PRFX_FIELD_0", "val0")

	if err := confmap.MapEnvs("prfx", &obj); err != nil {
		t.Errorf("error = %v", err)
	}

	equals(t, "val0", obj.Field0)
	equals(t, "def_val1", obj.Field1)
}
