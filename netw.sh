sudo wpa_supplicant -B -D nl80211 -i wlan0 -c /etc/wpa_supplicant/wpa_supplicant.conf 
sudo dhclient -r
sudo dhclient wlan0

