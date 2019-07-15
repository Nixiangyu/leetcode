package main

import "fmt"

// 楼梯算法 0ms fn = f(n-1) + f(n -2)
func climbStairs(n int) int {

	if n<=1 {
		return 1
	}
	dp := make([]int, n)
	dp[0]=1
	dp[1]=2
	for i:=2;i<n ;i++  {
		dp[i]=dp[i-1]+dp[i-2]
	}
	return dp[n-1]
}

func climbStairs1(n int) int {

	if n<=1 {
		return 1
	}

	return climbStairs1(n-1)+climbStairs1(n-2)
}
/*
func climbStairs(n int) int {
	//var temp []int
	var res int
	i := 0
	for {
		if i + 1 == n {
			res ++
			fmt.Println(" -> ",i)
			break
		}else if i + 2 == n{
			fmt.Println(" => ", i)
			res ++
		}

		i++
		fmt.Println(i)
	}

	*//*for j:= 0; j < len(temp); j++{
		i := temp[j]
		fmt.Println(i)
		for {
			if i + 1 == n {
				res ++
				fmt.Println(" -> ",i)
				break
			}else if i + 2 == n{
				fmt.Println(" => ", i)
				res ++
			}else if i + 2 < n{
				temp = append(temp, i + 2)
			}

			i++
		}

		temp = append(temp[:j], temp[j+1:]...)
		j--
	}*//*

	return res
}*/

func main() {
	fmt.Println("==",climbStairs(45))
}
