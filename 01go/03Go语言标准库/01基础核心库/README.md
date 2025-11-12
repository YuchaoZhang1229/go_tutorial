### 第一阶段：基础核心库
1. fmt库
- ```Print/Println/Printf``` 不同打印方式
- ```Sprintf``` 格式化字符串
- ```Scan``` 用户输入处理

2. strings库
- ```Contains/Index``` 查找操作
- ```Replace``` 替换
- ```ToUpper/ToLower``` 大小写转换
- ```Split/Join``` 分割和连接
- ```TrimSpace``` 修剪空格
- ```Builder``` 高效字符串构建

3. strconv库
- ```Atoi/Itoa``` 字符串转整数/整数转字符串
- ```ParseInt/FormatInt``` 按指定进制解析/按指定进制格式化
- ```ParseBool/FormatBool``` 字符串转布尔值/布尔值转字符串
- ```ParseFloat/FormatFloat``` 字符串转浮点数/浮点数转字符串（支持格式控制）
- ```Quote/Unquote()```  添加引号和转义字符/去除引号和反转义
- ```AppendXXX()``` 将值追加到字节切片，避免内存分配

4. errors库
- ```errors.New``` 创建自定义错误