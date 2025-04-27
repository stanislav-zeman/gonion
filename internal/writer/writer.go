package writer

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/stanislav-zeman/gonion/internal/convertor"
	"github.com/stanislav-zeman/gonion/internal/layers"
)

const defaultFilePermissions = 0o755

type Writer struct {
	directory string
}

func NewWriter(directory string) Writer {
	return Writer{
		directory: directory,
	}
}

// ----------------------------------------------------------------------------

func (w *Writer) WriteDomainEntity(service, name string, data []byte) error {
	return w.writeFile(service, layers.DomainLayer, "entity", name, data)
}

func (w *Writer) WriteDomainValue(service, name string, data []byte) error {
	return w.writeFile(service, layers.DomainLayer, "value", name, data)
}

func (w *Writer) WriteDomainInterface(service, name string, data []byte) error {
	return w.writeFile(service, layers.DomainLayer, "interface", name, data)
}

func (w *Writer) WriteDomainService(service, name string, data []byte) error {
	directory := "service"
	return w.writeFile(service, layers.DomainLayer, directory, name+"_"+directory, data)
}

func (w *Writer) WriteDomainRepository(service, name string, data []byte) error {
	directory := "repository"
	return w.writeFile(service, layers.DomainLayer, directory, name+"_"+directory, data)
}

// ----------------------------------------------------------------------------

func (w *Writer) WriteApplicationCommand(service, name string, data []byte) error {
	directory := "command"
	return w.writeFile(service, layers.ApplicationLayer, directory, name+"_"+directory, data)
}

func (w *Writer) WriteApplicationQuery(service, name string, data []byte) error {
	directory := "query"
	return w.writeFile(service, layers.ApplicationLayer, directory, name+"_"+directory, data)
}

func (w *Writer) WriteApplicationInterface(service, name string, data []byte) error {
	return w.writeFile(service, layers.ApplicationLayer, "interface", name, data)
}

func (w *Writer) WriteApplicationService(service, name string, data []byte) error {
	directory := "service"
	return w.writeFile(service, layers.ApplicationLayer, directory, name+"_"+directory, data)
}

// ----------------------------------------------------------------------------

func (w *Writer) WriteInfrastructureRepository(service, name, typ string, data []byte) error {
	return w.writeFile(service, layers.InfrastructureLayer, filepath.Join("persistence", typ), name+"_repository", data)
}

func (w *Writer) writeFile(service, layer, directory, name string, data []byte) error {
	fp := filepath.Join(
		w.directory,
		"internal",
		service,
		layer,
		directory,
		name+".go",
	)

	fp = convertor.ToSnakeCase(fp)

	err := os.MkdirAll(filepath.Dir(fp), defaultFilePermissions)
	if err != nil {
		return fmt.Errorf("failed creating directories: %w", err)
	}

	log.Printf("writing to path: %v", fp)

	err = os.WriteFile(fp, data, defaultFilePermissions)
	if err != nil {
		return fmt.Errorf("failed writing file: %w", err)
	}

	return nil
}
