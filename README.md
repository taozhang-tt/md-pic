# md-pic
上传系统剪切板图片到oss

## 使用
* clone 项目到本地
* 添加 config.json 文件，可以参考 config.json.sample
* sudo make
* `source ~/.zshrc` or `source ~/.bashrc`
* 截图值系统剪切板
* 执行 `mdpic`
上传成功以后的图片访问路径已经存储到系统剪切板中，只要 control + v 即可

## 2021-12-02 增加删除功能

命令：`mdpic rm img_name`

假如要删除`1638425206.png`，只需要执行 `mdpic rm 1638425206`

支持批量删除，多个图片名以空格间隔
