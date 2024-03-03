package generator

import (
	"gRPC/pb"

	"github.com/golang/protobuf/ptypes"
)

// return a new KeyBoard sample
func NewKeyBoard() *pb.Keyboard {
	return &pb.Keyboard{
		Layout:  randomKeyBoardLayout(),
		Backlit: randomBool(),
	}
}

// return a new CPU sample
func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numberCores := randomInt(2, 8)
	numberThreads := uint32(randomInt(numberCores, 16))
	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	return &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: numberThreads,
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
}

func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)
	memory := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb.Memory_GIGABYE,
	}

	return &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}
}

func NewRAM() *pb.Memory {
	value := randomInt(4, 64)
	return &pb.Memory{
		Value: uint64(value),
		Unit:  pb.Memory_GIGABYE,
	}
}

func NewSSD() *pb.Storage {
	return &pb.Storage{
		Dirver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  pb.Memory_GIGABYE,
		},
	}
}

func NewHDD() *pb.Storage {
	return &pb.Storage{
		Dirver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}
}

func NewScreen() *pb.Screen {
	return &pb.Screen{
		SizeInch:   float32(randomFloat64(18.0, 27.0)),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}
}

func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)
	return &pb.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Memory:   NewRAM(),
		Gpus:     []*pb.GPU{NewGPU()},
		Storages: []*pb.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyBoard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		Priceusd:    randomFloat64(1500.0, 3000.0),
		ReleaseYear: uint32(randomInt(2018, 2024)),
		UpdateAt:    ptypes.TimestampNow(),
	}
}
