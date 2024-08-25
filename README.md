# NetCD的GO语言版
## 1. 进度：9月之前写出后端的base
- [x] 播放音乐
- [x]  暂停
- [x] progressBar输出进度
- [x]  播放器切歌逻辑
- [x] 切换歌单
- [x] 通过channel监听网页模拟nfc的输入
- [x] 连接nfc
- [x] 当txt音乐文件不在本地时，跳过播放切换
- [x] 根据resources下的mp3文件初始化txt文件
- [x] NFC根据UID切换歌单
- [x] NFC实现各种操作
- [ ] 修复processBar会覆盖输入的bug
- [ ] 根据NFC卡里信息来切换歌单
- [ ] 一段时间内NFC只会被读取一次：避免一直切换导致不播放

## 2. 使用说明
### 2.1. 安装libNFC
 #### 2.1.1. windows安装libNFC的资料（不推荐，巨难装，而且安装了libNFC也还有问题）
https://www.jianshu.com/p/c02b7f7a7cfa
https://www.d-logic.com/zh/knowledge_base/libnfc-installation-on-windows-11/
https://blog.csdn.net/fengshuiyue/article/details/37921717
https://fanzheng.org/archives/29

 #### 2.1.2. linux安装libNFC（推荐）
https://github.com/nfc-tools/libnfc
#### 2.1.3. mac安装libNFC
推荐直接用brew安装，具体安装自行搜索

### 2.2 NFC-CD-GO
```
git clone git@github.com:Sprite137/NFC-CD-GO.git

cd NFC-CD-GO

cd music_opr

go build -o player

./player
```

## 3. NOTICE:
- resources存放了歌单(playList)和歌曲(music),如果要修改，可以放入其中
- music暂时只支持mp3格式，应该会补充的
- playList中请按已有的格式编辑或者新建自己的歌单

## 4. REF
1. https://github.com/nfc-tools/libnfc
2. https://github.com/xBlaz3kx/nfc-reader-example
3. https://github.com/0neSe7en/jikefm
4. https://github.com/faiface/beep
5. https://github.com/clausecker/nfc
6. https://github.com/lyswhut/lx-music-desktop
7. https://blog.csdn.net/u012915636/article/details/115414911

