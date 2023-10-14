package main

import "fmt"

type GroundVehicle interface {
	Drive() string
	OpenDoor() string
}
type Parts struct {
	Wheel int

}
type Car struct {
	Owner struct {
		Name string
		Age int
	}
	Type string
	Year string
	Brand string
	Parts
}
type MotorCycle struct {
	// This is nested structs
	Owner struct {
		Name string
		Age int
	}
	// This is embedded structs
	Parts
}

func (c Car) Drive() string {
	return "Wuuuuush wushhh";
}
func (c Car) OpenDoor() string {
	return "Blakkk"
}
func (m MotorCycle) Drive() string {
	return "Ngenggg ngennggg";
}
func (m MotorCycle) OpenDoor() string {
	return "Motor doesnt has door but it has step"
}

func main() {

	myCar := Car{Owner: struct{Name string; Age int}{Name: "Rizki", Age: 20}, Type: "Manual", Year: "2019", Brand: "Toyota"}
	myMotorCycle := MotorCycle{Owner: struct{Name string; Age int}{Name : "Rizakk", Age: 30}}
	myCar.Wheel = 10
	myMotorCycle.Wheel = 2
	onlyGroundVehicle(myMotorCycle)
	onlyGroundVehicle(myCar)

}

func onlyGroundVehicle(v GroundVehicle) {
	/*
		Type assertions with switch statements
		recommend if you only want to do assertion on many types that implements
		interfaces
	*/
	switch typeVehicle := v.(type) {
	case Car:
		fmt.Println("Naik Mobil")
		fmt.Println(typeVehicle.OpenDoor())
		fmt.Println(typeVehicle.Drive())
	case MotorCycle:
		fmt.Println("Naik motor")
		fmt.Println(typeVehicle.OpenDoor())
		fmt.Println(typeVehicle.Drive())
	default: 
		fmt.Println("Jenis kendaraan tidak ditemukan")
	}
	
	/* 
		Type assertions with if else statements
		recommend if you only want to do assertion on a few types that implements
		interfaces
	*/
	// if typeVehicle, ok := v.(Car); ok {
	// 	fmt.Println("Naik Mobil")
	// 	fmt.Println(typeVehicle.OpenDoor())
	// 	fmt.Println(typeVehicle.Drive())
	// } else if typeVehicle, ok := v.(MotorCycle); ok {
	// 	fmt.Println("Naik motor")
	// 	fmt.Println(typeVehicle.OpenDoor())
	// 	fmt.Println(typeVehicle.Drive())
	// }
	
}