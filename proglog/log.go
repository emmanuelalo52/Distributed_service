package server
import ("fmt"
"sync")

type Log struct{
	mu sync.Mutex
	records []Record
}

func NewLog() *Log{
	return &Log{}
}
func (c *Log) Append(record Record) (uint64, error){
	c.mu.Lock() // locks the mutex till the operation is mutually exclusive
	defer c.mu.Unlock() //unlocks when the condition of the sorrounding functions are met
	record.Offset = uint64(len(c.records)) //Sets the offset of the record to the current length of the records slice.
	c.records = append(c.records, record) //Appends the record to the records slice.
	return record.Offset,nil
}
func (c *Log) Read(offset uint64)(Record, error){
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)){ // Checks if the provided offset is valid; if not, it returns an empty record and the ErrOffsetNotFound error.
		return Record{},ErrOffsetNotFound
	}
	return c.records[offset], nil
}

type Record struct{
	Value []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")