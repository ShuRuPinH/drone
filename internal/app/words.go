package app

import (
	"fmt"
	"github.com/white-pony/go-fann"
)

func Words() {

	ann := fann.CreateFromFile("words.net")

	rez := ann.Run([]fann.FannType{2, 2, 1})

	for _, fannType := range rez {
		fmt.Print(fannType)
		fmt.Print(" * ")
	}

	ann.Destroy()
}

func WordsTrain() {
	const numLayers = 3
	const desiredError = 0.00001
	const maxEpochs = 50000
	const epochsBetweenReports = 1000

	ann := fann.CreateStandard(numLayers, []uint32{3, 10, 3})
	ann.SetActivationFunctionHidden(fann.SIGMOID_SYMMETRIC)
	ann.SetActivationFunctionOutput(fann.SIGMOID_SYMMETRIC)
	ann.TrainOnFile("datasets/words.data", maxEpochs, epochsBetweenReports, desiredError)
	ann.Save("words.net")

	rez := ann.Run([]fann.FannType{2, 2, 1})

	for _, fannType := range rez {
		fmt.Print(fannType)
		fmt.Print(" * ")
	}

	ann.Destroy()
}
