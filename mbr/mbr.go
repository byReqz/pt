package mbr
import (
	"fmt"
	"os"
)

const (
	MBRsize = 512

	BootstrapCodeArea1Size = 218
	BootstrapCodeArea1Offset = 0
	BootstrapCodeArea2Size = 216
	BootstrapCodeArea2Offset = 224
	LegacyBootstrapCodeAreaOffset = 0
	LegacyBootstrapCodeAreaSize = 446

	DiskTimestampOffset = 218
	DiskTimestampSize = 6
	DiskTimeStampSecondsOffset = 221
	DiskTimeStampSecondsSize = 1
	DiskTimeStampMinutesOffset = 222
	DiskTimeStampMinutesSize = 1
	DiskTimeStampHoursOffset = 223
	DiskTimeStampHoursSize = 1

	PartitionEntrySize = 16
	PartitionEntry1Offset = 446
	PartitionEntry2Offset = 462
	PartitionEntry3Offset = 478
	PartitionEntry4Offset = 494

	BootSignatureSize = 2
	BootSignatureOffset = 510
)

type MBR struct {
	BootstrapCode []byte
	TimeStamp string
	DiskSignature DiskSignature
	Partitions []Partition
	BootSignature int
}

type Partition struct {
	Status string
	Type string
	Sectors int
}

type DiskSignature struct {
	CopyProtected bool 
	Signature int32
}

func DumpMBR(f *os.File) ([]byte, error) {
	mbr := make([]byte, MBRsize)
	_, err := f.Read(mbr)
	if err != nil {
		return mbr, err
	}
	return mbr, nil
}

func GetMBR(mbr []byte) error {
	if len(mbr) < MBRsize {
		return fmt.Errorf("input does not contain full mbr")
	}

	fmt.Println(mbr[446:462]) // part entry 1
	fmt.Println(mbr[446]) // status  --  if 128 == bootable, if 0 == inactive
	fmt.Println(mbr[450]) // partition type https://en.wikipedia.org/wiki/Partition_type

	return nil
}