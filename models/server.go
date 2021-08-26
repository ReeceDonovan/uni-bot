package models

import "log"

// Get server object
func (server *Server) Get() error {
	log.Println("Getting server data: ", server.SID)
	return db.First(server).Error
}

// Create server object
func (server *Server) Create() error {
	log.Println("Creating server link: ", server)
	return db.Create(server).Error
}

// Put updates a server object
func (server *Server) Put(field string, value interface{}) {
	db.First(server).Update(field, value)
}

// Delete deletes a server's db entry
func (server *Server) Delete() error {
	return db.Delete(server).Error
}
