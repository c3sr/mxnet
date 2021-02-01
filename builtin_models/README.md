# builtin_models

This folder contains all the built-in model manifests.
The models are collected from [Gluon Model Zoo](https://gluon-cv.mxnet.io/model_zoo/index.html).

## How to Add a Model

No code is required to add a new model.

## Update the Built-in Model Catalog

After updating model manifests or adding new model manifests, run `make generate` in the root directory.

Note: Only recording first three detections
| Name                        | Image                                  | Label | Xmin  | Xmax  | Ymin  | Ymax  | Probability |
|:---------------------------:|:--------------------------------------:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----------:|
| Faster_RCNN_ResNet50_v1b_VOC| ./_fixtures/lane_control.jpg           | car   | 0.617 | 0.986 | 0.573 | 0.969 | 1.000       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.007 | 0.363 | 0.570 | 0.803 | 0.999       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.552 | 0.707 | 0.439 | 0.612 | 0.998       |
| SSD_300_VGG16_Atrous_COCO   | ./_fixtures/lane_control.jpg           | car   | 0.028 | 0.338 | 0.573 | 0.801 | 0.992       |
|                             | ./_fixtures/lane_control.jpg           | person| 0.451 | 0.498 | 0.478 | 0.696 | 0.992       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.293 | 0.437 | 0.501 | 0.605 | 0.991       |
| SSD_300_VGG16_Atrous_VOC    | ./_fixtures/lane_control.jpg           | car   | 0.561 | 0.703 | 0.439 | 0.625 | 1.000       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.292 | 0.447 | 0.498 | 0.598 | 1.000       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.033 | 0.337 | 0.571 | 0.801 | 1.000       |
| SSD_512_MobileNet_1.0_COCO  | ./_fixtures/lane_control.jpg           | car   | 0.022 | 0.337 | 0.575 | 0.807 | 0.993       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.648 | 0.962 | 0.575 | 0.952 | 0.970       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.291 | 0.435 | 0.501 | 0.595 | 0.941       |
| SSD_512_MobileNet_1.0_VOC   | ./_fixtures/lane_control.jpg           | car   | 0.611 | 1.007 | 0.575 | 0.990 | 0.998       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.033 | 0.340 | 0.582 | 0.807 | 0.998       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.554 | 0.706 | 0.430 | 0.631 | 0.994       |
| SSD_512_ResNet50_v1_COCO    | ./_fixtures/lane_control.jpg           | car   | 0.022 | 0.337 | 0.569 | 0.802 | 0.999       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.703 | 0.965 | 0.576 | 0.986 | 0.999       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.614 | 0.850 | 0.559 | 0.940 | 0.998       |
| SSD_512_ResNet50_v1_VOC     | ./_fixtures/lane_control.jpg           | car   | 0.555 | 0.703 | 0.437 | 0.623 | 1.000       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.019 | 0.343 | 0.565 | 0.807 | 1.000       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.622 | 0.930 | 0.551 | 0.976 | 1.000       |
| SSD_512_ResNet101_v2_VOC    | ./_fixtures/lane_control.jpg           | car   | 0.018 | 0.340 | 0.562 | 0.806 | 1.000       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.608 | 0.864 | 0.564 | 0.942 | 1.000       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.562 | 0.701 | 0.439 | 0.626 | 1.000       |
| SSD_512_VGG16_Atrous_COCO   | ./_fixtures/lane_control.jpg           | car   | 0.014 | 0.346 | 0.573 | 0.804 | 0.999       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.706 | 0.993 | 0.594 | 0.986 | 0.999       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.288 | 0.436 | 0.502 | 0.601 | 0.993       |
| SSD_512_VGG16_Atrous_VOC    | ./_fixtures/lane_control.jpg           | car   | 0.562 | 0.697 | 0.435 | 0.620 | 1.000       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.680 | 1.001 | 0.573 | 0.994 | 1.000       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.019 | 0.346 | 0.578 | 0.812 | 1.000       |
