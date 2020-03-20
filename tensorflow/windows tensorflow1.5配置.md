python 的版本和opencv的版本对应：
    from .cv2 import *
    ImportError: DLL load failed: 找不到指定的模块。
python:3.6.10 对应opencv_python-3.4.1.15-cp36-cp36m-win_amd64.whl


1、在windows里安装anaconda

2、在anaconda里创建tensorflow的env，conda create -n tf1 python=3.6 tensorflow=1.15如果有gpu，就吧tensorflow-gpu=1.15

3、然后安装下面的python包。
    conda install -c anaconda protobuf
    pip pillow lxml cython jupyter matplotlib pandas opencv-python

4、下载github.com/tensorflow/models 到一个工作目录下比如c:\code

5、cd tutorials/image/imagenet/
  python3 classify_image.py --image_file=/root/abc.jpg
  这个时候就可以利用已有的模型探测图片了，如果可以，说明tensorflow环境安装没问题

6、在windows的系统环境变量中加入
PYTHONPATH=c:\code\tensorflow1\models;c:\code\tensorflow1\models\research;c:\code\tensorflow1\models\research\slim

7、在c:\code\tensorflow1\models\research目录下执行下面语句：
protoc --python_out=. .\object_detection\protos\anchor_generator.proto .\object_detection\protos\argmax_matcher.proto .\object_detection\protos\bipartite_matcher.proto .\object_detection\protos\box_coder.proto .\object_detection\protos\box_predictor.proto .\object_detection\protos\eval.proto .\object_detection\protos\faster_rcnn.proto .\object_detection\protos\faster_rcnn_box_coder.proto .\object_detection\protos\grid_anchor_generator.proto .\object_detection\protos\hyperparams.proto .\object_detection\protos\image_resizer.proto .\object_detection\protos\input_reader.proto .\object_detection\protos\losses.proto .\object_detection\protos\matcher.proto .\object_detection\protos\mean_stddev_box_coder.proto .\object_detection\protos\model.proto .\object_detection\protos\optimizer.proto .\object_detection\protos\pipeline.proto .\object_detection\protos\post_processing.proto .\object_detection\protos\preprocessor.proto .\object_detection\protos\region_similarity_calculator.proto .\object_detection\protos\square_box_coder.proto .\object_detection\protos\ssd.proto .\object_detection\protos\ssd_anchor_generator.proto .\object_detection\protos\string_int_label_map.proto .\object_detection\protos\train.proto .\object_detection\protos\keypoint_box_coder.proto .\object_detection\protos\multiscale_anchor_generator.proto .\object_detection\protos\graph_rewriter.proto
python setup.py build
python setup.py install

8、据说这部可以执行成功python model_builder_test.py 

9、下载https://tzutalin.github.io/labelImg/，给图片标注。（快捷键：w标注、d下一张、ctrl+s保存

10、用xml2csv把labelimg输出的xml转换成tensorflow的csv文件，xml2csv.py
    （分训练集和测试集，这些文件下载后标注，然后位置就不要动了，放在object_detection\test_images目录下的test和train里

11、再将csv文件转换成TFRecord文件，generate_TFR.py
    python generate_tfrecord.py --csv_input=test_images/train_labels.csv --image_dir=train_images/train --output_path=train.record
    python generate_tfrecord.py --csv_input=test_images/test_labels.csv --image_dir=test_images/test    --output_path=test.record

12、配置labelmap.pbtxt，一个对象一段。

13、复制samples\configs\ssd_inception_v2_coco.config 到 object_detection\training目录下，并且修改好多地方
    
14、
   python object_detection/model_main.py --pipeline_config_path=object_detection/training/ssd_mobilenet_v1_coco.config --model_dir=object_detection/training --num_train_steps=50000 --num_eval_steps=2000 --alsologtostderr
或者
   python train.py --logtostderr --train_dir=training/ --pipeline_config_path=training/ssd_mobilenet_v1_coco.config

15、python export_inference_graph.py --input_type image_tensor --pipeline_config_path training\ssd_mobilenet_v1_coco.config   --trained_checkpoint_prefix training\model.ckpt-50000  --output_directory wdj_inference_graph

16、bject_detection_image1.py
