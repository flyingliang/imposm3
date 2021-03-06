package config

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

type Config struct {
	CacheDir           string  `json:"cachedir"`
	Connection         string  `json:"connection"`
	MappingFile        string  `json:"mapping"`
	LimitTo            string  `json:"limitto"`
	LimitToCacheBuffer float64 `json:"limitto_cache_buffer"`
	Srid               int     `json:"srid"`
}

const defaultSrid = 3857
const defaultCacheDir = "/tmp/imposm3"

var ImportFlags = flag.NewFlagSet("import", flag.ExitOnError)
var DiffFlags = flag.NewFlagSet("diff", flag.ExitOnError)

type _BaseOptions struct {
	Connection         string
	CacheDir           string
	MappingFile        string
	Srid               int
	LimitTo            string
	LimitToCacheBuffer float64
	ConfigFile         string
}

func (o *_BaseOptions) updateFromConfig() error {
	conf := &Config{
		CacheDir: defaultCacheDir,
		Srid:     defaultSrid,
	}

	if o.ConfigFile != "" {
		f, err := os.Open(o.ConfigFile)
		if err != nil {
			return err
		}
		decoder := json.NewDecoder(f)

		err = decoder.Decode(&conf)
		if err != nil {
			return err
		}
	}
	if o.Connection == "" {
		o.Connection = conf.Connection
	}
	if conf.Srid == 0 {
		conf.Srid = defaultSrid
	}
	if o.Srid != defaultSrid {
		o.Srid = conf.Srid
	}
	if o.MappingFile == "" {
		o.MappingFile = conf.MappingFile
	}
	if o.LimitTo == "" {
		o.LimitTo = conf.LimitTo
	}
	if o.LimitToCacheBuffer == 0.0 {
		o.LimitToCacheBuffer = conf.LimitToCacheBuffer
	}
	if o.CacheDir == defaultCacheDir {
		o.CacheDir = conf.CacheDir
	}
	return nil
}

func (o *_BaseOptions) check() []error {
	errs := []error{}
	if o.Srid != 3857 {
		errs = append(errs, errors.New("srid!=3857 not implemented"))
	}
	if o.MappingFile == "" {
		errs = append(errs, errors.New("missing mapping"))
	}
	return errs
}

type _ImportOptions struct {
	Cpuprofile       string
	Httpprofile      string
	Memprofile       string
	Overwritecache   bool
	Appendcache      bool
	Read             string
	Write            bool
	Optimize         bool
	Diff             bool
	DeployProduction bool
	RevertDeploy     bool
	RemoveBackup     bool
	Quiet            bool
}

var BaseOptions = _BaseOptions{}
var ImportOptions = _ImportOptions{}

func addBaseFlags(flags *flag.FlagSet) {
	flags.StringVar(&BaseOptions.Connection, "connection", "", "connection parameters")
	flags.StringVar(&BaseOptions.CacheDir, "cachedir", defaultCacheDir, "cache directory")
	flags.StringVar(&BaseOptions.MappingFile, "mapping", "", "mapping file")
	flags.IntVar(&BaseOptions.Srid, "srid", defaultSrid, "srs id")
	flags.StringVar(&BaseOptions.LimitTo, "limitto", "", "limit to geometries")
	flags.StringVar(&BaseOptions.ConfigFile, "config", "", "config (json)")
}

func UsageImport() {
	fmt.Fprintf(os.Stderr, "Usage: %s %s [args]\n\n", os.Args[0], os.Args[1])
	ImportFlags.PrintDefaults()
	os.Exit(2)
}

func UsageDiff() {
	fmt.Fprintf(os.Stderr, "Usage: %s %s [args] [.osc.gz, ...]\n\n", os.Args[0], os.Args[1])
	DiffFlags.PrintDefaults()
	os.Exit(2)
}

func init() {
	ImportFlags.Usage = UsageImport
	DiffFlags.Usage = UsageDiff

	addBaseFlags(DiffFlags)
	addBaseFlags(ImportFlags)
	ImportFlags.StringVar(&ImportOptions.Cpuprofile, "cpuprofile", "", "filename of cpu profile output")
	ImportFlags.StringVar(&ImportOptions.Httpprofile, "httpprofile", "", "bind address for profile server")
	ImportFlags.StringVar(&ImportOptions.Memprofile, "memprofile", "", "dir name of mem profile output and interval (fname:interval)")
	ImportFlags.BoolVar(&ImportOptions.Overwritecache, "overwritecache", false, "overwritecache")
	ImportFlags.BoolVar(&ImportOptions.Appendcache, "appendcache", false, "append cache")
	ImportFlags.StringVar(&ImportOptions.Read, "read", "", "read")
	ImportFlags.BoolVar(&ImportOptions.Write, "write", false, "write")
	ImportFlags.BoolVar(&ImportOptions.Optimize, "optimize", false, "optimize")
	ImportFlags.BoolVar(&ImportOptions.Diff, "diff", false, "enable diff support")
	ImportFlags.BoolVar(&ImportOptions.DeployProduction, "deployproduction", false, "deploy production")
	ImportFlags.BoolVar(&ImportOptions.RevertDeploy, "revertdeploy", false, "revert deploy to production")
	ImportFlags.BoolVar(&ImportOptions.RemoveBackup, "removebackup", false, "remove backups from deploy")
	ImportFlags.BoolVar(&ImportOptions.Quiet, "quiet", false, "quiet log output")
}

func ParseImport(args []string) {
	if len(args) == 0 {
		UsageImport()
	}
	err := ImportFlags.Parse(args)
	if err != nil {
		log.Fatal(err)
	}
	err = BaseOptions.updateFromConfig()
	if err != nil {
		log.Fatal(err)
	}
	errs := BaseOptions.check()
	if len(errs) != 0 {
		reportErrors(errs)
		UsageImport()
	}
}

func ParseDiffImport(args []string) {
	if len(args) == 0 {
		UsageDiff()
	}
	err := DiffFlags.Parse(args)
	if err != nil {
		log.Fatal(err)
	}

	err = BaseOptions.updateFromConfig()
	if err != nil {
		log.Fatal(err)
	}

	errs := BaseOptions.check()
	if len(errs) != 0 {
		reportErrors(errs)
		UsageDiff()
	}
}

func reportErrors(errs []error) {
	fmt.Println("errors in config/options:")
	for _, err := range errs {
		fmt.Printf("\t%s\n", err)
	}
	os.Exit(1)
}
