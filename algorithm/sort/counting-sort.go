package main

import "fmt"

func main()  {
	data := []int{34, 30, 3, 6, 5, 34, 35, 6, 22, 8,11}

	fmt.Print("排序前")
	fmt.Println(data)

	countSort(data)
}
//计数排序，a 是数组，n 是数组大小。假设数组中存储的都是非负整数。
func countSort(data []int)  {
	if len(data)==0 {
		return
	}
	//获取最大值
	max:=0
	for _, v:=range data {
		if v>max{
			max=v
		}
	}
	fmt.Println("最大值：",max)

	newData:=make([]int,max+1)
	//计算每个元素的个数
	for _, v:=range data {
		newData[v]++
	}
	fmt.Println(newData)
	// 依次累加，求得次序
	for i:=1;i<=max ;i++  {
		newData[i]+=newData[i-1]
	}
	fmt.Println(newData)

	tmp:=make([]int,len(data))
	n:=len(data)-1
	for  ;n>=0;n--{
		//fmt.Println(n,data[n],newData[data[n]])
		newData[data[n]]--
		tmp[newData[data[n]]]=data[n]
	}
	//data=tmp
	fmt.Println(tmp)
	fmt.Println(data)
}
