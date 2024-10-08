package config

import (
	"io"
	"os"

	"github.com/go-logr/logr"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"sigs.k8s.io/yaml"

	"github.com/kumahq/kuma/pkg/core"
)

type Loader struct {
	logger logr.Logger
	cfg    Config

	strict        bool
	includeEnv    bool
	validate      bool
	envVarsPrefix string
}

func NewLoader(cfg Config) *Loader {
	return &Loader{
		logger: core.Log.WithName("config"),
		cfg:    cfg,
	}
}

func (l *Loader) WithStrictParsing() *Loader {
	l.strict = true
	return l
}

func (l *Loader) WithEnvVarsLoading(envVarsPrefix string) *Loader {
	l.includeEnv = true
	l.envVarsPrefix = envVarsPrefix
	return l
}

func (l *Loader) WithValidation() *Loader {
	l.validate = true
	return l
}

func (l *Loader) Load(stdin io.Reader, content []byte, filename string) error {
	switch filename {
	case "-":
		return l.LoadReader(stdin)
	case "":
		return l.LoadBytes(content)
	default:
		return l.LoadFile(filename)
	}
}

func (l *Loader) LoadFile(filename string) error {
	if filename == "" {
		l.logger.Info("no configuration file provided, skipping file-based config loading")
		return l.postProcess()
	}

	if _, err := os.Stat(filename); err != nil {
		return errors.Errorf("unable to access configuration file '%s', please check if the file exists and has the correct permissions", filename)
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return errors.Wrapf(err, "reading configuration from file '%s' failed", filename)
	}

	if err := l.LoadBytes(content); err != nil {
		return errors.Wrapf(err, "parsing configuration from file '%s' failed", filename)
	}

	return nil
}

func (l *Loader) LoadReader(r io.Reader) error {
	if r == nil {
		l.logger.Info("no configuration reader provided, skipping reader-based config loading")
		return nil
	}

	content, err := io.ReadAll(r)
	if err != nil {
		return errors.Wrap(err, "reading configuration from reader failed")
	}

	return l.LoadBytes(content)
}

func (l *Loader) LoadBytes(content []byte) error {
	if content == nil {
		l.logger.Info("no configuration content provided, skipping byte-based config loading")
		return nil
	}

	if err := l.unmarshal(content); err != nil {
		return errors.Wrap(err, "unable to parse configuration")
	}

	return l.postProcess()
}

func (l *Loader) unmarshal(content []byte) error {
	if l.strict {
		return yaml.UnmarshalStrict(content, l.cfg)
	}

	return yaml.Unmarshal(content, l.cfg)
}

func (l *Loader) postProcess() error {
	if l.includeEnv {
		if err := envconfig.Process(l.envVarsPrefix, l.cfg); err != nil {
			return errors.Wrap(err, "processing environment variables failed")
		}
	}

	if err := l.cfg.PostProcess(); err != nil {
		return errors.Wrap(err, "configuration post-processing failed")
	}

	if l.validate {
		if err := l.cfg.Validate(); err != nil {
			return errors.Wrap(err, "configuration validation failed")
		}
	}

	return nil
}

func Load(file string, cfg Config) error {
	return NewLoader(cfg).WithEnvVarsLoading("").WithValidation().LoadFile(file)
}
