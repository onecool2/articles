1、在windows里安装anaconda

2、在anaconda里创建tensorflow的env，conda create -n tf1 python=3.6 tensorflow=1.15如果有gpu，就吧tensorflow-gpu=1.15

3、然后安装下面的python包。
    conda install -c anaconda protobuf
    pip pillow lxml cython jupyter matplotlib pandas opencv-python

4、下载github.com/tensorflow/models 到一个工作目录下比如c:\code

5、cd tutorials/image/imagenet/
  python3 classify_image.py --image_file=/root/abc.jpg
  这个时候就可以利用已有的模型探测图片了，如果可以，说明tensorflow环境安装没问题

6、
