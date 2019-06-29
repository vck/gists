import numpy as np
from skimage.color import rgb2gray
from skimage.data import imread
import matplotlib.pyplot as plt
import sys

def imshow(img):
   plt.imshow(img, cmap=plt.cm.gray)
   plt.show()

img = imread(sys.argv[1])
img_gray = rgb2gray(img)
img_grad = np.gradient(img_gray)

imshow(img_grad[0]+img_grad[1])
