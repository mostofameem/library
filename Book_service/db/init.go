package db

func init() {
	// initialize resource type repository
	initBookTypeRepo()
	initBorrowTypeRepo()
	// initialize audit log repository
	//initProductTypeRepo()

	//initUserTypeRepo()
}
