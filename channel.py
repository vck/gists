#not efective for the 14th channel
def frequency(channel):return (((channel-1)*5)*0.001) + 2.412
for i in range(1, 14):print i, frequency(i)
