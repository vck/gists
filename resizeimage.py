from PIL import Image
img = Image.open("/home/dasta/Downloads/f2017441_a153_437a_9f0f_0e4dbd49ae56.jpg")
img = img.resize((int(3456/2), int(5184/2)), Image.ANTIALIAS)
img.save('resized.jpg')
img.show()
