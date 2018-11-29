package redis

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

//HSet 将哈希表 key 中的域 field 的值设为 value 。
func (c *RedisClient) HSet(key string, field string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = c.do("HSET", key, field, data)
	return err
}

//HGet 返回哈希表 key 中给定域 field 的值。
func (c *RedisClient) HGet(key string, field string, value interface{}) error {
	buffer, err := redis.Bytes(c.do("HGET", key, field))
	if err != nil {
		return err
	}
	return json.Unmarshal(buffer, value)
}

//HGetAll 返回哈希表 key 中，所有的域和值。
func (c *RedisClient) HGetAll(key string, value []interface{}) error {
	buffer, err := redis.ByteSlices(c.do("HVALS", key))
	if err != nil {
		return err
	}
	if len(buffer) > 0 {
		for i := 0; i < len(buffer); i++ {
			err = json.Unmarshal(buffer[i], value[i])
			if err != nil {
				return err
			}
		}
	}
	return err
	//return json.Unmarshal(buffer,value)
}

//HDel 删除哈希表 key 中的一个或多个指定域，不存在的域将被忽略。
func (c *RedisClient) HDel(key string, field string) (int, error) {
	count, err := redis.Int(c.do("HDEL", key, field))
	if err != nil {
		return 0, err
	}
	return count, err
}



