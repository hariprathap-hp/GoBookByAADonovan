package main

import (
	"TheGoProgLangBook/GoRoutines_Channels/Section84to/Channels/Cakeshop/cake"
	"time"
)

func main() {
	caking := cake.Shop{
		Verbose:        true,
		Cakes:          5,               // number of cakes to bake
		BakeTime:       3 * time.Second, // time to bake one cake
		BakeStdDev:     1 * time.Second, // standard deviation of baking time
		BakeBuf:        2,               // buffer slots between baking and icing
		NumIcers:       3,               // number of cooks doing icing
		IceTime:        2 * time.Second, // time to ice one cake
		IceStdDev:      2 * time.Second, // standard deviation of icing time
		IceBuf:         2,               // buffer slots between icing and inscribing
		InscribeTime:   2 * time.Second, // time to inscribe one cake
		InscribeStdDev: 2 * time.Second, // standard deviation of inscribing time
	}

	caking.Work(5)
}
