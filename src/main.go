package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// To Do
// Directory Structure ?
// Nested Type List Factura - Documentos - Autorizaciones
// Obtener Detectado GET
// Registrar Venta POST
// Servicio Adicional - GET Extern

type ResultadoOperacion struct {
	Id        int
	Resultado int
}

type Detectado struct {
	IdDetectado            int
	IdTransaccion          string
	Identificador          string
	Imagen                 string
	IdTipoConsumidor       int
	TieneServicioAdicional int
	Estado                 int
	MensajePantalla        string
	MensajeImpresion       string
}

type RegistroVenta struct {
	Llave                 string
	IdVenta               int
	Volumen               float32
	PrecioUnitario        float32
	FechaHoraEnvio        int
	IdCodigoPais          int
	IdTipoConsumidor      int
	IdFormaPago           int
	IdProducto            int
	IdManguera            int
	NitConsumidor         int
	RazonSocialConsumidor string
	IdPreVenta            int
	Usuario               string
	ListaFacturas         []Factura
	ListaDocumentosVenta  []DocumentoVenta
	ListaAutorizaciones   []Autorizacion
}

type Factura struct {
	NitEmisor           int
	IdCodigoEmisor      int
	NroFactura          int
	NroAutorizacion     int
	FechaHoraEmision    int
	Importe             float32
	CodigoControl       string
	IdTipoFactura       int
	IdTipoCreditoFiscal int
	IdServicioAdicional int
}

type DocumentoVenta struct {
	NitEmisor        int
	NroDocumento     int
	Importe          float32
	FechaHoraEmision int
	IdTipoDocumento  int
	//ListaVales
	IdServicioAdicional int
}

type Autorizacion struct {
	IdAutorizacion     int
	CodigoAutorizacion string
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to Middleware 3.0!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Welcome to Middleware 3.0!\n")
}

func RegistrarVenta(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Escribir en BD 
	// Mandar a WS Core
	// 
	
	w.Header().Set(("Content-Type"), "application/json; charset=utf-8")
	r.ParseForm()

	var ve RegistroVenta
	err := json.NewDecoder(r.Body).Decode(&ve)

	if err != nil {
		log.Println("Error", err)
	}
	//log.Println(t.Test)

	fmt.Println("=========================")
	fmt.Println("Llave:", ve.Llave)
	fmt.Println("Nit Consumidor:", ve.NitConsumidor)
	fmt.Println("Razon social:", ve.RazonSocialConsumidor)
	fmt.Println("Usuario:", ve.Usuario)
	fmt.Println("Volumen:", ve.Volumen)
	fmt.Println("Precio Unitario:", ve.PrecioUnitario)
	fmt.Println("Fecha Hora:", ve.FechaHoraEnvio)

	if len(ve.ListaFacturas) > 0 {
		fmt.Println("Codigo Control:", ve.ListaFacturas[0].CodigoControl)
	}

	//value := "20141103103050"
	//layout := "Mon, 01/02/06, 03:04PM"
	//t, _ := time.Parse(layout, value)

	re := &ResultadoOperacion{Id: 1, Resultado: 1}
	var ob, _ = json.Marshal(re)
	fmt.Fprintf(w, "%s", ob)
}

func ObtenerDetectadosRfid(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Welcome to Middleware 3.0!\n")
}

func ObtenerDetectados(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set(("Content-Type"), "application/json; charset=utf-8")
	w.Header().Set(("Access-Control-Allow-Origin"), "*")
	w.Header().Set(("Access-Control-Allow-Methods"), "GET, POST")
	w.Header().Set(("Access-Control-Allow-Headers"), "Content-Type")
	w.Header().Set(("X-Powered-On"), "AGENCIA NACIONAL DE HIDROCARBUROS")
	w.Header().Set(("Cache-Control"), "no-cache, private, must-revalidate, max-stale=0, post-check=0, pre-check=0, no-store")
	w.Header().Set(("Pragma"), "no-cache")
	w.Header().Set(("Expires"), "0")
	w.Header().Set(("Vary"), "*")

	var pp = Detectado{IdDetectado: 1,
		IdTransaccion:          "3091FSF",
		Identificador:          p.ByName("identificador"),
		IdTipoConsumidor:       1,
		TieneServicioAdicional: 1,
		Estado:                 1,
		MensajePantalla:        "HOLA MUNDO",
		MensajeImpresion:       "HOLA"}

	var ob, _ = json.Marshal(pp)
	//fmt.Println()
	//fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
	fmt.Fprintf(w, "%s", string(ob))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	// # Detectados
	router.GET("/obtenerdetectados/:llave", ObtenerDetectadosRfid)
	router.GET("/obtenerdetectados/:llave/:identificador/:tipoconsulta", ObtenerDetectados)
	router.GET("/obtenerservicioadicional/:llave/:idtransaccion", Hello)

	// # Catalogos
	router.GET("/obtenercatalogos/:llave", Hello)
	router.GET("/obtenertanque/:llave", Hello)
	router.GET("/obtenermanguera/:llave", Hello)

	// # Registros
	router.POST("/registrarventa/", RegistrarVenta)
	router.POST("/registrarfacturaanulada/", Hello)
	router.POST("/registrardescarga/", Hello)
	router.POST("/registrarmedicion/", Hello)

	log.Fatal(http.ListenAndServe(":28321", router))
}
