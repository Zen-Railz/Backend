package nexus

import "zenrailz/service"

func NewStore(
	nativeService service.Native,
	railwayService service.Railway,
) *Store {
	return &Store{
		nativeSvc:  nativeService,
		railwaySvc: railwayService,
	}
}
