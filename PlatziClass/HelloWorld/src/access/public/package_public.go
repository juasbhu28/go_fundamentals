package public

// CarPblic is a package public struct
type CarPublic struct {
	// Propiedades publicas
	PublicBrand string
	PublicYear  int
}

//Es buena práctica usar documentación para cada paquete

//Si la primera letra en GO es mayuscula, es publica sino es privada
//Eso quiere decir que si queremos que una propiedad sea publica,
//la primera letra debe ser mayuscula
