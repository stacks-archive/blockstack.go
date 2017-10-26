package indexer

// import (
// 	"log"
// )

// processNameZonefileMap takes a map[name]zonefile and returns the go representation
// TODO: Check this
// func (i *Indexer) processNameZonefileMap() {
// 	for name := range i.NameZonefileHash {
// 		i.Domains = append(i.Domains, NewDomain(name, i.NameZonefileHash[name]))
// 	}
// }
//
// // fetchAllZonefiles fetches all the zonefiles from the startBlock to the CurrentBlock
// func (i *Indexer) fetchAllZonefiles() {
// 	numBlocks := i.CurrentBlock - i.StartBlock
// 	fetchBlocks := 100
// 	iter := (numBlocks/fetchBlocks + 1)
// 	for page := 0; page <= iter; page++ {
// 		st := i.StartBlock + (page * fetchBlocks)
// 		end := st + 100
// 		log.Printf("%v Fetching zonefiles from block %v to block %v", logPrefix, st, end)
// 		i.GetZonefilesByBlock(st, end)
// 	}
// }
//
// // GetZonefilesByBlock returns a map[name]zonefile
// // TODO: Parallelize - difficult, need to keep blocks in order. Maybe hold off on this
// func (i *Indexer) GetZonefilesByBlock(startBlock, endBlock int) {
// 	zfhrs := make([]blockstack.ZonefileHashResults, 0)
// 	iter := 0
// 	for iter < 100000 {
// 		// Fetch batch of Zonefiles by block
// 		res, err := i.client().GetZonefilesByBlock(startBlock, endBlock, (iter * 100), 100)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		iter++
//
// 		// Make a batch that maps to a get_zonefiles rpc call
// 		batch := make([]blockstack.ZonefileHashResult, 0)
// 		for _, zfhrs := range res.ZonefileInfo {
// 			batch = append(batch, zfhrs)
// 		}
//
// 		// Save the batch
// 		zfhrs = append(zfhrs, blockstack.ZonefileHashResults(batch))
//
// 		// If the batch doesn't have 100 records then stop
// 		if len(res.ZonefileInfo) != 100 {
// 			iter = 100000
// 		}
// 	}
//
// 	// Loop over the batches and call get_zonefile for each
// 	for _, batch := range zfhrs {
// 		res, err := i.client().GetZonefiles(batch.Zonefiles())
// 		if err != nil {
// 			log.Fatal(err)
// 		}
//
// 		// Decode the base64 encoded zonefiles
// 		dec := res.Decode()
//
// 		// Range over the decoded zonefiles and associate them with names
// 		for zfh := range dec {
// 			l := batch.LatestZonefileHash(zfh)
// 			i.NameZonefileHash[l.Name] = dec[zfh]
// 		}
// 	}
// 	i.stats.LastProcessed.Set(float64(endBlock))
// }
