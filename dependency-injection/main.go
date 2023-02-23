package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// Dependency injection involves four roles:

// - the service object(s) to be used - also known as "dependency"
// - the client object that is depending on the service(s) it uses
// - the interfaces that define how the client may use the services
// - the injector, which is responsible for constructing the services
//   and injecting them into the client

// As an analogy,

// service - an electric, gas, hybrid, or diesel car
// client - a driver who uses the car the same way regardless of the engine
// interface - automatic, ensures driver does not have to understand engine details like gears
// injector - the parent who bought the kid the car and decided which kind

var (
	errUsage = errors.New(`usage:
	set <key> <value>  Set specified key and value
	get <key>	   Get specified key`)
)

// main is the injector
func main() {
	runner := newRunner(newFileDatabase("database.txt"))
	if err := runner.run(os.Stdout, os.Args); err != nil {
		fmt.Println(err)
	}
}

// storage is the interface
type storage interface {
	Set(string, string) error
	Get(string) (string, error)
}

// runner is the client
type runner struct {
	database storage
}

func newRunner(db storage) runner {
	return runner{db}
}

func (r runner) run(output io.StringWriter, args []string) error {
	if len(args) < 3 {
		return errUsage
	}
	switch args[1] {
	case "set":
		if len(args) < 4 {
			return errUsage
		}
		if err := r.database.Set(args[2], args[3]); err != nil {
			return err
		}
	case "get":
		v, err := r.database.Get(args[2])
		if err != nil {
			return err
		}
		output.WriteString(v + "\n")
	default:
		return errUsage
	}

	return nil
}

// fileDatabase is the service
type fileDatabase struct {
	file string
}

func newFileDatabase(path string) fileDatabase {
	return fileDatabase{path}
}

func (db fileDatabase) Set(key, value string) error {
	f, err := os.OpenFile(db.file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("%s:%s\n", key, value)); err != nil {
		return err
	}

	return nil
}

func (db fileDatabase) Get(key string) (string, error) {
	f, err := os.OpenFile(db.file, os.O_RDONLY, 0600)
	if err != nil {
		return "", err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var last string
	for scanner.Scan() {
		row := scanner.Text()
		parts := strings.Split(row, ":")
		if len(parts) < 2 {
			return "", errors.New("invalid record")
		}
		if parts[0] == key {
			last = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return "", errors.Wrap(err, "scanner.Err")
	}

	if last != "" {
		return last, nil
	}

	return "", errors.New("not found")
}
