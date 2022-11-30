package histogram

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type BucketType interface {
	~string | ~int | ~int64
	constraints.Ordered
}

type Histogram[T BucketType] struct {
	title   string
	buckets map[T]int64
	total   int64
}

func New[T BucketType](title string) *Histogram[T] {
	return &Histogram[T]{
		title:   title,
		buckets: map[T]int64{},
	}
}

func (h *Histogram[T]) Observe(bucket T) {
	h.buckets[bucket]++
	h.total++
}

func (h *Histogram[T]) Print(w io.Writer) {
	fmt.Printf("%s:\n", h.title)
	sortedKeys := h.sortedKeys()
	len := len(sortedKeys)
	if len > 100 {
		len = 100
	}
	for _, k := range sortedKeys[:len] {
		fmt.Fprintf(w, "%v: %.02f%% (%d)\n", k, float64(h.buckets[k])/float64(h.total)*100, h.buckets[k])
	}
	fmt.Println()
}

func (h *Histogram[T]) ToCSV(filename string) error {
	var sb strings.Builder
	for _, k := range h.sortedKeys() {
		ratio := float64(h.buckets[k]) / float64(h.total) * 100
		if ratio < 0.01 {
			continue
		}
		sb.WriteString(fmt.Sprintf("%v,%.02f\n", k, ratio))
	}
	if err := os.WriteFile(filename, []byte(sb.String()), 0o755); err != nil {
		return fmt.Errorf("writing to file: %s", err)
	}
	log.Printf("saved %s\n", filename)

	return nil
}

func (h *Histogram[T]) sortedKeys() []T {
	var keys []T
	for k := range h.buckets {
		keys = append(keys, k)
	}
	slices.SortFunc(keys, func(a, b T) bool {
		return h.buckets[a] > h.buckets[b]
	})
	return keys
}
