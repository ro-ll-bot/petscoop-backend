package customer

type cacheRepository struct {
	// Just put a real redis plugin here to activate
	// an actual caching mechanism
	redis string
}

func (c *cacheRepository) Get(id string) *Customer {
	return nil
}

func (c *cacheRepository) Delete(id string) error {
	return nil	
}

func (c *cacheRepository) Update(customer Customer) error {
	return nil	
}

func (c *cacheRepository) Save(customer Customer) error {
	return nil	
}