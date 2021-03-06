package api

import (
	"bufio"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/cosmtrek/violet/engine/index"
	"github.com/cosmtrek/violet/pkg/analyzer"
	"github.com/cosmtrek/violet/pkg/utils"
	"github.com/pkg/errors"
)

// Indexer contains all indexes
type Indexer struct {
	Indexes   map[string]*index.Index
	Path      string
	Segmenter analyzer.Analyzer
}

// NewIndexer initializes indexer
func NewIndexer(path string, segmenter analyzer.Analyzer) (*Indexer, error) {
	indexer := &Indexer{
		Indexes: make(map[string]*index.Index),
		Path:    path,
	}
	var err error
	if segmenter == nil {
		seg, err := analyzer.New()
		if err != nil {
			return nil, errors.New("failed to new segmenter")
		}
		indexer.Segmenter = seg
	}
	if path == "/" || path == "./" {
		return nil, errors.New("indexer path error")
	}
	if utils.DirExists(path) {
		if err = os.RemoveAll(path); err != nil {
			return nil, err
		}
	}
	if err = os.MkdirAll(path, os.ModeDir|os.ModePerm); err != nil {
		return nil, err
	}
	return indexer, nil
}

// AddIndex initializes index meta
func (r *Indexer) AddIndex(name string, fields map[string]uint64) error {
	if _, ok := r.Indexes[name]; ok {
		return errors.New("index existed")
	}

	index, err := index.NewIndex(r.Path, name, r.Segmenter)
	if err != nil {
		return errors.New("failed to create index")
	}
	if err = index.IndexFields(fields); err != nil {
		return errors.New("failed to map index fields")
	}
	r.Indexes[name] = index
	return nil
}

// LoadDocumentsFromFile inserts documents into indexer
func (r *Indexer) LoadDocumentsFromFile(index string, file string, fieldType string, fields []string) error {
	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()
	log.Infof("load documents from file: %s\n", file)
	scanner := bufio.NewScanner(fd)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	doc := make(map[string]string)
	for scanner.Scan() {
		if fieldType == "text" {
			txt := scanner.Text()
			if txt == "" {
				continue
			}
			txts := strings.Split(txt, "  ")
			txtLen := len(txts)
			// map txts to fields
			for i, f := range fields {
				if txtLen <= i {
					doc[f] = ""
				} else {
					doc[f] = txts[i]
				}
			}
			if err = r.Indexes[index].AddDocument(doc); err != nil {
				log.Errorf("failed to add document %v into indexer, err: %s\n", doc, err.Error())
				return errors.Wrap(err, "failed to add document into indexer")
			}
		}
	}
	return r.Indexes[index].SyncToDisk()
}

// Search searches everything
func (r *Indexer) Search(index string, query string) ([]map[string]string, bool) {
	docs, found := r.Indexes[index].Search(query)
	if !found {
		return nil, false
	}
	var results []map[string]string
	for _, doc := range docs {
		d, ok := r.Indexes[index].GetDocument(doc.DocID)
		if ok {
			results = append(results, d)
		}
	}
	if len(results) == 0 {
		return nil, false
	}
	return results, true
}
