package file

import (
	"encoding/csv"
	"fmt"
	"image"
	"image/png"
	"os"
)

func ReadImage(path string) image.Image {
	imageFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer imageFile.Close()

	img, _, err := image.Decode(imageFile)
	if err != nil {
		panic(err)
	}

	return img
}

func WriteImage(img image.Image, path string, filename string) {
	createFolder(path)

	filePath := fmt.Sprintf("%s/%s.png", path, filename)

	imagePath, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer imagePath.Close()

	err = png.Encode(imagePath, img)
	if err != nil {
		panic(err)
	}
}

func WriteGradient(gradient []int, path string, filename string) {
	createFolder(path)
	filePath := fmt.Sprintf("%s/%s", path, filename)
	
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	length := len(gradient)

	_, err = file.WriteString(fmt.Sprintf("%d\n", length))
	if err != nil {
		panic(err)
	}

	for _, value := range gradient {
		_, err = file.WriteString(fmt.Sprintf("%d\n", value))
		if err != nil {
			panic(err)
		}
	}
}

func WriteHistogram(histogram []uint16, path string, filename string) {
	createFolder(path)
	filePath := fmt.Sprintf("%s/%s", path, filename)
	
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	length := len(histogram)

	_, err = file.WriteString(fmt.Sprintf("%d\n", length))
	if err != nil {
		panic(err)
	}

	for _, value := range histogram {
		_, err = file.WriteString(fmt.Sprintf("%d\n", value))
		if err != nil {
			panic(err)
		}
	}
}

func WriteCsvHistogram(histograms [][]uint16, header []string, path, filename string) {
	filePath := fmt.Sprintf("%s/%s", path, filename)
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	err = csvWriter.Write(header)
	if err != nil {
		panic(err)
	}

	maxColumn := len(histograms)
	maxRow := len(histograms[0])

	for row := 0; row < maxRow; row++ {
		var rowData []string
		rowData = append(rowData, fmt.Sprintf("%d", row))
		for column := 0; column < maxColumn; column++ {
			data := histograms[column][row]
			rowData = append(rowData, fmt.Sprintf("%d", data))
		}

		if err := csvWriter.Write(rowData); err != nil {
			panic(err)
		}
	}
}

func WriteCsvHistogramGradient(histogram []uint16, gradient []int, header []string, path, filename string) {
	filePath := fmt.Sprintf("%s/%s", path, filename)
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	err = csvWriter.Write(header)
	if err != nil {
		panic(err)
	}

	maxRow := len(histogram)

	for i := 0; i < maxRow; i++ {

		x := fmt.Sprintf("%d",i)
		h := fmt.Sprintf("%d",histogram[i])
		g := fmt.Sprintf("%d",gradient[i])

		row := []string{x, h, g}
		if err := csvWriter.Write(row); err != nil {
			panic(err)
		}
	}
}

func CreateFolder(folderPath string) {
	fmt.Println("CreateFolder: ", folderPath)
	createFolder(folderPath)
}

func createFolder(folderPath string) {
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}