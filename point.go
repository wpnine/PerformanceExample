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

//创建一个6000000000大小的整型切片，大概需要占用22G的进程虚拟空间。
//注：如果是32位系统的话需要将值改小，32操作系统的虚拟空间最大只有4G
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


func toSum5(result *int)  {
	k := *result
	data := GetData()
	dataLength := len(data)
	for i:=1;i< dataLength;i+=2{
		k += data[i] + data[i - 1]
	}
	if dataLength % 2 == 1{
		k += data[dataLength-1]
	}

	*result = k
}

func toSum6(result *int)  {
	k1 := 0
	k2 := 0
	data := GetData()
	dataLength := len(data)

	for i:=1;i< dataLength;i+=2{
		k1 += data[i]
		k2 += data[i - 1]
	}

	//如果是传入的数量是奇数，则单独对最后一个数进行累加
	if dataLength % 2 == 1{
		k1 += data[dataLength-1]
	}

	*result = k1 + k2
}

func toSum7(result *int)  {
	k1 := 0
	k2 := 0
	k3 := 0

	data := GetData()
	dataLength := len(data)

	for i:=0;i< dataLength-3;i+=3{
		k1 += data[i]
		k2 += data[i + 1]
		k3 += data[i + 2]
	}


	*result = k1 + k2 + k3
}