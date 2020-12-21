package config

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// ViperizeFlags checks whether viperConfig was specified, read it, and updates the configuration
// in the specified flag set accordingly. Must be called after framework.HandleFlags() and
// before framework.AfterReadingAllFlags(). If the configuration is empty, nothing happens.
func ViperizeFlags(viperConfig string, flags *flag.FlagSet) error {
	if viperConfig == "" {
		return nil
	}
	ext := filepath.Ext(viperConfig)
	viper.SetConfigName(filepath.Base(viperConfig[0 : len(viperConfig)-len(ext)]))
	viper.AddConfigPath(filepath.Dir(viperConfig))

	wrapError := func(err error) error {
		if err == nil {
			return nil
		}
		errorPrefix := fmt.Sprintf("input viper config %q, actually %q", viperConfig, viper.ConfigFileUsed())
		return errors.Wrap(err, errorPrefix)
	}

	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			return wrapError(errors.New("not found"))
		case viper.UnsupportedConfigError:
			return wrapError(errors.New("not using a support file format"))
		case viper.ConfigParseError:
			return wrapError(errors.New("failed to parse configuration file"))
		default:
			return wrapError(err)
		}
	}

	// Update all flag values not already set with values found
	// via Viper.
	return wrapError(viperUnmarshal(flags))
}

// viperUnmarshall updates all flags with the corresponding values found  via Viper
func viperUnmarshal(flags *flag.FlagSet) error {
	var err error
	set := make(map[string]bool)

	// Determine which values were already set explicitly via
	// flags. Those we don't overwrite because command line
	// flags have a higher priority.
	flags.Visit(func(f *flag.Flag) {
		set[f.Name] = true
	})

	flags.VisitAll(func(f *flag.Flag) {
		if set[f.Name] || err != nil || !viper.IsSet(f.Name) {
			return
		}
		str := fmt.Sprintf("%v", viper.Get(f.Name))
		if err := f.Value.Set(str); err != nil {
			err = fmt.Errorf("setting option %q from config file value: %s", f.Name, err)
		}

	})
	return err
}
