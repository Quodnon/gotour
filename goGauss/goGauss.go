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




func main ( ) {
    var inputFile *os.File;
    var inputError error;
    var readerError error;
    var inputReader *bufio.Reader;
    var inputString string;
    
    
    inputFile,inputError=os.OpenFile("linear.dat",os.O_RDONLY,0);
    
	if inputError!=nil {
		fmt.Printf("Error opening file");
		return ;
	}

    
    inputReader	= bufio.NewReader(inputFile);
    
    for {
        
        inputString,readerError = inputReader.ReadString(';');
        if readerError!=nil {
            
            return;    
        }
        
        fmt.Println(inputString);
    }
}
