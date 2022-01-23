package mbr

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