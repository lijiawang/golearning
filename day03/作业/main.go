package main

import "fmt"
/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	for name,num := range distribution{
		fmt.Println(name,"金币数量:",num)
	}
	fmt.Println("剩下:", left)
}
 
// dispatchCoin分规则金币，返回剩余金币
func dispatchCoin() int{
	//1.依次给每个人(拿到每个人的名字)
	// 拿到每个名字去判断规则
	for _, name := range users{
		// fmt.Println(name)
		// 2、某个人分配多少金币
		userNum := dispatchForUser(name)
		// 每个人分的多少金币
		distribution[name] = userNum
		// 计算剩下的金币
		coins = coins - userNum


	}
	fmt.Println(distribution)
	return coins
} 

func dispatchForUser(name string) int{
	userNum := 0
	// 某个人名的每一个字母
	for _,c := range name{ // Matthew 例如
		// 判断 是否包含分金币的字母
		switch c{
		case 'e','E':
			userNum += 1
		case 'i','I':
			userNum += 2
		case 'o','O':
			userNum += 3
		case 'u','U':
			userNum += 4
		}
	}
	return userNum
}