
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

上面的代码会出现问题， 因为flag 变量的初始化是在flag.Parse 的时候， flag.Parse 只能在所有变量都什么完了才适合，init 函数这个时候还不能获取到 flag 真正的值。

golang 的 package 相当于 module,  import 也表明了模块之间的依赖，如果不用flag,  他们的依赖关系是对的，初始化顺序也是对的，一切都很美好。 一旦你的模块需要可配置， 事情就麻烦起来了。

- 写死一个配置模块， 其它模块依赖配置模块来实现可配置化。
    `写死配置文件不是大问题， 问题是任何一个人想添加一项配置都要去修改这个配置模块， 随着时间的推移，这个配置模块最终会腐化。`

- 延迟初始化， 把所有需要配置的代码延迟到 main 执行之后。 可以手写， 也可以用一个统一的延迟初始 module 来解决。

- 修改标准的flag,  让它自己先去读取os.args， 不用等到 main 函数中才去Parse初始化。 

### 这个库用的是最后一种方式，flag 自己先初始化好自己， flag.String,  flag.StringVar 定义变量的时候，直接去获取到值，没有破坏标准的flag的接口， 只是替换初始化的顺序



## 选项
- flagenv 选型, 默认开启， flag 会从 .env 文件和环境变量中获取flag配置项, `微服务12条的建议是也是把环境当成配置`, 这个方式非常适合k8s环境下运行， 把ConfigMap 映射成环境变量就解决了
- flagfile 选型, 默认读取 flagfile 文件; 用户指定文件名的时候，可以根据文件后缀选择不同的解析器， 支持 Yaml, Toml, Json 和 flagfile 格式; 支持读取多个文件配置文件(使用逗号分隔多个文件)

## TODO
添加 flagout 功能， 可以把默认配置导出来， 方便调试