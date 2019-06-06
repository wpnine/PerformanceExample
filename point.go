package main

//测试数据
var testData = CreateTestData()


func GetValue(index int) int {
	return testData[index]
}

//获取测试数据长度
func GetDataLen() int{
	return len(testData)
}

//获取测试数据数组
func GetData() []int {
	return testData
}

//创建一个6000000000大小的整型切片
func CreateTestData()[]int  {
	data := make([]int,6000000000)
	for index,_ := range data{
		data[index] = index % 128
	}
	return data
}


func toSum1(result *int)  {
	for i:=0;i< GetDataLen();i++{
		*result += GetValue(i)
	}
}

func toSum2(result *int)  {
	dataLength := GetDataLen()
	for i:=0;i< dataLength;i++{
		*result += GetValue(i)
	}
}


func toSum3(result *int)  {
	data := GetData()
	dataLength := len(data)
	for i:=0;i< dataLength;i++{
		*result += data[i]
	}
}


func toSum4(result *int)  {
	k := *result
	data := GetData()
	dataLength := len(data)
	for i:=0;i< dataLength;i++{
		k += data[i]
	}
	*result = k
}