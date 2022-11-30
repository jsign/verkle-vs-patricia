package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/jsign/go-ethereum/common"
	"github.com/jsign/go-ethereum/core/rawdb"
	"github.com/jsign/go-ethereum/core/types"
	"github.com/jsign/go-ethereum/rlp"
	"github.com/jsign/go-ethereum/trie"
	"github.com/jsign/verkle-vs-patricia/cmd/analytics/histogram"
)

var emptyRoot = common.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")

func main() {
	snapshotPath := "/mnt/data/.ethereum/geth/chaindata"
	db, err := rawdb.NewLevelDBDatabase(snapshotPath, 1024, 2000, "", true)
	if err != nil {
		log.Fatalf("opening leveldb db: %s", err)
	}

	head := rawdb.ReadHeadBlock(db)
	if head == nil {
		log.Fatalf("get head block: %s", err)
	}

	triedb := trie.NewDatabase(db)

	t, err := trie.NewStateTrie(trie.StateTrieID(head.Root()), triedb)
	if err != nil {
		log.Fatalf("new state trie: %s", err)
	}

	ctx := context.Background()
	ctx, cls := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cls()

	analyzeTries(ctx, head.Root(), t, triedb)
}

func analyzeTries(ctx context.Context, trieRoot common.Hash, t *trie.StateTrie, triedb *trie.Database) {
	var leafNodes int
	var storageTries int64
	lastReport := time.Now()

	// Histograms for the State Trie.
	histStateTrieDepths := histogram.New[int]("State Trie - Depths")
	histStatePathTypes := histogram.New[string]("State Trie - Path types")

	// Histograms for Storage Tries.
	histStorageTrieDepths := histogram.New[int]("Storage Trie - Depths")
	histStorageTriePathTypes := histogram.New[string]("Storage Trie - Path types")
	histStorageTriesNumSlots := histogram.New[int64]("Storage Trie - Number of non-zero slots")

	iter := t.NodeIterator(nil)
	for iter.Next(true) {
		if iter.Leaf() {
			leafNodes++

			// State Trie analysis.
			pathNodeTypes := iter.Stack()
			histStateTrieDepths.Observe(len(pathNodeTypes))
			histStatePathTypes.Observe(toShortPathTypes(pathNodeTypes))

			// Storage tries analysis.
			var acc types.StateAccount
			if err := rlp.DecodeBytes(iter.LeafBlob(), &acc); err != nil {
				log.Fatalf("invalid account encountered during traversal: %s", err)
			}
			if acc.Root != emptyRoot {
				storageTries++
				id := trie.StorageTrieID(trieRoot, common.BytesToHash(iter.LeafKey()), acc.Root)
				storageTrie, err := trie.NewStateTrie(id, triedb)
				if err != nil {
					log.Fatalf("failed to open storage trie: %s", err)
				}

				var storageTriesNumSlots int64
				storageIter := storageTrie.NodeIterator(nil)
				for storageIter.Next(true) {
					if storageIter.Leaf() {
						if ctx.Err() != nil {
							break
						}
						storageTriesNumSlots += 1

						pathNodeTypes := storageIter.Stack()
						histStorageTrieDepths.Observe(len(pathNodeTypes))
						histStorageTriePathTypes.Observe(toShortPathTypes(pathNodeTypes))
					}
				}
				histStorageTriesNumSlots.Observe(storageTriesNumSlots)

				if storageIter.Error() != nil {
					log.Fatalf("Failed to traverse storage trie: %s", err)
				}
			}
		}

		if time.Since(lastReport) > time.Second*5 {
			// State Trie stdout reports.
			fmt.Printf("Walked %d (EOA + SC) accounts:\n", leafNodes)
			histStateTrieDepths.Print(os.Stdout)
			histStatePathTypes.Print(os.Stdout)
			fmt.Println()

			// Storage tries stdout reports.
			fmt.Printf("Walked %d Storage Tries:\n", storageTries)
			histStorageTrieDepths.Print(os.Stdout)
			histStorageTriePathTypes.Print(os.Stdout)
			fmt.Printf("-----\n\n")

			lastReport = time.Now()
		}

		if ctx.Err() != nil {
			// State Trie.
			histStateTrieDepths.ToCSV("statetrie_depth.csv")
			histStatePathTypes.ToCSV("statetrie_pathtypes.csv")

			// Storage Tries.
			histStorageTrieDepths.ToCSV("storagetrie_depth.csv")
			histStorageTriePathTypes.ToCSV("storagetrie_pathtypes.csv")
			histStorageTriesNumSlots.ToCSV("storagetrie_numslots.csv")

			return
		}
	}
	if iter.Error() != nil {
		log.Fatalf("iterating trie: %s", iter.Error())
	}
}

func toShortPathTypes(nodeTypes []string) string {
	var sb strings.Builder
	for _, nodeType := range nodeTypes {
		switch nodeType {
		case "trie.valueNode":
			sb.WriteString("L.")
		case "*trie.shortNode":
			sb.WriteString("E.")
		case "*trie.fullNode":
			sb.WriteString("B.")
		default:
			panic("unkown node type")
		}
	}
	strPath := sb.String()

	return strPath[:len(strPath)-1]
}
