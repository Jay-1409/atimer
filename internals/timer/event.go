package event;

import {
	"net/http"
}

type TimerEventHandler struct {
ID        string
}

fun shootEvent (url, taskId int) {
	resp, err := http.Post(urls, "application/json");
	if err != nil {
    	log.Fatalf("An Error Occured %v", err);
   	}
	else {
		log.Printf("Shot task with ID", taskId);
   }
}





