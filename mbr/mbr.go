package mbr
import (
	"fmt"
	"os"
)

type MBR struct {
	BootstrapCode []byte
	TimeStamp string
	DiskSignature DiskSignature
	Partitions []Partition
	BootSignature int
}

type Partition struct {
	Position int
	Status string
	StatusCode int
	Type string
	TypeCode int
	Sectors int
}

type DiskSignature struct {
	CopyProtected bool 
	Signature int32
}

func Dump(f *os.File) ([]byte, error) {
	mbr := make([]byte, MBRsize)
	_, err := f.Read(mbr)
	if err != nil {
		return mbr, err
	}
	return mbr, nil
}

func ParseStatus(extract byte) string {
	var status string
	if extract == 128 {
		status = "bootable"
	} else if extract > 0 && extract < 128 {
		status = "invalid"
	} else {
		status = "inactive"
	}
	return status
}

func ParsePartition(entry []byte) (Partition, error) {
	var partition Partition
	if len(entry) < PartitionEntrySize {
		return partition, fmt.Errorf("input does not contain a valid partition entry")
	}
	partition.StatusCode = int(entry[0])
	partition.Status = ParseStatus(entry[0])

	partition.TypeCode = int(entry[4])


	return partition, nil
}

func ParsePartitions(entries []byte) ([]Partition, error) {
	if len(entries) < PartitionEntrySize * 4 {
		return make([]Partition, 0), fmt.Errorf("input does not contain all 4 partition entries")
	}
	partitions := make([]Partition, 0)

	part1, err := ParsePartition(entries[0:16])
	if err != nil {
		return partitions, fmt.Errorf("Could not parse partition entry 1: %s", err)
	}
	part1.Position = 1
	partitions = append(partitions, part1)

	part2, err := ParsePartition(entries[16:32])
	if err != nil {
		return partitions, fmt.Errorf("Could not parse partition entry 2: %s", err)
	}
	part2.Position = 2
	partitions = append(partitions, part2)

	part3, err := ParsePartition(entries[32:48])
	if err != nil {
		return partitions, fmt.Errorf("Could not parse partition entry 3: %s", err)
	}
	part3.Position = 3
	partitions = append(partitions, part3)
	
	part4, err := ParsePartition(entries[48:64])
	if err != nil {
		return partitions, fmt.Errorf("Could not parse partition entry 4: %s", err)
	}
	part4.Position = 4
	partitions = append(partitions, part4)

	return partitions, nil
}

func Parse(mbr []byte) (MBR, error) {
	var MBR MBR
	if len(mbr) < MBRsize {
		return MBR, fmt.Errorf("input does not contain full mbr")
	}

	partitions, err := ParsePartitions(mbr[446:510])
	if err != nil {
		return MBR, fmt.Errorf("Could not parse partition entries: %s", err)
	}
	MBR.Partitions = partitions

	return MBR, nil
}