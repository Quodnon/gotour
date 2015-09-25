package main

import (
    "bufio"
    "os"
    "fmt"
    )
//TODO get data from file  separator x , separator a : end of line ; float .
//TODO convert string to arrays
//TODO first step
//TODO print result of step;

func getDataFromfile  (fileName string) string {
    
    var inputFile *os.File;
    var inputError error;
    var readerError error;
    var inputReader *bufio.Reader;
    var inputString string;
    var readenString string;
    
    inputFile,inputError=os.OpenFile(fileName,os.O_RDONLY,0);
    
	if inputError!=nil {
		fmt.Printf("Error opening file");
		return "Error opening file";
	}

    
    inputReader	= bufio.NewReader(inputFile);
    
    for {
        
        readenString,readerError = inputReader.ReadString(';');
        if readerError!=nil {
            
            break;   
        }
        
        inputString+=readenString;
        
    }
    return inputString;
}

 
func main ( ) {
    linearEquations := getDataFromfile("linear.dat")
    fmt.Println(linearEquations)
    
    matrB:= []float64 {-2,-3,-4,-44.3};
    matrX:= [][]float64 {{4,4,3,2,4,4},{4,1,1,1,1},{32,3,21.2323,443}};
    fmt.Println(matrX,matrB)
    
}
