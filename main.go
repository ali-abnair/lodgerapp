package lodgerapp

import "log"

func LogInfo( mess string){
	log.Println("INFO - ", mess)
}

func LogWarning( mess string){
	log.Println("WARNING - ",  mess)
}

func LogError( mess string){
	log.Println("ERROR - ", mess)
}