package nexus

import "zenrailz/service"

type Store struct {
	nativeSvc  service.Native
	railwaySvc service.Railway
}

func NewStore(
	nativeService service.Native,
	railwayService service.Railway,
) *Store {
	return &Store{
		nativeSvc:  nativeService,
		railwaySvc: railwayService,
	}
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
