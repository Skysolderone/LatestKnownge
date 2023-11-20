package main

import (
	"fmt"

	"github.com/google/uuid"
)

// UUID的优势和用途
// 全局唯一性：UUID能够在全球范围内保证唯一性，不会重复。
// 分布式系统：在分布式系统中，UUID可用于唯一标识分布式节点、事务等。
// 数据库主键：UUID可以作为数据库表的主键，避免主键冲突。
// 安全性：版本4的UUID是完全随机生成的，可以用于密码重置、令牌等场景，提高安全性。
// 可读性：版本1的UUID基于时间戳生成，可以用于记录日志、跟踪数据变化等场景。
func main() {
	result := uuid.New()
	fmt.Println(result)
}
