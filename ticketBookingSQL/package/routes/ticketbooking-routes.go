package routes
import(
	"github.com/prantodev/go-ticket/package/controllers"
	"github.com/gorilla/mux"

)
var  RegisterTicketsRoutes= func (router *mux.Router)  {
	router.HandleFunc("/",controllers.Index)
	router.HandleFunc("/ticket/",controllers.CreateTickets).Methods("POST") 
	router.HandleFunc("/ticket/get", controllers.GetGuest).Methods("GET") 
	// router.HandleFunc("/ticket/{userID}}",controllers.BookTicket).Methods("POST")
	// router.HandleFunc("/ticket/{userID}}",controllers.UpdateGuest).Methods("PUT")
	// router.HandleFunc("/ticket/{userID}}",controllers.DeleteTicket).Methods("DELETE")
} 