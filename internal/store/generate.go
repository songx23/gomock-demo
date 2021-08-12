package store

//go:generate mockgen -destination=../../test/mocks/store/reflect/clients.go -package=mock_client -mock_names Client=MockDarkMagicClient,Villain=MockDarkVillain . Client,Villain

