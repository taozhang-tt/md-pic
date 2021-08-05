# md-pic
上传系统剪切板图片到oss

## 使用
* clone 项目到本地
* 添加 config.json 文件，可以参考 config.json.sample
* go build main.go 生成 main 文件
* 执行 `ln main /usr/local/bin/mdpic` 将 main 文件软链到 /usr/local/bin/mdpic
* `source ~/.zshrc`
* 截图
* 执行 `mdpic`
上传成功以后的图片访问路径已经存储到系统剪切板中，只要 control + v 即可
