package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (C App) Sueldoutm (S float64) float64{
	utm := 47729.0
	var Sutm = S/utm
	return Sutm
}



func (c App) FormulaSueldoLiquido(S float64) float64{
	sueldobrutoi := (S)*0.176
	utm := 47729.0
	sueldoliquidoi:= S- sueldobrutoi - c.Impuesto(c.Sueldoutm(S)) * utm
	return sueldoliquidoi
	
}
	
func (c App) Impuesto (SU float64) float64{
	var imp float64
	if SU > 0 && SU  < 13.5{
		imp = 0}
	if SU > 13.5 && SU  <30{
		imp = SU * 0.04 - 0.54}
	if SU > 30 && SU  <50{
		imp = SU * 0.08 - 1.74}
	if SU > 50 && SU  <70{
		imp = SU * 0.135 - 4.49}
	if SU > 70 &&SU  <90{
		imp = SU * 0.23 - 11.14}
	if SU > 90 &&SU  <120{
		imp = SU * 0.304 - 17.8}
	if SU > 120 {
		imp = SU * 0.35 - 23.32}
	return imp
}

func (c App) Index() revel.Result {
	greeting:="Calculadora de Sueldo Líquido y Bruto"
	return c.Render(greeting)
}

func (c App) Sueldo(S float64) revel.Result {
		utm := 47729.0
		sueldoliquido3:=c.FormulaSueldoLiquido(S)
		impuesto:= c.Impuesto(c.Sueldoutm(S)) * utm
		sueldobrutoi := S-(S * 0.0176)
		cesantia:= S * 0.006
		salud := S*0.07
		afp := S*0.104
		sueldobrutoi2 := S *0.82
		return c.Render(sueldoliquido3, sueldobrutoi, cesantia, salud, afp, sueldobrutoi2, impuesto)
}

func (c App) Sueldo2(L float64) revel.Result {
	utm := 47729.0
	impuestoI := c.ImpuestoI(c.Sueldoutm(L)) * utm
	sbruto:= (impuestoI + L) / 0.824
	return c.Render (impuestoI, sbruto)
	}

func (c App) ImpuestoI (SUT float64) float64{
	var impi float64
	if SUT < 13.5{
		impi = 0}
	if SUT >= 13.5 && SUT <29.34{
		impi = ((SUT-13.5)/24)}
	if SUT >=29.34 && SUT<47.739{
		impi = ((SUT-29.34/11.5)+0.66)}
	if SUT >=47.739 && SUT < 65.0399{
		impi = (((SUT-47.739)/6.40742) + 2.26)}
	if SUT >=65.0399 && SUT <80.44{
		impi = (((SUT-65.0399)/3.3478) + 4.96)}
	if SUT >= 80.44 && SUT <101.319{
		impi = (((SUT - 80.44)/2.28947) + 9.56)}
	if SUT >=101.319{
		impi = (((SUT - 103.319)/1.857141) + 18.68)}
	return impi
}