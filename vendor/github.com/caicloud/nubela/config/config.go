package config

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
	"time"
	"unicode"
	"unicode/utf8"
)

// The command line flags all get stored in a private flag set. The
// developer of the E2E test suite decides how they are exposed. Options
// include:
// - exposing as normal flags in the actual command line:
//   CopyFlags(Flags, flag.CommandLine)
// - populate via test/e2e/framework/viperconfig:
//   viperconfig.ViperizeFlags("my-config.yaml", "", Flags)
// - a combination of both:
//   CopyFlags(Flags, flag.CommandLine)
//   viperconfig.ViperizeFlags("my-config.yaml", "", flag.CommandLine)

// Flags is the flag set that AddOptions adds to. Test authors should
// also use it instead of directly adding to the global command line.
var Flags = flag.NewFlagSet("", flag.ContinueOnError)

// AddOptions analyzes the options value and creates the necessary
// flags to populate it.
//
// The prefix can be used to root the options deeper in the overall
// set of options, with a dot separating different levels.
//
// The function always returns true, to enable this simplified
// registration of options:
// _ = AddOptions(...)
//
// It panics when it encounters an error, like unsupported types
// or option name conflicts.

func AddOptions(options interface{}, prefix string) bool {
	return AddOptionsToSet(Flags, options, prefix)
}

// AddOptionsToSet is the same as AddOption, except that it allows choosing the flag set.
func AddOptionsToSet(flags *flag.FlagSet, options interface{}, prefix string) bool {
	optionsType := reflect.TypeOf(options)
	if optionsType == nil {
		panic("options parameter without a type - nil?!")
	}
	if optionsType.Kind() != reflect.Ptr || optionsType.Elem().Kind() != reflect.Struct {
		panic(fmt.Sprintf("need a pointer to a struct, got instead: %T", options))
	}
	addStructFields(flags, optionsType.Elem(), reflect.Indirect(reflect.ValueOf(options)), prefix)
	return true
}

func addStructFields(flags *flag.FlagSet, structType reflect.Type, structValue reflect.Value, prefix string) {
	for i := 0; i < structValue.NumField(); i++ {
		entry := structValue.Field(i)
		addr := entry.Addr()
		structField := structType.Field(i)
		name := structField.Name
		r, n := utf8.DecodeRuneInString(name)
		name = string(unicode.ToLower(r)) + name[n:]
		usage := structField.Tag.Get("usage")
		def := structField.Tag.Get("default")
		if prefix != "" {
			name = prefix + "." + name
		}
		if structField.PkgPath != "" {
			panic(fmt.Sprintf("struct entry should %q not exported, please make it to lower case", name))
		}
		ptr := addr.Interface()
		if structField.Anonymous {
			// Entries in embedded fields are treated like
			// entries, in the struct itself, i.e. we add
			// them with the same prefix.
			addStructFields(flags, structField.Type, entry, prefix)
			continue
		}
		if structField.Type.Kind() == reflect.Struct {
			// Add nested options.
			addStructFields(flags, structField.Type, entry, name)
			continue
		}
		// We could switch based on structField.Type. Doing a
		// switch after getting an interface holding the
		// pointer to the entry has the advantage that we
		// immediately have something that we can add as flag
		// variable.
		//
		// Perhaps generics will make this entire switch redundant someday...
		switch ptr := ptr.(type) {
		case *bool:
			var defValue bool
			parseDefault(&defValue, name, def)
			flags.BoolVar(ptr, name, defValue, usage)
		case *time.Duration:
			var defValue time.Duration
			parseDefault(&defValue, name, def)
			flags.DurationVar(ptr, name, defValue, usage)
		case *float64:
			var defValue float64
			parseDefault(&defValue, name, def)
			flags.Float64Var(ptr, name, defValue, usage)
		case *string:
			flags.StringVar(ptr, name, def, usage)
		case *int:
			var defValue int
			parseDefault(&defValue, name, def)
			flags.IntVar(ptr, name, defValue, usage)
		case *int64:
			var defValue int64
			parseDefault(&defValue, name, def)
			flags.Int64Var(ptr, name, defValue, usage)
		case *uint:
			var defValue uint
			parseDefault(&defValue, name, def)
			flags.UintVar(ptr, name, defValue, usage)
		case *uint64:
			var defValue uint64
			parseDefault(&defValue, name, def)
			flags.Uint64Var(ptr, name, defValue, usage)
		default:
			panic(fmt.Sprintf("unsupported struct entry type %q: %T", name, entry.Interface()))
		}
	}
}

// parseDefault is necessary because "flag" wants the default in the
// actual type and cannot take a string. It would be nice to reuse the
// existing code for parsing from the "flag" package, but it isn't
// exported.
func parseDefault(value interface{}, name, def string) {
	if def == "" {
		return
	}
	checkErr := func(err error, value interface{}) {
		if err != nil {
			panic(fmt.Sprintf("invalid default %q for %T entry %s: %s", def, value, name, err))
		}
	}
	switch value := value.(type) {
	case *bool:
		v, err := strconv.ParseBool(def)
		checkErr(err, *value)
		*value = v
	case *time.Duration:
		v, err := time.ParseDuration(def)
		checkErr(err, *value)
		*value = v
	case *float64:
		v, err := strconv.ParseFloat(def, 64)
		checkErr(err, *value)
		*value = v
	case *int:
		v, err := strconv.Atoi(def)
		checkErr(err, *value)
		*value = v
	case *int64:
		v, err := strconv.ParseInt(def, 0, 64)
		checkErr(err, *value)
		*value = v
	case *uint:
		v, err := strconv.ParseUint(def, 0, strconv.IntSize)
		checkErr(err, *value)
		*value = uint(v)
	case *uint64:
		v, err := strconv.ParseUint(def, 0, 64)
		checkErr(err, *value)
		*value = v
	default:
		panic(fmt.Sprintf("%q: setting defaults not supported for type %T", name, value))
	}
}
