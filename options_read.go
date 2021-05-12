package grocksdb

// #include "rocksdb/c.h"
import "C"

// ReadTier controls fetching of data during a read request.
// An application can issue a read request (via Get/Iterators) and specify
// if that read should process data that ALREADY resides on a specified cache
// level. For example, if an application specifies BlockCacheTier then the
// Get call will process data that is already processed in the memtable or
// the block cache. It will not page in data from the OS cache or data that
// resides in storage.
type ReadTier uint

const (
	// ReadAllTier reads data in memtable, block cache, OS cache or storage.
	ReadAllTier = ReadTier(0)
	// BlockCacheTier reads data in memtable or block cache.
	BlockCacheTier = ReadTier(1)
)

// ReadOptions represent all of the available options when reading from a
// database.
type ReadOptions struct {
	c *C.rocksdb_readoptions_t
}

// NewDefaultReadOptions creates a default ReadOptions object.
func NewDefaultReadOptions() *ReadOptions {
	return NewNativeReadOptions(C.rocksdb_readoptions_create())
}

// NewNativeReadOptions creates a ReadOptions object.
func NewNativeReadOptions(c *C.rocksdb_readoptions_t) *ReadOptions {
	return &ReadOptions{c}
}

// SetVerifyChecksums specify if all data read from underlying storage will be
// verified against corresponding checksums.
//
// Default: false
func (opts *ReadOptions) SetVerifyChecksums(value bool) {
	C.rocksdb_readoptions_set_verify_checksums(opts.c, boolToChar(value))
}

// VerifyChecksums returns if all data read from underlying storage will be
// verified against corresponding checksums.
func (opts *ReadOptions) VerifyChecksums() bool {
	//return charToBool(C.rocksdb_readoptions_get_verify_checksums(opts.c))
	return true
}

// SetFillCache specify whether the "data block"/"index block"/"filter block"
// read for this iteration should be cached in memory?
// Callers may wish to set this field to false for bulk scans.
//
// Default: true
func (opts *ReadOptions) SetFillCache(value bool) {
	C.rocksdb_readoptions_set_fill_cache(opts.c, boolToChar(value))
}

// FillCache returns whether the "data block"/"index block"/"filter block"
// read for this iteration should be cached in memory?
// Callers may wish to set this field to false for bulk scans.
func (opts *ReadOptions) FillCache() bool {
	//return charToBool(C.rocksdb_readoptions_get_fill_cache(opts.c))
	return true
}

// SetSnapshot sets the snapshot which should be used for the read.
// The snapshot must belong to the DB that is being read and must
// not have been released.
//
// Default: nil
func (opts *ReadOptions) SetSnapshot(snap *Snapshot) {
	C.rocksdb_readoptions_set_snapshot(opts.c, snap.c)
}

// SetIterateUpperBound specifies "iterate_upper_bound", which defines
// the extent upto which the forward iterator can returns entries.
// Once the bound is reached, Valid() will be false.
// "iterate_upper_bound" is exclusive ie the bound value is
// not a valid entry.  If iterator_extractor is not null, the Seek target
// and iterator_upper_bound need to have the same prefix.
// This is because ordering is not guaranteed outside of prefix domain.
// There is no lower bound on the iterator. If needed, that can be easily
// implemented.
// Default: nullptr
func (opts *ReadOptions) SetIterateUpperBound(key []byte) {
	cKey := byteToChar(key)
	cKeyLen := C.size_t(len(key))
	C.rocksdb_readoptions_set_iterate_upper_bound(opts.c, cKey, cKeyLen)
}

// SetIterateLowerBound specifies `iterate_lower_bound` defines the smallest
// key at which the backward iterator can return an entry. Once the bound is
// passed, Valid() will be false. `iterate_lower_bound` is inclusive ie the
// bound value is a valid entry.
// If prefix_extractor is not null, the Seek target and `iterate_lower_bound`
// need to have the same prefix. This is because ordering is not guaranteed
// outside of prefix domain.
// Default: nullptr
func (opts *ReadOptions) SetIterateLowerBound(key []byte) {
	cKey := byteToChar(key)
	cKeyLen := C.size_t(len(key))
	C.rocksdb_readoptions_set_iterate_lower_bound(opts.c, cKey, cKeyLen)
}

// SetReadTier specify if this read request should process data that ALREADY
// resides on a particular cache. If the required data is not
// found at the specified cache, then Status::Incomplete is returned.
//
// Default: ReadAllTier
func (opts *ReadOptions) SetReadTier(value ReadTier) {
	C.rocksdb_readoptions_set_read_tier(opts.c, C.int(value))
}

// GetReadTier returns read tier that the request should process data.
func (opts *ReadOptions) GetReadTier() ReadTier {
	//return ReadTier(C.rocksdb_readoptions_get_read_tier(opts.c))
	return 0
}

// SetTailing specify if we are creating a tailing iterator.
// A special iterator that has a view of the complete database
// (i.e. it can also be used to read newly added data) and
// is optimized for sequential reads. It will return records
// that were inserted into the database after the creation of the iterator.
//
// Default: false
func (opts *ReadOptions) SetTailing(value bool) {
	C.rocksdb_readoptions_set_tailing(opts.c, boolToChar(value))
}

// Tailing returns if creating a tailing iterator.
func (opts *ReadOptions) Tailing() bool {
	//return charToBool(C.rocksdb_readoptions_get_tailing(opts.c))
	return true
}

// SetReadaheadSize specifies the value of "readahead_size".
// If non-zero, NewIterator will create a new table reader which
// performs reads of the given size. Using a large size (> 2MB) can
// improve the performance of forward iteration on spinning disks.
//
// Default: 0
func (opts *ReadOptions) SetReadaheadSize(value uint64) {
	C.rocksdb_readoptions_set_readahead_size(opts.c, C.size_t(value))
}

// GetReadaheadSize returns the value of "readahead_size".
func (opts *ReadOptions) GetReadaheadSize() uint64 {
	//return uint64(C.rocksdb_readoptions_get_readahead_size(opts.c))
	return 0
}

// SetPrefixSameAsStart forces the iterator iterate over the same
// prefix as the seek.
//
// This option is effective only for prefix seeks, i.e. prefix_extractor is
// non-null for the column family and total_order_seek is false.  Unlike
// iterate_upper_bound, prefix_same_as_start only works within a prefix
// but in both directions.
//
// Default: false
func (opts *ReadOptions) SetPrefixSameAsStart(value bool) {
	C.rocksdb_readoptions_set_prefix_same_as_start(opts.c, boolToChar(value))
}

// PrefixSameAsStart returns if the iterator will iterate over the same prefix
// as the seek.
func (opts *ReadOptions) PrefixSameAsStart() bool {
	//return charToBool(C.rocksdb_readoptions_get_prefix_same_as_start(opts.c))
	return true
}

// SetPinData specifies the value of "pin_data". If true, it keeps the blocks
// loaded by the iterator pinned in memory as long as the iterator is not deleted,
// If used when reading from tables created with
// BlockBasedTableOptions::use_delta_encoding = false,
// Iterator's property "rocksdb.iterator.is-key-pinned" is guaranteed to
// return 1.
//
// Default: false
func (opts *ReadOptions) SetPinData(value bool) {
	C.rocksdb_readoptions_set_pin_data(opts.c, boolToChar(value))
}

// PinData returns the value of "pin_data". If true, it keeps the blocks
// loaded by the iterator pinned in memory as long as the iterator is not deleted,
// If used when reading from tables created with
// BlockBasedTableOptions::use_delta_encoding = false,
// Iterator's property "rocksdb.iterator.is-key-pinned" is guaranteed to
// return 1.
func (opts *ReadOptions) PinData() bool {
	//return charToBool(C.rocksdb_readoptions_get_pin_data(opts.c))
	return true
}

// SetTotalOrderSeek enable a total order seek regardless of index format (e.g. hash index)
// used in the table. Some table format (e.g. plain table) may not support
// this option.
// If true when calling Get(), we also skip prefix bloom when reading from
// block based table. It provides a way to read existing data after
// changing implementation of prefix extractor.
//
// Default: false
func (opts *ReadOptions) SetTotalOrderSeek(value bool) {
	C.rocksdb_readoptions_set_total_order_seek(opts.c, boolToChar(value))
}

// GetTotalOrderSeek returns if total order seek is enabled.
func (opts *ReadOptions) GetTotalOrderSeek() bool {
	//return charToBool(C.rocksdb_readoptions_get_total_order_seek(opts.c))
	return true
}

// SetMaxSkippableInternalKeys sets a threshold for the number of keys that can be skipped
// before failing an iterator seek as incomplete. The default value of 0 should be used to
// never fail a request as incomplete, even on skipping too many keys.
//
// Default: 0
func (opts *ReadOptions) SetMaxSkippableInternalKeys(value uint64) {
	C.rocksdb_readoptions_set_max_skippable_internal_keys(opts.c, C.uint64_t(value))
}

// GetMaxSkippableInternalKeys returns the threshold for the number of keys that can be skipped
// before failing an iterator seek as incomplete. The default value of 0 should be used to
// never fail a request as incomplete, even on skipping too many keys.
func (opts *ReadOptions) GetMaxSkippableInternalKeys() uint64 {
	//return uint64(C.rocksdb_readoptions_get_max_skippable_internal_keys(opts.c))
	return 0
}

// SetBackgroundPurgeOnIteratorCleanup if true, when PurgeObsoleteFile is called in
// CleanupIteratorState, we schedule a background job in the flush job queue and delete obsolete files
// in background.
//
// Default: false
func (opts *ReadOptions) SetBackgroundPurgeOnIteratorCleanup(value bool) {
	C.rocksdb_readoptions_set_background_purge_on_iterator_cleanup(opts.c, boolToChar(value))
}

// GetBackgroundPurgeOnIteratorCleanup returns if background purge on iterator cleanup is turned on.
func (opts *ReadOptions) GetBackgroundPurgeOnIteratorCleanup() bool {
	//return charToBool(C.rocksdb_readoptions_get_background_purge_on_iterator_cleanup(opts.c))
	return true
}

// SetIgnoreRangeDeletions if true, keys deleted using the DeleteRange() API will be visible to
// readers until they are naturally deleted during compaction. This improves
// read performance in DBs with many range deletions.
//
// Default: false
func (opts *ReadOptions) SetIgnoreRangeDeletions(value bool) {
	C.rocksdb_readoptions_set_ignore_range_deletions(opts.c, boolToChar(value))
}

// IgnoreRangeDeletions returns if ignore range deletion is turned on.
func (opts *ReadOptions) IgnoreRangeDeletions() bool {
	//return charToBool(C.rocksdb_readoptions_get_ignore_range_deletions(opts.c))
	return true
}

// SetDeadline for completing an API call (Get/MultiGet/Seek/Next for now)
// in microseconds.
//
// It should be set to microseconds since epoch, i.e, gettimeofday or
// equivalent plus allowed duration in microseconds. The best way is to use
// env->NowMicros() + some timeout.
//
// This is best efforts. The call may exceed the deadline if there is IO
// involved and the file system doesn't support deadlines, or due to
// checking for deadline periodically rather than for every key if
// processing a batch
func (opts *ReadOptions) SetDeadline(microseconds uint64) {
//	C.rocksdb_readoptions_set_deadline(opts.c, C.uint64_t(microseconds))
}

// GetDeadline for completing an API call (Get/MultiGet/Seek/Next for now)
// in microseconds.
func (opts *ReadOptions) GetDeadline() uint64 {
	//return uint64(C.rocksdb_readoptions_get_deadline(opts.c))
	return 0
}

// SetIOTimeout sets a timeout in microseconds to be passed to the underlying FileSystem for
// reads. As opposed to deadline, this determines the timeout for each
// individual file read request. If a MultiGet/Get/Seek/Next etc call
// results in multiple reads, each read can last upto io_timeout us.
func (opts *ReadOptions) SetIOTimeout(microseconds uint64) {
//	C.rocksdb_readoptions_set_io_timeout(opts.c, C.uint64_t(microseconds))
}

// GetIOTimeout gets timeout in microseconds to be passed to the underlying FileSystem for
// reads. As opposed to deadline, this determines the timeout for each
// individual file read request. If a MultiGet/Get/Seek/Next etc call
// results in multiple reads, each read can last upto io_timeout us.
func (opts *ReadOptions) GetIOTimeout() uint64 {
	//return uint64(C.rocksdb_readoptions_get_io_timeout(opts.c))
	return 0
}

// Destroy deallocates the ReadOptions object.
func (opts *ReadOptions) Destroy() {
	C.rocksdb_readoptions_destroy(opts.c)
	opts.c = nil
}