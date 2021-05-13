
# flag

## 示例

module db

```golang
package cache

import "github.com/go-redis/redis/v8"
import "flag"

var Db redis.Client

var g_redis_addr = flag.String("redis-addr", "redis://user:pwd@127.0.0.1:6379/db",  "redis address")
var g_timeout = flag.Int("redis-timeout", 100, "redis operator timeout")

func init()  {
	opt, err := redis.ParseURL(*g_redis_addr)
	if err != nil {
		panic(err)
	}
	Db = redis.NewClient(opt);
}

func Get(key string, val interface{}) error {
	.....
}

```

上面的代码会出错， 因为flag 变量的初始化是在flag.Parse 的时候， flag.Parse 只能在所有变量都定义完了才能调用，init 函数这个时候还不能获取到flag真正的值。

golang 的 package 就是module,  import 表明模块之间的依赖。如果不用flag, 他们的依赖关系是对的，初始化顺序也是对的，一切都很美好。 一旦你的模块需要可配置， 事情就麻烦起来了。

这个库会先初始化自己， 当用flag.String,  flag.StringVar 定义变量的时候， 直接获取到值；这样就没有破坏标准的flag接口， 只是改变初始化顺序



## 选项
- flagenv[=true] 

	flag 从 .env 文件和环境变量中获取flag配置项, `微服务12条建议把环境当成配置`; k8s环境下运行， 可以把ConfigMap 映射成环境变量
- flagfile[=flagfile]

    读取 flagfile 配置文件; 用户指定文件名的时候，可以根据文件后缀选择不同的解析器， 支持 Yaml, Toml, Json, Jsonnet, Ini, env, flagfile格式; 支持读取多个文件配置文件(使用逗号分隔多个文件)

- flagdump[=flag|json|yaml|toml|ini|env]

	输出flag的默认值