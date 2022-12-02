package main

import (
	"fmt"
    "encoding/csv"
    "log"
	"io"
    "os"
	"strconv"

	"golang.org/x/text/encoding/japanese"
    "golang.org/x/text/transform"
)

func main() {
	var input string 
	var output string
	var columnNum int 
	var columnName string 
	var columnNumber int 
	
	var nextColumnName string

	print("抽出対象のcsvのpathを入力してください：")
	fmt.Scan(&input);
	print("対象のカラム番号を入力してください：")
	fmt.Scan(&columnNum)
	print("対象のカラム名を入力してください：")
	fmt.Scan(&columnName)
	print("出力先のcsvのファイル名を入力してください：")
	fmt.Scan(&output)
	print("出力先のファイルの他のカラムの数を入力してください：")
	fmt.Scan(&columnNumber)

	var columnNameList []string
	columnNameList = append(columnNameList,columnName)
	for i :=0;i<columnNumber;i++{
		print("出力先の他の" +strconv.Itoa(i+1)+ "つ目のカラム名を入力してください：")
		fmt.Scan(&nextColumnName)
		columnNameList =append(columnNameList,nextColumnName)
	}

    f, err := os.Open(input)
    if err != nil {
        log.Fatal(err)
    }

    r := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
    var recordList [] string 
	var uniqueList [] string
    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
		recordList = append(recordList, record[columnNum -1])
		uniqueList = sliceUnique(recordList);
    }

	
	var outputRecords [][]string

	fmt.Println(uniqueList);
	
	for _,v:= range uniqueList{
		if v == columnName {
          outputRecords = append(outputRecords,columnNameList);
		}else{
		  var a = []string{v};
		  for i:=0;i< columnNumber; i++{
			a = append(a,"");
		  }
		  outputRecords = append(outputRecords,a);
		}
	}
	fmt.Println(outputRecords);

	f2, err2 := os.Create(output);
	if err2 != nil{
		fmt.Println(err2);
	}
	w := csv.NewWriter(transform.NewWriter(f2, japanese.ShiftJIS.NewEncoder()));
	w.WriteAll(outputRecords);
}

/*
　配列から重複を省きます。
*/
func sliceUnique(target []string) (unique []string) {
    m := map[string]bool{}

    for _, v := range target {
        if !m[v] {
            m[v] = true
            unique = append(unique, v)
        }
    }
    
    return unique
}