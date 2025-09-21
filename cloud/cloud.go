package cloud

type CloudDb struct {
	url string
}

func NewCLoudDb(url string) *CloudDb {
	return &CloudDb{
		url: url,
	}
}

func (db *CloudDb) Read() ([]byte, error) {
	return []byte{}, nil
}

func (db *CloudDb) Write(content []byte) {
}
