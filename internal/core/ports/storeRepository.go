package ports

type StoreRepository interface {
	GetStoreById(id string)
	GetStores()
}
