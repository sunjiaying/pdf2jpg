# PDF2JPG

需要先安装 ImageMagick 和 Ghostscript 工具

## Windows下安装
建议先安装scoop，在powershell下面运行

```
iex (new-object net.webclient).downloadstring('https://get.scoop.sh')
```

然后通过scoop安装ghostscript和imagemagick，注意ghostscript安装32位的

```
scoop install ghostscript --arch 32bit
scoop install imagemagick
```

## Linux下安装
