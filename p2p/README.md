در آرايه peers فايل node.go بوت نودهاي IPFS لوکال را تعريف کنيد. به منظور تعريف نود بوتسترپ سفارشي IPFS مراحل زير را دنبال کنيد:

1) دانلود و نصب Go

1.1) خارج کردن فايل دانلود شده از حالت فشرده

sudo tar -C /usr/local -xzf go1.19.1.linux-amd64.tar.gz


1.2) اضافه کردن مسير نصب به PATH

export PATH=$PATH:/usr/local/go/bin


1.3) اعمال تغييرات

source $HOME/.profile


2) نصب IPFS 


2.1) اجراي فرمان هاي زير:

sudo apt-get update

wget https://dist.ipfs.io/go-ipfs/v0.17.0/go-ipfs_v0.17.0_linux-amd64.tar.gz

tar xvfz go-ipfs_v0.17.0_linux-amd64.tar.gz

sudo mv go-ipfs/ipfs /usr/local/bin/ipfs

2.2) کنترل صحت نصب

ipfs version

3) مقدار دهي اوليه نود IPFS

IPFS_PATH=~/.ipfs ipfs init

4) بوتسترپ کردن نود

4.1) ابتدا تمام نودهاي بوتسترپ پيش فرض را حذف کنيد

IPFS_PATH=~/.ipfs ipfs bootstrap rm --all

4.2) پيدا کردن Peer IP و همچنين  Peer Identity

4.2.1) پيدا کردن IP

hostname -I

4.2.2) پيدا کردن Peer Identity

IPFS_PATH=~/.ipfs ipfs config show | grep "PeerID"

4.3) اضافه کردن نود بوتسترپ به IPFS

IPFS_PATH=~/.ipfs ipfs bootstrap add /ip4/<ip address of bootnode>/tcp/4001/ipfs/<peer identity hash of bootnode>

دستورات بندهاي 3 و 4.1 و 4.3 بايد روي ساير نودها اجرا شود

5) استارت شبکه

IPFS_PATH=~/.ipfs ipfs daemon & 


در صورتي که تنظيمات به درستي انجام شده باشد پس از اجراي فرمان زير، اطلاعات ID و IP هاي نودهاي ديگر در شبکه نمايش داده مي شود
IPFS_PATH=~/.ipfs ipfs swarm peers

  
تنظيمات مربوط به نودهاي بوتسترپ پشت NAT
  
  1) ابتدا در قسمت Port forwarding  يا Virtual server مودم روتر، پورت 4001 را به سمت IP نود بوتسترپ باز کنيد.
  2) فايل config موجود در مسير
  .ipfs/
  
  را به شرح زير تغيير دهيد:
  
      "AppendAnnounce": [
      "/ip4/x.x.x.x/tcp/4001",
      "/ip4/x.x.x.x/udp/4001/quic",
      "/ip4/x.x.x.x/udp/4001/quic/webtransport"
    ],
  
  آدرس IP نود بوتسترپ را با x.x.x.x جايگزين کنيد


