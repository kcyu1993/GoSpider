package util

import (
	"fmt"
	"testing"
)

func TestUtil(t *testing.T) {
	// test int to string
	i := 2
	if "2" == IS(i) {
		t.Log("Test IS:int to string")
	}

	// test string to int
	v, err := SI("e2")
	if err == nil && v == i {
		t.Log("Test SI:string to int")
	} else {
		t.Log("Test SI:" + err.Error())
	}

	// caller dir
	t.Log("Test CurDir:" + CurDir())

	// create dir
	err = MakeDir("../data")
	if err != nil {
		t.Log("Test MakeDir:" + err.Error())
	} else {
		t.Log("Test MakeDir:dir already exist")
	}

	// create dir by filename
	filename := "../data/testutil.txt"
	err = MakeDirByFile(filename)
	if err != nil {
		t.Log("Test MakeDirByFile:" + err.Error())
	} else {
		t.Log("Test MakeDirByFile: dir already exist")
	}

	// save bytes into file
	err = SaveToFile(filename, []byte("testutil"))
	if err != nil {
		t.Log("Test SaveToFile" + err.Error())
	}

	// read bytes from file
	filebytes, err := ReadfromFile(filename)
	if err != nil {
		t.Error("Test ReadfromFile:" + err.Error())
	} else {
		t.Log("Test ReadfromFile:" + string(filebytes))
	}

	// file exist?
	t.Logf("%v", FileExist("../r.txt"))

	// find the go file in some dir
	filenames, err := ListDir(`G:\smartdogo\src\github.com\hunterhug\go_tool`, "go")
	t.Logf("%v:%v", filenames, err)

	// devide a string list into severy string list
	stringlist := []string{"2", "3", "4", "5", "6"}
	num := 3
	result, err := DevideStringList(stringlist, num)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%#v", result)
	}

	finfo, err := GetFilenameInfo(`G:\smartdogo\src\github.com\hunterhug\go_tool\README.md`)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(finfo.Name())
	}

	t.Log(Substr("123456", 0, 15))
}

func TestValidFileName(t *testing.T) {
	s := "*sdvdsv*|sdvsd>sdv<sdvds-\""
	fmt.Println(s)
	s1 := ValidFileName(s)
	fmt.Println(s1)
	s2 := ValidBackFileName(s1)
	fmt.Println(s2)
}

func TestTodayString(t *testing.T) {
	// times format
	t.Log(TodayString(3))

	// now secord times from January 1, 1970 UTC.
	secord := GetSecondTimes()
	t.Log(secord)

	// now date string by secord
	timestring := GetSecond2DateTimes(secord)
	t.Log(timestring)

	// change back
	t.Log(GetDateTimes2Secord(timestring))
}

func TestGetCurrentPath(t *testing.T) {
	// 二进制文件所在位置
	path, e := GetBinaryCurrentPath()
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println(path)
	}

	// 执行命令所在位置
	path, e = GetCurrentPath()
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println(path)
	}

	// 永远都是该项目文件所在目录
	fmt.Println(CurDir())
}
