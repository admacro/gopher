package main

import "fmt"

type Car struct {
	model string
}

func (c *Car) register() {
	fmt.Println("start engine...")
}

func (c Car) recycle() {
	fmt.Println("recycle car...")
}

type MedicalRecord struct {
	vision     float32
	impairment bool
}

func (mi MedicalRecord) updateMedicalRecord() {
	fmt.Printf("update medical info [vision: %f, impairment: %v]\n", mi.vision, mi.impairment)
}

func (mi *MedicalRecord) expireMedicalRecord() {
	fmt.Printf("expire medical info [vision: %f, impairment: %v]\n", mi.vision, mi.impairment)
}

type Driver struct {
	*Car
	MedicalRecord
}

func main() {
	driver := Driver{
		MedicalRecord: MedicalRecord{
			vision:     2.0,
			impairment: false,
		},
		Car: &Car{
			model: "Toyota RAV4",
		},
	}
	fmt.Printf("%#v\n", driver.Car)
	driver.register()
	driver.recycle()
	driver.Car.register()
	driver.updateMedicalRecord()
	driver.expireMedicalRecord()
}
