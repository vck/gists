import PIL
from PIL.ExifTags import TAGS

def getdata(img):
	ret {}
	i = Image.open(img)
	info = i._getexif()
	try:
		for tag, value in info.items():
			decoded = TAGS.get(tag, tag)
			ret[decoded] = value
	except:
		print 'cannot read EXIF data'
	return ret

print getdata(pathtoimgfile/file)




