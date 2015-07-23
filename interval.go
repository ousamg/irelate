package irelate

import (
	"strconv"
	"strings"
)

const empty = ""

// Interval satisfies the Relatable interface.
type Interval struct {
	// chrom, start, end, line, source, related[]
	chrom   string
	start   uint32
	end     uint32
	Fields  []string
	source  uint32
	related []Relatable
}

func (i *Interval) Chrom() string        { return i.chrom }
func (i *Interval) Start() uint32        { return i.start }
func (i *Interval) End() uint32          { return i.end }
func (i *Interval) Related() []Relatable { return i.related }
func (i *Interval) AddRelated(b Relatable) {
	if i.related == nil {
		i.related = make([]Relatable, 1, 4)
		i.related[0] = b
	} else {
		i.related = append(i.related, b)
	}
}
func (i *Interval) Source() uint32       { return i.source }
func (i *Interval) SetSource(src uint32) { i.source = src }

func (i *Interval) String() string {
	return strings.Join(i.Fields, "\t")
}

func IntervalFromBedLine(line string) (Relatable, error) {
	fields := strings.Split(line, "\t")
	start, err := strconv.ParseUint(fields[1], 10, 32)
	if err != nil {
		return nil, err
	}
	end, err := strconv.ParseUint(fields[2], 10, 32)
	if err != nil {
		return nil, err
	}
	i := Interval{chrom: fields[0], start: uint32(start), end: uint32(end), related: nil, Fields: fields}
	return &i, nil
}
