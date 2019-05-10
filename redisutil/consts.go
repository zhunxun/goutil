package redis

//部分常用的redis命令
const (
	//基本命令
	Auth   = "auth"
	Ping   = "ping"
	Select = "select"

	//key的操作命令
	Del       = "Del"       //DEL key   删除key
	Dump      = "Dump"      //DUMP key  序列化key ，并返回被序列化的值
	Exists    = "Exists"    //EXISTS key  检查key是否存在
	Expire    = "Expire"    //EXPIRE key seconds  设置过期时间，以秒计
	ExpireAt  = "ExpireAt"  //EXPIREAT key timestamp  设置过期时间，指定时间戳(秒级)
	Pexpire   = "Pexpire"   //PEXPIRE key milliseconds   设置过期时间以毫秒计
	PexpireAt = "PexpireAt" //PEXPIREAT key milliseconds-timestamp  设置过期时间，指定时间戳(毫秒级)
	Keys      = "Keys"      //KEYS pattern 查找所有符合给定模式( pattern)的key
	Move      = "Move"      //MOVE key db 将当前数据库的 key 移动到给定的数据库db
	Persist   = "Persist"   //PERSIST key 移除 key 的过期时间，key 将持久保持
	Pttl      = "Pttl"      //PTTL key 以毫秒为单位返回 key 的剩余的过期时间
	Ttl       = "ttl"       //TTL key 以秒为单位，返回给定 key 的剩余生存时间(TTL, time to live)
	Rename    = "Rename"    //RENAME key newkey 修改 key 的名称
	RenameNx  = "RenameNx"  //仅当 newkey 不存在时，将 key 改名为 newkey
	Type      = "type"      //TYPE key 返回 key 所储存的值的类型

	//字符串的操作命令
	Set          = "set"          //Set key value
	Get          = "get"          //Get key
	GetRange     = "GetRange"     //GETRANGE key start end  返回 key 中字符串值的子字符
	GetSet       = "GetSet"       //GETSET key value 将给定 key 的值设为 value ，并返回 key 的旧值(old value)
	GetBit       = "GetBit"       //GETBIT key offset 对 key 所储存的字符串值，获取指定偏移量上的位(bit)
	Mget         = "Mget"         //MGET key1 [key2..] 获取所有(一个或多个)给定 key 的值
	SetBit       = "SetBit"       //SETBIT key offset value 对 key 所储存的字符串值，设置或清除指定偏移量上的位(bit)
	SetEx        = "SetEx"        //SETEX key seconds value
	SetNx        = "SetNx"        //SETNX key value 只有在 key 不存在时设置 key 的值
	Incr         = "incr"         //INCR key   将 key 中储存的数字值增一
	Decr         = "decr"         //DECR key  将 key 中储存的数字值减一
	HIncrByFloat = "hIncrByFloat" //HINCRBYFLOAT key field increment 为哈希表 key 中的指定字段的浮点数值加上增量 increment

	//hash 命令
	Hdel    = "hdel"    //HDEL key field1 [field2]   删除一个或多个哈希表字段
	Hexists = "hexists" //HEXISTS key field 查看哈希表 key 中，指定字段是否存在
	Hget    = "hget"    //HGET key field
	HgetAll = "hgetall" //HGETALL key   获取在哈希表中指定 key 的所有字段和值
	Hkeys   = "hkeys"   //HKEYS key  获取所有哈希表中的字段
	Hlen    = "hlen"    //HLEN key 获取哈希表中字段的数量
	Hmget   = "hmget"   //HMGET key field1 [field2] 获取所有给定字段的值
	Hset    = "hset"    //HSET key field value
	Hmset   = "hmset"   //HMSET key field1 value1 [field2 value2 ] 同时set多个 field-value
	Hvals   = "hvals"   //HVALS key  获取哈希表中所有值

	//list 命令
	Blpop   = "blpop"   //BLPOP key1 [key2 ] timeout 移出并获取列表的第一个元素 不存在则阻塞超时
	Brpop   = "brpop"   //BRPOP key1 [key2 ] timeout 移出并获取列表的最后一个元素 不存在则阻塞超时
	Lindex  = "lindex"  //LINDEX key index  通过索引获取列表中的元素
	Linsert = "linsert" //LINSERT key BEFORE|AFTER pivot value 在列表的元素前或者后插入元素
	Llen    = "llen"    //LLEN key 获取列表长度
	Lpop    = "lpop"    //LPOP key 移出并获取列表的第一个元素
	Lpush   = "lpush"   // LPUSH key value1 [value2] 将一个或多个值插入到列表头部
	Lrange  = "lrange"  //LRANGE key start stop 获取列表指定范围内的元素
	Rpop    = "rpop"    //RPOP key 移除列表的最后一个元素，返回值为移除的元素
	Rpush   = "rpush"   //RPUSH key value1 [value2] 在列表尾添加一个或多个值

	//set 命令
	Sadd      = "sadd"      //SADD key member1 [member2] 向集合添加一个或多个成员
	Scard     = "scard"     //SCARD key 获取集合的成员数(个数)
	Sdiff     = "sdiff"     //SDIFF key1 [key2]  返回给定所有集合的差集
	Sinter    = "sinter"    //SINTER key1 [key2] 返回给定所有集合的交集
	SisMember = "sismember" //SISMEMBER key member 判断 member 元素是否是集合 key 的成员
	Smembers  = "smembers"  //SMEMBERS key 返回集合中的所有成员
	Srem      = "srem"      //SREM key member1 [member2] 移除集合中一个或多个成员
	Sunion    = "sunion"    //SUNION key1 [key2] 返回所有给定集合的并集

	//zset 部分常用命令
	Zadd             = "zadd"             //ZADD key score1 member1 [score2 member2] 向有序集合添加一个或多个成员，或者更新已存在成员的分数
	Zcard            = "zcard"            //ZCARD key 获取有序集合的成员数
	Zcount           = "zcount"           //ZCOUNT key min max 计算在有序集合中指定区间分数的成员数
	Zrange           = "zrange"           //ZRANGE key start stop [WITHSCORES] 通过索引区间返回有序集合成指定区间内的成员
	Zrank            = "zrank"            //ZRANK key member 返回有序集合中指定成员的索引(排行)
	Zrem             = "zrem"             //ZREM key member [member ...] 移除有序集合中的一个或多个成员
	ZremRangeByLex   = "ZremRangeByLex"   //ZREMRANGEBYLEX key min max  移除有序集合中给定的字典区间的所有成员
	ZremRangebyRank  = "ZremRangebyRank"  //ZREMRANGEBYRANK key start stop 移除有序集合中给定的排名区间的所有成员
	ZremRangebyScore = "ZremRangebyScore" //移除有序集合中给定的分数区间的所有成员
	ZrevRange        = "ZrevRange"        //ZREVRANGE key start stop [WITHSCORES] 返回有序集中指定区间内的成员，通过索引，分数从高到底
	ZrevRangeByScore = "ZrevRangeByScore" //ZREVRANGEBYSCORE key max min [WITHSCORES]  返回有序集中指定分数区间内的成员，分数从高到低排序
	ZrevRank         = "ZrevRank"         //ZREVRANK key member 返回有序集合中指定成员的排名，有序集成员按分数值递减(从大到小)排序
	Zscore           = "Zscore"           //ZSCORE key member  返回有序集中，成员的分数值

	//发布订阅命令
	Publish     = "publish"     //PUBLISH channel message  将信息发送到指定的频道
	Subscribe   = "subscribe"   //SUBSCRIBE channel [channel ...]  订阅给定的一个或多个频道的信息
	UnSubscribe = "unSubscribe" //UNSUBSCRIBE channel [channel ...]  退订指定的频道

	//事务命令
	Multi   = "multi"   //标记事务开始
	DisCard = "discard" //取消事务 放弃执行事务块内的所有命令
	Exec    = "exec"    //执行所有事务块内的命令

	//script 脚本

	//管道
)
