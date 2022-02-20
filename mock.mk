mock: mock/mock_store

mock/mock_store: mock/mock_store/client.go

mock/mock_store/client.go: store/client.go
	$(GOGEN_MOCKERY) store/client.go