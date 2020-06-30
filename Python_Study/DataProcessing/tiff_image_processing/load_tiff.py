import numpy as np
import cv2
import matplotlib.pyplot as plt

img = cv2.imread('meta_data\\tiff_test.tiff')
# print(img.shape)  # [height, width, color_channels]
gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
# print(gray.shape)  # [height, width], only one color channel
gray = gray[0:img.shape[1], 0:img.shape[1]]  # cut into a square picture
# cv2.imwrite('img.png', img)  # save to .png

# cv2.imshow("img",gray)
# cv2.waitKey(0)
# cv2.destroyAllWindows()

# fig = plt.figure("Image")  # 图像窗口名称
# ax = plt.gca()
# ax.xaxis.set_ticks_position('top')  # 将x轴显示在上方
plt.imshow(gray, cmap='gray')
# plt.axis('on')  # 关掉坐标轴为 off
# plt.title('image')  # 图像题目
plt.show()
