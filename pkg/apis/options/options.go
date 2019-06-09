package options

import (
	"fmt"
	"io"
	"strings"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Options contains configuration options for the OAuth2 Proxy
type Options struct {
	Cookie *CookieOptions
}

// New creates a new deafulted copy of the Options struct
func New() *Options {
	options := &Options{}
	err := defaultStruct(options)
	if err != nil {
		// If we get an error here, there must be a code error
		panic(err)
	}
	return options
}

// Load reads a config file, flag arguments and the environment to set the
// correct configuration options
func Load(config io.Reader, configType string, args []string) (*Options, error) {
	flagSet := flag.NewFlagSet("oauth2-proxy", flag.ExitOnError)
	flagSet.SetNormalizeFunc(wordSepNormalizeFunc)

	// Add FlagSets to main flagSet
	flagSet.AddFlagSet(cookieFlagSet)

	flagSet.Parse(args)

	// Create a viper for binding config
	v := viper.New()

	// Bind flags to viper
	err := v.BindPFlags(flagSet)
	if err != nil {
		return nil, err
	}

	// Configure loading of environment variables
	// All flag options are prefixed by the EnvPrefix
	v.SetEnvPrefix("OAUTH2_PROXY")
	// Substitute "-" for "_" so `FOO_BAR` matches the config `foo.bar`
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// Read the configuration file
	if config != nil {
		if configType == "" {
			return nil, fmt.Errorf("config-type not set")
		}
		v.SetConfigType(configType)
		v.ReadConfig(config)
	}

	// Unmarhsal the config into the Options struct
	options := New()
	err = v.Unmarshal(&options)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %v", err)
	}
	return options, nil
}

// wordSepNormalizeFunc replaces "-" in flags entered with "."
// This ensures that flags are mapped to the correct values in the Options struct
func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
	from := []string{"-"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return flag.NormalizedName(name)
}
