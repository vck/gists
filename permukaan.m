pkg load image

[x, cmap_bmp] = imread("~/apel.jpeg");
x = rgb2gray(x);

s = 0;
v = 0;
d_old = 0;
cal_v = 4.381/247.667; % calibration in vertical direction
cal_h = 5.675/299.333; % calibration in horizontal direction
for xx = 1:225,
   d = 0;
   s1 = 0;
   v1 = 0;
   for yy = 1:225q,
      if x(yy, xx) == 1
         d = d+1;
      end
   end

   s1 = pi*(d_old+d)/2*cal_v*sqrt((d/2-d_old/2)^2*cal_v^2+cal_h^2)
   v1 = pi/12*cal_h*(cal_v^2)*(d_old^2+d_old*d+d^2)
   s = s+s1;
   v = v+v1;
   d_old = d;
end

